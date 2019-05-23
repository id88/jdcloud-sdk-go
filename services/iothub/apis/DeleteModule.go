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

type DeleteModuleRequest struct {

    core.JDCloudRequest

    /* 边缘计算节点编号  */
    EdgeId string `json:"edgeId"`

    /* 边缘计算模块编号  */
    ModuleId string `json:"moduleId"`
}

/*
 * param edgeId: 边缘计算节点编号 (Required)
 * param moduleId: 边缘计算模块编号 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteModuleRequest(
    edgeId string,
    moduleId string,
) *DeleteModuleRequest {

	return &DeleteModuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/edge/{edgeId}/module/{moduleId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        EdgeId: edgeId,
        ModuleId: moduleId,
	}
}

/*
 * param edgeId: 边缘计算节点编号 (Required)
 * param moduleId: 边缘计算模块编号 (Required)
 */
func NewDeleteModuleRequestWithAllParams(
    edgeId string,
    moduleId string,
) *DeleteModuleRequest {

    return &DeleteModuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/edge/{edgeId}/module/{moduleId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        EdgeId: edgeId,
        ModuleId: moduleId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteModuleRequestWithoutParam() *DeleteModuleRequest {

    return &DeleteModuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/edge/{edgeId}/module/{moduleId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param edgeId: 边缘计算节点编号(Required) */
func (r *DeleteModuleRequest) SetEdgeId(edgeId string) {
    r.EdgeId = edgeId
}

/* param moduleId: 边缘计算模块编号(Required) */
func (r *DeleteModuleRequest) SetModuleId(moduleId string) {
    r.ModuleId = moduleId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteModuleRequest) GetRegionId() string {
    return ""
}

type DeleteModuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteModuleResult `json:"result"`
}

type DeleteModuleResult struct {
}