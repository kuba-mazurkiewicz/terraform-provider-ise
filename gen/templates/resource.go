//go:build ignore
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
	"net/url"
	"strings"
	"sync"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/floatplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-ise"
	"github.com/CiscoDevNet/terraform-provider-ise/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
)
//template:end imports

//template:begin header

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &{{camelCase .Name}}Resource{}
{{- if not .NoImport}}
var _ resource.ResourceWithImportState = &{{camelCase .Name}}Resource{}
{{- end}}

func New{{camelCase .Name}}Resource() resource.Resource {
	return &{{camelCase .Name}}Resource{}
}

type {{camelCase .Name}}Resource struct {
	client *ise.Client
}

func (r *{{camelCase .Name}}Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_{{snakeCase .Name}}"
}
//template:end header

//template:begin model
func (r *{{camelCase .Name}}Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("{{.ResDescription}}").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			{{- range  .Attributes}}
			{{- if not .Value}}
			"{{.TfName}}": schema.{{if isNestedListSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}")
					{{- if len .EnumValues -}}
					.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
					{{- end -}}
					{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
					.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
					{{- end -}}
					{{- if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0) -}}
					.AddFloatRangeDescription({{.MinFloat}}, {{.MaxFloat}})
					{{- end -}}
					{{- if .DefaultValue -}}
					.AddDefaultValueDescription("{{.DefaultValue}}")
					{{- end -}}
					.String,
				{{- if isListSet .}}
				ElementType:         types.{{.ElementType}}Type,
				{{- else if eq .Type "Map"}}
				ElementType:         types.StringType,
				{{- end}}
				{{- if or .Reference .Mandatory}}
				Required:            true,
				{{- else}}
				Optional:            true,
				{{- end}}
				{{- if or (len .DefaultValue) .Computed}}
				Computed:            true,
				{{- end}}
				{{- if len .EnumValues}}
				Validators: []validator.String{
					stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
				},
				{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0) }}
				Validators: []validator.String{
					{{- if or (ne .StringMinLength 0) (ne .StringMaxLength 0)}}
					stringvalidator.LengthBetween({{.StringMinLength}}, {{.StringMaxLength}}),
					{{- end}}
					{{- range .StringPatterns}}
					stringvalidator.RegexMatches(regexp.MustCompile(`{{.}}`), ""),
					{{- end}}
				},
				{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
				Validators: []validator.Int64{
					int64validator.Between({{.MinInt}}, {{.MaxInt}}),
				},
				{{- else if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0)}}
				Validators: []validator.Float64{
					float64validator.Between({{.MinFloat}}, {{.MaxFloat}}),
				},
				{{- end}}
				{{- if and (len .DefaultValue) (eq .Type "Int64")}}
				Default:             int64default.StaticInt64({{.DefaultValue}}),
				{{- else if and (len .DefaultValue) (eq .Type "Bool")}}
				Default:             booldefault.StaticBool({{.DefaultValue}}),
				{{- else if and (len .DefaultValue) (eq .Type "String")}}
				Default:             stringdefault.StaticString("{{.DefaultValue}}"),
				{{- end}}
				{{- if or .Id .Reference .RequiresReplace}}
				PlanModifiers: []planmodifier.{{.Type}}{
					{{snakeCase .Type}}planmodifier.RequiresReplace(),
				},
				{{- end}}
				{{- if .Computed}}
				PlanModifiers: []planmodifier.{{.Type}}{
					{{snakeCase .Type}}planmodifier.UseStateForUnknown(),
				},
				{{- end}}
				{{- if isNestedListSet .}}
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						{{- range  .Attributes}}
						{{- if not .Value}}
						"{{.TfName}}": schema.{{if isNestedListSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}")
								{{- if len .EnumValues -}}
								.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
								{{- end -}}
								{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
								.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
								{{- end -}}
								{{- if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0) -}}
								.AddFloatRangeDescription({{.MinFloat}}, {{.MaxFloat}})
								{{- end -}}
								{{- if .DefaultValue -}}
								.AddDefaultValueDescription("{{.DefaultValue}}")
								{{- end -}}
								.String,
							{{- if isListSet .}}
							ElementType:         types.{{.ElementType}}Type,
							{{- else if eq .Type "Map"}}
							ElementType:         types.StringType,
							{{- end}}
							{{- if or .Reference .Mandatory}}
							Required:            true,
							{{- else}}
							Optional:            true,
							{{- end}}
							{{- if or (len .DefaultValue) .Computed}}
							Computed:            true,
							{{- end}}
							{{- if len .EnumValues}}
							Validators: []validator.String{
								stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
							},
							{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0) }}
							Validators: []validator.String{
								{{- if or (ne .StringMinLength 0) (ne .StringMaxLength 0)}}
								stringvalidator.LengthBetween({{.StringMinLength}}, {{.StringMaxLength}}),
								{{- end}}
								{{- range .StringPatterns}}
								stringvalidator.RegexMatches(regexp.MustCompile(`{{.}}`), ""),
								{{- end}}
							},
							{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
							Validators: []validator.Int64{
								int64validator.Between({{.MinInt}}, {{.MaxInt}}),
							},
							{{- else if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0)}}
							Validators: []validator.Float64{
								float64validator.Between({{.MinFloat}}, {{.MaxFloat}}),
							},
							{{- end}}
							{{- if and (len .DefaultValue) (eq .Type "Int64")}}
							Default:             int64default.StaticInt64({{.DefaultValue}}),
							{{- else if and (len .DefaultValue) (eq .Type "Bool")}}
							Default:             booldefault.StaticBool({{.DefaultValue}}),
							{{- else if and (len .DefaultValue) (eq .Type "String")}}
							Default:             stringdefault.StaticString("{{.DefaultValue}}"),
							{{- end}}
							{{- if .RequiresReplace}}
							PlanModifiers: []planmodifier.{{.Type}}{
								{{snakeCase .Type}}planmodifier.RequiresReplace(),
							},
							{{- end}}
							{{- if .Computed}}
							PlanModifiers: []planmodifier.{{.Type}}{
								{{snakeCase .Type}}planmodifier.UseStateForUnknown(),
							},
							{{- end}}
							{{- if isNestedListSet .}}
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									{{- range  .Attributes}}
									{{- if not .Value}}
									"{{.TfName}}": schema.{{if isNestedListSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}")
											{{- if len .EnumValues -}}
											.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
											{{- end -}}
											{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
											.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
											{{- end -}}
											{{- if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0) -}}
											.AddFloatRangeDescription({{.MinFloat}}, {{.MaxFloat}})
											{{- end -}}
											{{- if .DefaultValue -}}
											.AddDefaultValueDescription("{{.DefaultValue}}")
											{{- end -}}
											.String,
										{{- if isListSet .}}
										ElementType:         types.{{.ElementType}}Type,
										{{- else if eq .Type "Map"}}
										ElementType:         types.StringType,
										{{- end}}
										{{- if or .Reference .Mandatory}}
										Required:            true,
										{{- else}}
										Optional:            true,
										{{- end}}
										{{- if or (len .DefaultValue) .Computed}}
										Computed:            true,
										{{- end}}
										{{- if len .EnumValues}}
										Validators: []validator.String{
											stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
										},
										{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0) }}
										Validators: []validator.String{
											{{- if or (ne .StringMinLength 0) (ne .StringMaxLength 0)}}
											stringvalidator.LengthBetween({{.StringMinLength}}, {{.StringMaxLength}}),
											{{- end}}
											{{- range .StringPatterns}}
											stringvalidator.RegexMatches(regexp.MustCompile(`{{.}}`), ""),
											{{- end}}
										},
										{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
										Validators: []validator.Int64{
											int64validator.Between({{.MinInt}}, {{.MaxInt}}),
										},
										{{- else if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0)}}
										Validators: []validator.Float64{
											float64validator.Between({{.MinFloat}}, {{.MaxFloat}}),
										},
										{{- end}}
										{{- if and (len .DefaultValue) (eq .Type "Int64")}}
										Default:             int64default.StaticInt64({{.DefaultValue}}),
										{{- else if and (len .DefaultValue) (eq .Type "Bool")}}
										Default:             booldefault.StaticBool({{.DefaultValue}}),
										{{- else if and (len .DefaultValue) (eq .Type "String")}}
										Default:             stringdefault.StaticString("{{.DefaultValue}}"),
										{{- end}}
										{{- if .RequiresReplace}}
										PlanModifiers: []planmodifier.{{.Type}}{
											{{snakeCase .Type}}planmodifier.RequiresReplace(),
										},
										{{- end}}
										{{- if .Computed}}
										PlanModifiers: []planmodifier.{{.Type}}{
											{{snakeCase .Type}}planmodifier.UseStateForUnknown(),
										},
										{{- end}}
										{{- if isNestedListSet .}}
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												{{- range  .Attributes}}
												{{- if not .Value}}
												"{{.TfName}}": schema.{{if isNestedListSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}")
														{{- if len .EnumValues -}}
														.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
														{{- end -}}
														{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
														.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
														{{- end -}}
														{{- if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0) -}}
														.AddFloatRangeDescription({{.MinFloat}}, {{.MaxFloat}})
														{{- end -}}
														{{- if .DefaultValue -}}
														.AddDefaultValueDescription("{{.DefaultValue}}")
														{{- end -}}
														.String,
													{{- if isListSet .}}
													ElementType:         types.{{.ElementType}}Type,
													{{- else if eq .Type "Map"}}
													ElementType:         types.StringType,
													{{- end}}
													{{- if or .Reference .Mandatory}}
													Required:            true,
													{{- else}}
													Optional:            true,
													{{- end}}
													{{- if or (len .DefaultValue) .Computed}}
													Computed:            true,
													{{- end}}
													{{- if len .EnumValues}}
													Validators: []validator.String{
														stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
													},
													{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0) }}
													Validators: []validator.String{
														{{- if or (ne .StringMinLength 0) (ne .StringMaxLength 0)}}
														stringvalidator.LengthBetween({{.StringMinLength}}, {{.StringMaxLength}}),
														{{- end}}
														{{- range .StringPatterns}}
														stringvalidator.RegexMatches(regexp.MustCompile(`{{.}}`), ""),
														{{- end}}
													},
													{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
													Validators: []validator.Int64{
														int64validator.Between({{.MinInt}}, {{.MaxInt}}),
													},
													{{- else if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0)}}
													Validators: []validator.Float64{
														float64validator.Between({{.MinFloat}}, {{.MaxFloat}}),
													},
													{{- end}}
													{{- if and (len .DefaultValue) (eq .Type "Int64")}}
													Default:             int64default.StaticInt64({{.DefaultValue}}),
													{{- else if and (len .DefaultValue) (eq .Type "Bool")}}
													Default:             booldefault.StaticBool({{.DefaultValue}}),
													{{- else if and (len .DefaultValue) (eq .Type "String")}}
													Default:             stringdefault.StaticString("{{.DefaultValue}}"),
													{{- end}}
													{{- if .RequiresReplace}}
													PlanModifiers: []planmodifier.{{.Type}}{
														{{snakeCase .Type}}planmodifier.RequiresReplace(),
													},
													{{- end}}
													{{- if .Computed}}
													PlanModifiers: []planmodifier.{{.Type}}{
														{{snakeCase .Type}}planmodifier.UseStateForUnknown(),
													},
													{{- end}}
												},
												{{- end}}
												{{- end}}
											},
										},
										{{- if or (ne .MinList 0) (ne .MaxList 0)}}
										Validators: []validator.List{
											{{- if ne .MinList 0}}
											listvalidator.SizeAtLeast({{.MinList}}),
											{{- end}}
											{{- if ne .MaxList 0}}
											listvalidator.SizeAtMost({{.MaxList}}),
											{{- end}}
										},
										{{- end}}
										{{- if .RequiresReplace}}
										PlanModifiers: []planmodifier.{{.Type}}{
											{{snakeCase .Type}}planmodifier.RequiresReplace(),
										},
										{{- end}}
										{{- end}}
									},
									{{- end}}
									{{- end}}
								},
							},
							{{- if or (ne .MinList 0) (ne .MaxList 0)}}
							Validators: []validator.List{
								{{- if ne .MinList 0}}
								listvalidator.SizeAtLeast({{.MinList}}),
								{{- end}}
								{{- if ne .MaxList 0}}
								listvalidator.SizeAtMost({{.MaxList}}),
								{{- end}}
							},
							{{- end}}
							{{- if .RequiresReplace}}
							PlanModifiers: []planmodifier.{{.Type}}{
								{{snakeCase .Type}}planmodifier.RequiresReplace(),
							},
							{{- end}}
							{{- end}}
						},
						{{- end}}
						{{- end}}
					},
				},
				{{- if or (ne .MinList 0) (ne .MaxList 0)}}
				Validators: []validator.List{
					{{- if ne .MinList 0}}
					listvalidator.SizeAtLeast({{.MinList}}),
					{{- end}}
					{{- if ne .MaxList 0}}
					listvalidator.SizeAtMost({{.MaxList}}),
					{{- end}}
				},
				{{- end}}
				{{- if .RequiresReplace}}
				PlanModifiers: []planmodifier.{{.Type}}{
					{{snakeCase .Type}}planmodifier.RequiresReplace(),
				},
				{{- end}}
				{{- end}}
			},
			{{- end}}
			{{- end}}
		},
	}
}
//template:end model

