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
    jdfusion "github.com/jdcloud-api/jdcloud-sdk-go/services/jdfusion/models"
)

type RegistCloudInfoRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /*   */
    Cloud *jdfusion.CloudInfo `json:"cloud"`
}

/*
 * param regionId: 地域ID (Required)
 * param cloud:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewRegistCloudInfoRequest(
    regionId string,
    cloud *jdfusion.CloudInfo,
) *RegistCloudInfoRequest {

	return &RegistCloudInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/cloud_info",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Cloud: cloud,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param cloud:  (Required)
 */
func NewRegistCloudInfoRequestWithAllParams(
    regionId string,
    cloud *jdfusion.CloudInfo,
) *RegistCloudInfoRequest {

    return &RegistCloudInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cloud_info",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Cloud: cloud,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewRegistCloudInfoRequestWithoutParam() *RegistCloudInfoRequest {

    return &RegistCloudInfoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cloud_info",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *RegistCloudInfoRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param cloud: (Required) */
func (r *RegistCloudInfoRequest) SetCloud(cloud *jdfusion.CloudInfo) {
    r.Cloud = cloud
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RegistCloudInfoRequest) GetRegionId() string {
    return r.RegionId
}

type RegistCloudInfoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RegistCloudInfoResult `json:"result"`
}

type RegistCloudInfoResult struct {
    Cloud jdfusion.CloudInfo `json:"cloud"`
}