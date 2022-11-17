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
    iotlink "github.com/jdcloud-api/jdcloud-sdk-go/services/iotlink/models"
)

type CloseIotFlowRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 物联网卡号码列表(单次提交最多不超过200个号码)  */
    Iccids []string `json:"iccids"`
}

/*
 * param regionId: Region ID (Required)
 * param iccids: 物联网卡号码列表(单次提交最多不超过200个号码) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCloseIotFlowRequest(
    regionId string,
    iccids []string,
) *CloseIotFlowRequest {

	return &CloseIotFlowRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/closeIotFlow",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Iccids: iccids,
	}
}

/*
 * param regionId: Region ID (Required)
 * param iccids: 物联网卡号码列表(单次提交最多不超过200个号码) (Required)
 */
func NewCloseIotFlowRequestWithAllParams(
    regionId string,
    iccids []string,
) *CloseIotFlowRequest {

    return &CloseIotFlowRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/closeIotFlow",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Iccids: iccids,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCloseIotFlowRequestWithoutParam() *CloseIotFlowRequest {

    return &CloseIotFlowRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/closeIotFlow",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CloseIotFlowRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param iccids: 物联网卡号码列表(单次提交最多不超过200个号码)(Required) */
func (r *CloseIotFlowRequest) SetIccids(iccids []string) {
    r.Iccids = iccids
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CloseIotFlowRequest) GetRegionId() string {
    return r.RegionId
}

type CloseIotFlowResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CloseIotFlowResult `json:"result"`
}

type CloseIotFlowResult struct {
    Status string `json:"status"`
    Message string `json:"message"`
    Result []iotlink.OperationIotlinkResp `json:"result"`
}