//template:begin configure
func (r *{{camelCase .Name}}Resource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*IseProviderData).Client
}
//template:end configure

//template:begin create
func (r *{{camelCase .Name}}Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan {{camelCase .Name}}
	{{- if strContains (camelCase .Name) "UpdateRank" }}
	var existingData {{strReplace (camelCase .Name) "UpdateRank" "" -1}}
	{{- end}}

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	{{- if .IsBulk}}
	{{- $id := getId .Attributes}}
	// Create request is split to multiple requests, where just subset of them may be successful
	state := {{camelCase .Name}}{}
	state.Items = make([]{{camelCase .Name}}Items, len(plan.Items))
	state.Id = types.StringValue(fmt.Sprint(plan.{{toGoName $id.TfName}}.Value{{$id.Type}}()))
	// Create object
	plan, diags = r.createSubresources(ctx, state, plan)
	resp.Diagnostics.Append(diags...)
	{{- else}}
	{{- if strContains (camelCase .Name) "UpdateRank" }}
	// Read existing attributes from the API
	{{- if strContains (camelCase .Name) "Rule" }}
	res, err := r.client.Get(plan.getPath() + "/" + url.QueryEscape(plan.RuleId.ValueString()))
	{{- else}}
	res, err := r.client.Get(plan.getPath() + "/" + url.QueryEscape(plan.PolicySetId.ValueString()))
	{{- end}}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}
	existingData.fromBody(ctx, res)

	// Use the `toBody` function to construct the body from existingData
	body := existingData.toBody(ctx, existingData)

	// Update rank
	{{- if strContains (camelCase .Name) "Rule" }}
	body, _ = sjson.Set(body, "rule.rank", plan.Rank.ValueInt64())
	res, err = r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.RuleId.ValueString()), body)
	{{- else}}
	body, _ = sjson.Set(body, "rank", plan.Rank.ValueInt64())
	res, err = r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.PolicySetId.ValueString()), body)
	{{- end}}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}
	{{- if strContains (camelCase .Name) "Rule" }}
	plan.Id = types.StringValue(fmt.Sprint(plan.RuleId.ValueString()))
	{{- else}}
	plan.Id = types.StringValue(fmt.Sprint(plan.PolicySetId.ValueString()))
	{{- end}}

	{{- else}}

	// Create object
	body := plan.toBody(ctx, {{camelCase .Name}}{})

	{{- if .UpdateDefault}}
	if plan.Name.ValueString() != "Default" {
	{{- end}}
	{{- if .PutCreate}}
	params := ""
	{{- if .PutIdQueryPath}}
	params += "/" + url.QueryEscape(gjson.Get(body, "{{.PutIdIncludePath}}.id").String())
	{{- end}}
	res, err := r.client.Put(plan.getPath() + params, body)
	{{- else if and (isErs .RestEndpoint) (not .IdPath) (not (hasId .Attributes))}}
	res, location, err := r.client.Post(plan.getPath(), body)
	{{- else}}
	res, _, err := r.client.Post(plan.getPath(), body)
	{{- end}}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object ({{- if .PutCreate}}PUT{{else}}POST{{end}}), got error: %s, %s", err, res.String()))
		return
	}
	{{- if .IdPath}}
	plan.Id = types.StringValue(res.Get("{{.IdPath}}").String())
	{{- else if and (hasId .Attributes) (not .PutIdQueryPath)}}
		{{- $id := getId .Attributes}}
	plan.Id = types.StringValue(fmt.Sprint(plan.{{toGoName $id.TfName}}.Value{{$id.Type}}()))
	{{- else if .PutIdQueryPath}}
	plan.Id = types.StringValue(gjson.Get(body, "{{.PutIdIncludePath}}.id").String())
	{{- else}}
	locationElements := strings.Split(location, "/")
	plan.Id = types.StringValue(locationElements[len(locationElements)-1])
	{{- end}}
	{{- if .UpdateDefault}}
	} else {
		res, err := r.client.Get(plan.getPath())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
			return
		}
		// Find id
		{{- if not (strContains (camelCase .Name) "Rule") }}
		res = res.Get("response.#(name==\"" + plan.Name.ValueString() + "\")")
		plan.Id = types.StringValue(res.Get("id").String())
		{{- else}}
		res = res.Get("response.#(rule.name==\"" + plan.Name.ValueString() + "\")")
		plan.Id = types.StringValue(res.Get("rule.id").String())
		{{- end}}

		// Read existing attributes from the API
		res, err = r.client.Get(plan.getPath() + "/" + url.QueryEscape(plan.Id.ValueString()))
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
			return
		}
		var existingData {{camelCase .Name}}

		// Populate existingData with current state from the API
		existingData.fromBody(ctx, res)

		if plan.Name.ValueString() == "Default" {
			{{- if not (strContains (camelCase .Name) "Rule") }}
			// Set Rank and Description in the request body from the existing data for Default Policy Set
			// Description on Default policy set cannot be modify, hence reading that form existing resource and setting to body
			body, _ = sjson.Set(body, "description", existingData.Description.ValueString())
			body, _ = sjson.Set(body, "rank", existingData.Rank.ValueInt64())
			{{- else}}
			// Set Rank in the request body from the existing data for Default Auth and Authz Rules
			body, _ = sjson.Set(body, "rule.rank", existingData.Rank.ValueInt64())	
			{{- end}}
		}

		res, err = r.client.Put(plan.getPath()+"/"+plan.Id.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}
	{{- end}}
	{{- end}}
	{{- end}}
	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}
