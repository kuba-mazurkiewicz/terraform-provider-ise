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
func TestAccDataSourceIseIdentitySourceSequence(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_identity_source_sequence.test", "name", "Sequence1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_identity_source_sequence.test", "description", "My identity source sequence"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_identity_source_sequence.test", "break_on_store_fail", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_identity_source_sequence.test", "certificate_authentication_profile", "Preloaded_Certificate_Profile"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_identity_source_sequence.test", "identity_sources.0.name", "Internal Users"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_identity_source_sequence.test", "identity_sources.0.order", "1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIseIdentitySourceSequenceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

//template:end testAccDataSource

//template:begin testPrerequisites
//template:end testPrerequisites

//template:begin testAccDataSourceConfig
func testAccDataSourceIseIdentitySourceSequenceConfig() string {
	config := `resource "ise_identity_source_sequence" "test" {` + "\n"
	config += `	name = "Sequence1"` + "\n"
	config += `	description = "My identity source sequence"` + "\n"
	config += `	break_on_store_fail = true` + "\n"
	config += `	certificate_authentication_profile = "Preloaded_Certificate_Profile"` + "\n"
	config += `	identity_sources = [{` + "\n"
	config += `	  name = "Internal Users"` + "\n"
	config += `	  order = 1` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "ise_identity_source_sequence" "test" {
			id = ise_identity_source_sequence.test.id
		}
	`
	return config
}

//template:end testAccDataSourceConfig
