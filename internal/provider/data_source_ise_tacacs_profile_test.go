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
func TestAccDataSourceIseTACACSProfile(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_tacacs_profile.test", "name", "Profile1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_tacacs_profile.test", "description", "My TACACS profile"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_tacacs_profile.test", "session_attributes.0.type", "MANDATORY"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_tacacs_profile.test", "session_attributes.0.name", "attr1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_tacacs_profile.test", "session_attributes.0.value", "value"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIseTACACSProfileConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

//template:end testAccDataSource

//template:begin testPrerequisites
//template:end testPrerequisites

//template:begin testAccDataSourceConfig
func testAccDataSourceIseTACACSProfileConfig() string {
	config := `resource "ise_tacacs_profile" "test" {` + "\n"
	config += `	name = "Profile1"` + "\n"
	config += `	description = "My TACACS profile"` + "\n"
	config += `	session_attributes = [{` + "\n"
	config += `	  type = "MANDATORY"` + "\n"
	config += `	  name = "attr1"` + "\n"
	config += `	  value = "value"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "ise_tacacs_profile" "test" {
			id = ise_tacacs_profile.test.id
		}
	`
	return config
}

//template:end testAccDataSourceConfig