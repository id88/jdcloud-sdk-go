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

type ZoneTrafficDateHistogramRequest struct {

    core.JDCloudRequest

    /* 域名标识  */
    Zone_identifier string `json:"zone_identifier"`

    /* all - 所有
normal - 业务
mitigation - 缓解
cache - 缓存
origin - 回源
  */
    QueryMode string `json:"queryMode"`

    /* 域名  */
    ZoneName string `json:"zoneName"`

    /* 开始时间  */
    Since string `json:"since"`

    /* 结束时间  */
    Until string `json:"until"`
}

/*
 * param zone_identifier: 域名标识 (Required)
 * param queryMode: all - 所有
normal - 业务
mitigation - 缓解
cache - 缓存
origin - 回源
 (Required)
 * param zoneName: 域名 (Required)
 * param since: 开始时间 (Required)
 * param until: 结束时间 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewZoneTrafficDateHistogramRequest(
    zone_identifier string,
    queryMode string,
    zoneName string,
    since string,
    until string,
) *ZoneTrafficDateHistogramRequest {

	return &ZoneTrafficDateHistogramRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/zoneTrafficDateHistogram",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Zone_identifier: zone_identifier,
        QueryMode: queryMode,
        ZoneName: zoneName,
        Since: since,
        Until: until,
	}
}

/*
 * param zone_identifier: 域名标识 (Required)
 * param queryMode: all - 所有
normal - 业务
mitigation - 缓解
cache - 缓存
origin - 回源
 (Required)
 * param zoneName: 域名 (Required)
 * param since: 开始时间 (Required)
 * param until: 结束时间 (Required)
 */
func NewZoneTrafficDateHistogramRequestWithAllParams(
    zone_identifier string,
    queryMode string,
    zoneName string,
    since string,
    until string,
) *ZoneTrafficDateHistogramRequest {

    return &ZoneTrafficDateHistogramRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/zoneTrafficDateHistogram",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Zone_identifier: zone_identifier,
        QueryMode: queryMode,
        ZoneName: zoneName,
        Since: since,
        Until: until,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewZoneTrafficDateHistogramRequestWithoutParam() *ZoneTrafficDateHistogramRequest {

    return &ZoneTrafficDateHistogramRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/zoneTrafficDateHistogram",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param zone_identifier: 域名标识(Required) */
func (r *ZoneTrafficDateHistogramRequest) SetZone_identifier(zone_identifier string) {
    r.Zone_identifier = zone_identifier
}
/* param queryMode: all - 所有
normal - 业务
mitigation - 缓解
cache - 缓存
origin - 回源
(Required) */
func (r *ZoneTrafficDateHistogramRequest) SetQueryMode(queryMode string) {
    r.QueryMode = queryMode
}
/* param zoneName: 域名(Required) */
func (r *ZoneTrafficDateHistogramRequest) SetZoneName(zoneName string) {
    r.ZoneName = zoneName
}
/* param since: 开始时间(Required) */
func (r *ZoneTrafficDateHistogramRequest) SetSince(since string) {
    r.Since = since
}
/* param until: 结束时间(Required) */
func (r *ZoneTrafficDateHistogramRequest) SetUntil(until string) {
    r.Until = until
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ZoneTrafficDateHistogramRequest) GetRegionId() string {
    return ""
}

type ZoneTrafficDateHistogramResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ZoneTrafficDateHistogramResult `json:"result"`
}

type ZoneTrafficDateHistogramResult struct {
    Data starshield.SimpleDateHistogram `json:"data"`
}