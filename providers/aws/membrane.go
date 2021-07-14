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

	"github.com/nitric-dev/membrane/membrane"
	document "github.com/nitric-dev/membrane/plugins/document/dynamodb"
	eventing "github.com/nitric-dev/membrane/plugins/eventing/sns"
	httpGateway "github.com/nitric-dev/membrane/plugins/gateway/ecs"
	lambdaGateway "github.com/nitric-dev/membrane/plugins/gateway/lambda"
	kv_dynamodb "github.com/nitric-dev/membrane/plugins/kv/dynamodb"
	queue "github.com/nitric-dev/membrane/plugins/queue/sqs"
	storage "github.com/nitric-dev/membrane/plugins/storage/s3"
	"github.com/nitric-dev/membrane/sdk"
	"github.com/nitric-dev/membrane/utils"
)

func main() {
	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	signal.Notify(term, os.Interrupt, syscall.SIGINT)

	gatewayEnv := utils.GetEnv("GATEWAY_ENVIRONMENT", "lambda")

	// Load the appropriate gateway based on the environment.
	var gatewayPlugin sdk.GatewayService
	switch gatewayEnv {
	case "lambda":
		gatewayPlugin, _ = lambdaGateway.New()
	default:
		gatewayPlugin, _ = httpGateway.New()
	}
	documentPlugin, _ := document.New()
	eventingPlugin, _ := eventing.New()
	keyValuePlugin, _ := kv_dynamodb.New()
	queuePlugin, _ := queue.New()
	storagePlugin, _ := storage.New()

	m, err := membrane.New(&membrane.MembraneOptions{
		DocumentPlugin: documentPlugin,
		EventingPlugin: eventingPlugin,
		GatewayPlugin:  gatewayPlugin,
		KvPlugin:       keyValuePlugin,
		QueuePlugin:    queuePlugin,
		StoragePlugin:  storagePlugin,
	})

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
		fmt.Println(fmt.Sprintf("Membrane Error: %v, exiting", membraneError))
	case sigTerm := <-term:
		fmt.Println(fmt.Sprintf("Received %v, exiting", sigTerm))
	}

	m.Stop()
}
