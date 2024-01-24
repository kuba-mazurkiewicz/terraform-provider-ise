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
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

//template:end imports

//template:begin types
type UserIdentityGroup struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Parent      types.String `tfsdk:"parent"`
}

//template:end types

//template:begin getPath
func (data UserIdentityGroup) getPath() string {
	return "/ers/config/identitygroup"
}

//template:end getPath

//template:begin getPathPut

//template:end getPathPut

//template:begin toBody
func (data UserIdentityGroup) toBody(ctx context.Context, state UserIdentityGroup) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "IdentityGroup.name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "IdentityGroup.description", data.Description.ValueString())
	}
	if !data.Parent.IsNull() {
		body, _ = sjson.Set(body, "IdentityGroup.parent", data.Parent.ValueString())
	}
	return body
}

//template:end toBody

//template:begin fromBody
func (data *UserIdentityGroup) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("IdentityGroup.name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("IdentityGroup.description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("IdentityGroup.parent"); value.Exists() {
		data.Parent = types.StringValue(value.String())
	} else {
		data.Parent = types.StringNull()
	}
}

//template:end fromBody

//template:begin updateFromBody
func (data *UserIdentityGroup) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("IdentityGroup.name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("IdentityGroup.description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("IdentityGroup.parent"); value.Exists() && !data.Parent.IsNull() {
		data.Parent = types.StringValue(value.String())
	} else {
		data.Parent = types.StringNull()
	}
}

//template:end updateFromBody
