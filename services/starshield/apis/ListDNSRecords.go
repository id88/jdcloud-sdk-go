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

type ListDNSRecordsRequest struct {

    core.JDCloudRequest

    /*   */
    Zone_identifier string `json:"zone_identifier"`

    /* 是否匹配所有搜索要求或至少一个（任何） (Optional) */
    Match *string `json:"match"`

    /* DNS record name (Optional) */
    Name *string `json:"name"`

    /* 用于排序的字段 (Optional) */
    Order *string `json:"order"`

    /* 分页结果的页码 (Optional) */
    Page *int `json:"page"`

    /* 每页的DNS记录数 (Optional) */
    Per_page *int `json:"per_page"`

    /* DNS记录内容 (Optional) */
    Content *string `json:"content"`

    /* DNS记录类型 (Optional) */
    Type *string `json:"type"`

    /* DNS记录代理状态 (Optional) */
    Proxied *bool `json:"proxied"`

    /* asc - 升序；desc - 降序 (Optional) */
    Direction *string `json:"direction"`
}

/*
 * param zone_identifier:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewListDNSRecordsRequest(
    zone_identifier string,
) *ListDNSRecordsRequest {

	return &ListDNSRecordsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/dns_records",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Zone_identifier: zone_identifier,
	}
}

/*
 * param zone_identifier:  (Required)
 * param match: 是否匹配所有搜索要求或至少一个（任何） (Optional)
 * param name: DNS record name (Optional)
 * param order: 用于排序的字段 (Optional)
 * param page: 分页结果的页码 (Optional)
 * param per_page: 每页的DNS记录数 (Optional)
 * param content: DNS记录内容 (Optional)
 * param type_: DNS记录类型 (Optional)
 * param proxied: DNS记录代理状态 (Optional)
 * param direction: asc - 升序；desc - 降序 (Optional)
 */
func NewListDNSRecordsRequestWithAllParams(
    zone_identifier string,
    match *string,
    name *string,
    order *string,
    page *int,
    per_page *int,
    content *string,
    type_ *string,
    proxied *bool,
    direction *string,
) *ListDNSRecordsRequest {

    return &ListDNSRecordsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/dns_records",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Zone_identifier: zone_identifier,
        Match: match,
        Name: name,
        Order: order,
        Page: page,
        Per_page: per_page,
        Content: content,
        Type: type_,
        Proxied: proxied,
        Direction: direction,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewListDNSRecordsRequestWithoutParam() *ListDNSRecordsRequest {

    return &ListDNSRecordsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/dns_records",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param zone_identifier: (Required) */
func (r *ListDNSRecordsRequest) SetZone_identifier(zone_identifier string) {
    r.Zone_identifier = zone_identifier
}
/* param match: 是否匹配所有搜索要求或至少一个（任何）(Optional) */
func (r *ListDNSRecordsRequest) SetMatch(match string) {
    r.Match = &match
}
/* param name: DNS record name(Optional) */
func (r *ListDNSRecordsRequest) SetName(name string) {
    r.Name = &name
}
/* param order: 用于排序的字段(Optional) */
func (r *ListDNSRecordsRequest) SetOrder(order string) {
    r.Order = &order
}
/* param page: 分页结果的页码(Optional) */
func (r *ListDNSRecordsRequest) SetPage(page int) {
    r.Page = &page
}
/* param per_page: 每页的DNS记录数(Optional) */
func (r *ListDNSRecordsRequest) SetPer_page(per_page int) {
    r.Per_page = &per_page
}
/* param content: DNS记录内容(Optional) */
func (r *ListDNSRecordsRequest) SetContent(content string) {
    r.Content = &content
}
/* param type_: DNS记录类型(Optional) */
func (r *ListDNSRecordsRequest) SetType(type_ string) {
    r.Type = &type_
}
/* param proxied: DNS记录代理状态(Optional) */
func (r *ListDNSRecordsRequest) SetProxied(proxied bool) {
    r.Proxied = &proxied
}
/* param direction: asc - 升序；desc - 降序(Optional) */
func (r *ListDNSRecordsRequest) SetDirection(direction string) {
    r.Direction = &direction
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ListDNSRecordsRequest) GetRegionId() string {
    return ""
}

type ListDNSRecordsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ListDNSRecordsResult `json:"result"`
}

type ListDNSRecordsResult struct {
    DataList []starshield.DnsRecord `json:"dataList"`
    CurrentCount int `json:"currentCount"`
    TotalCount int `json:"totalCount"`
    TotalPage int `json:"totalPage"`
}