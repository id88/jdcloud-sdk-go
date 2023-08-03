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
    monitor "github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/models"
)

type BatchDescribeMetricDataRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 聚合方式，用于不同时间轴上的聚合。如balance产品同一个resourceId下存在port=80和port=8080等多种维度。可选值参考:sum、avg、min、max (Optional) */
    AggrType *string `json:"aggrType"`

    /* 采样方式，用于在时间轴维度上将聚合周期内的数据聚合为一个点。可选值参考：sum(聚合周期内的数据求和)、avg(求平均)、last(最新值)、min(最小值)、max(最大值) (Optional) */
    DownSampleType *string `json:"downSampleType"`

    /* 查询时间范围的开始时间， UTC时间，格式：2016-12-11T00:00:00+0800（注意在url中+要转译为%2B故url中为2016-12-11T00:00:00%2B0800） (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询时间范围的结束时间， UTC时间，格式：2016-12-11T00:00:00+0800（为空时，将由startTime与timeInterval计算得出）（注意在url中+要转译为%2B故url中为2016-12-11T00:00:00%2B0800） (Optional) */
    EndTime *string `json:"endTime"`

    /* 时间间隔：1h，6h，12h，1d，3d，7d，14d，固定时间间隔，timeInterval默认为1h，当前时间往 前1h (Optional) */
    TimeInterval *string `json:"timeInterval"`

    /* 监控指标数据的维度信息,根据tags来筛选指标数据不同的维度 (Optional) */
    Tags []monitor.TagFilter `json:"tags"`

    /* 是否对查询的tags分组 (Optional) */
    GroupBy *bool `json:"groupBy"`

    /* 是否求速率 (Optional) */
    Rate *bool `json:"rate"`

    /* 资源的类型，取值vm, lb, ip, database 等,<a href="https://docs.jdcloud.com/cn/monitoring/api/describeservices?content=API&SOP=JDCloud">describeServices</a>：查询己接入云监控的产品线列表 (Optional) */
    ServiceCode *string `json:"serviceCode"`

    /* 资源的维度。查询serviceCode下可用的维度请使用describeServices接口 (Optional) */
    Dimension *string `json:"dimension"`

    /* 资源的uuid  */
    ResourceId string `json:"resourceId"`

    /* 是否跨资源查询，默认为false。当该字段为false时，取resourceId字段进行查询；当该子弹为true时，忽略resourceId字段，从tags中取resourceId作为实际的多资源id处理。 (Optional) */
    MultiResources *bool `json:"multiResources"`

    /* 自定义过滤标签，查询时必须在filters中指定要查询的metric，支持多个metric。如：  name='metric',values=["metric1","metric2"]  */
    Filters []monitor.Filter `json:"filters"`
}

/*
 * param regionId: 地域 Id (Required)
 * param resourceId: 资源的uuid (Required)
 * param filters: 自定义过滤标签，查询时必须在filters中指定要查询的metric，支持多个metric。如：  name='metric',values=["metric1","metric2"] (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewBatchDescribeMetricDataRequest(
    regionId string,
    resourceId string,
    filters []monitor.Filter,
) *BatchDescribeMetricDataRequest {

	return &BatchDescribeMetricDataRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/ydMetricsData",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        ResourceId: resourceId,
        Filters: filters,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param aggrType: 聚合方式，用于不同时间轴上的聚合。如balance产品同一个resourceId下存在port=80和port=8080等多种维度。可选值参考:sum、avg、min、max (Optional)
 * param downSampleType: 采样方式，用于在时间轴维度上将聚合周期内的数据聚合为一个点。可选值参考：sum(聚合周期内的数据求和)、avg(求平均)、last(最新值)、min(最小值)、max(最大值) (Optional)
 * param startTime: 查询时间范围的开始时间， UTC时间，格式：2016-12-11T00:00:00+0800（注意在url中+要转译为%2B故url中为2016-12-11T00:00:00%2B0800） (Optional)
 * param endTime: 查询时间范围的结束时间， UTC时间，格式：2016-12-11T00:00:00+0800（为空时，将由startTime与timeInterval计算得出）（注意在url中+要转译为%2B故url中为2016-12-11T00:00:00%2B0800） (Optional)
 * param timeInterval: 时间间隔：1h，6h，12h，1d，3d，7d，14d，固定时间间隔，timeInterval默认为1h，当前时间往 前1h (Optional)
 * param tags: 监控指标数据的维度信息,根据tags来筛选指标数据不同的维度 (Optional)
 * param groupBy: 是否对查询的tags分组 (Optional)
 * param rate: 是否求速率 (Optional)
 * param serviceCode: 资源的类型，取值vm, lb, ip, database 等,<a href="https://docs.jdcloud.com/cn/monitoring/api/describeservices?content=API&SOP=JDCloud">describeServices</a>：查询己接入云监控的产品线列表 (Optional)
 * param dimension: 资源的维度。查询serviceCode下可用的维度请使用describeServices接口 (Optional)
 * param resourceId: 资源的uuid (Required)
 * param multiResources: 是否跨资源查询，默认为false。当该字段为false时，取resourceId字段进行查询；当该子弹为true时，忽略resourceId字段，从tags中取resourceId作为实际的多资源id处理。 (Optional)
 * param filters: 自定义过滤标签，查询时必须在filters中指定要查询的metric，支持多个metric。如：  name='metric',values=["metric1","metric2"] (Required)
 */
