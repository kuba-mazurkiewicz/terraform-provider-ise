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

	"github.com/CiscoDevNet/terraform-provider-ise/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

//template:end imports

//template:begin types
type NetworkAccessAuthorizationGlobalExceptionRule struct {
	Id                       types.String                                            `tfsdk:"id"`
	Name                     types.String                                            `tfsdk:"name"`
	Rank                     types.Int64                                             `tfsdk:"rank"`
	State                    types.String                                            `tfsdk:"state"`
	ConditionType            types.String                                            `tfsdk:"condition_type"`
	ConditionId              types.String                                            `tfsdk:"condition_id"`
	ConditionIsNegate        types.Bool                                              `tfsdk:"condition_is_negate"`
	ConditionAttributeName   types.String                                            `tfsdk:"condition_attribute_name"`
	ConditionAttributeValue  types.String                                            `tfsdk:"condition_attribute_value"`
	ConditionDictionaryName  types.String                                            `tfsdk:"condition_dictionary_name"`
	ConditionDictionaryValue types.String                                            `tfsdk:"condition_dictionary_value"`
	ConditionOperator        types.String                                            `tfsdk:"condition_operator"`
	Children                 []NetworkAccessAuthorizationGlobalExceptionRuleChildren `tfsdk:"children"`
	Profiles                 types.Set                                               `tfsdk:"profiles"`
	SecurityGroup            types.String                                            `tfsdk:"security_group"`
}

type NetworkAccessAuthorizationGlobalExceptionRuleChildren struct {
	ConditionType   types.String                                                    `tfsdk:"condition_type"`
	Id              types.String                                                    `tfsdk:"id"`
	IsNegate        types.Bool                                                      `tfsdk:"is_negate"`
	AttributeName   types.String                                                    `tfsdk:"attribute_name"`
	AttributeValue  types.String                                                    `tfsdk:"attribute_value"`
	DictionaryName  types.String                                                    `tfsdk:"dictionary_name"`
	DictionaryValue types.String                                                    `tfsdk:"dictionary_value"`
	Operator        types.String                                                    `tfsdk:"operator"`
	Children        []NetworkAccessAuthorizationGlobalExceptionRuleChildrenChildren `tfsdk:"children"`
}

type NetworkAccessAuthorizationGlobalExceptionRuleChildrenChildren struct {
	ConditionType   types.String `tfsdk:"condition_type"`
	Id              types.String `tfsdk:"id"`
	IsNegate        types.Bool   `tfsdk:"is_negate"`
	AttributeName   types.String `tfsdk:"attribute_name"`
	AttributeValue  types.String `tfsdk:"attribute_value"`
	DictionaryName  types.String `tfsdk:"dictionary_name"`
	DictionaryValue types.String `tfsdk:"dictionary_value"`
	Operator        types.String `tfsdk:"operator"`
}

//template:end types

//template:begin getPath
func (data NetworkAccessAuthorizationGlobalExceptionRule) getPath() string {
	return "/api/v1/policy/network-access/policy-set/global-exception"
}

//template:end getPath

//template:begin getPathDelete

//template:end getPathDelete

//template:begin toBody
func (data NetworkAccessAuthorizationGlobalExceptionRule) toBody(ctx context.Context, state NetworkAccessAuthorizationGlobalExceptionRule) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "rule.name", data.Name.ValueString())
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
	if !data.ConditionId.IsNull() {
		body, _ = sjson.Set(body, "rule.condition.id", data.ConditionId.ValueString())
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
	if len(data.Children) > 0 {
		body, _ = sjson.Set(body, "rule.condition.children", []interface{}{})
		for _, item := range data.Children {
			itemBody := ""
			if !item.ConditionType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "conditionType", item.ConditionType.ValueString())
			}
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.IsNegate.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "isNegate", item.IsNegate.ValueBool())
			}
			if !item.AttributeName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "attributeName", item.AttributeName.ValueString())
			}
			if !item.AttributeValue.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "attributeValue", item.AttributeValue.ValueString())
			}
			if !item.DictionaryName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dictionaryName", item.DictionaryName.ValueString())
			}
			if !item.DictionaryValue.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dictionaryValue", item.DictionaryValue.ValueString())
			}
			if !item.Operator.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "operator", item.Operator.ValueString())
			}
			if len(item.Children) > 0 {
				itemBody, _ = sjson.Set(itemBody, "children", []interface{}{})
				for _, childItem := range item.Children {
					itemChildBody := ""
					if !childItem.ConditionType.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "conditionType", childItem.ConditionType.ValueString())
					}
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.IsNegate.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "isNegate", childItem.IsNegate.ValueBool())
					}
					if !childItem.AttributeName.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "attributeName", childItem.AttributeName.ValueString())
					}
					if !childItem.AttributeValue.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "attributeValue", childItem.AttributeValue.ValueString())
					}
					if !childItem.DictionaryName.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "dictionaryName", childItem.DictionaryName.ValueString())
					}
					if !childItem.DictionaryValue.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "dictionaryValue", childItem.DictionaryValue.ValueString())
					}
					if !childItem.Operator.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "operator", childItem.Operator.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "children.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "rule.condition.children.-1", itemBody)
		}
	}
	if !data.Profiles.IsNull() {
		var values []string
		data.Profiles.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "profile", values)
	}
	if !data.SecurityGroup.IsNull() {
		body, _ = sjson.Set(body, "securityGroup", data.SecurityGroup.ValueString())
	}
	return body
}

