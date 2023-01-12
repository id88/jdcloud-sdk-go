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

type CopyPolicyRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 策略ID  */
    PolicyId string `json:"policyId"`

    /* 目标地域  */
    TargetRegion string `json:"targetRegion"`
}

/*
 * param regionId: 地域ID (Required)
 * param policyId: 策略ID (Required)
 * param targetRegion: 目标地域 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCopyPolicyRequest(
    regionId string,
    policyId string,
    targetRegion string,
) *CopyPolicyRequest {

	return &CopyPolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/policy/{policyId}:copy",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        PolicyId: policyId,
        TargetRegion: targetRegion,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param policyId: 策略ID (Required)
 * param targetRegion: 目标地域 (Required)
 */
func NewCopyPolicyRequestWithAllParams(
    regionId string,
    policyId string,
    targetRegion string,
) *CopyPolicyRequest {

    return &CopyPolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/policy/{policyId}:copy",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PolicyId: policyId,
        TargetRegion: targetRegion,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCopyPolicyRequestWithoutParam() *CopyPolicyRequest {

    return &CopyPolicyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/policy/{policyId}:copy",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CopyPolicyRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param policyId: 策略ID(Required) */
func (r *CopyPolicyRequest) SetPolicyId(policyId string) {
    r.PolicyId = policyId
}
/* param targetRegion: 目标地域(Required) */
func (r *CopyPolicyRequest) SetTargetRegion(targetRegion string) {
    r.TargetRegion = targetRegion
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CopyPolicyRequest) GetRegionId() string {
    return r.RegionId
}

type CopyPolicyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CopyPolicyResult `json:"result"`
}

type CopyPolicyResult struct {
    PolicyId string `json:"policyId"`
}