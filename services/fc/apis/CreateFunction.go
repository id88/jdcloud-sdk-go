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
    fc "github.com/jdcloud-api/jdcloud-sdk-go/services/fc/models"
)

type CreateFunctionRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Service Name  */
    ServiceName string `json:"serviceName"`

    /* 函数 创建参数  */
    FunctionSpec *fc.FunctionSpec `json:"functionSpec"`

    /* 保证请求幂等性的字符串；最大长度64个ASCII字符 (Optional) */
    ClientToken *string `json:"clientToken"`
}

/*
 * param regionId: Region ID (Required)
 * param serviceName: Service Name (Required)
 * param functionSpec: 函数 创建参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateFunctionRequest(
    regionId string,
    serviceName string,
    functionSpec *fc.FunctionSpec,
) *CreateFunctionRequest {

	return &CreateFunctionRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/services/{serviceName}/functions",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ServiceName: serviceName,
        FunctionSpec: functionSpec,
	}
}

/*
 * param regionId: Region ID (Required)
 * param serviceName: Service Name (Required)
 * param functionSpec: 函数 创建参数 (Required)
 * param clientToken: 保证请求幂等性的字符串；最大长度64个ASCII字符 (Optional)
 */
func NewCreateFunctionRequestWithAllParams(
    regionId string,
    serviceName string,
    functionSpec *fc.FunctionSpec,
    clientToken *string,
) *CreateFunctionRequest {

    return &CreateFunctionRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/services/{serviceName}/functions",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ServiceName: serviceName,
        FunctionSpec: functionSpec,
        ClientToken: clientToken,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateFunctionRequestWithoutParam() *CreateFunctionRequest {

    return &CreateFunctionRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/services/{serviceName}/functions",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateFunctionRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param serviceName: Service Name(Required) */
func (r *CreateFunctionRequest) SetServiceName(serviceName string) {
    r.ServiceName = serviceName
}
/* param functionSpec: 函数 创建参数(Required) */
func (r *CreateFunctionRequest) SetFunctionSpec(functionSpec *fc.FunctionSpec) {
    r.FunctionSpec = functionSpec
}
/* param clientToken: 保证请求幂等性的字符串；最大长度64个ASCII字符(Optional) */
func (r *CreateFunctionRequest) SetClientToken(clientToken string) {
    r.ClientToken = &clientToken
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateFunctionRequest) GetRegionId() string {
    return r.RegionId
}

type CreateFunctionResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateFunctionResult `json:"result"`
}

type CreateFunctionResult struct {
    FunctionName string `json:"functionName"`
}