// Copyright Nitric Pty Ltd.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package deploy

import (
	"encoding/json"
	"fmt"
	"runtime/debug"

	deploy "github.com/nitrictech/nitric/core/pkg/api/nitric/deploy/v1"
)

type DownStreamMessageWriter struct {
	stream deploy.DeployService_DownServer
}

func (s *DownStreamMessageWriter) Write(bytes []byte) (int, error) {
	err := s.stream.Send(&deploy.DeployDownEvent{
		Content: &deploy.DeployDownEvent_Message{
			Message: &deploy.DeployEventMessage{
				Message: string(bytes),
			},
		},
	})
	if err != nil {
		return 0, err
	}

	return len(bytes), nil
}

func (d *DeployServer) Down(request *deploy.DeployDownRequest, stream deploy.DeployService_DownServer) error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			stack := string(debug.Stack())
			err = fmt.Errorf("recovered panic: %+v\n Stack: %s", r, stack)
		}
	}()

	reqJson, err := json.MarshalIndent(request, "", "    ")
	if err != nil {
		return err
	}

	fmt.Println(reqJson)

	return nil
}