//template:end create

//template:begin read
func (r *{{camelCase .Name}}Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state {{camelCase .Name}}

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	{{- if not .NoRead}}
	res, err := r.client.Get(state.getPath(){{if not .GetNoId}} + "/" + url.QueryEscape(state.Id.ValueString()){{end}})
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	// If every attribute is set to null we are dealing with an import operation and therefore reading all attributes
	if state.isNull(ctx, res) {
		state.fromBody(ctx, res)
	} else {
		state.updateFromBody(ctx, res)
	}
	{{- end}}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
//template:end read

//template:begin update
func (r *{{camelCase .Name}}Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state {{camelCase .Name}}
	{{- if strContains (camelCase .Name) "UpdateRank" }}
	var existingData {{strReplace (camelCase .Name) "UpdateRank" "" -1}}
	{{- end}}

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	{{- if strContains (camelCase .Name) "UpdateRank" }}
	
	// Read existing attributes from the API
	{{- if strContains (camelCase .Name) "Rule" }}
	res, err := r.client.Get(plan.getPath() + "/" + url.QueryEscape(plan.RuleId.ValueString()))
	{{- else}}
	res, err := r.client.Get(plan.getPath() + "/" + url.QueryEscape(plan.PolicySetId.ValueString()))
	{{- end}}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}
	existingData.fromBody(ctx, res)

	// Use the `toBody` function to construct the body from existingData
	body := existingData.toBody(ctx, existingData)

	// Update rank
	{{- if strContains (camelCase .Name) "Rule" }}
	body, _ = sjson.Set(body, "rule.rank", plan.Rank.ValueInt64())
	{{- else}}
	body, _ = sjson.Set(body, "rank", plan.Rank.ValueInt64())
	{{- end}}

	res, err = r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}
	{{- else}}
	{{- if not .NoUpdate}}
	body := plan.toBody(ctx, state)

	{{- if hasAttribute .Attributes "rank"}}
	// Check if resource has rank attribute
	// Check if plan.Rank is null (i.e., not provided) and set Rank to body from existingData to avoid reordering the rule during update
	if plan.Rank.IsNull() {
		var existingData {{camelCase .Name}}
		// Fetch existing data from the API
		res, err := r.client.Get(plan.getPath() + "/" + url.QueryEscape(plan.Id.ValueString()))
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
			return
		}
		// Populate existingData with current state from the API
		existingData.fromBody(ctx, res)
		// Set Rank in the request body from the existing data if it's missing from the plan
		{{- if strContains (camelCase .Name) "Rule" }}
		body, _ = sjson.Set(body, "rule.rank", existingData.Rank.ValueInt64())
		{{- else}}
		body, _ = sjson.Set(body, "rank", existingData.Rank.ValueInt64())
		// Description on Default policy set cannot be modify, hence reading that form existing resource and setting to body
		if plan.Name.ValueString() == "Default" {
			body, _ = sjson.Set(body, "description", existingData.Description.ValueString())
		}
		{{- end}}
	}
	{{- end}}
	{{if .PostUpdate}}
	res, _, err := r.client.Post(plan.getPath(), body)
	{{- else}}
	res, err := r.client.Put(plan.getPath() + "/" + url.QueryEscape(plan.Id.ValueString()), body)
	{{- end}}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}
	{{- end}}
	{{- end}}
	
	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}
