// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
)

type OpenFirewallBotRequest struct {

    core.JDCloudRequest

    /*   */
    ZoneId string `json:"zoneId"`
}

/*
 * param zoneId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewOpenFirewallBotRequest(
    zoneId string,
) *OpenFirewallBotRequest {

	return &OpenFirewallBotRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/bot/{zoneId}",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        ZoneId: zoneId,
	}
}

/*
 * param zoneId:  (Required)
 */
func NewOpenFirewallBotRequestWithAllParams(
    zoneId string,
) *OpenFirewallBotRequest {

    return &OpenFirewallBotRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/bot/{zoneId}",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        ZoneId: zoneId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewOpenFirewallBotRequestWithoutParam() *OpenFirewallBotRequest {

    return &OpenFirewallBotRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/bot/{zoneId}",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param zoneId: (Required) */
func (r *OpenFirewallBotRequest) SetZoneId(zoneId string) {
    r.ZoneId = zoneId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r OpenFirewallBotRequest) GetRegionId() string {
    return ""
}

type OpenFirewallBotResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result OpenFirewallBotResult `json:"result"`
}

type OpenFirewallBotResult struct {
    Success bool `json:"success"`
}