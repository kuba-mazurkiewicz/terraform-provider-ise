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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-ise"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &AllowedProtocolsDataSource{}
	_ datasource.DataSourceWithConfigure = &AllowedProtocolsDataSource{}
)

func NewAllowedProtocolsDataSource() datasource.DataSource {
	return &AllowedProtocolsDataSource{}
}

type AllowedProtocolsDataSource struct {
	client *ise.Client
}

func (d *AllowedProtocolsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_allowed_protocols"
}

func (d *AllowedProtocolsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read an allowed protocols policy element.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the allowed protocols",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description",
				Computed:            true,
			},
			"process_host_lookup": schema.BoolAttribute{
				MarkdownDescription: "Process host lookup",
				Computed:            true,
			},
			"allow_pap_ascii": schema.BoolAttribute{
				MarkdownDescription: "Allow PAP ASCII",
				Computed:            true,
			},
			"allow_chap": schema.BoolAttribute{
				MarkdownDescription: "Allow CHAP",
				Computed:            true,
			},
			"allow_ms_chap_v1": schema.BoolAttribute{
				MarkdownDescription: "Allow MS CHAP v1",
				Computed:            true,
			},
			"allow_ms_chap_v2": schema.BoolAttribute{
				MarkdownDescription: "Allow MS CHAP v2",
				Computed:            true,
			},
			"allow_eap_md5": schema.BoolAttribute{
				MarkdownDescription: "Allow EAP MD5",
				Computed:            true,
			},
			"allow_leap": schema.BoolAttribute{
				MarkdownDescription: "Allow LEAP",
				Computed:            true,
			},
			"allow_eap_tls": schema.BoolAttribute{
				MarkdownDescription: "Allow EAP TLS",
				Computed:            true,
			},
			"allow_eap_ttls": schema.BoolAttribute{
				MarkdownDescription: "Allow EAP TTLS",
				Computed:            true,
			},
			"allow_eap_fast": schema.BoolAttribute{
				MarkdownDescription: "Allow EAP Fast",
				Computed:            true,
			},
			"allow_peap": schema.BoolAttribute{
				MarkdownDescription: "Allow PEAP",
				Computed:            true,
			},
			"allow_teap": schema.BoolAttribute{
				MarkdownDescription: "Allow TEAP",
				Computed:            true,
			},
			"allow_preferred_eap_protocol": schema.BoolAttribute{
				MarkdownDescription: "Allow preferred EAP protocol",
				Computed:            true,
			},
			"preferred_eap_protocol": schema.StringAttribute{
				MarkdownDescription: "Preferred EAP protocol",
				Computed:            true,
			},
			"eap_tls_l_bit": schema.BoolAttribute{
				MarkdownDescription: "EAP TLS L-Bit",
				Computed:            true,
			},
			"allow_weak_ciphers_for_eap": schema.BoolAttribute{
				MarkdownDescription: "Allow weak ciphers for EAP",
				Computed:            true,
			},
			"require_message_auth": schema.BoolAttribute{
				MarkdownDescription: "Require message authentication",
				Computed:            true,
			},
			"eap_tls_allow_auth_of_expired_certs": schema.BoolAttribute{
				MarkdownDescription: "Allow authentication of expired certificates",
				Computed:            true,
			},
			"eap_tls_enable_stateless_session_resume": schema.BoolAttribute{
				MarkdownDescription: "Enable stateless session resume",
				Computed:            true,
			},
			"eap_tls_session_ticket_ttl": schema.Int64Attribute{
				MarkdownDescription: "Session ticket TTL. Is required only if `eap_tls_enable_stateless_session_resume` is `true`.",
				Computed:            true,
			},
			"eap_tls_session_ticket_ttl_unit": schema.StringAttribute{
				MarkdownDescription: "Session ticket TTL unit. Is required only if `eap_tls_enable_stateless_session_resume` is `true`.",
				Computed:            true,
			},
			"eap_tls_session_ticket_percentage": schema.Int64Attribute{
				MarkdownDescription: "Session ticket percentage. Is required only if `eap_tls_enable_stateless_session_resume` is `true`.",
				Computed:            true,
			},
			"peap_allow_peap_eap_ms_chap_v2": schema.BoolAttribute{
				MarkdownDescription: "Allow PEAP EAP MS CHAP v2",
				Computed:            true,
			},
			"peap_allow_peap_eap_ms_chap_v2_pwd_change": schema.BoolAttribute{
				MarkdownDescription: "Allow PEAP EAP MS CHAP v2 password change. Is required only if `allow_peap_eap_ms_chap_v2` is `true`.",
				Computed:            true,
			},
			"peap_allow_peap_eap_ms_chap_v2_pwd_change_retries": schema.Int64Attribute{
				MarkdownDescription: "Allow PEAP EAP MS CHAP v2 password change retries. Is required only if `allow_peap_eap_ms_chap_v2` is `true`.",
				Computed:            true,
			},
			"peap_allow_peap_eap_gtc": schema.BoolAttribute{
				MarkdownDescription: "Allow PEAP EAP GTC",
				Computed:            true,
			},
			"peap_allow_peap_eap_gtc_pwd_change": schema.BoolAttribute{
				MarkdownDescription: "Allow PEAP EAP GTC password change. Is required only if `allow_peap_eap_gtc` is `true`.",
				Computed:            true,
			},
			"peap_allow_peap_eap_gtc_pwd_change_retries": schema.Int64Attribute{
				MarkdownDescription: "PEAP EAP GTC password change retries. Is required only if `allow_peap_eap_gtc` is `true`.",
				Computed:            true,
			},
			"peap_allow_peap_eap_tls": schema.BoolAttribute{
				MarkdownDescription: "Allow PEAP EAP TLS",
				Computed:            true,
			},
			"peap_allow_peap_eap_tls_auth_of_expired_certs": schema.BoolAttribute{
				MarkdownDescription: "Allow PEAP EAP TLS authentication of expired certificates. Is required only if `peap_allow_peap_eap_tls` is `true`.",
				Computed:            true,
			},
			"require_cryptobinding": schema.BoolAttribute{
				MarkdownDescription: "Require cryptobinding",
				Computed:            true,
			},
			"peap_peap_v0": schema.BoolAttribute{
				MarkdownDescription: "Allow PEAP v0",
				Computed:            true,
			},
			"eap_ttls_pap_ascii": schema.BoolAttribute{
				MarkdownDescription: "Allow PAP ASCII",
				Computed:            true,
			},
			"eap_ttls_chap": schema.BoolAttribute{
				MarkdownDescription: "Allow CHAP",
				Computed:            true,
			},
			"eap_ttls_ms_chap_v1": schema.BoolAttribute{
				MarkdownDescription: "Allow MS CHAP v1",
				Computed:            true,
			},
			"eap_ttls_ms_chap_v2": schema.BoolAttribute{
				MarkdownDescription: "Allow MS CHAP v2",
				Computed:            true,
			},
			"eap_ttls_eap_md5": schema.BoolAttribute{
				MarkdownDescription: "Allow EAP MD5",
				Computed:            true,
			},
			"eap_ttls_eap_ms_chap_v2": schema.BoolAttribute{
				MarkdownDescription: "Allow EAP MS CHAP v2",
				Computed:            true,
			},
			"eap_ttls_eap_ms_chap_v2_pwd_change": schema.BoolAttribute{
				MarkdownDescription: "Allow EAP MS CHAP v2 password change. Is required only if `eap_ttls_eap_ms_chap_v2` is `true`.",
				Computed:            true,
			},
			"eap_ttls_eap_ms_chap_v2_pwd_change_retries": schema.Int64Attribute{
				MarkdownDescription: "EAP MS CHAP v2 password change retries. Is required only if `eap_ttls_eap_ms_chap_v2` is `true`.",
				Computed:            true,
			},
			"eap_fast_eap_ms_chap_v2": schema.BoolAttribute{
				MarkdownDescription: "Allow EAP MS CHAP v2",
				Computed:            true,
			},
			"eap_fast_eap_ms_chap_v2_pwd_change": schema.BoolAttribute{
				MarkdownDescription: "Allow EAP MS CHAP v2 password change. Is required only if `eap_fast_eap_ms_chap_v2` is `true`.",
				Computed:            true,
			},
			"eap_fast_eap_ms_chap_v2_pwd_change_retries": schema.Int64Attribute{
				MarkdownDescription: "EAP MS CHAP v2 password change retries. Is required only if `eap_fast_eap_ms_chap_v2` is `true`.",
				Computed:            true,
			},
			"eap_fast_eap_gtc": schema.BoolAttribute{
				MarkdownDescription: "Allow EAP GTC",
				Computed:            true,
			},
			"eap_fast_eap_gtc_pwd_change": schema.BoolAttribute{
				MarkdownDescription: "Allow EAP GTC password change. Is required only if `eap_fast_eap_gtc` is `true`.",
				Computed:            true,
			},
			"eap_fast_eap_gtc_pwd_change_retries": schema.Int64Attribute{
				MarkdownDescription: "EAP GTC password change retries. Is required only if `eap_fast_eap_gtc` is `true`.",
				Computed:            true,
			},
			"eap_fast_eap_tls": schema.BoolAttribute{
				MarkdownDescription: "Allow EAP TLS",
				Computed:            true,
			},
			"eap_fast_eap_tls_auth_of_expired_certs": schema.BoolAttribute{
				MarkdownDescription: "Allow EAP TLS authentication of expired certificates. Is required only if `eap_fast_eap_tls` is `true`.",
				Computed:            true,
			},
			"eap_fast_enable_eap_chaining": schema.BoolAttribute{
				MarkdownDescription: "Enable EAP chaining",
				Computed:            true,
			},
			"eap_fast_use_pacs": schema.BoolAttribute{
				MarkdownDescription: "Use PACs",
				Computed:            true,
			},
			"eap_fast_pacs_tunnel_pac_ttl": schema.Int64Attribute{
				MarkdownDescription: "PACs tunnel PAC time to live. Is required only if `eap_fast_use_pacs` is `true`.",
				Computed:            true,
			},
			"eap_fast_pacs_tunnel_pac_ttl_unit": schema.StringAttribute{
				MarkdownDescription: "PACs tunnel PAC time to live unit. Is required only if `eap_fast_use_pacs` is `true`.",
				Computed:            true,
			},
			"eap_fast_pacs_use_proactive_pac_update_percentage": schema.Int64Attribute{
				MarkdownDescription: "Use proactive pac update percentage. Is required only if `eap_fast_use_pacs` is `true`.",
				Computed:            true,
			},
			"eap_fast_pacs_allow_anonymous_provisioning": schema.BoolAttribute{
				MarkdownDescription: "Allow anonymous provisioning. Is required only if `eap_fast_use_pacs` is `true`.",
				Computed:            true,
			},
			"eap_fast_pacs_allow_authenticated_provisioning": schema.BoolAttribute{
				MarkdownDescription: "Allow authenticated provisioning. Is required only if `eap_fast_use_pacs` is `true`.",
				Computed:            true,
			},
			"eap_fast_pacs_server_returns": schema.BoolAttribute{
				MarkdownDescription: "Server returns access accept after authenticated provisioning. Is required only if `eap_fast_pacs_allow_authenticated_provisioning` is `true`.",
				Computed:            true,
			},
			"eap_fast_pacs_allow_client_cert": schema.BoolAttribute{
				MarkdownDescription: "Accept client certification for provisioning. Is required only if `eap_fast_pacs_allow_authenticated_provisioning` is `true`.",
				Computed:            true,
			},
			"eap_fast_pacs_allow_machine_authentication": schema.BoolAttribute{
				MarkdownDescription: "Allow machine authentication. Is required only if `eap_fast_use_pacs` is `true`.",
				Computed:            true,
			},
			"eap_fast_pacs_machine_pac_ttl": schema.Int64Attribute{
				MarkdownDescription: "Machine PAC TTL. Is required only if `eap_fast_pacs_allow_machine_authentication` is `true`.",
				Computed:            true,
			},
			"eap_fast_pacs_machine_pac_ttl_unit": schema.StringAttribute{
				MarkdownDescription: "Machine PAC TTL unit. Is required only if `eap_fast_pacs_allow_machine_authentication` is `true`.",
				Computed:            true,
			},
			"eap_fast_pacs_stateless_session_resume": schema.BoolAttribute{
				MarkdownDescription: "Stateless session resume. Is required only if `eap_fast_use_pacs` is `true`.",
				Computed:            true,
			},
			"eap_fast_pacs_authorization_pac_ttl": schema.Int64Attribute{
				MarkdownDescription: "Authorization PAC TTL. Is required only if `eap_fast_pacs_stateless_session_resume` is `true`.",
				Computed:            true,
			},
			"eap_fast_pacs_authorization_pac_ttl_unit": schema.StringAttribute{
				MarkdownDescription: "Authorization PAC TTL unit. Is required only if `eap_fast_pacs_stateless_session_resume` is `true`.",
				Computed:            true,
			},
			"eap_fast_accept_client_cert": schema.BoolAttribute{
				MarkdownDescription: "Accept client certificates. Is required only if `eap_fast_use_pacs` is `false`.",
				Computed:            true,
			},
			"eap_fast_allow_machine_authentication": schema.BoolAttribute{
				MarkdownDescription: "Allow machine authentication. Is required only if `eap_fast_use_pacs` is `false`.",
				Computed:            true,
			},
			"teap_eap_ms_chap_v2": schema.BoolAttribute{
				MarkdownDescription: "Allow EAP MS CHAP v2",
				Computed:            true,
			},
			"teap_eap_ms_chap_v2_pwd_change": schema.BoolAttribute{
				MarkdownDescription: "Allow EAP MS CHAP v2 password change. Is required only if `teap_eap_ms_chap_v2` is `true`.",
				Computed:            true,
			},
			"teap_eap_ms_chap_v2_pwd_change_retries": schema.Int64Attribute{
				MarkdownDescription: "EAP MS CHAP v2 password change retries. Is required only if `teap_eap_ms_chap_v2` is `true`.",
				Computed:            true,
			},
			"teap_eap_tls": schema.BoolAttribute{
				MarkdownDescription: "Allow EAP TLS",
				Computed:            true,
			},
			"teap_eap_tls_auth_of_expired_certs": schema.BoolAttribute{
				MarkdownDescription: "Allow EAP TLS authentication of expired certs. Is required only if `teap_eap_tls` is `true`.",
				Computed:            true,
			},
			"teap_eap_accept_client_cert_during_tunnel_est": schema.BoolAttribute{
				MarkdownDescription: "Accept client certificate during tunnel establishment",
				Computed:            true,
			},
			"teap_eap_chaining": schema.BoolAttribute{
				MarkdownDescription: "Allow EAP chaining",
				Computed:            true,
			},
			"teap_downgrade_msk": schema.BoolAttribute{
				MarkdownDescription: "Allow downgrade to MSK",
				Computed:            true,
			},
			"teap_request_basic_pwd_auth": schema.BoolAttribute{
				MarkdownDescription: "Request basic password authentication",
				Computed:            true,
			},
			"allow_5g": schema.BoolAttribute{
				MarkdownDescription: "Allow 5G. This field is only supported from ISE 3.2.",
				Computed:            true,
			},
		},
	}
}

func (d *AllowedProtocolsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*IseProviderData).Client
}

func (d *AllowedProtocolsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config AllowedProtocols

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get("/ers/config/allowedprotocols" + "/" + config.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
