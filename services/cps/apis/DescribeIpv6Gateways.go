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
    cps "github.com/jdcloud-api/jdcloud-sdk-go/services/cps/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeIpv6GatewaysRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeRegions）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为20；取值范围[20, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* IPv6网关名称 (Optional) */
    Ipv6GatewayName *string `json:"ipv6GatewayName"`

    /* ipv6GatewayId - IPv6网关实例ID，精确匹配，支持多个
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域ID，可调用接口（describeRegions）获取云物理服务器支持的地域 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeIpv6GatewaysRequest(
    regionId string,
) *DescribeIpv6GatewaysRequest {

	return &DescribeIpv6GatewaysRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/ipv6Gateways",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeRegions）获取云物理服务器支持的地域 (Required)
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为20；取值范围[20, 100] (Optional)
 * param ipv6GatewayName: IPv6网关名称 (Optional)
 * param filters: ipv6GatewayId - IPv6网关实例ID，精确匹配，支持多个
 (Optional)
 */
func NewDescribeIpv6GatewaysRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    ipv6GatewayName *string,
    filters []common.Filter,
) *DescribeIpv6GatewaysRequest {

    return &DescribeIpv6GatewaysRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ipv6Gateways",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Ipv6GatewayName: ipv6GatewayName,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeIpv6GatewaysRequestWithoutParam() *DescribeIpv6GatewaysRequest {

    return &DescribeIpv6GatewaysRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ipv6Gateways",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeRegions）获取云物理服务器支持的地域(Required) */
func (r *DescribeIpv6GatewaysRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribeIpv6GatewaysRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为20；取值范围[20, 100](Optional) */
func (r *DescribeIpv6GatewaysRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param ipv6GatewayName: IPv6网关名称(Optional) */
func (r *DescribeIpv6GatewaysRequest) SetIpv6GatewayName(ipv6GatewayName string) {
    r.Ipv6GatewayName = &ipv6GatewayName
}

/* param filters: ipv6GatewayId - IPv6网关实例ID，精确匹配，支持多个
(Optional) */
func (r *DescribeIpv6GatewaysRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeIpv6GatewaysRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeIpv6GatewaysResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeIpv6GatewaysResult `json:"result"`
}

type DescribeIpv6GatewaysResult struct {
    Ipv6Gateways []cps.Ipv6Gateway `json:"ipv6Gateways"`
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TotalCount int `json:"totalCount"`
}