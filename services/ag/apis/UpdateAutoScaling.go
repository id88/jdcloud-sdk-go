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
    ag "github.com/jdcloud-api/jdcloud-sdk-go/services/ag/models"
)

type UpdateAutoScalingRequest struct {

    core.JDCloudRequest

    /* 地域  */
    RegionId string `json:"regionId"`

    /* 高可用组 ID  */
    AgId string `json:"agId"`

    /* 伸缩组详细信息  */
    AutoscalingSpec *ag.AutoscalingSpecByUpdate `json:"autoscalingSpec"`
}

/*
 * param regionId: 地域 (Required)
 * param agId: 高可用组 ID (Required)
 * param autoscalingSpec: 伸缩组详细信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateAutoScalingRequest(
    regionId string,
    agId string,
    autoscalingSpec *ag.AutoscalingSpecByUpdate,
) *UpdateAutoScalingRequest {

	return &UpdateAutoScalingRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/autoScaling/{agId}:updateAutoScaling",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        AgId: agId,
        AutoscalingSpec: autoscalingSpec,
	}
}

/*
 * param regionId: 地域 (Required)
 * param agId: 高可用组 ID (Required)
 * param autoscalingSpec: 伸缩组详细信息 (Required)
 */
func NewUpdateAutoScalingRequestWithAllParams(
    regionId string,
    agId string,
    autoscalingSpec *ag.AutoscalingSpecByUpdate,
) *UpdateAutoScalingRequest {

    return &UpdateAutoScalingRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/autoScaling/{agId}:updateAutoScaling",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AgId: agId,
        AutoscalingSpec: autoscalingSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateAutoScalingRequestWithoutParam() *UpdateAutoScalingRequest {

    return &UpdateAutoScalingRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/autoScaling/{agId}:updateAutoScaling",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域(Required) */
func (r *UpdateAutoScalingRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param agId: 高可用组 ID(Required) */
func (r *UpdateAutoScalingRequest) SetAgId(agId string) {
    r.AgId = agId
}
/* param autoscalingSpec: 伸缩组详细信息(Required) */
func (r *UpdateAutoScalingRequest) SetAutoscalingSpec(autoscalingSpec *ag.AutoscalingSpecByUpdate) {
    r.AutoscalingSpec = autoscalingSpec
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateAutoScalingRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateAutoScalingResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateAutoScalingResult `json:"result"`
}

type UpdateAutoScalingResult struct {
}