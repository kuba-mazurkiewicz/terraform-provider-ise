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
type DeviceAdminCondition struct {
	Id              types.String                   `tfsdk:"id"`
	Name            types.String                   `tfsdk:"name"`
	Description     types.String                   `tfsdk:"description"`
	ConditionType   types.String                   `tfsdk:"condition_type"`
	IsNegate        types.Bool                     `tfsdk:"is_negate"`
	AttributeName   types.String                   `tfsdk:"attribute_name"`
	AttributeValue  types.String                   `tfsdk:"attribute_value"`
	DictionaryName  types.String                   `tfsdk:"dictionary_name"`
	DictionaryValue types.String                   `tfsdk:"dictionary_value"`
	Operator        types.String                   `tfsdk:"operator"`
	Children        []DeviceAdminConditionChildren `tfsdk:"children"`
}

type DeviceAdminConditionChildren struct {
	Name            types.String                           `tfsdk:"name"`
	Description     types.String                           `tfsdk:"description"`
	ConditionType   types.String                           `tfsdk:"condition_type"`
	Id              types.String                           `tfsdk:"id"`
	IsNegate        types.Bool                             `tfsdk:"is_negate"`
	AttributeName   types.String                           `tfsdk:"attribute_name"`
	AttributeValue  types.String                           `tfsdk:"attribute_value"`
	DictionaryName  types.String                           `tfsdk:"dictionary_name"`
	DictionaryValue types.String                           `tfsdk:"dictionary_value"`
	Operator        types.String                           `tfsdk:"operator"`
	Children        []DeviceAdminConditionChildrenChildren `tfsdk:"children"`
}

type DeviceAdminConditionChildrenChildren struct {
	Name            types.String `tfsdk:"name"`
	Description     types.String `tfsdk:"description"`
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
func (data DeviceAdminCondition) getPath() string {
	return "/api/v1/policy/device-admin/condition"
}

//template:end getPath

//template:begin getPathDelete

//template:end getPathDelete

//template:begin toBody
func (data DeviceAdminCondition) toBody(ctx context.Context, state DeviceAdminCondition) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.ConditionType.IsNull() {
		body, _ = sjson.Set(body, "conditionType", data.ConditionType.ValueString())
	}
	if !data.IsNegate.IsNull() {
		body, _ = sjson.Set(body, "isNegate", data.IsNegate.ValueBool())
	}
	if !data.AttributeName.IsNull() {
		body, _ = sjson.Set(body, "attributeName", data.AttributeName.ValueString())
	}
	if !data.AttributeValue.IsNull() {
		body, _ = sjson.Set(body, "attributeValue", data.AttributeValue.ValueString())
	}
	if !data.DictionaryName.IsNull() {
		body, _ = sjson.Set(body, "dictionaryName", data.DictionaryName.ValueString())
	}
	if !data.DictionaryValue.IsNull() {
		body, _ = sjson.Set(body, "dictionaryValue", data.DictionaryValue.ValueString())
	}
	if !data.Operator.IsNull() {
		body, _ = sjson.Set(body, "operator", data.Operator.ValueString())
	}
	if len(data.Children) > 0 {
		body, _ = sjson.Set(body, "children", []interface{}{})
		for _, item := range data.Children {
			itemBody := ""
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.Description.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "description", item.Description.ValueString())
			}
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
					if !childItem.Name.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "name", childItem.Name.ValueString())
					}
					if !childItem.Description.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "description", childItem.Description.ValueString())
					}
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
			body, _ = sjson.SetRaw(body, "children.-1", itemBody)
		}
	}
	return body
}

//template:end toBody