//template:end toBody

//template:begin fromBody
func (data *NetworkAccessAuthorizationGlobalExceptionRule) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.rule.name"); value.Exists() && value.Type != gjson.Null {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("response.rule.rank"); value.Exists() && value.Type != gjson.Null {
		data.Rank = types.Int64Value(value.Int())
	} else {
		data.Rank = types.Int64Null()
	}
	if value := res.Get("response.rule.state"); value.Exists() && value.Type != gjson.Null {
		data.State = types.StringValue(value.String())
	} else {
		data.State = types.StringNull()
	}
	if value := res.Get("response.rule.condition.conditionType"); value.Exists() && value.Type != gjson.Null {
		data.ConditionType = types.StringValue(value.String())
	} else {
		data.ConditionType = types.StringNull()
	}
	if value := res.Get("response.rule.condition.id"); value.Exists() && value.Type != gjson.Null {
		data.ConditionId = types.StringValue(value.String())
	} else {
		data.ConditionId = types.StringNull()
	}
	if value := res.Get("response.rule.condition.isNegate"); value.Exists() && value.Type != gjson.Null {
		data.ConditionIsNegate = types.BoolValue(value.Bool())
	} else {
		data.ConditionIsNegate = types.BoolNull()
	}
	if value := res.Get("response.rule.condition.attributeName"); value.Exists() && value.Type != gjson.Null {
		data.ConditionAttributeName = types.StringValue(value.String())
	} else {
		data.ConditionAttributeName = types.StringNull()
	}
	if value := res.Get("response.rule.condition.attributeValue"); value.Exists() && value.Type != gjson.Null {
		data.ConditionAttributeValue = types.StringValue(value.String())
	} else {
		data.ConditionAttributeValue = types.StringNull()
	}
	if value := res.Get("response.rule.condition.dictionaryName"); value.Exists() && value.Type != gjson.Null {
		data.ConditionDictionaryName = types.StringValue(value.String())
	} else {
		data.ConditionDictionaryName = types.StringNull()
	}
	if value := res.Get("response.rule.condition.dictionaryValue"); value.Exists() && value.Type != gjson.Null {
		data.ConditionDictionaryValue = types.StringValue(value.String())
	} else {
		data.ConditionDictionaryValue = types.StringNull()
	}
	if value := res.Get("response.rule.condition.operator"); value.Exists() && value.Type != gjson.Null {
		data.ConditionOperator = types.StringValue(value.String())
	} else {
		data.ConditionOperator = types.StringNull()
	}
	if value := res.Get("response.rule.condition.children"); value.Exists() {
		data.Children = make([]NetworkAccessAuthorizationGlobalExceptionRuleChildren, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := NetworkAccessAuthorizationGlobalExceptionRuleChildren{}
			if cValue := v.Get("conditionType"); cValue.Exists() && cValue.Type != gjson.Null {
				item.ConditionType = types.StringValue(cValue.String())
			} else {
				item.ConditionType = types.StringNull()
			}
			if cValue := v.Get("id"); cValue.Exists() && cValue.Type != gjson.Null {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("isNegate"); cValue.Exists() && cValue.Type != gjson.Null {
				item.IsNegate = types.BoolValue(cValue.Bool())
			} else {
				item.IsNegate = types.BoolNull()
			}
			if cValue := v.Get("attributeName"); cValue.Exists() && cValue.Type != gjson.Null {
				item.AttributeName = types.StringValue(cValue.String())
			} else {
				item.AttributeName = types.StringNull()
			}
			if cValue := v.Get("attributeValue"); cValue.Exists() && cValue.Type != gjson.Null {
				item.AttributeValue = types.StringValue(cValue.String())
			} else {
				item.AttributeValue = types.StringNull()
			}
			if cValue := v.Get("dictionaryName"); cValue.Exists() && cValue.Type != gjson.Null {
				item.DictionaryName = types.StringValue(cValue.String())
			} else {
				item.DictionaryName = types.StringNull()
			}
			if cValue := v.Get("dictionaryValue"); cValue.Exists() && cValue.Type != gjson.Null {
				item.DictionaryValue = types.StringValue(cValue.String())
			} else {
				item.DictionaryValue = types.StringNull()
			}
			if cValue := v.Get("operator"); cValue.Exists() && cValue.Type != gjson.Null {
				item.Operator = types.StringValue(cValue.String())
			} else {
				item.Operator = types.StringNull()
			}
			if cValue := v.Get("children"); cValue.Exists() {
				item.Children = make([]NetworkAccessAuthorizationGlobalExceptionRuleChildrenChildren, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := NetworkAccessAuthorizationGlobalExceptionRuleChildrenChildren{}
					if ccValue := cv.Get("conditionType"); ccValue.Exists() && ccValue.Type != gjson.Null {
						cItem.ConditionType = types.StringValue(ccValue.String())
					} else {
						cItem.ConditionType = types.StringNull()
					}
					if ccValue := cv.Get("id"); ccValue.Exists() && ccValue.Type != gjson.Null {
						cItem.Id = types.StringValue(ccValue.String())
					} else {
						cItem.Id = types.StringNull()
					}
					if ccValue := cv.Get("isNegate"); ccValue.Exists() && ccValue.Type != gjson.Null {
						cItem.IsNegate = types.BoolValue(ccValue.Bool())
					} else {
						cItem.IsNegate = types.BoolNull()
					}
					if ccValue := cv.Get("attributeName"); ccValue.Exists() && ccValue.Type != gjson.Null {
						cItem.AttributeName = types.StringValue(ccValue.String())
					} else {
						cItem.AttributeName = types.StringNull()
					}
					if ccValue := cv.Get("attributeValue"); ccValue.Exists() && ccValue.Type != gjson.Null {
						cItem.AttributeValue = types.StringValue(ccValue.String())
					} else {
						cItem.AttributeValue = types.StringNull()
					}
					if ccValue := cv.Get("dictionaryName"); ccValue.Exists() && ccValue.Type != gjson.Null {
						cItem.DictionaryName = types.StringValue(ccValue.String())
					} else {
						cItem.DictionaryName = types.StringNull()
					}
					if ccValue := cv.Get("dictionaryValue"); ccValue.Exists() && ccValue.Type != gjson.Null {
						cItem.DictionaryValue = types.StringValue(ccValue.String())
					} else {
						cItem.DictionaryValue = types.StringNull()
					}
					if ccValue := cv.Get("operator"); ccValue.Exists() && ccValue.Type != gjson.Null {
						cItem.Operator = types.StringValue(ccValue.String())
					} else {
						cItem.Operator = types.StringNull()
					}
					item.Children = append(item.Children, cItem)
					return true
				})
			}
			data.Children = append(data.Children, item)
			return true
		})
	}
	if value := res.Get("response.profile"); value.Exists() {
		data.Profiles = helpers.GetStringSet(value.Array())
	} else {
		data.Profiles = types.SetNull(types.StringType)
	}
	if value := res.Get("response.securityGroup"); value.Exists() && value.Type != gjson.Null {
		data.SecurityGroup = types.StringValue(value.String())
	} else {
		data.SecurityGroup = types.StringNull()
	}
}

