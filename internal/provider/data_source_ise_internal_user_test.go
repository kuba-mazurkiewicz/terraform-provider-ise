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
func TestAccDataSourceIseInternalUser(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_internal_user.test", "name", "UserTF"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_internal_user.test", "change_password", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_internal_user.test", "email", "aaa@cisco.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_internal_user.test", "account_name_alias", "User 1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_internal_user.test", "enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_internal_user.test", "password_never_expires", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_internal_user.test", "first_name", "John"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_internal_user.test", "last_name", "Doe"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_internal_user.test", "password_id_store", "Internal Users"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_internal_user.test", "description", "My first Terraform user"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIseInternalUserConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

//template:end testAccDataSource

//template:begin testPrerequisites
//template:end testPrerequisites

//template:begin testAccDataSourceConfig
func testAccDataSourceIseInternalUserConfig() string {
	config := `resource "ise_internal_user" "test" {` + "\n"
	config += `	name = "UserTF"` + "\n"
	config += `	password = "Cisco123"` + "\n"
	config += `	change_password = true` + "\n"
	config += `	email = "aaa@cisco.com"` + "\n"
	config += `	account_name_alias = "User 1"` + "\n"
	config += `	enable_password = "Cisco123"` + "\n"
	config += `	enabled = true` + "\n"
	config += `	password_never_expires = false` + "\n"
	config += `	first_name = "John"` + "\n"
	config += `	last_name = "Doe"` + "\n"
	config += `	password_id_store = "Internal Users"` + "\n"
	config += `	description = "My first Terraform user"` + "\n"
	config += `}` + "\n"

	config += `
		data "ise_internal_user" "test" {
			id = ise_internal_user.test.id
		}
	`
	return config
}

//template:end testAccDataSourceConfig
