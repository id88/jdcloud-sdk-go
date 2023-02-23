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

type BandwidthTrendRequest struct {

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
    Direction *string `json:"direction"`
}

/*
 * param zone_identifier:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewBandwidthTrendRequest(
    zone_identifier string,
) *BandwidthTrendRequest {

	return &BandwidthTrendRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/analytics$$bandwidthTrend",
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
 * param direction:  (Optional)
 */
func NewBandwidthTrendRequestWithAllParams(
    zone_identifier string,
    zoneName *string,
    since *string,
    until *string,
    direction *string,
) *BandwidthTrendRequest {

    return &BandwidthTrendRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/analytics$$bandwidthTrend",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Zone_identifier: zone_identifier,
        ZoneName: zoneName,
        Since: since,
        Until: until,
        Direction: direction,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewBandwidthTrendRequestWithoutParam() *BandwidthTrendRequest {

    return &BandwidthTrendRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/analytics$$bandwidthTrend",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param zone_identifier: (Required) */
func (r *BandwidthTrendRequest) SetZone_identifier(zone_identifier string) {
    r.Zone_identifier = zone_identifier
}
/* param zoneName: (Optional) */
func (r *BandwidthTrendRequest) SetZoneName(zoneName string) {
    r.ZoneName = &zoneName
}
/* param since: (Optional) */
func (r *BandwidthTrendRequest) SetSince(since string) {
    r.Since = &since
}
/* param until: (Optional) */
func (r *BandwidthTrendRequest) SetUntil(until string) {
    r.Until = &until
}
/* param direction: (Optional) */
func (r *BandwidthTrendRequest) SetDirection(direction string) {
    r.Direction = &direction
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r BandwidthTrendRequest) GetRegionId() string {
    return ""
}

type BandwidthTrendResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result BandwidthTrendResult `json:"result"`
}

type BandwidthTrendResult struct {
    Data starshield.SimpleDateHistogram `json:"data"`
}