//template:end update

//template:begin delete
func (r *{{camelCase .Name}}Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state {{camelCase .Name}}

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	{{- if not .NoDelete}}
	{{- if .PutDelete}}
	body := state.toBody(ctx, state)
	res, err := r.client.Put(state.getPathDelete(), body)
	{{- else}}
	res, err := r.client.Delete({{if .DeleteRestEndpoint}}state.getPathDelete(){{else}}state.getPath(){{end}} + "/" + url.QueryEscape(state.Id.ValueString()))
	{{- end}}
	if err != nil{{if .IgnoreDeleteError}} && !strings.Contains(res.String(), "{{.IgnoreDeleteError}}"){{end}} {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}
	{{- end}}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}
//template:end delete

//template:begin import
{{- if not .NoImport}}
func (r *{{camelCase .Name}}Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	{{- if hasReference .Attributes}}
	idParts := strings.Split(req.ID, ",")

    if len(idParts) != {{importParts .Attributes}}{{range $index, $attr := .Attributes}}{{if $attr.Reference}} || idParts[{{$index}}] == ""{{end}}{{end}}  || idParts[{{subtract (importParts .Attributes) 1}}] == "" {
        resp.Diagnostics.AddError(
            "Unexpected Import Identifier",
            fmt.Sprintf("Expected import identifier with format: {{range $index, $attr := .Attributes}}{{if $attr.Reference}}{{if $index}},{{end}}<{{$attr.TfName}}>{{end}}{{end}},<id>. Got: %q", req.ID),
        )
        return
    }

	{{- range $index, $attr := .Attributes}}
	{{- if or $attr.Reference $attr.Id}}
    resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("{{$attr.TfName}}"), idParts[{{$index}}])...)
	{{- end}}
	{{- end}}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[{{subtract (importParts .Attributes) 1}}])...)
	{{- else}}
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	{{- end}}
}
{{- end}}
//template:end import

