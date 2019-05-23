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
    antipro "github.com/jdcloud-api/jdcloud-sdk-go/services/antipro/models"
)

type DescribeOperationRecordsRequest struct {

    core.JDCloudRequest

    /* 页码 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 开始时间, UTC 时间, 格式：yyyy-MM-dd'T'HH:mm:ssZ  */
    StartTime string `json:"startTime"`

    /* 结束时间, UTC 时间, 格式：yyyy-MM-dd'T'HH:mm:ssZ  */
    EndTime string `json:"endTime"`

    /* 操作类型, 默认查全部.
- 0: 全部
- 1: 套餐变更
- 2: 防护规则变更
- 3: 防护对象变更
- 4: IP 地址变更
- 5: 防护包名称变更
 (Optional) */
    Action *int `json:"action"`

    /* 防护包名称, 支持模糊匹配 (Optional) */
    Name *string `json:"name"`
}

/*
 * param startTime: 开始时间, UTC 时间, 格式：yyyy-MM-dd'T'HH:mm:ssZ (Required)
 * param endTime: 结束时间, UTC 时间, 格式：yyyy-MM-dd'T'HH:mm:ssZ (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeOperationRecordsRequest(
    startTime string,
    endTime string,
) *DescribeOperationRecordsRequest {

	return &DescribeOperationRecordsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/operationRecords",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        StartTime: startTime,
        EndTime: endTime,
	}
}

/*
 * param pageNumber: 页码 (Optional)
 * param pageSize: 分页大小 (Optional)
 * param startTime: 开始时间, UTC 时间, 格式：yyyy-MM-dd'T'HH:mm:ssZ (Required)
 * param endTime: 结束时间, UTC 时间, 格式：yyyy-MM-dd'T'HH:mm:ssZ (Required)
 * param action: 操作类型, 默认查全部.
- 0: 全部
- 1: 套餐变更
- 2: 防护规则变更
- 3: 防护对象变更
- 4: IP 地址变更
- 5: 防护包名称变更
 (Optional)
 * param name: 防护包名称, 支持模糊匹配 (Optional)
 */
func NewDescribeOperationRecordsRequestWithAllParams(
    pageNumber *int,
    pageSize *int,
    startTime string,
    endTime string,
    action *int,
    name *string,
) *DescribeOperationRecordsRequest {

    return &DescribeOperationRecordsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/operationRecords",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNumber: pageNumber,
        PageSize: pageSize,
        StartTime: startTime,
        EndTime: endTime,
        Action: action,
        Name: name,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeOperationRecordsRequestWithoutParam() *DescribeOperationRecordsRequest {

    return &DescribeOperationRecordsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/operationRecords",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNumber: 页码(Optional) */
func (r *DescribeOperationRecordsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小(Optional) */
func (r *DescribeOperationRecordsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param startTime: 开始时间, UTC 时间, 格式：yyyy-MM-dd'T'HH:mm:ssZ(Required) */
func (r *DescribeOperationRecordsRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 结束时间, UTC 时间, 格式：yyyy-MM-dd'T'HH:mm:ssZ(Required) */
func (r *DescribeOperationRecordsRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

/* param action: 操作类型, 默认查全部.
- 0: 全部
- 1: 套餐变更
- 2: 防护规则变更
- 3: 防护对象变更
- 4: IP 地址变更
- 5: 防护包名称变更
(Optional) */
func (r *DescribeOperationRecordsRequest) SetAction(action int) {
    r.Action = &action
}

/* param name: 防护包名称, 支持模糊匹配(Optional) */
func (r *DescribeOperationRecordsRequest) SetName(name string) {
    r.Name = &name
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeOperationRecordsRequest) GetRegionId() string {
    return ""
}

type DescribeOperationRecordsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeOperationRecordsResult `json:"result"`
}

type DescribeOperationRecordsResult struct {
    DataList []antipro.OperationRecord `json:"dataList"`
    CurrentCount int `json:"currentCount"`
    TotalCount int `json:"totalCount"`
    TotalPage int `json:"totalPage"`
}