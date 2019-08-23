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
    rds "github.com/jdcloud-api/jdcloud-sdk-go/services/rds/models"
)

type DescribeSlowLogsRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* RDS 实例ID，唯一标识一个RDS实例  */
    InstanceId string `json:"instanceId"`

    /* 慢日志开始时间，格式为：YYYY-MM-DD HH:mm:ss，开始时间到当前时间不能大于 7 天，开始时间不能大于结束时间，结束时间不能大于当前时间  */
    StartTime string `json:"startTime"`

    /* 慢日志结束时间，格式为：YYYY-MM-DD HH:mm:ss，开始时间到当前时间不能大于 7 天，开始时间不能大于结束时间，结束时间不能大于当前时间  */
    EndTime string `json:"endTime"`

    /* 查询哪个数据库的慢日志，不填表示返回所有数据库的慢日志 (Optional) */
    DbName *string `json:"dbName"`

    /* 显示数据的页码，默认为1，取值范围：[-1,1000)。pageNumber为-1时，返回所有数据页码；超过总页数时，显示最后一页。 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 每页显示的数据条数，默认为10，取值范围：10、20、30、50、100 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param startTime: 慢日志开始时间，格式为：YYYY-MM-DD HH:mm:ss，开始时间到当前时间不能大于 7 天，开始时间不能大于结束时间，结束时间不能大于当前时间 (Required)
 * param endTime: 慢日志结束时间，格式为：YYYY-MM-DD HH:mm:ss，开始时间到当前时间不能大于 7 天，开始时间不能大于结束时间，结束时间不能大于当前时间 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeSlowLogsRequest(
    regionId string,
    instanceId string,
    startTime string,
    endTime string,
) *DescribeSlowLogsRequest {

	return &DescribeSlowLogsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/performance:describeSlowLogs",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        StartTime: startTime,
        EndTime: endTime,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param startTime: 慢日志开始时间，格式为：YYYY-MM-DD HH:mm:ss，开始时间到当前时间不能大于 7 天，开始时间不能大于结束时间，结束时间不能大于当前时间 (Required)
 * param endTime: 慢日志结束时间，格式为：YYYY-MM-DD HH:mm:ss，开始时间到当前时间不能大于 7 天，开始时间不能大于结束时间，结束时间不能大于当前时间 (Required)
 * param dbName: 查询哪个数据库的慢日志，不填表示返回所有数据库的慢日志 (Optional)
 * param pageNumber: 显示数据的页码，默认为1，取值范围：[-1,1000)。pageNumber为-1时，返回所有数据页码；超过总页数时，显示最后一页。 (Optional)
 * param pageSize: 每页显示的数据条数，默认为10，取值范围：10、20、30、50、100 (Optional)
 */
func NewDescribeSlowLogsRequestWithAllParams(
    regionId string,
    instanceId string,
    startTime string,
    endTime string,
    dbName *string,
    pageNumber *int,
    pageSize *int,
) *DescribeSlowLogsRequest {

    return &DescribeSlowLogsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/performance:describeSlowLogs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        StartTime: startTime,
        EndTime: endTime,
        DbName: dbName,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeSlowLogsRequestWithoutParam() *DescribeSlowLogsRequest {

    return &DescribeSlowLogsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/performance:describeSlowLogs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *DescribeSlowLogsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: RDS 实例ID，唯一标识一个RDS实例(Required) */
func (r *DescribeSlowLogsRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param startTime: 慢日志开始时间，格式为：YYYY-MM-DD HH:mm:ss，开始时间到当前时间不能大于 7 天，开始时间不能大于结束时间，结束时间不能大于当前时间(Required) */
func (r *DescribeSlowLogsRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 慢日志结束时间，格式为：YYYY-MM-DD HH:mm:ss，开始时间到当前时间不能大于 7 天，开始时间不能大于结束时间，结束时间不能大于当前时间(Required) */
func (r *DescribeSlowLogsRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

/* param dbName: 查询哪个数据库的慢日志，不填表示返回所有数据库的慢日志(Optional) */
func (r *DescribeSlowLogsRequest) SetDbName(dbName string) {
    r.DbName = &dbName
}

/* param pageNumber: 显示数据的页码，默认为1，取值范围：[-1,1000)。pageNumber为-1时，返回所有数据页码；超过总页数时，显示最后一页。(Optional) */
func (r *DescribeSlowLogsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 每页显示的数据条数，默认为10，取值范围：10、20、30、50、100(Optional) */
func (r *DescribeSlowLogsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeSlowLogsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeSlowLogsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeSlowLogsResult `json:"result"`
}

type DescribeSlowLogsResult struct {
    SlowLogs []rds.SlowLogDigest `json:"slowLogs"`
    TotalCount int `json:"totalCount"`
}