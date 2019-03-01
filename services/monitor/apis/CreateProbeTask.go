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

type CreateProbeTaskRequest struct {

    core.JDCloudRequest

    /* 幂等性校验参数,最长36位  */
    ClientToken string `json:"clientToken"`

    /*   */
    CreateProbeTaskSpec *monitor.CreateProbeTaskParam `json:"createProbeTaskSpec"`
}

/*
 * param clientToken: 幂等性校验参数,最长36位 (Required)
 * param createProbeTaskSpec:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateProbeTaskRequest(
    clientToken string,
    createProbeTaskSpec *monitor.CreateProbeTaskParam,
) *CreateProbeTaskRequest {

	return &CreateProbeTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/am/probeTask",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        ClientToken: clientToken,
        CreateProbeTaskSpec: createProbeTaskSpec,
	}
}

/*
 * param clientToken: 幂等性校验参数,最长36位 (Required)
 * param createProbeTaskSpec:  (Required)
 */
func NewCreateProbeTaskRequestWithAllParams(
    clientToken string,
    createProbeTaskSpec *monitor.CreateProbeTaskParam,
) *CreateProbeTaskRequest {

    return &CreateProbeTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/am/probeTask",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        ClientToken: clientToken,
        CreateProbeTaskSpec: createProbeTaskSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateProbeTaskRequestWithoutParam() *CreateProbeTaskRequest {

    return &CreateProbeTaskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/am/probeTask",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param clientToken: 幂等性校验参数,最长36位(Required) */
func (r *CreateProbeTaskRequest) SetClientToken(clientToken string) {
    r.ClientToken = clientToken
}

/* param createProbeTaskSpec: (Required) */
func (r *CreateProbeTaskRequest) SetCreateProbeTaskSpec(createProbeTaskSpec *monitor.CreateProbeTaskParam) {
    r.CreateProbeTaskSpec = createProbeTaskSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateProbeTaskRequest) GetRegionId() string {
    return ""
}

type CreateProbeTaskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateProbeTaskResult `json:"result"`
}

type CreateProbeTaskResult struct {
    Suc bool `json:"suc"`
    TaskId string `json:"taskId"`
}