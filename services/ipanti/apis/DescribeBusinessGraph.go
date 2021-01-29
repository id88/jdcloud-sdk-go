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
)

type DescribeBusinessGraphRequest struct {

    core.JDCloudRequest

    /* 区域 ID, 高防不区分区域, 传 cn-north-1 即可  */
    RegionId string `json:"regionId"`

    /* 开始时间, 只能查询最近 90 天以内的数据, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ  */
    StartTime string `json:"startTime"`

    /* 查询的结束时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ (Optional) */
    EndTime *string `json:"endTime"`

    /* 高防实例 Id 列表 (Optional) */
    InstanceId []string `json:"instanceId"`
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param startTime: 开始时间, 只能查询最近 90 天以内的数据, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeBusinessGraphRequest(
    regionId string,
    startTime string,
) *DescribeBusinessGraphRequest {

	return &DescribeBusinessGraphRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/charts:businessGraph",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        StartTime: startTime,
	}
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param startTime: 开始时间, 只能查询最近 90 天以内的数据, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ (Required)
 * param endTime: 查询的结束时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ (Optional)
 * param instanceId: 高防实例 Id 列表 (Optional)
 */
func NewDescribeBusinessGraphRequestWithAllParams(
    regionId string,
    startTime string,
    endTime *string,
    instanceId []string,
) *DescribeBusinessGraphRequest {

    return &DescribeBusinessGraphRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/charts:businessGraph",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        StartTime: startTime,
        EndTime: endTime,
        InstanceId: instanceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeBusinessGraphRequestWithoutParam() *DescribeBusinessGraphRequest {

    return &DescribeBusinessGraphRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/charts:businessGraph",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可(Required) */
func (r *DescribeBusinessGraphRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param startTime: 开始时间, 只能查询最近 90 天以内的数据, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ(Required) */
func (r *DescribeBusinessGraphRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 查询的结束时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ(Optional) */
func (r *DescribeBusinessGraphRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param instanceId: 高防实例 Id 列表(Optional) */
func (r *DescribeBusinessGraphRequest) SetInstanceId(instanceId []string) {
    r.InstanceId = instanceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeBusinessGraphRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeBusinessGraphResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeBusinessGraphResult `json:"result"`
}

type DescribeBusinessGraphResult struct {
    InTraffic []float64 `json:"inTraffic"`
    OutTraffic []float64 `json:"outTraffic"`
    Time []string `json:"time"`
    Unit string `json:"unit"`
}