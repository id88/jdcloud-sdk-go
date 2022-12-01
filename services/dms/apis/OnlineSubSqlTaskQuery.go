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
    dms "github.com/jdcloud-api/jdcloud-sdk-go/services/dms/models"
)

type OnlineSubSqlTaskQueryRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* 任务id (Optional) */
    TaskId *int `json:"taskId"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewOnlineSubSqlTaskQueryRequest(
    regionId string,
) *OnlineSubSqlTaskQueryRequest {

	return &OnlineSubSqlTaskQueryRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/subsqltask:query",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param taskId: 任务id (Optional)
 */
func NewOnlineSubSqlTaskQueryRequestWithAllParams(
    regionId string,
    taskId *int,
) *OnlineSubSqlTaskQueryRequest {

    return &OnlineSubSqlTaskQueryRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/subsqltask:query",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        TaskId: taskId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewOnlineSubSqlTaskQueryRequestWithoutParam() *OnlineSubSqlTaskQueryRequest {

    return &OnlineSubSqlTaskQueryRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/subsqltask:query",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *OnlineSubSqlTaskQueryRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param taskId: 任务id(Optional) */
func (r *OnlineSubSqlTaskQueryRequest) SetTaskId(taskId int) {
    r.TaskId = &taskId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r OnlineSubSqlTaskQueryRequest) GetRegionId() string {
    return r.RegionId
}

type OnlineSubSqlTaskQueryResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result OnlineSubSqlTaskQueryResult `json:"result"`
}

type OnlineSubSqlTaskQueryResult struct {
    DmsOnlineSubSqlTasks []dms.DmsOnlineSubSqlTask `json:"dmsOnlineSubSqlTasks"`
}