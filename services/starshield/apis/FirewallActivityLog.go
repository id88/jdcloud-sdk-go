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

type FirewallActivityLogRequest struct {

    core.JDCloudRequest

    /*   */
    Zone_identifier string `json:"zone_identifier"`

    /*  (Optional) */
    ZoneName *string `json:"zoneName"`

    /*  (Optional) */
    Since *string `json:"since"`

    /*  (Optional) */
    Until *string `json:"until"`

    /*  (Optional) */
    PageNumber *int `json:"pageNumber"`

    /*  (Optional) */
    PageSize *int `json:"pageSize"`

    /*  (Optional) */
    Filters []starshield.AnalyticsReportingFilter `json:"filters"`
}

/*
 * param zone_identifier:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewFirewallActivityLogRequest(
    zone_identifier string,
) *FirewallActivityLogRequest {

	return &FirewallActivityLogRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/analytics$$firewallActivityLog",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Zone_identifier: zone_identifier,
	}
}

/*
 * param zone_identifier:  (Required)
 * param zoneName:  (Optional)
 * param since:  (Optional)
 * param until:  (Optional)
 * param pageNumber:  (Optional)
 * param pageSize:  (Optional)
 * param filters:  (Optional)
 */
func NewFirewallActivityLogRequestWithAllParams(
    zone_identifier string,
    zoneName *string,
    since *string,
    until *string,
    pageNumber *int,
    pageSize *int,
    filters []starshield.AnalyticsReportingFilter,
) *FirewallActivityLogRequest {

    return &FirewallActivityLogRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/analytics$$firewallActivityLog",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Zone_identifier: zone_identifier,
        ZoneName: zoneName,
        Since: since,
        Until: until,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewFirewallActivityLogRequestWithoutParam() *FirewallActivityLogRequest {

    return &FirewallActivityLogRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/analytics$$firewallActivityLog",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param zone_identifier: (Required) */
func (r *FirewallActivityLogRequest) SetZone_identifier(zone_identifier string) {
    r.Zone_identifier = zone_identifier
}
/* param zoneName: (Optional) */
func (r *FirewallActivityLogRequest) SetZoneName(zoneName string) {
    r.ZoneName = &zoneName
}
/* param since: (Optional) */
func (r *FirewallActivityLogRequest) SetSince(since string) {
    r.Since = &since
}
/* param until: (Optional) */
func (r *FirewallActivityLogRequest) SetUntil(until string) {
    r.Until = &until
}
/* param pageNumber: (Optional) */
func (r *FirewallActivityLogRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}
/* param pageSize: (Optional) */
func (r *FirewallActivityLogRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}
/* param filters: (Optional) */
func (r *FirewallActivityLogRequest) SetFilters(filters []starshield.AnalyticsReportingFilter) {
    r.Filters = filters
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r FirewallActivityLogRequest) GetRegionId() string {
    return ""
}

type FirewallActivityLogResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result FirewallActivityLogResult `json:"result"`
}

type FirewallActivityLogResult struct {
    Total int `json:"total"`
    ActivityLogs []starshield.ActivityLog `json:"activityLogs"`
}