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

type QueryRefreshTaskByIdRequest struct {

    core.JDCloudRequest

    /* 域名组id  */
    TaskId string `json:"taskId"`
}

/*
 * param taskId: 域名组id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryRefreshTaskByIdRequest(
    taskId string,
) *QueryRefreshTaskByIdRequest {

	return &QueryRefreshTaskByIdRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/task/{taskId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        TaskId: taskId,
	}
}

/*
 * param taskId: 域名组id (Required)
 */
func NewQueryRefreshTaskByIdRequestWithAllParams(
    taskId string,
) *QueryRefreshTaskByIdRequest {

    return &QueryRefreshTaskByIdRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/task/{taskId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        TaskId: taskId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryRefreshTaskByIdRequestWithoutParam() *QueryRefreshTaskByIdRequest {

    return &QueryRefreshTaskByIdRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/task/{taskId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param taskId: 域名组id(Required) */
func (r *QueryRefreshTaskByIdRequest) SetTaskId(taskId string) {
    r.TaskId = taskId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryRefreshTaskByIdRequest) GetRegionId() string {
    return ""
}

type QueryRefreshTaskByIdResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryRefreshTaskByIdResult `json:"result"`
}

type QueryRefreshTaskByIdResult struct {
    Task cdn.RefreshTask `json:"task"`
}