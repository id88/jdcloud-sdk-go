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
    nativecontainer "github.com/jdcloud-api/jdcloud-sdk-go/services/nativecontainer/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeInstanceTypesRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* instanceTypes - 实例规格，精确匹配，支持多个
az - 可用区，精确匹配，支持多个
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeInstanceTypesRequest(
    regionId string,
) *DescribeInstanceTypesRequest {

	return &DescribeInstanceTypesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instanceTypes",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param filters: instanceTypes - 实例规格，精确匹配，支持多个
az - 可用区，精确匹配，支持多个
 (Optional)
 */
func NewDescribeInstanceTypesRequestWithAllParams(
    regionId string,
    filters []common.Filter,
) *DescribeInstanceTypesRequest {

    return &DescribeInstanceTypesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instanceTypes",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeInstanceTypesRequestWithoutParam() *DescribeInstanceTypesRequest {

    return &DescribeInstanceTypesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instanceTypes",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribeInstanceTypesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param filters: instanceTypes - 实例规格，精确匹配，支持多个
az - 可用区，精确匹配，支持多个
(Optional) */
func (r *DescribeInstanceTypesRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeInstanceTypesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeInstanceTypesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeInstanceTypesResult `json:"result"`
}

type DescribeInstanceTypesResult struct {
    InstanceTypes []nativecontainer.InstanceType `json:"instanceTypes"`
    SpecificInstanceTypes []nativecontainer.InstanceType `json:"specificInstanceTypes"`
    TotalCount int `json:"totalCount"`
}