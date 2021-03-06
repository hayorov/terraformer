// Copyright 2018 The Terraformer Authors.
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

package aws

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
)

var route53AllowEmptyValues = []string{}

var route53AdditionalFields = map[string]interface{}{}

type Route53Generator struct {
	AWSService
}

func (g *Route53Generator) createZonesResources(svc *route53.Client) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	p := route53.NewListHostedZonesPaginator(svc.ListHostedZonesRequest(&route53.ListHostedZonesInput{}))
	for p.Next(context.Background()) {
		for _, zone := range p.CurrentPage().HostedZones {
			zoneID := cleanZoneID(aws.StringValue(zone.Id))
			resources = append(resources, terraform_utils.NewResource(
				zoneID,
				zoneID+"_"+strings.TrimSuffix(aws.StringValue(zone.Name), "."),
				"aws_route53_zone",
				"aws",
				map[string]string{
					"name":          aws.StringValue(zone.Name),
					"force_destroy": "false",
				},
				route53AllowEmptyValues,
				route53AdditionalFields,
			))
			records := g.createRecordsResources(svc, zoneID)
			resources = append(resources, records...)
		}
	}
	if err := p.Err(); err != nil {
		log.Println(err)
	}
	return resources
}

func (Route53Generator) createRecordsResources(svc *route53.Client, zoneID string) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	listParams := &route53.ListResourceRecordSetsInput{
		HostedZoneId: aws.String(zoneID),
	}

	p := route53.NewListResourceRecordSetsPaginator(svc.ListResourceRecordSetsRequest(listParams))
	for p.Next(context.Background()) {
		for _, record := range p.CurrentPage().ResourceRecordSets {
			recordName := wildcardUnescape(aws.StringValue(record.Name))
			typeString, _ := record.Type.MarshalValue()
			resources = append(resources, terraform_utils.NewResource(
				fmt.Sprintf("%s_%s_%s_%s", zoneID, recordName, typeString, aws.StringValue(record.SetIdentifier)),
				fmt.Sprintf("%s_%s_%s_%s", zoneID, recordName, typeString, aws.StringValue(record.SetIdentifier)),
				"aws_route53_record",
				"aws",
				map[string]string{
					"name":           strings.TrimSuffix(recordName, "."),
					"zone_id":        zoneID,
					"type":           typeString,
					"set_identifier": aws.StringValue(record.SetIdentifier),
				},
				route53AllowEmptyValues,
				route53AdditionalFields,
			))
		}
	}
	if err := p.Err(); err != nil {
		log.Println(err)
		return []terraform_utils.Resource{}
	}
	return resources
}

// Generate TerraformResources from AWS API,
// create terraform resource for each zone + each record
func (g *Route53Generator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := route53.New(config)

	g.Resources = g.createZonesResources(svc)
	return nil
}

func (g *Route53Generator) PostConvertHook() error {
	for i, resourceRecord := range g.Resources {
		if resourceRecord.InstanceInfo.Type == "aws_route53_zone" {
			continue
		}
		item := resourceRecord.Item
		zoneID := item["zone_id"].(string)
		for _, resourceZone := range g.Resources {
			if resourceZone.InstanceInfo.Type != "aws_route53_zone" {
				continue
			}
			if zoneID == resourceZone.InstanceState.ID {
				g.Resources[i].Item["zone_id"] = "${aws_route53_zone." + resourceZone.ResourceName + ".zone_id}"
			}
		}
		if _, aliasExist := resourceRecord.Item["alias"]; aliasExist {
			if _, ttlExist := resourceRecord.Item["ttl"]; ttlExist {
				delete(g.Resources[i].Item, "ttl")
			}
		}
	}
	return nil
}

func wildcardUnescape(s string) string {
	if strings.Contains(s, "\\052") {
		s = strings.Replace(s, "\\052", "*", 1)
	}
	return s
}

// cleanZoneID is used to remove the leading /hostedzone/
func cleanZoneID(ID string) string {
	return cleanPrefix(ID, "/hostedzone/")
}

// cleanPrefix removes a string prefix from an ID
func cleanPrefix(ID, prefix string) string {
	return strings.TrimPrefix(ID, prefix)
}
