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

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Repository struct {
	Id         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	Protocol   types.String `tfsdk:"protocol"`
	Path       types.String `tfsdk:"path"`
	ServerName types.String `tfsdk:"server_name"`
	UserName   types.String `tfsdk:"user_name"`
	Password   types.String `tfsdk:"password"`
	EnablePki  types.Bool   `tfsdk:"enable_pki"`
}

func (data Repository) toBody(ctx context.Context, state Repository) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Protocol.IsNull() {
		body, _ = sjson.Set(body, "protocol", data.Protocol.ValueString())
	}
	if !data.Path.IsNull() {
		body, _ = sjson.Set(body, "path", data.Path.ValueString())
	}
	if !data.ServerName.IsNull() {
		body, _ = sjson.Set(body, "serverName", data.ServerName.ValueString())
	}
	if !data.UserName.IsNull() {
		body, _ = sjson.Set(body, "userName", data.UserName.ValueString())
	}
	if !data.Password.IsNull() {
		body, _ = sjson.Set(body, "password", data.Password.ValueString())
	}
	if !data.EnablePki.IsNull() {
		body, _ = sjson.Set(body, "enablePki", data.EnablePki.ValueBool())
	}
	return body
}

func (data *Repository) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("response.protocol"); value.Exists() {
		data.Protocol = types.StringValue(value.String())
	} else {
		data.Protocol = types.StringNull()
	}
	if value := res.Get("response.path"); value.Exists() {
		data.Path = types.StringValue(value.String())
	} else {
		data.Path = types.StringNull()
	}
	if value := res.Get("response.serverName"); value.Exists() {
		data.ServerName = types.StringValue(value.String())
	} else {
		data.ServerName = types.StringNull()
	}
	if value := res.Get("response.userName"); value.Exists() {
		data.UserName = types.StringValue(value.String())
	} else {
		data.UserName = types.StringNull()
	}
	if value := res.Get("response.enablePki"); value.Exists() {
		data.EnablePki = types.BoolValue(value.Bool())
	} else {
		data.EnablePki = types.BoolNull()
	}
}

func (data *Repository) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("response.protocol"); value.Exists() && !data.Protocol.IsNull() {
		data.Protocol = types.StringValue(value.String())
	} else {
		data.Protocol = types.StringNull()
	}
	if value := res.Get("response.path"); value.Exists() && !data.Path.IsNull() {
		data.Path = types.StringValue(value.String())
	} else {
		data.Path = types.StringNull()
	}
	if value := res.Get("response.serverName"); value.Exists() && !data.ServerName.IsNull() {
		data.ServerName = types.StringValue(value.String())
	} else {
		data.ServerName = types.StringNull()
	}
	if value := res.Get("response.userName"); value.Exists() && !data.UserName.IsNull() {
		data.UserName = types.StringValue(value.String())
	} else {
		data.UserName = types.StringNull()
	}
	if value := res.Get("response.enablePki"); value.Exists() && !data.EnablePki.IsNull() {
		data.EnablePki = types.BoolValue(value.Bool())
	} else {
		data.EnablePki = types.BoolNull()
	}
}
