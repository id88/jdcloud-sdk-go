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
    starshield "github.com/jdcloud-api/jdcloud-sdk-go/services/starshield/models"
)

type ListSSLConfigurationsRequest struct {

    core.JDCloudRequest

    /*   */
    Zone_identifier string `json:"zone_identifier"`

    /* 域的自定义SSL的状态 (Optional) */
    Status *string `json:"status"`

    /* 分页结果的页码 (Optional) */
    Page *int `json:"page"`

    /* 每页的域数 (Optional) */
    Per_page *int `json:"per_page"`

    /* 是否匹配所有搜索要求或至少一个（任何） (Optional) */
    Match *string `json:"match"`
}

/*
 * param zone_identifier:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewListSSLConfigurationsRequest(
    zone_identifier string,
) *ListSSLConfigurationsRequest {

	return &ListSSLConfigurationsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/custom_certificates",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Zone_identifier: zone_identifier,
	}
}

/*
 * param zone_identifier:  (Required)
 * param status: 域的自定义SSL的状态 (Optional)
 * param page: 分页结果的页码 (Optional)
 * param per_page: 每页的域数 (Optional)
 * param match: 是否匹配所有搜索要求或至少一个（任何） (Optional)
 */
func NewListSSLConfigurationsRequestWithAllParams(
    zone_identifier string,
    status *string,
    page *int,
    per_page *int,
    match *string,
) *ListSSLConfigurationsRequest {

    return &ListSSLConfigurationsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/custom_certificates",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Zone_identifier: zone_identifier,
        Status: status,
        Page: page,
        Per_page: per_page,
        Match: match,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewListSSLConfigurationsRequestWithoutParam() *ListSSLConfigurationsRequest {

    return &ListSSLConfigurationsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/custom_certificates",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param zone_identifier: (Required) */
func (r *ListSSLConfigurationsRequest) SetZone_identifier(zone_identifier string) {
    r.Zone_identifier = zone_identifier
}
/* param status: 域的自定义SSL的状态(Optional) */
func (r *ListSSLConfigurationsRequest) SetStatus(status string) {
    r.Status = &status
}
/* param page: 分页结果的页码(Optional) */
func (r *ListSSLConfigurationsRequest) SetPage(page int) {
    r.Page = &page
}
/* param per_page: 每页的域数(Optional) */
func (r *ListSSLConfigurationsRequest) SetPer_page(per_page int) {
    r.Per_page = &per_page
}
/* param match: 是否匹配所有搜索要求或至少一个（任何）(Optional) */
func (r *ListSSLConfigurationsRequest) SetMatch(match string) {
    r.Match = &match
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ListSSLConfigurationsRequest) GetRegionId() string {
    return ""
}

type ListSSLConfigurationsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ListSSLConfigurationsResult `json:"result"`
}

type ListSSLConfigurationsResult struct {
    DataList []starshield.CustomSSL `json:"dataList"`
}