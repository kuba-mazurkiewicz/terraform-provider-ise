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
	"github.com/netascode/terraform-provider-ise/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type NetworkDevice struct {
	Id                                                   types.String       `tfsdk:"id"`
	Name                                                 types.String       `tfsdk:"name"`
	Description                                          types.String       `tfsdk:"description"`
	AuthenticationEnableKeyWrap                          types.Bool         `tfsdk:"authentication_enable_key_wrap"`
	AuthenticationEncryptionKey                          types.String       `tfsdk:"authentication_encryption_key"`
	AuthenticationEncryptionKeyFormat                    types.String       `tfsdk:"authentication_encryption_key_format"`
	AuthenticationMessageAuthenticatorCodeKey            types.String       `tfsdk:"authentication_message_authenticator_code_key"`
	AuthenticationNetworkProtocol                        types.String       `tfsdk:"authentication_network_protocol"`
	AuthenticationRadiusSharedSecret                     types.String       `tfsdk:"authentication_radius_shared_secret"`
	AuthenticationEnableMultiSecret                      types.Bool         `tfsdk:"authentication_enable_multi_secret"`
	AuthenticationSecondRadiusSharedSecret               types.String       `tfsdk:"authentication_second_radius_shared_secret"`
	AuthenticationDtlsRequired                           types.Bool         `tfsdk:"authentication_dtls_required"`
	CoaPort                                              types.Int64        `tfsdk:"coa_port"`
	DtlsDnsName                                          types.String       `tfsdk:"dtls_dns_name"`
	Ips                                                  []NetworkDeviceIps `tfsdk:"ips"`
	NetworkDeviceGroups                                  types.List         `tfsdk:"network_device_groups"`
	ModelName                                            types.String       `tfsdk:"model_name"`
	SoftwareVersion                                      types.String       `tfsdk:"software_version"`
	ProfileName                                          types.String       `tfsdk:"profile_name"`
	SnmpLinkTrapQuery                                    types.Bool         `tfsdk:"snmp_link_trap_query"`
	SnmpMacTrapQuery                                     types.Bool         `tfsdk:"snmp_mac_trap_query"`
	SnmpOriginatingPolicyServiceNode                     types.String       `tfsdk:"snmp_originating_policy_service_node"`
	SnmpPollingInterval                                  types.Int64        `tfsdk:"snmp_polling_interval"`
	SnmpRoCommunity                                      types.String       `tfsdk:"snmp_ro_community"`
	SnmpVersion                                          types.String       `tfsdk:"snmp_version"`
	TacacsConnectModeOptions                             types.String       `tfsdk:"tacacs_connect_mode_options"`
	TacacsSharedSecret                                   types.String       `tfsdk:"tacacs_shared_secret"`
	TrustsecDeviceId                                     types.String       `tfsdk:"trustsec_device_id"`
	TrustsecDevicePassword                               types.String       `tfsdk:"trustsec_device_password"`
	TrustsecRestApiUsername                              types.String       `tfsdk:"trustsec_rest_api_username"`
	TrustsecRestApiPassword                              types.String       `tfsdk:"trustsec_rest_api_password"`
	TrustsecEnableModePassword                           types.String       `tfsdk:"trustsec_enable_mode_password"`
	TrustsecExecModePassword                             types.String       `tfsdk:"trustsec_exec_mode_password"`
	TrustsecExecModeUsername                             types.String       `tfsdk:"trustsec_exec_mode_username"`
	TrustsecIncludeWhenDeployingSgtUpdates               types.Bool         `tfsdk:"trustsec_include_when_deploying_sgt_updates"`
	TrustsecDownloadEnviromentDataEveryXSeconds          types.Int64        `tfsdk:"trustsec_download_enviroment_data_every_x_seconds"`
	TrustsecDownloadPeerAuthorizationPolicyEveryXSeconds types.Int64        `tfsdk:"trustsec_download_peer_authorization_policy_every_x_seconds"`
	TrustsecDownloadSgaclListsEveryXSeconds              types.Int64        `tfsdk:"trustsec_download_sgacl_lists_every_x_seconds"`
	TrustsecOtherSgaDevicesToTrustThisDevice             types.Bool         `tfsdk:"trustsec_other_sga_devices_to_trust_this_device"`
	TrustsecReAuthenticationEveryXSeconds                types.Int64        `tfsdk:"trustsec_re_authentication_every_x_seconds"`
	TrustsecSendConfigurationToDevice                    types.Bool         `tfsdk:"trustsec_send_configuration_to_device"`
	TrustsecSendConfigurationToDeviceUsing               types.String       `tfsdk:"trustsec_send_configuration_to_device_using"`
	TrustsecCoaSourceHost                                types.String       `tfsdk:"trustsec_coa_source_host"`
}

