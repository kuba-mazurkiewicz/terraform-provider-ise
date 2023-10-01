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
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type NetworkAccessAuthenticationRule struct {
	Id                       types.String `tfsdk:"id"`
	PolicySetId              types.String `tfsdk:"policy_set_id"`
	Name                     types.String `tfsdk:"name"`
	Default                  types.Bool   `tfsdk:"default"`
	Rank                     types.Int64  `tfsdk:"rank"`
	State                    types.String `tfsdk:"state"`
	ConditionType            types.String `tfsdk:"condition_type"`
	ConditionIsNegate        types.Bool   `tfsdk:"condition_is_negate"`
	ConditionAttributeName   types.String `tfsdk:"condition_attribute_name"`
	ConditionAttributeValue  types.String `tfsdk:"condition_attribute_value"`
	ConditionDictionaryName  types.String `tfsdk:"condition_dictionary_name"`
	ConditionDictionaryValue types.String `tfsdk:"condition_dictionary_value"`
	ConditionOperator        types.String `tfsdk:"condition_operator"`
	IdentitySourceName       types.String `tfsdk:"identity_source_name"`
	IfAuthFail               types.String `tfsdk:"if_auth_fail"`
	IfProcessFail            types.String `tfsdk:"if_process_fail"`
	IfUserNotFound           types.String `tfsdk:"if_user_not_found"`
}

func (data NetworkAccessAuthenticationRule) getPath() string {
	return fmt.Sprintf("/api/v1/policy/network-access/policy-set/%v/authentication", data.PolicySetId.ValueString())
}

func (data NetworkAccessAuthenticationRule) toBody(ctx context.Context, state NetworkAccessAuthenticationRule) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "rule.name", data.Name.ValueString())
	}
	if !data.Default.IsNull() {
		body, _ = sjson.Set(body, "rule.default", data.Default.ValueBool())
	}
	if !data.Rank.IsNull() {
		body, _ = sjson.Set(body, "rule.rank", data.Rank.ValueInt64())
	}
	if !data.State.IsNull() {
		body, _ = sjson.Set(body, "rule.state", data.State.ValueString())
	}
	if !data.ConditionType.IsNull() {
		body, _ = sjson.Set(body, "rule.condition.conditionType", data.ConditionType.ValueString())
	}
	if !data.ConditionIsNegate.IsNull() {
		body, _ = sjson.Set(body, "rule.condition.isNegate", data.ConditionIsNegate.ValueBool())
	}
	if !data.ConditionAttributeName.IsNull() {
		body, _ = sjson.Set(body, "rule.condition.attributeName", data.ConditionAttributeName.ValueString())
	}
	if !data.ConditionAttributeValue.IsNull() {
		body, _ = sjson.Set(body, "rule.condition.attributeValue", data.ConditionAttributeValue.ValueString())
	}
	if !data.ConditionDictionaryName.IsNull() {
		body, _ = sjson.Set(body, "rule.condition.dictionaryName", data.ConditionDictionaryName.ValueString())
	}
	if !data.ConditionDictionaryValue.IsNull() {
		body, _ = sjson.Set(body, "rule.condition.dictionaryValue", data.ConditionDictionaryValue.ValueString())
	}
	if !data.ConditionOperator.IsNull() {
		body, _ = sjson.Set(body, "rule.condition.operator", data.ConditionOperator.ValueString())
	}
	if !data.IdentitySourceName.IsNull() {
		body, _ = sjson.Set(body, "identitySourceName", data.IdentitySourceName.ValueString())
	}
	if !data.IfAuthFail.IsNull() {
		body, _ = sjson.Set(body, "ifAuthFail", data.IfAuthFail.ValueString())
	}
	if !data.IfProcessFail.IsNull() {
		body, _ = sjson.Set(body, "ifProcessFail", data.IfProcessFail.ValueString())
	}
	if !data.IfUserNotFound.IsNull() {
		body, _ = sjson.Set(body, "ifUserNotFound", data.IfUserNotFound.ValueString())
	}
	return body
}

