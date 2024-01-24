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
type TrustSecIPToSGTMapping struct {
	Id           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	Description  types.String `tfsdk:"description"`
	DeployTo     types.String `tfsdk:"deploy_to"`
	DeployType   types.String `tfsdk:"deploy_type"`
	HostName     types.String `tfsdk:"host_name"`
	HostIp       types.String `tfsdk:"host_ip"`
	Sgt          types.String `tfsdk:"sgt"`
	MappingGroup types.String `tfsdk:"mapping_group"`
}

//template:end types

//template:begin getPath
func (data TrustSecIPToSGTMapping) getPath() string {
	return "/ers/config/sgmapping"
}

//template:end getPath

//template:begin getPathPut

//template:end getPathPut

//template:begin toBody
func (data TrustSecIPToSGTMapping) toBody(ctx context.Context, state TrustSecIPToSGTMapping) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "SGMapping.name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "SGMapping.description", data.Description.ValueString())
	}
	if !data.DeployTo.IsNull() {
		body, _ = sjson.Set(body, "SGMapping.deployTo", data.DeployTo.ValueString())
	}
	if !data.DeployType.IsNull() {
		body, _ = sjson.Set(body, "SGMapping.deployType", data.DeployType.ValueString())
	}
	if !data.HostName.IsNull() {
		body, _ = sjson.Set(body, "SGMapping.hostName", data.HostName.ValueString())
	}
	if !data.HostIp.IsNull() {
		body, _ = sjson.Set(body, "SGMapping.hostIp", data.HostIp.ValueString())
	}
	if !data.Sgt.IsNull() {
		body, _ = sjson.Set(body, "SGMapping.sgt", data.Sgt.ValueString())
	}
	if !data.MappingGroup.IsNull() {
		body, _ = sjson.Set(body, "SGMapping.mappingGroup", data.MappingGroup.ValueString())
	}
	return body
}

//template:end toBody

//template:begin fromBody
func (data *TrustSecIPToSGTMapping) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("SGMapping.name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("SGMapping.description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("SGMapping.deployTo"); value.Exists() {
		data.DeployTo = types.StringValue(value.String())
	} else {
		data.DeployTo = types.StringNull()
	}
	if value := res.Get("SGMapping.deployType"); value.Exists() {
		data.DeployType = types.StringValue(value.String())
	} else {
		data.DeployType = types.StringNull()
	}
	if value := res.Get("SGMapping.hostName"); value.Exists() {
		data.HostName = types.StringValue(value.String())
	} else {
		data.HostName = types.StringNull()
	}
	if value := res.Get("SGMapping.hostIp"); value.Exists() {
		data.HostIp = types.StringValue(value.String())
	} else {
		data.HostIp = types.StringNull()
	}
	if value := res.Get("SGMapping.sgt"); value.Exists() {
		data.Sgt = types.StringValue(value.String())
	} else {
		data.Sgt = types.StringNull()
	}
	if value := res.Get("SGMapping.mappingGroup"); value.Exists() {
		data.MappingGroup = types.StringValue(value.String())
	} else {
		data.MappingGroup = types.StringNull()
	}
}

//template:end fromBody

//template:begin updateFromBody
func (data *TrustSecIPToSGTMapping) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("SGMapping.name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("SGMapping.description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("SGMapping.deployTo"); value.Exists() && !data.DeployTo.IsNull() {
		data.DeployTo = types.StringValue(value.String())
	} else {
		data.DeployTo = types.StringNull()
	}
	if value := res.Get("SGMapping.deployType"); value.Exists() && !data.DeployType.IsNull() {
		data.DeployType = types.StringValue(value.String())
	} else {
		data.DeployType = types.StringNull()
	}
	if value := res.Get("SGMapping.hostName"); value.Exists() && !data.HostName.IsNull() {
		data.HostName = types.StringValue(value.String())
	} else {
		data.HostName = types.StringNull()
	}
	if value := res.Get("SGMapping.hostIp"); value.Exists() && !data.HostIp.IsNull() {
		data.HostIp = types.StringValue(value.String())
	} else {
		data.HostIp = types.StringNull()
	}
	if value := res.Get("SGMapping.sgt"); value.Exists() && !data.Sgt.IsNull() {
		data.Sgt = types.StringValue(value.String())
	} else {
		data.Sgt = types.StringNull()
	}
	if value := res.Get("SGMapping.mappingGroup"); value.Exists() && !data.MappingGroup.IsNull() {
		data.MappingGroup = types.StringValue(value.String())
	} else {
		data.MappingGroup = types.StringNull()
	}
}

//template:end updateFromBody