type NetworkDeviceIps struct {
	Ipaddress        types.String `tfsdk:"ipaddress"`
	IpaddressExclude types.String `tfsdk:"ipaddress_exclude"`
	Mask             types.String `tfsdk:"mask"`
}

func (data NetworkDevice) getPath() string {
	return "/ers/config/networkdevice"
}

func (data NetworkDevice) toBody(ctx context.Context, state NetworkDevice) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.description", data.Description.ValueString())
	}
	if !data.AuthenticationEnableKeyWrap.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.authenticationSettings.enableKeyWrap", data.AuthenticationEnableKeyWrap.ValueBool())
	}
	if !data.AuthenticationEncryptionKey.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.authenticationSettings.keyEncryptionKey", data.AuthenticationEncryptionKey.ValueString())
	}
	if !data.AuthenticationEncryptionKeyFormat.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.authenticationSettings.keyInputFormat", data.AuthenticationEncryptionKeyFormat.ValueString())
	}
	if !data.AuthenticationMessageAuthenticatorCodeKey.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.authenticationSettings.messageAuthenticatorCodeKey", data.AuthenticationMessageAuthenticatorCodeKey.ValueString())
	}
	if !data.AuthenticationNetworkProtocol.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.authenticationSettings.networkProtocol", data.AuthenticationNetworkProtocol.ValueString())
	}
	if !data.AuthenticationRadiusSharedSecret.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.authenticationSettings.radiusSharedSecret", data.AuthenticationRadiusSharedSecret.ValueString())
	}
	if !data.AuthenticationEnableMultiSecret.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.authenticationSettings.enableMultiSecret", data.AuthenticationEnableMultiSecret.ValueBool())
	}
	if !data.AuthenticationSecondRadiusSharedSecret.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.authenticationSettings.secondRadiusSharedSecret", data.AuthenticationSecondRadiusSharedSecret.ValueString())
	}
	if !data.AuthenticationDtlsRequired.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.authenticationSettings.dtlsRequired", data.AuthenticationDtlsRequired.ValueBool())
	}
	if !data.CoaPort.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.coaPort", data.CoaPort.ValueInt64())
	}
	if !data.DtlsDnsName.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.dtlsDnsName", data.DtlsDnsName.ValueString())
	}
	if len(data.Ips) > 0 {
		body, _ = sjson.Set(body, "NetworkDevice.NetworkDeviceIPList", []interface{}{})
		for _, item := range data.Ips {
			itemBody := ""
			if !item.Ipaddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipaddress", item.Ipaddress.ValueString())
			}
			if !item.IpaddressExclude.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipaddressExclude", item.IpaddressExclude.ValueString())
			}
			if !item.Mask.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "mask", item.Mask.ValueString())
			}
			body, _ = sjson.SetRaw(body, "NetworkDevice.NetworkDeviceIPList.-1", itemBody)
		}
	}
	if !data.NetworkDeviceGroups.IsNull() {
		var values []string
		data.NetworkDeviceGroups.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "NetworkDevice.NetworkDeviceGroupList", values)
	}
	if !data.ModelName.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.modelName", data.ModelName.ValueString())
	}
	if !data.SoftwareVersion.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.softwareVersion", data.SoftwareVersion.ValueString())
	}
	if !data.ProfileName.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.profileName", data.ProfileName.ValueString())
	}
	if !data.SnmpLinkTrapQuery.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.snmpsettings.linkTrapQuery", data.SnmpLinkTrapQuery.ValueBool())
	}
	if !data.SnmpMacTrapQuery.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.snmpsettings.macTrapQuery", data.SnmpMacTrapQuery.ValueBool())
	}
	if !data.SnmpOriginatingPolicyServiceNode.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.snmpsettings.originatingPolicyServicesNode", data.SnmpOriginatingPolicyServiceNode.ValueString())
	}
	if !data.SnmpPollingInterval.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.snmpsettings.pollingInterval", data.SnmpPollingInterval.ValueInt64())
	}
	if !data.SnmpRoCommunity.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.snmpsettings.roCommunity", data.SnmpRoCommunity.ValueString())
	}
	if !data.SnmpVersion.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.snmpsettings.version", data.SnmpVersion.ValueString())
	}
	if !data.TacacsConnectModeOptions.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.tacacsSettings.connectModeOptions", data.TacacsConnectModeOptions.ValueString())
	}
	if !data.TacacsSharedSecret.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.tacacsSettings.sharedSecret", data.TacacsSharedSecret.ValueString())
	}
	if !data.TrustsecDeviceId.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.trustsecsettings.deviceAuthenticationSettings.sgaDeviceId", data.TrustsecDeviceId.ValueString())
	}
	if !data.TrustsecDevicePassword.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.trustsecsettings.deviceAuthenticationSettings.sgaDevicePassword", data.TrustsecDevicePassword.ValueString())
	}
	if !data.TrustsecRestApiUsername.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.trustsecsettings.deviceAuthenticationSettings.restApiUsername", data.TrustsecRestApiUsername.ValueString())
	}
	if !data.TrustsecRestApiPassword.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.trustsecsettings.deviceAuthenticationSettings.restApiPassword", data.TrustsecRestApiPassword.ValueString())
	}
	if !data.TrustsecEnableModePassword.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.trustsecsettings.deviceConfigurationDeployment.enableModePassword", data.TrustsecEnableModePassword.ValueString())
	}
	if !data.TrustsecExecModePassword.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.trustsecsettings.deviceConfigurationDeployment.execModePassword", data.TrustsecExecModePassword.ValueString())
	}
	if !data.TrustsecExecModeUsername.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.trustsecsettings.deviceConfigurationDeployment.execModeUsername", data.TrustsecExecModeUsername.ValueString())
	}
	if !data.TrustsecIncludeWhenDeployingSgtUpdates.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.trustsecsettings.deviceConfigurationDeployment.includeWhenDeployingSGTUpdates", data.TrustsecIncludeWhenDeployingSgtUpdates.ValueBool())
	}
	if !data.TrustsecDownloadEnviromentDataEveryXSeconds.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.trustsecsettings.sgaNotificationAndUpdates.downlaodEnvironmentDataEveryXSeconds", data.TrustsecDownloadEnviromentDataEveryXSeconds.ValueInt64())
	}
	if !data.TrustsecDownloadPeerAuthorizationPolicyEveryXSeconds.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.trustsecsettings.sgaNotificationAndUpdates.downlaodPeerAuthorizationPolicyEveryXSeconds", data.TrustsecDownloadPeerAuthorizationPolicyEveryXSeconds.ValueInt64())
	}
	if !data.TrustsecDownloadSgaclListsEveryXSeconds.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.trustsecsettings.sgaNotificationAndUpdates.downloadSGACLListsEveryXSeconds", data.TrustsecDownloadSgaclListsEveryXSeconds.ValueInt64())
	}
	if !data.TrustsecOtherSgaDevicesToTrustThisDevice.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.trustsecsettings.sgaNotificationAndUpdates.otherSGADevicesToTrustThisDevice", data.TrustsecOtherSgaDevicesToTrustThisDevice.ValueBool())
	}
	if !data.TrustsecReAuthenticationEveryXSeconds.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.trustsecsettings.sgaNotificationAndUpdates.reAuthenticationEveryXSeconds", data.TrustsecReAuthenticationEveryXSeconds.ValueInt64())
	}
	if !data.TrustsecSendConfigurationToDevice.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.trustsecsettings.sgaNotificationAndUpdates.sendConfigurationToDevice", data.TrustsecSendConfigurationToDevice.ValueBool())
	}
	if !data.TrustsecSendConfigurationToDeviceUsing.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.trustsecsettings.sgaNotificationAndUpdates.sendConfigurationToDeviceUsing", data.TrustsecSendConfigurationToDeviceUsing.ValueString())
	}
	if !data.TrustsecCoaSourceHost.IsNull() {
		body, _ = sjson.Set(body, "NetworkDevice.trustsecsettings.sgaNotificationAndUpdates.coaSourceHost", data.TrustsecCoaSourceHost.ValueString())
	}
	return body
}

