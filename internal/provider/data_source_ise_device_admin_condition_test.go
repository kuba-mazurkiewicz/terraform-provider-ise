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
func TestAccDataSourceIseDeviceAdminCondition(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_device_admin_condition.test", "name", "Cond1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_device_admin_condition.test", "description", "My description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_device_admin_condition.test", "condition_type", "LibraryConditionAttributes"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_device_admin_condition.test", "is_negate", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_device_admin_condition.test", "attribute_name", "EapAuthentication"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_device_admin_condition.test", "attribute_value", "EAP-TLS"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_device_admin_condition.test", "dictionary_name", "Network Access"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_device_admin_condition.test", "operator", "equals"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIseDeviceAdminConditionConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

//template:end testAccDataSource

//template:begin testPrerequisites
//template:end testPrerequisites

//template:begin testAccDataSourceConfig
func testAccDataSourceIseDeviceAdminConditionConfig() string {
	config := `resource "ise_device_admin_condition" "test" {` + "\n"
	config += `	name = "Cond1"` + "\n"
	config += `	description = "My description"` + "\n"
	config += `	condition_type = "LibraryConditionAttributes"` + "\n"
	config += `	is_negate = false` + "\n"
	config += `	attribute_name = "EapAuthentication"` + "\n"
	config += `	attribute_value = "EAP-TLS"` + "\n"
	config += `	dictionary_name = "Network Access"` + "\n"
	config += `	operator = "equals"` + "\n"
	config += `}` + "\n"

	config += `
		data "ise_device_admin_condition" "test" {
			id = ise_device_admin_condition.test.id
		}
	`
	return config
}

//template:end testAccDataSourceConfig