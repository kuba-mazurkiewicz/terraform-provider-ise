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

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-ise"
	"github.com/tidwall/gjson"
)

//template:end imports

//template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &NetworkAccessAuthorizationRuleDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworkAccessAuthorizationRuleDataSource{}
)

func NewNetworkAccessAuthorizationRuleDataSource() datasource.DataSource {
	return &NetworkAccessAuthorizationRuleDataSource{}
}

type NetworkAccessAuthorizationRuleDataSource struct {
	client *ise.Client
}

func (d *NetworkAccessAuthorizationRuleDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_access_authorization_rule"
}

func (d *NetworkAccessAuthorizationRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Network Access Authorization Rule.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Optional:            true,
				Computed:            true,
			},
			"policy_set_id": schema.StringAttribute{
				MarkdownDescription: "Policy set ID",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]",
				Optional:            true,
				Computed:            true,
			},
			"default": schema.BoolAttribute{
				MarkdownDescription: "Indicates if this rule is the default one",
				Computed:            true,
			},
			"rank": schema.Int64Attribute{
				MarkdownDescription: "The rank (priority) in relation to other rules. Lower rank is higher priority.",
				Computed:            true,
			},
			"state": schema.StringAttribute{
				MarkdownDescription: "The state that the rule is in. A disabled rule cannot be matched.",
				Computed:            true,
			},
			"condition_type": schema.StringAttribute{
				MarkdownDescription: "Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.",
				Computed:            true,
			},
			"condition_id": schema.StringAttribute{
				MarkdownDescription: "UUID for condition",
				Computed:            true,
			},
			"condition_is_negate": schema.BoolAttribute{
				MarkdownDescription: "Indicates whereas this condition is in negate mode",
				Computed:            true,
			},
			"condition_attribute_name": schema.StringAttribute{
				MarkdownDescription: "Dictionary attribute name",
				Computed:            true,
			},
			"condition_attribute_value": schema.StringAttribute{
				MarkdownDescription: "Attribute value for condition. Value type is specified in dictionary object.",
				Computed:            true,
			},
			"condition_dictionary_name": schema.StringAttribute{
				MarkdownDescription: "Dictionary name",
				Computed:            true,
			},
			"condition_dictionary_value": schema.StringAttribute{
				MarkdownDescription: "Dictionary value",
				Computed:            true,
			},
			"condition_operator": schema.StringAttribute{
				MarkdownDescription: "Equality operator",
				Computed:            true,
			},
			"children": schema.ListNestedAttribute{
				MarkdownDescription: "List of child conditions. `condition_type` must be one of `ConditionAndBlock`, `ConditionOrBlock`, `ConditionAttributes` or `ConditionReference`.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"condition_type": schema.StringAttribute{
							MarkdownDescription: "Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.",
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: "UUID for condition",
							Computed:            true,
						},
						"is_negate": schema.BoolAttribute{
							MarkdownDescription: "Indicates whereas this condition is in negate mode",
							Computed:            true,
						},
						"attribute_name": schema.StringAttribute{
							MarkdownDescription: "Dictionary attribute name",
							Computed:            true,
						},
						"attribute_value": schema.StringAttribute{
							MarkdownDescription: "Attribute value for condition. Value type is specified in dictionary object.",
							Computed:            true,
						},
						"dictionary_name": schema.StringAttribute{
							MarkdownDescription: "Dictionary name",
							Computed:            true,
						},
						"dictionary_value": schema.StringAttribute{
							MarkdownDescription: "Dictionary value",
							Computed:            true,
						},
						"operator": schema.StringAttribute{
							MarkdownDescription: "Equality operator",
							Computed:            true,
						},
						"children": schema.ListNestedAttribute{
							MarkdownDescription: "List of child conditions. `condition_type` must be one of `ConditionAndBlock`, `ConditionOrBlock`, `ConditionAttributes` or `ConditionReference`.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"condition_type": schema.StringAttribute{
										MarkdownDescription: "Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.",
										Computed:            true,
									},
									"id": schema.StringAttribute{
										MarkdownDescription: "UUID for condition",
										Computed:            true,
									},
									"is_negate": schema.BoolAttribute{
										MarkdownDescription: "Indicates whereas this condition is in negate mode",
										Computed:            true,
									},
									"attribute_name": schema.StringAttribute{
										MarkdownDescription: "Dictionary attribute name",
										Computed:            true,
									},
									"attribute_value": schema.StringAttribute{
										MarkdownDescription: "Attribute value for condition. Value type is specified in dictionary object.",
										Computed:            true,
									},
									"dictionary_name": schema.StringAttribute{
										MarkdownDescription: "Dictionary name",
										Computed:            true,
									},
									"dictionary_value": schema.StringAttribute{
										MarkdownDescription: "Dictionary value",
										Computed:            true,
									},
									"operator": schema.StringAttribute{
										MarkdownDescription: "Equality operator",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"profile": schema.ListAttribute{
				MarkdownDescription: "The authorization profile(s)",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"security_group": schema.StringAttribute{
				MarkdownDescription: "Security group used in authorization policies",
				Computed:            true,
			},
		},
	}
}
func (d *NetworkAccessAuthorizationRuleDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *NetworkAccessAuthorizationRuleDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*IseProviderData).Client
}

//template:end model

//template:begin read
func (d *NetworkAccessAuthorizationRuleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config NetworkAccessAuthorizationRule

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))
	if config.Id.IsNull() && !config.Name.IsNull() {
		res, err := d.client.Get(config.getPath())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
			return
		}
		if value := res.Get("response"); len(value.Array()) > 0 {
			value.ForEach(func(k, v gjson.Result) bool {
				if config.Name.ValueString() == v.Get("name").String() {
					config.Id = types.StringValue(v.Get("id").String())
					tflog.Debug(ctx, fmt.Sprintf("%s: Found object with name '%v', id: %v", config.Id.String(), config.Name.ValueString(), config.Id.String()))
					return false
				}
				return true
			})
		}
		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with name: %s", config.Name.ValueString()))
			return
		}
	}

	res, err := d.client.Get(config.getPath() + "/" + config.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

//template:end read