func (data *NetworkDevice) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("NetworkDevice.name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("NetworkDevice.description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("NetworkDevice.authenticationSettings.enableKeyWrap"); value.Exists() {
		data.AuthenticationEnableKeyWrap = types.BoolValue(value.Bool())
	} else {
		data.AuthenticationEnableKeyWrap = types.BoolNull()
	}
	if value := res.Get("NetworkDevice.authenticationSettings.keyEncryptionKey"); value.Exists() {
		data.AuthenticationEncryptionKey = types.StringValue(value.String())
	} else {
		data.AuthenticationEncryptionKey = types.StringNull()
	}
	if value := res.Get("NetworkDevice.authenticationSettings.keyInputFormat"); value.Exists() {
		data.AuthenticationEncryptionKeyFormat = types.StringValue(value.String())
	} else {
		data.AuthenticationEncryptionKeyFormat = types.StringNull()
	}
	if value := res.Get("NetworkDevice.authenticationSettings.messageAuthenticatorCodeKey"); value.Exists() {
		data.AuthenticationMessageAuthenticatorCodeKey = types.StringValue(value.String())
	} else {
		data.AuthenticationMessageAuthenticatorCodeKey = types.StringNull()
	}
	if value := res.Get("NetworkDevice.authenticationSettings.networkProtocol"); value.Exists() {
		data.AuthenticationNetworkProtocol = types.StringValue(value.String())
	} else {
		data.AuthenticationNetworkProtocol = types.StringNull()
	}
	if value := res.Get("NetworkDevice.authenticationSettings.radiusSharedSecret"); value.Exists() {
		data.AuthenticationRadiusSharedSecret = types.StringValue(value.String())
	} else {
		data.AuthenticationRadiusSharedSecret = types.StringNull()
	}
	if value := res.Get("NetworkDevice.authenticationSettings.enableMultiSecret"); value.Exists() {
		data.AuthenticationEnableMultiSecret = types.BoolValue(value.Bool())
	} else {
		data.AuthenticationEnableMultiSecret = types.BoolNull()
	}
	if value := res.Get("NetworkDevice.authenticationSettings.secondRadiusSharedSecret"); value.Exists() {
		data.AuthenticationSecondRadiusSharedSecret = types.StringValue(value.String())
	} else {
		data.AuthenticationSecondRadiusSharedSecret = types.StringNull()
	}
	if value := res.Get("NetworkDevice.authenticationSettings.dtlsRequired"); value.Exists() {
		data.AuthenticationDtlsRequired = types.BoolValue(value.Bool())
	} else {
		data.AuthenticationDtlsRequired = types.BoolNull()
	}
	if value := res.Get("NetworkDevice.coaPort"); value.Exists() {
		data.CoaPort = types.Int64Value(value.Int())
	} else {
		data.CoaPort = types.Int64Null()
	}
	if value := res.Get("NetworkDevice.dtlsDnsName"); value.Exists() {
		data.DtlsDnsName = types.StringValue(value.String())
	} else {
		data.DtlsDnsName = types.StringNull()
	}
	if value := res.Get("NetworkDevice.NetworkDeviceIPList"); value.Exists() {
		data.Ips = make([]NetworkDeviceIps, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := NetworkDeviceIps{}
			if cValue := v.Get("ipaddress"); cValue.Exists() {
				item.Ipaddress = types.StringValue(cValue.String())
			} else {
				item.Ipaddress = types.StringNull()
			}
			if cValue := v.Get("ipaddressExclude"); cValue.Exists() {
				item.IpaddressExclude = types.StringValue(cValue.String())
			} else {
				item.IpaddressExclude = types.StringNull()
			}
			if cValue := v.Get("mask"); cValue.Exists() {
				item.Mask = types.StringValue(cValue.String())
			} else {
				item.Mask = types.StringNull()
			}
			data.Ips = append(data.Ips, item)
			return true
		})
	}
	if value := res.Get("NetworkDevice.NetworkDeviceGroupList"); value.Exists() {
		data.NetworkDeviceGroups = helpers.GetStringList(value.Array())
	} else {
		data.NetworkDeviceGroups = types.ListNull(types.StringType)
	}
	if value := res.Get("NetworkDevice.modelName"); value.Exists() {
		data.ModelName = types.StringValue(value.String())
	} else {
		data.ModelName = types.StringNull()
	}
	if value := res.Get("NetworkDevice.softwareVersion"); value.Exists() {
		data.SoftwareVersion = types.StringValue(value.String())
	} else {
		data.SoftwareVersion = types.StringNull()
	}
	if value := res.Get("NetworkDevice.profileName"); value.Exists() {
		data.ProfileName = types.StringValue(value.String())
	} else {
		data.ProfileName = types.StringNull()
	}
	if value := res.Get("NetworkDevice.snmpsettings.linkTrapQuery"); value.Exists() {
		data.SnmpLinkTrapQuery = types.BoolValue(value.Bool())
	} else {
		data.SnmpLinkTrapQuery = types.BoolNull()
	}
	if value := res.Get("NetworkDevice.snmpsettings.macTrapQuery"); value.Exists() {
		data.SnmpMacTrapQuery = types.BoolValue(value.Bool())
	} else {
		data.SnmpMacTrapQuery = types.BoolNull()
	}
	if value := res.Get("NetworkDevice.snmpsettings.originatingPolicyServicesNode"); value.Exists() {
		data.SnmpOriginatingPolicyServiceNode = types.StringValue(value.String())
	} else {
		data.SnmpOriginatingPolicyServiceNode = types.StringNull()
	}
	if value := res.Get("NetworkDevice.snmpsettings.pollingInterval"); value.Exists() {
		data.SnmpPollingInterval = types.Int64Value(value.Int())
	} else {
		data.SnmpPollingInterval = types.Int64Null()
	}
	if value := res.Get("NetworkDevice.snmpsettings.roCommunity"); value.Exists() {
		data.SnmpRoCommunity = types.StringValue(value.String())
	} else {
		data.SnmpRoCommunity = types.StringNull()
	}
	if value := res.Get("NetworkDevice.snmpsettings.version"); value.Exists() {
		data.SnmpVersion = types.StringValue(value.String())
	} else {
		data.SnmpVersion = types.StringNull()
	}
	if value := res.Get("NetworkDevice.tacacsSettings.connectModeOptions"); value.Exists() {
		data.TacacsConnectModeOptions = types.StringValue(value.String())
	} else {
		data.TacacsConnectModeOptions = types.StringNull()
	}
	if value := res.Get("NetworkDevice.tacacsSettings.sharedSecret"); value.Exists() {
		data.TacacsSharedSecret = types.StringValue(value.String())
	} else {
		data.TacacsSharedSecret = types.StringNull()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.deviceAuthenticationSettings.sgaDeviceId"); value.Exists() {
		data.TrustsecDeviceId = types.StringValue(value.String())
	} else {
		data.TrustsecDeviceId = types.StringNull()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.deviceAuthenticationSettings.sgaDevicePassword"); value.Exists() {
		data.TrustsecDevicePassword = types.StringValue(value.String())
	} else {
		data.TrustsecDevicePassword = types.StringNull()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.deviceAuthenticationSettings.restApiUsername"); value.Exists() {
		data.TrustsecRestApiUsername = types.StringValue(value.String())
	} else {
		data.TrustsecRestApiUsername = types.StringNull()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.deviceConfigurationDeployment.enableModePassword"); value.Exists() {
		data.TrustsecEnableModePassword = types.StringValue(value.String())
	} else {
		data.TrustsecEnableModePassword = types.StringNull()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.deviceConfigurationDeployment.execModePassword"); value.Exists() {
		data.TrustsecExecModePassword = types.StringValue(value.String())
	} else {
		data.TrustsecExecModePassword = types.StringNull()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.deviceConfigurationDeployment.execModeUsername"); value.Exists() {
		data.TrustsecExecModeUsername = types.StringValue(value.String())
	} else {
		data.TrustsecExecModeUsername = types.StringNull()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.deviceConfigurationDeployment.includeWhenDeployingSGTUpdates"); value.Exists() {
		data.TrustsecIncludeWhenDeployingSgtUpdates = types.BoolValue(value.Bool())
	} else {
		data.TrustsecIncludeWhenDeployingSgtUpdates = types.BoolNull()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.sgaNotificationAndUpdates.downlaodEnvironmentDataEveryXSeconds"); value.Exists() {
		data.TrustsecDownloadEnviromentDataEveryXSeconds = types.Int64Value(value.Int())
	} else {
		data.TrustsecDownloadEnviromentDataEveryXSeconds = types.Int64Null()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.sgaNotificationAndUpdates.downlaodPeerAuthorizationPolicyEveryXSeconds"); value.Exists() {
		data.TrustsecDownloadPeerAuthorizationPolicyEveryXSeconds = types.Int64Value(value.Int())
	} else {
		data.TrustsecDownloadPeerAuthorizationPolicyEveryXSeconds = types.Int64Null()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.sgaNotificationAndUpdates.downloadSGACLListsEveryXSeconds"); value.Exists() {
		data.TrustsecDownloadSgaclListsEveryXSeconds = types.Int64Value(value.Int())
	} else {
		data.TrustsecDownloadSgaclListsEveryXSeconds = types.Int64Null()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.sgaNotificationAndUpdates.otherSGADevicesToTrustThisDevice"); value.Exists() {
		data.TrustsecOtherSgaDevicesToTrustThisDevice = types.BoolValue(value.Bool())
	} else {
		data.TrustsecOtherSgaDevicesToTrustThisDevice = types.BoolNull()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.sgaNotificationAndUpdates.reAuthenticationEveryXSeconds"); value.Exists() {
		data.TrustsecReAuthenticationEveryXSeconds = types.Int64Value(value.Int())
	} else {
		data.TrustsecReAuthenticationEveryXSeconds = types.Int64Null()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.sgaNotificationAndUpdates.sendConfigurationToDevice"); value.Exists() {
		data.TrustsecSendConfigurationToDevice = types.BoolValue(value.Bool())
	} else {
		data.TrustsecSendConfigurationToDevice = types.BoolNull()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.sgaNotificationAndUpdates.sendConfigurationToDeviceUsing"); value.Exists() {
		data.TrustsecSendConfigurationToDeviceUsing = types.StringValue(value.String())
	} else {
		data.TrustsecSendConfigurationToDeviceUsing = types.StringNull()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.sgaNotificationAndUpdates.coaSourceHost"); value.Exists() {
		data.TrustsecCoaSourceHost = types.StringValue(value.String())
	} else {
		data.TrustsecCoaSourceHost = types.StringNull()
	}
}

