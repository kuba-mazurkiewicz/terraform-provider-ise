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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-ise"
	"github.com/netascode/terraform-provider-ise/internal/provider/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &CertificateAuthenticationProfileResource{}
var _ resource.ResourceWithImportState = &CertificateAuthenticationProfileResource{}

func NewCertificateAuthenticationProfileResource() resource.Resource {
	return &CertificateAuthenticationProfileResource{}
}

type CertificateAuthenticationProfileResource struct {
	client *ise.Client
}

func (r *CertificateAuthenticationProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_certificate_authentication_profile"
}

func (r *CertificateAuthenticationProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Certificate Authentication Profile.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the certificate profile").String,
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description").String,
				Optional:            true,
			},
			"allowed_as_user_name": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow as username").String,
				Optional:            true,
			},
			"external_identity_store_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Referred IDStore name for the Certificate Profile or `[not applicable]` in case no identity store is chosen").String,
				Optional:            true,
			},
			"certificate_attribute_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Attribute name of the Certificate Profile - used only when CERTIFICATE is chosen in `username_from`.").AddStringEnumDescription("SUBJECT_COMMON_NAME", "SUBJECT_ALTERNATIVE_NAME", "SUBJECT_SERIAL_NUMBER", "SUBJECT", "SUBJECT_ALTERNATIVE_NAME_OTHER_NAME", "SUBJECT_ALTERNATIVE_NAME_EMAIL", "SUBJECT_ALTERNATIVE_NAME_DNS").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("SUBJECT_COMMON_NAME", "SUBJECT_ALTERNATIVE_NAME", "SUBJECT_SERIAL_NUMBER", "SUBJECT", "SUBJECT_ALTERNATIVE_NAME_OTHER_NAME", "SUBJECT_ALTERNATIVE_NAME_EMAIL", "SUBJECT_ALTERNATIVE_NAME_DNS"),
				},
			},
			"match_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Match mode of the Certificate Profile. Allowed values: NEVER, RESOLVE_IDENTITY_AMBIGUITY, BINARY_COMPARISON").AddStringEnumDescription("NEVER", "RESOLVE_IDENTITY_AMBIGUITY", "BINARY_COMPARISON").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("NEVER", "RESOLVE_IDENTITY_AMBIGUITY", "BINARY_COMPARISON"),
				},
			},
			"username_from": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The attribute in the certificate where the user name should be taken from. Allowed values: `CERTIFICATE` (for a specific attribute as defined in certificateAttributeName), `UPN` (for using any Subject or Alternative Name Attributes in the Certificate - an option only in AD)").AddStringEnumDescription("CERTIFICATE", "UPN").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("CERTIFICATE", "UPN"),
				},
			},
		},
	}
}

func (r *CertificateAuthenticationProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*IseProviderData).Client
}

func (r *CertificateAuthenticationProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CertificateAuthenticationProfile

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, CertificateAuthenticationProfile{})
	res, location, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}
	locationElements := strings.Split(location, "/")
	plan.Id = types.StringValue(locationElements[len(locationElements)-1])

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CertificateAuthenticationProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CertificateAuthenticationProfile

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	res, err := r.client.Get(state.getPath() + "/" + state.Id.ValueString())
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.updateFromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *CertificateAuthenticationProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CertificateAuthenticationProfile

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

	body := plan.toBody(ctx, state)

	res, err := r.client.Put(plan.getPath()+"/"+plan.Id.ValueString(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CertificateAuthenticationProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CertificateAuthenticationProfile

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath() + "/" + state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *CertificateAuthenticationProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