//template:begin fromBody
func (data *DeviceAdminCondition) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.name"); value.Exists() && value.Type != gjson.Null {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("response.description"); value.Exists() && value.Type != gjson.Null {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("response.conditionType"); value.Exists() && value.Type != gjson.Null {
		data.ConditionType = types.StringValue(value.String())
	} else {
		data.ConditionType = types.StringNull()
	}
	if value := res.Get("response.isNegate"); value.Exists() && value.Type != gjson.Null {
		data.IsNegate = types.BoolValue(value.Bool())
	} else {
		data.IsNegate = types.BoolNull()
	}
	if value := res.Get("response.attributeName"); value.Exists() && value.Type != gjson.Null {
		data.AttributeName = types.StringValue(value.String())
	} else {
		data.AttributeName = types.StringNull()
	}
	if value := res.Get("response.attributeValue"); value.Exists() && value.Type != gjson.Null {
		data.AttributeValue = types.StringValue(value.String())
	} else {
		data.AttributeValue = types.StringNull()
	}
	if value := res.Get("response.dictionaryName"); value.Exists() && value.Type != gjson.Null {
		data.DictionaryName = types.StringValue(value.String())
	} else {
		data.DictionaryName = types.StringNull()
	}
	if value := res.Get("response.dictionaryValue"); value.Exists() && value.Type != gjson.Null {
		data.DictionaryValue = types.StringValue(value.String())
	} else {
		data.DictionaryValue = types.StringNull()
	}
	if value := res.Get("response.operator"); value.Exists() && value.Type != gjson.Null {
		data.Operator = types.StringValue(value.String())
	} else {
		data.Operator = types.StringNull()
	}
	if value := res.Get("response.children"); value.Exists() {
		data.Children = make([]DeviceAdminConditionChildren, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := DeviceAdminConditionChildren{}
			if cValue := v.Get("name"); cValue.Exists() && cValue.Type != gjson.Null {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("description"); cValue.Exists() && cValue.Type != gjson.Null {
				item.Description = types.StringValue(cValue.String())
			} else {
				item.Description = types.StringNull()
			}
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
				item.Children = make([]DeviceAdminConditionChildrenChildren, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := DeviceAdminConditionChildrenChildren{}
					if ccValue := cv.Get("name"); ccValue.Exists() && ccValue.Type != gjson.Null {
						cItem.Name = types.StringValue(ccValue.String())
					} else {
						cItem.Name = types.StringNull()
					}
					if ccValue := cv.Get("description"); ccValue.Exists() && ccValue.Type != gjson.Null {
						cItem.Description = types.StringValue(ccValue.String())
					} else {
						cItem.Description = types.StringNull()
					}
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
}

//template:end fromBody

//template:begin updateFromBody
func (data *DeviceAdminCondition) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("response.description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("response.conditionType"); value.Exists() && !data.ConditionType.IsNull() {
		data.ConditionType = types.StringValue(value.String())
	} else {
		data.ConditionType = types.StringNull()
	}
	if value := res.Get("response.isNegate"); value.Exists() && !data.IsNegate.IsNull() {
		data.IsNegate = types.BoolValue(value.Bool())
	} else {
		data.IsNegate = types.BoolNull()
	}
	if value := res.Get("response.attributeName"); value.Exists() && !data.AttributeName.IsNull() {
		data.AttributeName = types.StringValue(value.String())
	} else {
		data.AttributeName = types.StringNull()
	}
	if value := res.Get("response.attributeValue"); value.Exists() && !data.AttributeValue.IsNull() {
		data.AttributeValue = types.StringValue(value.String())
	} else {
		data.AttributeValue = types.StringNull()
	}
	if value := res.Get("response.dictionaryName"); value.Exists() && !data.DictionaryName.IsNull() {
		data.DictionaryName = types.StringValue(value.String())
	} else {
		data.DictionaryName = types.StringNull()
	}
	if value := res.Get("response.dictionaryValue"); value.Exists() && !data.DictionaryValue.IsNull() {
		data.DictionaryValue = types.StringValue(value.String())
	} else {
		data.DictionaryValue = types.StringNull()
	}
	if value := res.Get("response.operator"); value.Exists() && !data.Operator.IsNull() {
		data.Operator = types.StringValue(value.String())
	} else {
		data.Operator = types.StringNull()
	}
	if value := res.Get("response.children"); value.Exists() {
		data.Children = make([]DeviceAdminConditionChildren, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := DeviceAdminConditionChildren{}
			if cValue := v.Get("name"); cValue.Exists() && cValue.Type != gjson.Null {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("description"); cValue.Exists() && cValue.Type != gjson.Null {
				item.Description = types.StringValue(cValue.String())
			} else {
				item.Description = types.StringNull()
			}
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
				item.Children = make([]DeviceAdminConditionChildrenChildren, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := DeviceAdminConditionChildrenChildren{}
					if ccValue := cv.Get("name"); ccValue.Exists() && ccValue.Type != gjson.Null {
						cItem.Name = types.StringValue(ccValue.String())
					} else {
						cItem.Name = types.StringNull()
					}
					if ccValue := cv.Get("description"); ccValue.Exists() && ccValue.Type != gjson.Null {
						cItem.Description = types.StringValue(ccValue.String())
					} else {
						cItem.Description = types.StringNull()
					}
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
}

//template:end updateFromBody

//template:begin isNull
func (data *DeviceAdminCondition) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.Description.IsNull() {
		return false
	}
	if !data.ConditionType.IsNull() {
		return false
	}
	if !data.IsNegate.IsNull() {
		return false
	}
	if !data.AttributeName.IsNull() {
		return false
	}
	if !data.AttributeValue.IsNull() {
		return false
	}
	if !data.DictionaryName.IsNull() {
		return false
	}
	if !data.DictionaryValue.IsNull() {
		return false
	}
	if !data.Operator.IsNull() {
		return false
	}
	if len(data.Children) > 0 {
		return false
	}
	return true
}

//template:end isNull
