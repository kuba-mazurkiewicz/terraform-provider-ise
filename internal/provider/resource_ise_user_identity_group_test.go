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
func TestAccIseUserIdentityGroup(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("ise_user_identity_group.test", "name", "Group1"))
	checks = append(checks, resource.TestCheckResourceAttr("ise_user_identity_group.test", "description", "My endpoint identity group"))
	checks = append(checks, resource.TestCheckResourceAttr("ise_user_identity_group.test", "parent", "NAC Group:NAC:IdentityGroups:User Identity Groups"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIseUserIdentityGroupConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIseUserIdentityGroupConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "ise_user_identity_group.test",
		ImportState:  true,
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
func testAccIseUserIdentityGroupConfig_minimum() string {
	config := `resource "ise_user_identity_group" "test" {` + "\n"
	config += `	name = "Group1"` + "\n"
	config += `}` + "\n"
	return config
}

//template:end testAccConfigMinimal

//template:begin testAccConfigAll
func testAccIseUserIdentityGroupConfig_all() string {
	config := `resource "ise_user_identity_group" "test" {` + "\n"
	config += `	name = "Group1"` + "\n"
	config += `	description = "My endpoint identity group"` + "\n"
	config += `	parent = "NAC Group:NAC:IdentityGroups:User Identity Groups"` + "\n"
	config += `}` + "\n"
	return config
}

//template:end testAccConfigAll
