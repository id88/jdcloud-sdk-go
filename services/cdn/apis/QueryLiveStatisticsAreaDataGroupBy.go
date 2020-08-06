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

type QueryLiveStatisticsAreaDataGroupByRequest struct {

    core.JDCloudRequest

    /* 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    EndTime *string `json:"endTime"`

    /* 需要查询的域名, 必须为用户pin下有权限的域名 (Optional) */
    Domain *string `json:"domain"`

    /* app名 (Optional) */
    AppName *string `json:"appName"`

    /* 需要查询的字段 (Optional) */
    Fields *string `json:"fields"`

    /*  (Optional) */
    Area *string `json:"area"`

    /*  (Optional) */
    Isp *string `json:"isp"`

    /*  (Optional) */
    StreamName *string `json:"streamName"`

    /* 时间粒度，可选值:[oneMin,fiveMin,followTime],followTime只会返回一个汇总后的数据 (Optional) */
    Period *string `json:"period"`

    /* 分组依据 (Optional) */
    GroupBy *string `json:"groupBy"`

    /*  (Optional) */
    SubDomain *string `json:"subDomain"`

    /* 查询的流协议 (Optional) */
    Scheme *string `json:"scheme"`

    /*  (Optional) */
    ReqMethod *string `json:"reqMethod"`

    /* cacheLevel (Optional) */
    CacheLevel *string `json:"cacheLevel"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryLiveStatisticsAreaDataGroupByRequest(
) *QueryLiveStatisticsAreaDataGroupByRequest {

	return &QueryLiveStatisticsAreaDataGroupByRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/liveStatistics:groupByArea",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 * param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 * param domain: 需要查询的域名, 必须为用户pin下有权限的域名 (Optional)
 * param appName: app名 (Optional)
 * param fields: 需要查询的字段 (Optional)
 * param area:  (Optional)
 * param isp:  (Optional)
 * param streamName:  (Optional)
 * param period: 时间粒度，可选值:[oneMin,fiveMin,followTime],followTime只会返回一个汇总后的数据 (Optional)
 * param groupBy: 分组依据 (Optional)
 * param subDomain:  (Optional)
 * param scheme: 查询的流协议 (Optional)
 * param reqMethod:  (Optional)
 * param cacheLevel: cacheLevel (Optional)
 */
func NewQueryLiveStatisticsAreaDataGroupByRequestWithAllParams(
    startTime *string,
    endTime *string,
    domain *string,
    appName *string,
    fields *string,
    area *string,
    isp *string,
    streamName *string,
    period *string,
    groupBy *string,
    subDomain *string,
    scheme *string,
    reqMethod *string,
    cacheLevel *string,
) *QueryLiveStatisticsAreaDataGroupByRequest {

    return &QueryLiveStatisticsAreaDataGroupByRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveStatistics:groupByArea",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        StartTime: startTime,
        EndTime: endTime,
        Domain: domain,
        AppName: appName,
        Fields: fields,
        Area: area,
        Isp: isp,
        StreamName: streamName,
        Period: period,
        GroupBy: groupBy,
        SubDomain: subDomain,
        Scheme: scheme,
        ReqMethod: reqMethod,
        CacheLevel: cacheLevel,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryLiveStatisticsAreaDataGroupByRequestWithoutParam() *QueryLiveStatisticsAreaDataGroupByRequest {

    return &QueryLiveStatisticsAreaDataGroupByRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveStatistics:groupByArea",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryLiveStatisticsAreaDataGroupByRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryLiveStatisticsAreaDataGroupByRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param domain: 需要查询的域名, 必须为用户pin下有权限的域名(Optional) */
func (r *QueryLiveStatisticsAreaDataGroupByRequest) SetDomain(domain string) {
    r.Domain = &domain
}

/* param appName: app名(Optional) */
func (r *QueryLiveStatisticsAreaDataGroupByRequest) SetAppName(appName string) {
    r.AppName = &appName
}

/* param fields: 需要查询的字段(Optional) */
func (r *QueryLiveStatisticsAreaDataGroupByRequest) SetFields(fields string) {
    r.Fields = &fields
}

/* param area: (Optional) */
func (r *QueryLiveStatisticsAreaDataGroupByRequest) SetArea(area string) {
    r.Area = &area
}

/* param isp: (Optional) */
func (r *QueryLiveStatisticsAreaDataGroupByRequest) SetIsp(isp string) {
    r.Isp = &isp
}

/* param streamName: (Optional) */
func (r *QueryLiveStatisticsAreaDataGroupByRequest) SetStreamName(streamName string) {
    r.StreamName = &streamName
}

/* param period: 时间粒度，可选值:[oneMin,fiveMin,followTime],followTime只会返回一个汇总后的数据(Optional) */
func (r *QueryLiveStatisticsAreaDataGroupByRequest) SetPeriod(period string) {
    r.Period = &period
}

/* param groupBy: 分组依据(Optional) */
func (r *QueryLiveStatisticsAreaDataGroupByRequest) SetGroupBy(groupBy string) {
    r.GroupBy = &groupBy
}

/* param subDomain: (Optional) */
func (r *QueryLiveStatisticsAreaDataGroupByRequest) SetSubDomain(subDomain string) {
    r.SubDomain = &subDomain
}

/* param scheme: 查询的流协议(Optional) */
func (r *QueryLiveStatisticsAreaDataGroupByRequest) SetScheme(scheme string) {
    r.Scheme = &scheme
}

/* param reqMethod: (Optional) */
func (r *QueryLiveStatisticsAreaDataGroupByRequest) SetReqMethod(reqMethod string) {
    r.ReqMethod = &reqMethod
}

/* param cacheLevel: cacheLevel(Optional) */
func (r *QueryLiveStatisticsAreaDataGroupByRequest) SetCacheLevel(cacheLevel string) {
    r.CacheLevel = &cacheLevel
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryLiveStatisticsAreaDataGroupByRequest) GetRegionId() string {
    return ""
}

type QueryLiveStatisticsAreaDataGroupByResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryLiveStatisticsAreaDataGroupByResult `json:"result"`
}

type QueryLiveStatisticsAreaDataGroupByResult struct {
    StartTime string `json:"startTime"`
    EndTime string `json:"endTime"`
    Domain string `json:"domain"`
    Statistics []cdn.StatisticsWithAreaGroupDetail `json:"statistics"`
}