func (data *NetworkDevice) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("NetworkDevice.name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("NetworkDevice.description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("NetworkDevice.authenticationSettings.enableKeyWrap"); value.Exists() && !data.AuthenticationEnableKeyWrap.IsNull() {
		data.AuthenticationEnableKeyWrap = types.BoolValue(value.Bool())
	} else {
		data.AuthenticationEnableKeyWrap = types.BoolNull()
	}
	if value := res.Get("NetworkDevice.authenticationSettings.keyEncryptionKey"); value.Exists() && !data.AuthenticationEncryptionKey.IsNull() {
		data.AuthenticationEncryptionKey = types.StringValue(value.String())
	} else {
		data.AuthenticationEncryptionKey = types.StringNull()
	}
	if value := res.Get("NetworkDevice.authenticationSettings.keyInputFormat"); value.Exists() && !data.AuthenticationEncryptionKeyFormat.IsNull() {
		data.AuthenticationEncryptionKeyFormat = types.StringValue(value.String())
	} else {
		data.AuthenticationEncryptionKeyFormat = types.StringNull()
	}
	if value := res.Get("NetworkDevice.authenticationSettings.messageAuthenticatorCodeKey"); value.Exists() && !data.AuthenticationMessageAuthenticatorCodeKey.IsNull() {
		data.AuthenticationMessageAuthenticatorCodeKey = types.StringValue(value.String())
	} else {
		data.AuthenticationMessageAuthenticatorCodeKey = types.StringNull()
	}
	if value := res.Get("NetworkDevice.authenticationSettings.networkProtocol"); value.Exists() && !data.AuthenticationNetworkProtocol.IsNull() {
		data.AuthenticationNetworkProtocol = types.StringValue(value.String())
	} else {
		data.AuthenticationNetworkProtocol = types.StringNull()
	}
	if value := res.Get("NetworkDevice.authenticationSettings.radiusSharedSecret"); value.Exists() && !data.AuthenticationRadiusSharedSecret.IsNull() {
		data.AuthenticationRadiusSharedSecret = types.StringValue(value.String())
	} else {
		data.AuthenticationRadiusSharedSecret = types.StringNull()
	}
	if value := res.Get("NetworkDevice.authenticationSettings.enableMultiSecret"); value.Exists() && !data.AuthenticationEnableMultiSecret.IsNull() {
		data.AuthenticationEnableMultiSecret = types.BoolValue(value.Bool())
	} else {
		data.AuthenticationEnableMultiSecret = types.BoolNull()
	}
	if value := res.Get("NetworkDevice.authenticationSettings.secondRadiusSharedSecret"); value.Exists() && !data.AuthenticationSecondRadiusSharedSecret.IsNull() {
		data.AuthenticationSecondRadiusSharedSecret = types.StringValue(value.String())
	} else {
		data.AuthenticationSecondRadiusSharedSecret = types.StringNull()
	}
	if value := res.Get("NetworkDevice.authenticationSettings.dtlsRequired"); value.Exists() && !data.AuthenticationDtlsRequired.IsNull() {
		data.AuthenticationDtlsRequired = types.BoolValue(value.Bool())
	} else {
		data.AuthenticationDtlsRequired = types.BoolNull()
	}
	if value := res.Get("NetworkDevice.coaPort"); value.Exists() && !data.CoaPort.IsNull() {
		data.CoaPort = types.Int64Value(value.Int())
	} else {
		data.CoaPort = types.Int64Null()
	}
	if value := res.Get("NetworkDevice.dtlsDnsName"); value.Exists() && !data.DtlsDnsName.IsNull() {
		data.DtlsDnsName = types.StringValue(value.String())
	} else {
		data.DtlsDnsName = types.StringNull()
	}
	for i := range data.Ips {
		keys := [...]string{"ipaddress"}
		keyValues := [...]string{data.Ips[i].Ipaddress.ValueString()}

		var r gjson.Result
		res.Get("NetworkDevice.NetworkDeviceIPList").ForEach(
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
		if value := r.Get("ipaddress"); value.Exists() && !data.Ips[i].Ipaddress.IsNull() {
			data.Ips[i].Ipaddress = types.StringValue(value.String())
		} else {
			data.Ips[i].Ipaddress = types.StringNull()
		}
		if value := r.Get("ipaddressExclude"); value.Exists() && !data.Ips[i].IpaddressExclude.IsNull() {
			data.Ips[i].IpaddressExclude = types.StringValue(value.String())
		} else {
			data.Ips[i].IpaddressExclude = types.StringNull()
		}
		if value := r.Get("mask"); value.Exists() && !data.Ips[i].Mask.IsNull() {
			data.Ips[i].Mask = types.StringValue(value.String())
		} else {
			data.Ips[i].Mask = types.StringNull()
		}
	}
	if value := res.Get("NetworkDevice.NetworkDeviceGroupList"); value.Exists() && !data.NetworkDeviceGroups.IsNull() {
		data.NetworkDeviceGroups = helpers.GetStringList(value.Array())
	} else {
		data.NetworkDeviceGroups = types.ListNull(types.StringType)
	}
	if value := res.Get("NetworkDevice.modelName"); value.Exists() && !data.ModelName.IsNull() {
		data.ModelName = types.StringValue(value.String())
	} else {
		data.ModelName = types.StringNull()
	}
	if value := res.Get("NetworkDevice.softwareVersion"); value.Exists() && !data.SoftwareVersion.IsNull() {
		data.SoftwareVersion = types.StringValue(value.String())
	} else {
		data.SoftwareVersion = types.StringNull()
	}
	if value := res.Get("NetworkDevice.profileName"); value.Exists() && !data.ProfileName.IsNull() {
		data.ProfileName = types.StringValue(value.String())
	} else {
		data.ProfileName = types.StringNull()
	}
	if value := res.Get("NetworkDevice.snmpsettings.linkTrapQuery"); value.Exists() && !data.SnmpLinkTrapQuery.IsNull() {
		data.SnmpLinkTrapQuery = types.BoolValue(value.Bool())
	} else {
		data.SnmpLinkTrapQuery = types.BoolNull()
	}
	if value := res.Get("NetworkDevice.snmpsettings.macTrapQuery"); value.Exists() && !data.SnmpMacTrapQuery.IsNull() {
		data.SnmpMacTrapQuery = types.BoolValue(value.Bool())
	} else {
		data.SnmpMacTrapQuery = types.BoolNull()
	}
	if value := res.Get("NetworkDevice.snmpsettings.originatingPolicyServicesNode"); value.Exists() && !data.SnmpOriginatingPolicyServiceNode.IsNull() {
		data.SnmpOriginatingPolicyServiceNode = types.StringValue(value.String())
	} else {
		data.SnmpOriginatingPolicyServiceNode = types.StringNull()
	}
	if value := res.Get("NetworkDevice.snmpsettings.pollingInterval"); value.Exists() && !data.SnmpPollingInterval.IsNull() {
		data.SnmpPollingInterval = types.Int64Value(value.Int())
	} else {
		data.SnmpPollingInterval = types.Int64Null()
	}
	if value := res.Get("NetworkDevice.snmpsettings.roCommunity"); value.Exists() && !data.SnmpRoCommunity.IsNull() {
		data.SnmpRoCommunity = types.StringValue(value.String())
	} else {
		data.SnmpRoCommunity = types.StringNull()
	}
	if value := res.Get("NetworkDevice.snmpsettings.version"); value.Exists() && !data.SnmpVersion.IsNull() {
		data.SnmpVersion = types.StringValue(value.String())
	} else {
		data.SnmpVersion = types.StringNull()
	}
	if value := res.Get("NetworkDevice.tacacsSettings.connectModeOptions"); value.Exists() && !data.TacacsConnectModeOptions.IsNull() {
		data.TacacsConnectModeOptions = types.StringValue(value.String())
	} else {
		data.TacacsConnectModeOptions = types.StringNull()
	}
	if value := res.Get("NetworkDevice.tacacsSettings.sharedSecret"); value.Exists() && !data.TacacsSharedSecret.IsNull() {
		data.TacacsSharedSecret = types.StringValue(value.String())
	} else {
		data.TacacsSharedSecret = types.StringNull()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.deviceAuthenticationSettings.sgaDeviceId"); value.Exists() && !data.TrustsecDeviceId.IsNull() {
		data.TrustsecDeviceId = types.StringValue(value.String())
	} else {
		data.TrustsecDeviceId = types.StringNull()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.deviceAuthenticationSettings.sgaDevicePassword"); value.Exists() && !data.TrustsecDevicePassword.IsNull() {
		data.TrustsecDevicePassword = types.StringValue(value.String())
	} else {
		data.TrustsecDevicePassword = types.StringNull()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.deviceAuthenticationSettings.restApiUsername"); value.Exists() && !data.TrustsecRestApiUsername.IsNull() {
		data.TrustsecRestApiUsername = types.StringValue(value.String())
	} else {
		data.TrustsecRestApiUsername = types.StringNull()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.deviceConfigurationDeployment.enableModePassword"); value.Exists() && !data.TrustsecEnableModePassword.IsNull() {
		data.TrustsecEnableModePassword = types.StringValue(value.String())
	} else {
		data.TrustsecEnableModePassword = types.StringNull()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.deviceConfigurationDeployment.execModePassword"); value.Exists() && !data.TrustsecExecModePassword.IsNull() {
		data.TrustsecExecModePassword = types.StringValue(value.String())
	} else {
		data.TrustsecExecModePassword = types.StringNull()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.deviceConfigurationDeployment.execModeUsername"); value.Exists() && !data.TrustsecExecModeUsername.IsNull() {
		data.TrustsecExecModeUsername = types.StringValue(value.String())
	} else {
		data.TrustsecExecModeUsername = types.StringNull()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.deviceConfigurationDeployment.includeWhenDeployingSGTUpdates"); value.Exists() && !data.TrustsecIncludeWhenDeployingSgtUpdates.IsNull() {
		data.TrustsecIncludeWhenDeployingSgtUpdates = types.BoolValue(value.Bool())
	} else {
		data.TrustsecIncludeWhenDeployingSgtUpdates = types.BoolNull()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.sgaNotificationAndUpdates.downlaodEnvironmentDataEveryXSeconds"); value.Exists() && !data.TrustsecDownloadEnviromentDataEveryXSeconds.IsNull() {
		data.TrustsecDownloadEnviromentDataEveryXSeconds = types.Int64Value(value.Int())
	} else {
		data.TrustsecDownloadEnviromentDataEveryXSeconds = types.Int64Null()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.sgaNotificationAndUpdates.downlaodPeerAuthorizationPolicyEveryXSeconds"); value.Exists() && !data.TrustsecDownloadPeerAuthorizationPolicyEveryXSeconds.IsNull() {
		data.TrustsecDownloadPeerAuthorizationPolicyEveryXSeconds = types.Int64Value(value.Int())
	} else {
		data.TrustsecDownloadPeerAuthorizationPolicyEveryXSeconds = types.Int64Null()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.sgaNotificationAndUpdates.downloadSGACLListsEveryXSeconds"); value.Exists() && !data.TrustsecDownloadSgaclListsEveryXSeconds.IsNull() {
		data.TrustsecDownloadSgaclListsEveryXSeconds = types.Int64Value(value.Int())
	} else {
		data.TrustsecDownloadSgaclListsEveryXSeconds = types.Int64Null()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.sgaNotificationAndUpdates.otherSGADevicesToTrustThisDevice"); value.Exists() && !data.TrustsecOtherSgaDevicesToTrustThisDevice.IsNull() {
		data.TrustsecOtherSgaDevicesToTrustThisDevice = types.BoolValue(value.Bool())
	} else {
		data.TrustsecOtherSgaDevicesToTrustThisDevice = types.BoolNull()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.sgaNotificationAndUpdates.reAuthenticationEveryXSeconds"); value.Exists() && !data.TrustsecReAuthenticationEveryXSeconds.IsNull() {
		data.TrustsecReAuthenticationEveryXSeconds = types.Int64Value(value.Int())
	} else {
		data.TrustsecReAuthenticationEveryXSeconds = types.Int64Null()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.sgaNotificationAndUpdates.sendConfigurationToDevice"); value.Exists() && !data.TrustsecSendConfigurationToDevice.IsNull() {
		data.TrustsecSendConfigurationToDevice = types.BoolValue(value.Bool())
	} else {
		data.TrustsecSendConfigurationToDevice = types.BoolNull()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.sgaNotificationAndUpdates.sendConfigurationToDeviceUsing"); value.Exists() && !data.TrustsecSendConfigurationToDeviceUsing.IsNull() {
		data.TrustsecSendConfigurationToDeviceUsing = types.StringValue(value.String())
	} else {
		data.TrustsecSendConfigurationToDeviceUsing = types.StringNull()
	}
	if value := res.Get("NetworkDevice.trustsecsettings.sgaNotificationAndUpdates.coaSourceHost"); value.Exists() && !data.TrustsecCoaSourceHost.IsNull() {
		data.TrustsecCoaSourceHost = types.StringValue(value.String())
	} else {
		data.TrustsecCoaSourceHost = types.StringNull()
	}
}