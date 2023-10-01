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

type LicenseTierState struct {
	Id       types.String               `tfsdk:"id"`
	Licenses []LicenseTierStateLicenses `tfsdk:"licenses"`
}

type LicenseTierStateLicenses struct {
	Name   types.String `tfsdk:"name"`
	Status types.String `tfsdk:"status"`
}

func (data LicenseTierState) getPath() string {
	return "/api/v1/license/system/tier-state"
}

func (data LicenseTierState) toBody(ctx context.Context, state LicenseTierState) string {
	body := "[]"
	if len(data.Licenses) > 0 {
		body, _ = sjson.Set(body, "", []interface{}{})
		for _, item := range data.Licenses {
			itemBody := ""
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.Status.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "status", item.Status.ValueString())
			}
			body, _ = sjson.SetRaw(body, "-1", itemBody)
		}
	}
	return body
}

func (data *LicenseTierState) fromBody(ctx context.Context, res gjson.Result) {
	if value := res; value.Exists() {
		data.Licenses = make([]LicenseTierStateLicenses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := LicenseTierStateLicenses{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("status"); cValue.Exists() {
				item.Status = types.StringValue(cValue.String())
			} else {
				item.Status = types.StringNull()
			}
			data.Licenses = append(data.Licenses, item)
			return true
		})
	}
}

func (data *LicenseTierState) updateFromBody(ctx context.Context, res gjson.Result) {
	for i := range data.Licenses {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Licenses[i].Name.ValueString()}

		var r gjson.Result
		res.ForEach(
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
		if value := r.Get("name"); value.Exists() && !data.Licenses[i].Name.IsNull() {
			data.Licenses[i].Name = types.StringValue(value.String())
		} else {
			data.Licenses[i].Name = types.StringNull()
		}
		if value := r.Get("status"); value.Exists() && !data.Licenses[i].Status.IsNull() {
			data.Licenses[i].Status = types.StringValue(value.String())
		} else {
			data.Licenses[i].Status = types.StringNull()
		}
	}
}
