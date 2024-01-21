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
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

//template:end imports

//template:begin types
type ActiveDirectoryJoinDomainWithAllNodes struct {
	Id             types.String                                          `tfsdk:"id"`
	JoinPointId    types.String                                          `tfsdk:"join_point_id"`
	AdditionalData []ActiveDirectoryJoinDomainWithAllNodesAdditionalData `tfsdk:"additional_data"`
}

type ActiveDirectoryJoinDomainWithAllNodesAdditionalData struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

//template:end types

//template:begin getPath
func (data ActiveDirectoryJoinDomainWithAllNodes) getPath() string {
	return fmt.Sprintf("/ers/config/activedirectory/%v/joinAllNodes", data.JoinPointId.ValueString())
}

//template:end getPath

//template:begin toBody
func (data ActiveDirectoryJoinDomainWithAllNodes) toBody(ctx context.Context, state ActiveDirectoryJoinDomainWithAllNodes) string {
	body := ""
	if len(data.AdditionalData) > 0 {
		body, _ = sjson.Set(body, "OperationAdditionalData.additionalData", []interface{}{})
		for _, item := range data.AdditionalData {
			itemBody := ""
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.Value.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "value", item.Value.ValueString())
			}
			body, _ = sjson.SetRaw(body, "OperationAdditionalData.additionalData.-1", itemBody)
		}
	}
	return body
}

//template:end toBody

//template:begin fromBody
func (data *ActiveDirectoryJoinDomainWithAllNodes) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("OperationAdditionalData.additionalData"); value.Exists() {
		data.AdditionalData = make([]ActiveDirectoryJoinDomainWithAllNodesAdditionalData, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ActiveDirectoryJoinDomainWithAllNodesAdditionalData{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("value"); cValue.Exists() {
				item.Value = types.StringValue(cValue.String())
			} else {
				item.Value = types.StringNull()
			}
			data.AdditionalData = append(data.AdditionalData, item)
			return true
		})
	}
}

//template:end fromBody

//template:begin updateFromBody
func (data *ActiveDirectoryJoinDomainWithAllNodes) updateFromBody(ctx context.Context, res gjson.Result) {
	for i := range data.AdditionalData {
		keys := [...]string{"name"}
		keyValues := [...]string{data.AdditionalData[i].Name.ValueString()}

		var r gjson.Result
		res.Get("OperationAdditionalData.additionalData").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("name"); value.Exists() && !data.AdditionalData[i].Name.IsNull() {
			data.AdditionalData[i].Name = types.StringValue(value.String())
		} else {
			data.AdditionalData[i].Name = types.StringNull()
		}
		if value := r.Get("value"); value.Exists() && !data.AdditionalData[i].Value.IsNull() {
			data.AdditionalData[i].Value = types.StringValue(value.String())
		} else {
			data.AdditionalData[i].Value = types.StringNull()
		}
	}
}

//template:end updateFromBody
