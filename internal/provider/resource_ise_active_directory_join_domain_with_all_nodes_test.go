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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

//template:end imports

//template:begin testAcc
func TestAccIseActiveDirectoryJoinDomainWithAllNodes(t *testing.T) {
	if os.Getenv("AD") == "" {
		t.Skip("skipping test, set environment variable AD")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("ise_active_directory_join_domain_with_all_nodes.test", "join_point_id", "73808580-b6e6-11ee-8960-de6d7692bc40"))
	checks = append(checks, resource.TestCheckResourceAttr("ise_active_directory_join_domain_with_all_nodes.test", "additional_data.0.name", "username"))
	checks = append(checks, resource.TestCheckResourceAttr("ise_active_directory_join_domain_with_all_nodes.test", "additional_data.0.value", "administrator"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIseActiveDirectoryJoinDomainWithAllNodesConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIseActiveDirectoryJoinDomainWithAllNodesConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

//template:end testAcc

//template:begin testPrerequisites
//template:end testPrerequisites

//template:begin testAccConfigMinimal
func testAccIseActiveDirectoryJoinDomainWithAllNodesConfig_minimum() string {
	config := `resource "ise_active_directory_join_domain_with_all_nodes" "test" {` + "\n"
	config += `	join_point_id = "73808580-b6e6-11ee-8960-de6d7692bc40"` + "\n"
	config += `	additional_data = [{` + "\n"
	config += `	  name = "username"` + "\n"
	config += `	  value = "administrator"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

//template:end testAccConfigMinimal

//template:begin testAccConfigAll
func testAccIseActiveDirectoryJoinDomainWithAllNodesConfig_all() string {
	config := `resource "ise_active_directory_join_domain_with_all_nodes" "test" {` + "\n"
	config += `	join_point_id = "73808580-b6e6-11ee-8960-de6d7692bc40"` + "\n"
	config += `	additional_data = [{` + "\n"
	config += `	  name = "username"` + "\n"
	config += `	  value = "administrator"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

//template:end testAccConfigAll