{{- if .IsBulk}}
//template:begin createSubresources

// createSubresources takes list of objects and creates them one by one
// We want to save the state after each create event, to be able track already created resources
func (r *{{camelCase .Name}}Resource) createSubresources(ctx context.Context, state, plan {{camelCase .Name}}) ({{camelCase .Name}}, diag.Diagnostics) {	

	tflog.Debug(ctx, fmt.Sprintf("%s: One-by-one creation mode ({{.Name}})", state.Id.ValueString()))
	var tmpObject {{camelCase .Name}}
    tmpObject.Items = make([]{{camelCase .Name}}Items, len(plan.Items))
	for k, v := range plan.Items {
		tmpObject.Items[k] = v

		body := tmpObject.toBody(ctx, state)
		res, _, err := r.client.Post(plan.getPath(), body)
		if err != nil {
			return state, diag.Diagnostics{
				diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("%s: Failed to create object (POST) id %s, got error: %s, %s", state.Id.ValueString(), v.Id.ValueString(), err, res.String())),
			}
		}

		// fromBodyUnknowns expect result to be listed under "items" key
		body, _ = sjson.SetRaw("{items:[]}", "items.-1", res.String())
		res = gjson.Parse(body)

		// Read computed values
		tmpObject.fromBodyUnknowns(ctx, res)

		// Save object to plan
		state.Items[k] = tmpObject.Items[k]

        // Clear tmpObject.Items
        tmpObject.Items[k] = {{camelCase .Name}}Items{}

	}

	return state, nil
}
//template:end createSubresources
{{- end}}