func NewBatchDescribeMetricDataRequestWithAllParams(
    regionId string,
    aggrType *string,
    downSampleType *string,
    startTime *string,
    endTime *string,
    timeInterval *string,
    tags []monitor.TagFilter,
    groupBy *bool,
    rate *bool,
    serviceCode *string,
    dimension *string,
    resourceId string,
    multiResources *bool,
    filters []monitor.Filter,
) *BatchDescribeMetricDataRequest {

    return &BatchDescribeMetricDataRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ydMetricsData",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AggrType: aggrType,
        DownSampleType: downSampleType,
        StartTime: startTime,
        EndTime: endTime,
        TimeInterval: timeInterval,
        Tags: tags,
        GroupBy: groupBy,
        Rate: rate,
        ServiceCode: serviceCode,
        Dimension: dimension,
        ResourceId: resourceId,
        MultiResources: multiResources,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewBatchDescribeMetricDataRequestWithoutParam() *BatchDescribeMetricDataRequest {

    return &BatchDescribeMetricDataRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ydMetricsData",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *BatchDescribeMetricDataRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param aggrType: 聚合方式，用于不同时间轴上的聚合。如balance产品同一个resourceId下存在port=80和port=8080等多种维度。可选值参考:sum、avg、min、max(Optional) */
func (r *BatchDescribeMetricDataRequest) SetAggrType(aggrType string) {
    r.AggrType = &aggrType
}
/* param downSampleType: 采样方式，用于在时间轴维度上将聚合周期内的数据聚合为一个点。可选值参考：sum(聚合周期内的数据求和)、avg(求平均)、last(最新值)、min(最小值)、max(最大值)(Optional) */
func (r *BatchDescribeMetricDataRequest) SetDownSampleType(downSampleType string) {
    r.DownSampleType = &downSampleType
}
/* param startTime: 查询时间范围的开始时间， UTC时间，格式：2016-12-11T00:00:00+0800（注意在url中+要转译为%2B故url中为2016-12-11T00:00:00%2B0800）(Optional) */
func (r *BatchDescribeMetricDataRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}
/* param endTime: 查询时间范围的结束时间， UTC时间，格式：2016-12-11T00:00:00+0800（为空时，将由startTime与timeInterval计算得出）（注意在url中+要转译为%2B故url中为2016-12-11T00:00:00%2B0800）(Optional) */
func (r *BatchDescribeMetricDataRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}
/* param timeInterval: 时间间隔：1h，6h，12h，1d，3d，7d，14d，固定时间间隔，timeInterval默认为1h，当前时间往 前1h(Optional) */
func (r *BatchDescribeMetricDataRequest) SetTimeInterval(timeInterval string) {
    r.TimeInterval = &timeInterval
}
/* param tags: 监控指标数据的维度信息,根据tags来筛选指标数据不同的维度(Optional) */
func (r *BatchDescribeMetricDataRequest) SetTags(tags []monitor.TagFilter) {
    r.Tags = tags
}
/* param groupBy: 是否对查询的tags分组(Optional) */
func (r *BatchDescribeMetricDataRequest) SetGroupBy(groupBy bool) {
    r.GroupBy = &groupBy
}
/* param rate: 是否求速率(Optional) */
func (r *BatchDescribeMetricDataRequest) SetRate(rate bool) {
    r.Rate = &rate
}
/* param serviceCode: 资源的类型，取值vm, lb, ip, database 等,<a href="https://docs.jdcloud.com/cn/monitoring/api/describeservices?content=API&SOP=JDCloud">describeServices</a>：查询己接入云监控的产品线列表(Optional) */
func (r *BatchDescribeMetricDataRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}
/* param dimension: 资源的维度。查询serviceCode下可用的维度请使用describeServices接口(Optional) */
func (r *BatchDescribeMetricDataRequest) SetDimension(dimension string) {
    r.Dimension = &dimension
}
/* param resourceId: 资源的uuid(Required) */
func (r *BatchDescribeMetricDataRequest) SetResourceId(resourceId string) {
    r.ResourceId = resourceId
}
/* param multiResources: 是否跨资源查询，默认为false。当该字段为false时，取resourceId字段进行查询；当该子弹为true时，忽略resourceId字段，从tags中取resourceId作为实际的多资源id处理。(Optional) */
func (r *BatchDescribeMetricDataRequest) SetMultiResources(multiResources bool) {
    r.MultiResources = &multiResources
}
/* param filters: 自定义过滤标签，查询时必须在filters中指定要查询的metric，支持多个metric。如：  name='metric',values=["metric1","metric2"](Required) */
func (r *BatchDescribeMetricDataRequest) SetFilters(filters []monitor.Filter) {
    r.Filters = filters
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r BatchDescribeMetricDataRequest) GetRegionId() string {
    return r.RegionId
}

type BatchDescribeMetricDataResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result BatchDescribeMetricDataResult `json:"result"`
}

type BatchDescribeMetricDataResult struct {
    MetricDatas []monitor.MetricData `json:"metricDatas"`
}