//template:end fromBody

//template:begin updateFromBody
func (data *NetworkAccessAuthorizationGlobalExceptionRule) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.rule.name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
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
	if value := res.Get("response.rule.condition.id"); value.Exists() && !data.ConditionId.IsNull() {
		data.ConditionId = types.StringValue(value.String())
	} else {
		data.ConditionId = types.StringNull()
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
	if value := res.Get("response.rule.condition.children"); value.Exists() {
		data.Children = make([]NetworkAccessAuthorizationGlobalExceptionRuleChildren, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := NetworkAccessAuthorizationGlobalExceptionRuleChildren{}
			if cValue := v.Get("conditionType"); cValue.Exists() && cValue.Type != gjson.Null {
				item.ConditionType = types.StringValue(cValue.String())
			} else {
				item.ConditionType = types.StringNull()
			}
			if cValue := v.Get("id"); cValue.Exists() && cValue.Type != gjson.Null {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("isNegate"); cValue.Exists() && cValue.Type != gjson.Null {
				item.IsNegate = types.BoolValue(cValue.Bool())
			} else {
				item.IsNegate = types.BoolNull()
			}
			if cValue := v.Get("attributeName"); cValue.Exists() && cValue.Type != gjson.Null {
				item.AttributeName = types.StringValue(cValue.String())
			} else {
				item.AttributeName = types.StringNull()
			}
			if cValue := v.Get("attributeValue"); cValue.Exists() && cValue.Type != gjson.Null {
				item.AttributeValue = types.StringValue(cValue.String())
			} else {
				item.AttributeValue = types.StringNull()
			}
			if cValue := v.Get("dictionaryName"); cValue.Exists() && cValue.Type != gjson.Null {
				item.DictionaryName = types.StringValue(cValue.String())
			} else {
				item.DictionaryName = types.StringNull()
			}
			if cValue := v.Get("dictionaryValue"); cValue.Exists() && cValue.Type != gjson.Null {
				item.DictionaryValue = types.StringValue(cValue.String())
			} else {
				item.DictionaryValue = types.StringNull()
			}
			if cValue := v.Get("operator"); cValue.Exists() && cValue.Type != gjson.Null {
				item.Operator = types.StringValue(cValue.String())
			} else {
				item.Operator = types.StringNull()
			}
			if cValue := v.Get("children"); cValue.Exists() {
				item.Children = make([]NetworkAccessAuthorizationGlobalExceptionRuleChildrenChildren, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := NetworkAccessAuthorizationGlobalExceptionRuleChildrenChildren{}
					if ccValue := cv.Get("conditionType"); ccValue.Exists() && ccValue.Type != gjson.Null {
						cItem.ConditionType = types.StringValue(ccValue.String())
					} else {
						cItem.ConditionType = types.StringNull()
					}
					if ccValue := cv.Get("id"); ccValue.Exists() && ccValue.Type != gjson.Null {
						cItem.Id = types.StringValue(ccValue.String())
					} else {
						cItem.Id = types.StringNull()
					}
					if ccValue := cv.Get("isNegate"); ccValue.Exists() && ccValue.Type != gjson.Null {
						cItem.IsNegate = types.BoolValue(ccValue.Bool())
					} else {
						cItem.IsNegate = types.BoolNull()
					}
					if ccValue := cv.Get("attributeName"); ccValue.Exists() && ccValue.Type != gjson.Null {
						cItem.AttributeName = types.StringValue(ccValue.String())
					} else {
						cItem.AttributeName = types.StringNull()
					}
					if ccValue := cv.Get("attributeValue"); ccValue.Exists() && ccValue.Type != gjson.Null {
						cItem.AttributeValue = types.StringValue(ccValue.String())
					} else {
						cItem.AttributeValue = types.StringNull()
					}
					if ccValue := cv.Get("dictionaryName"); ccValue.Exists() && ccValue.Type != gjson.Null {
						cItem.DictionaryName = types.StringValue(ccValue.String())
					} else {
						cItem.DictionaryName = types.StringNull()
					}
					if ccValue := cv.Get("dictionaryValue"); ccValue.Exists() && ccValue.Type != gjson.Null {
						cItem.DictionaryValue = types.StringValue(ccValue.String())
					} else {
						cItem.DictionaryValue = types.StringNull()
					}
					if ccValue := cv.Get("operator"); ccValue.Exists() && ccValue.Type != gjson.Null {
						cItem.Operator = types.StringValue(ccValue.String())
					} else {
						cItem.Operator = types.StringNull()
					}
					item.Children = append(item.Children, cItem)
					return true
				})
			}
			data.Children = append(data.Children, item)
			return true
		})
	}
	if value := res.Get("response.profile"); value.Exists() && !data.Profiles.IsNull() {
		data.Profiles = helpers.GetStringSet(value.Array())
	} else {
		data.Profiles = types.SetNull(types.StringType)
	}
	if value := res.Get("response.securityGroup"); value.Exists() && !data.SecurityGroup.IsNull() {
		data.SecurityGroup = types.StringValue(value.String())
	} else {
		data.SecurityGroup = types.StringNull()
	}
}

//template:end updateFromBody

//template:begin isNull
func (data *NetworkAccessAuthorizationGlobalExceptionRule) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.Rank.IsNull() {
		return false
	}
	if !data.State.IsNull() {
		return false
	}
	if !data.ConditionType.IsNull() {
		return false
	}
	if !data.ConditionId.IsNull() {
		return false
	}
	if !data.ConditionIsNegate.IsNull() {
		return false
	}
	if !data.ConditionAttributeName.IsNull() {
		return false
	}
	if !data.ConditionAttributeValue.IsNull() {
		return false
	}
	if !data.ConditionDictionaryName.IsNull() {
		return false
	}
	if !data.ConditionDictionaryValue.IsNull() {
		return false
	}
	if !data.ConditionOperator.IsNull() {
		return false
	}
	if len(data.Children) > 0 {
		return false
	}
	if !data.Profiles.IsNull() {
		return false
	}
	if !data.SecurityGroup.IsNull() {
		return false
	}
	return true
}

//template:end isNull
