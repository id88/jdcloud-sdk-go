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
    cdn "github.com/jdcloud-api/jdcloud-sdk-go/services/cdn/models"
)

type QueryStreamInfoRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    EndTime *string `json:"endTime"`

    /* app名 (Optional) */
    AppName *string `json:"appName"`

    /* 流名 (Optional) */
    StreamName *string `json:"streamName"`

    /* 页码，不传默认1 (Optional) */
    PageNum *int `json:"pageNum"`

    /* 页size,不传默认100,最大值1000 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryStreamInfoRequest(
    domain string,
) *QueryStreamInfoRequest {

	return &QueryStreamInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain/{domain}/liveStatistics:streamInfo",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 * param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 * param appName: app名 (Optional)
 * param streamName: 流名 (Optional)
 * param pageNum: 页码，不传默认1 (Optional)
 * param pageSize: 页size,不传默认100,最大值1000 (Optional)
 */
func NewQueryStreamInfoRequestWithAllParams(
    domain string,
    startTime *string,
    endTime *string,
    appName *string,
    streamName *string,
    pageNum *int,
    pageSize *int,
) *QueryStreamInfoRequest {

    return &QueryStreamInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/liveStatistics:streamInfo",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        StartTime: startTime,
        EndTime: endTime,
        AppName: appName,
        StreamName: streamName,
        PageNum: pageNum,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryStreamInfoRequestWithoutParam() *QueryStreamInfoRequest {

    return &QueryStreamInfoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/liveStatistics:streamInfo",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *QueryStreamInfoRequest) SetDomain(domain string) {
    r.Domain = domain
}
/* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryStreamInfoRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}
/* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryStreamInfoRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}
/* param appName: app名(Optional) */
func (r *QueryStreamInfoRequest) SetAppName(appName string) {
    r.AppName = &appName
}
/* param streamName: 流名(Optional) */
func (r *QueryStreamInfoRequest) SetStreamName(streamName string) {
    r.StreamName = &streamName
}
/* param pageNum: 页码，不传默认1(Optional) */
func (r *QueryStreamInfoRequest) SetPageNum(pageNum int) {
    r.PageNum = &pageNum
}
/* param pageSize: 页size,不传默认100,最大值1000(Optional) */
func (r *QueryStreamInfoRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryStreamInfoRequest) GetRegionId() string {
    return ""
}

type QueryStreamInfoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryStreamInfoResult `json:"result"`
}

type QueryStreamInfoResult struct {
    Total int `json:"total"`
    PageNum int `json:"pageNum"`
    PageSize int `json:"pageSize"`
    StreamInfoList []cdn.StatisticsLiveStreamInfo `json:"streamInfoList"`
}