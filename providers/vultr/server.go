// Copyright 2019 The Terraformer Authors.
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

package vultr

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/vultr/govultr"
)

type ServerGenerator struct {
	VultrService
}

func (g ServerGenerator) createResources(serverList []govultr.Server) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for _, server := range serverList {
		resources = append(resources, terraform_utils.NewSimpleResource(
			server.InstanceID,
			server.InstanceID,
			"vultr_server",
			"vultr",
			[]string{}))
	}
	return resources
}

func (g *ServerGenerator) InitResources() error {
	client := g.generateClient()
	output, err := client.Server.List(context.Background())
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
