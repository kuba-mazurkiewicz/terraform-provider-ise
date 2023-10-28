// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

//template:begin imports
import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

//template:end imports

//template:begin testAccDataSource
func TestAccDataSourceIseTrustSecIPToSGTMappingGroup(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_trustsec_ip_to_sgt_mapping_group.test", "name", "groupA"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_trustsec_ip_to_sgt_mapping_group.test", "deploy_type", "ALL"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_trustsec_ip_to_sgt_mapping_group.test", "sgt", "93e1bf00-8c01-11e6-996c-525400b48521"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIseTrustSecIPToSGTMappingGroupPrerequisitesConfig + testAccDataSourceIseTrustSecIPToSGTMappingGroupConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

//template:end testAccDataSource

//template:begin testPrerequisites
const testAccDataSourceIseTrustSecIPToSGTMappingGroupPrerequisitesConfig = `
resource "ise_trustsec_security_group" "test" {
  name              = "SGT1234"
  description       = "My SGT"
  value             = 1234
  propogate_to_apic = true
  is_read_only      = false
}

`

//template:end testPrerequisites

//template:begin testAccDataSourceConfig
func testAccDataSourceIseTrustSecIPToSGTMappingGroupConfig() string {
	config := `resource "ise_trustsec_ip_to_sgt_mapping_group" "test" {` + "\n"
	config += `	name = "groupA"` + "\n"
	config += `	deploy_type = "ALL"` + "\n"
	config += `	sgt = "93e1bf00-8c01-11e6-996c-525400b48521"` + "\n"
	config += `}` + "\n"

	config += `
		data "ise_trustsec_ip_to_sgt_mapping_group" "test" {
			id = ise_trustsec_ip_to_sgt_mapping_group.test.id
		}
	`
	return config
}

//template:end testAccDataSourceConfig
