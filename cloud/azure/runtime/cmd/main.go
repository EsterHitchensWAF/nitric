// Copyright 2021 Nitric Pty Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nitrictech/nitric/cloud/azure/runtime/core"
	azqueue_service "github.com/nitrictech/nitric/cloud/azure/runtime/queue"

	mongodb_service "github.com/nitrictech/nitric/cloud/azure/runtime/document"
	event_grid "github.com/nitrictech/nitric/cloud/azure/runtime/events"
	http_service "github.com/nitrictech/nitric/cloud/azure/runtime/gateway"
	key_vault "github.com/nitrictech/nitric/cloud/azure/runtime/secret"
	azblob_service "github.com/nitrictech/nitric/cloud/azure/runtime/storage"
	"github.com/nitrictech/nitric/core/pkg/membrane"
)

func main() {
	var err error
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	signal.Notify(term, os.Interrupt, syscall.SIGINT)

	provider, err := core.New()
	if err != nil {
		log.Fatalf("could not create core azure provider: %v", err)
	}

	membraneOpts := membrane.DefaultMembraneOptions()

	membraneOpts.DocumentPlugin, err = mongodb_service.New()
	if err != nil {
		log.Default().Println("Failed to load document plugin:", err.Error())
	}

	membraneOpts.EventsPlugin, err = event_grid.New(provider)
	if err != nil {
		log.Default().Println("Failed to load event plugin:", err.Error())
	}

	membraneOpts.GatewayPlugin, err = http_service.New(provider)
	if err != nil {
		log.Default().Println("Failed to load gateway plugin:", err.Error())
	}

	membraneOpts.QueuePlugin, _ = azqueue_service.New()
	if err != nil {
		log.Default().Println("Failed to load queue plugin:", err.Error())
	}

	membraneOpts.StoragePlugin, err = azblob_service.New()
	if err != nil {
		log.Default().Println("Failed to load storage plugin:", err.Error())
	}

	membraneOpts.SecretPlugin, err = key_vault.New()
	if err != nil {
		log.Default().Println("Failed to load secret plugin:", err.Error())
	}

	membraneOpts.ResourcesPlugin = provider

	m, err := membrane.New(membraneOpts)
	if err != nil {
		log.Fatalf("There was an error initialising the membrane server: %v", err)
	}

	errChan := make(chan error)
	// Start the Membrane server
	go func(chan error) {
		errChan <- m.Start()
	}(errChan)

	select {
	case membraneError := <-errChan:
		fmt.Printf("Membrane Error: %v, exiting\n", membraneError)
	case sigTerm := <-term:
		fmt.Printf("Received %v, exiting\n", sigTerm)
	}

	m.Stop()
}