func (data *NetworkAccessAuthenticationRule) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.rule.name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("response.rule.default"); value.Exists() {
		data.Default = types.BoolValue(value.Bool())
	} else {
		data.Default = types.BoolNull()
	}
	if value := res.Get("response.rule.rank"); value.Exists() {
		data.Rank = types.Int64Value(value.Int())
	} else {
		data.Rank = types.Int64Null()
	}
	if value := res.Get("response.rule.state"); value.Exists() {
		data.State = types.StringValue(value.String())
	} else {
		data.State = types.StringNull()
	}
	if value := res.Get("response.rule.condition.conditionType"); value.Exists() {
		data.ConditionType = types.StringValue(value.String())
	} else {
		data.ConditionType = types.StringNull()
	}
	if value := res.Get("response.rule.condition.isNegate"); value.Exists() {
		data.ConditionIsNegate = types.BoolValue(value.Bool())
	} else {
		data.ConditionIsNegate = types.BoolNull()
	}
	if value := res.Get("response.rule.condition.attributeName"); value.Exists() {
		data.ConditionAttributeName = types.StringValue(value.String())
	} else {
		data.ConditionAttributeName = types.StringNull()
	}
	if value := res.Get("response.rule.condition.attributeValue"); value.Exists() {
		data.ConditionAttributeValue = types.StringValue(value.String())
	} else {
		data.ConditionAttributeValue = types.StringNull()
	}
	if value := res.Get("response.rule.condition.dictionaryName"); value.Exists() {
		data.ConditionDictionaryName = types.StringValue(value.String())
	} else {
		data.ConditionDictionaryName = types.StringNull()
	}
	if value := res.Get("response.rule.condition.dictionaryValue"); value.Exists() {
		data.ConditionDictionaryValue = types.StringValue(value.String())
	} else {
		data.ConditionDictionaryValue = types.StringNull()
	}
	if value := res.Get("response.rule.condition.operator"); value.Exists() {
		data.ConditionOperator = types.StringValue(value.String())
	} else {
		data.ConditionOperator = types.StringNull()
	}
	if value := res.Get("response.identitySourceName"); value.Exists() {
		data.IdentitySourceName = types.StringValue(value.String())
	} else {
		data.IdentitySourceName = types.StringNull()
	}
	if value := res.Get("response.ifAuthFail"); value.Exists() {
		data.IfAuthFail = types.StringValue(value.String())
	} else {
		data.IfAuthFail = types.StringNull()
	}
	if value := res.Get("response.ifProcessFail"); value.Exists() {
		data.IfProcessFail = types.StringValue(value.String())
	} else {
		data.IfProcessFail = types.StringNull()
	}
	if value := res.Get("response.ifUserNotFound"); value.Exists() {
		data.IfUserNotFound = types.StringValue(value.String())
	} else {
		data.IfUserNotFound = types.StringNull()
	}
}

func (data *NetworkAccessAuthenticationRule) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.rule.name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("response.rule.default"); value.Exists() && !data.Default.IsNull() {
		data.Default = types.BoolValue(value.Bool())
	} else {
		data.Default = types.BoolNull()
	}
	if value := res.Get("response.rule.rank"); value.Exists() && !data.Rank.IsNull() {
		data.Rank = types.Int64Value(value.Int())
	} else {
		data.Rank = types.Int64Null()
	}
	if value := res.Get("response.rule.state"); value.Exists() && !data.State.IsNull() {
		data.State = types.StringValue(value.String())
	} else {
		data.State = types.StringNull()
	}
	if value := res.Get("response.rule.condition.conditionType"); value.Exists() && !data.ConditionType.IsNull() {
		data.ConditionType = types.StringValue(value.String())
	} else {
		data.ConditionType = types.StringNull()
	}
	if value := res.Get("response.rule.condition.isNegate"); value.Exists() && !data.ConditionIsNegate.IsNull() {
		data.ConditionIsNegate = types.BoolValue(value.Bool())
	} else {
		data.ConditionIsNegate = types.BoolNull()
	}
	if value := res.Get("response.rule.condition.attributeName"); value.Exists() && !data.ConditionAttributeName.IsNull() {
		data.ConditionAttributeName = types.StringValue(value.String())
	} else {
		data.ConditionAttributeName = types.StringNull()
	}
	if value := res.Get("response.rule.condition.attributeValue"); value.Exists() && !data.ConditionAttributeValue.IsNull() {
		data.ConditionAttributeValue = types.StringValue(value.String())
	} else {
		data.ConditionAttributeValue = types.StringNull()
	}
	if value := res.Get("response.rule.condition.dictionaryName"); value.Exists() && !data.ConditionDictionaryName.IsNull() {
		data.ConditionDictionaryName = types.StringValue(value.String())
	} else {
		data.ConditionDictionaryName = types.StringNull()
	}
	if value := res.Get("response.rule.condition.dictionaryValue"); value.Exists() && !data.ConditionDictionaryValue.IsNull() {
		data.ConditionDictionaryValue = types.StringValue(value.String())
	} else {
		data.ConditionDictionaryValue = types.StringNull()
	}
	if value := res.Get("response.rule.condition.operator"); value.Exists() && !data.ConditionOperator.IsNull() {
		data.ConditionOperator = types.StringValue(value.String())
	} else {
		data.ConditionOperator = types.StringNull()
	}
	if value := res.Get("response.identitySourceName"); value.Exists() && !data.IdentitySourceName.IsNull() {
		data.IdentitySourceName = types.StringValue(value.String())
	} else {
		data.IdentitySourceName = types.StringNull()
	}
	if value := res.Get("response.ifAuthFail"); value.Exists() && !data.IfAuthFail.IsNull() {
		data.IfAuthFail = types.StringValue(value.String())
	} else {
		data.IfAuthFail = types.StringNull()
	}
	if value := res.Get("response.ifProcessFail"); value.Exists() && !data.IfProcessFail.IsNull() {
		data.IfProcessFail = types.StringValue(value.String())
	} else {
		data.IfProcessFail = types.StringNull()
	}
	if value := res.Get("response.ifUserNotFound"); value.Exists() && !data.IfUserNotFound.IsNull() {
		data.IfUserNotFound = types.StringValue(value.String())
	} else {
		data.IfUserNotFound = types.StringNull()
	}
}