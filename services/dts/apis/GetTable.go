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
    dts "github.com/jdcloud-api/jdcloud-sdk-go/services/dts/models"
)

type GetTableRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》]  */
    RegionId string `json:"regionId"`

    /* DTS连接模板ID  */
    EndpointId string `json:"endpointId"`

    /* table名称  */
    TableName string `json:"tableName"`

    /* 数据传输任务ID (Optional) */
    TaskId *string `json:"taskId"`

    /* 数据库名称，此项选填。 (Optional) */
    DatabaseName *string `json:"databaseName"`

    /* 模式名称，此项选填。 (Optional) */
    SchemaName *string `json:"schemaName"`

    /* 连接模板信息  */
    Endpoint *dts.TransmissionEndpoint `json:"endpoint"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param endpointId: DTS连接模板ID (Required)
 * param tableName: table名称 (Required)
 * param endpoint: 连接模板信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetTableRequest(
    regionId string,
    endpointId string,
    tableName string,
    endpoint *dts.TransmissionEndpoint,
) *GetTableRequest {

	return &GetTableRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/endpoint/{endpointId}/table/{tableName}",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        EndpointId: endpointId,
        TableName: tableName,
        Endpoint: endpoint,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param endpointId: DTS连接模板ID (Required)
 * param tableName: table名称 (Required)
 * param taskId: 数据传输任务ID (Optional)
 * param databaseName: 数据库名称，此项选填。 (Optional)
 * param schemaName: 模式名称，此项选填。 (Optional)
 * param endpoint: 连接模板信息 (Required)
 */
func NewGetTableRequestWithAllParams(
    regionId string,
    endpointId string,
    tableName string,
    taskId *string,
    databaseName *string,
    schemaName *string,
    endpoint *dts.TransmissionEndpoint,
) *GetTableRequest {

    return &GetTableRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/endpoint/{endpointId}/table/{tableName}",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        EndpointId: endpointId,
        TableName: tableName,
        TaskId: taskId,
        DatabaseName: databaseName,
        SchemaName: schemaName,
        Endpoint: endpoint,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetTableRequestWithoutParam() *GetTableRequest {

    return &GetTableRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/endpoint/{endpointId}/table/{tableName}",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](Required) */
func (r *GetTableRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param endpointId: DTS连接模板ID(Required) */
func (r *GetTableRequest) SetEndpointId(endpointId string) {
    r.EndpointId = endpointId
}
/* param tableName: table名称(Required) */
func (r *GetTableRequest) SetTableName(tableName string) {
    r.TableName = tableName
}
/* param taskId: 数据传输任务ID(Optional) */
func (r *GetTableRequest) SetTaskId(taskId string) {
    r.TaskId = &taskId
}
/* param databaseName: 数据库名称，此项选填。(Optional) */
func (r *GetTableRequest) SetDatabaseName(databaseName string) {
    r.DatabaseName = &databaseName
}
/* param schemaName: 模式名称，此项选填。(Optional) */
func (r *GetTableRequest) SetSchemaName(schemaName string) {
    r.SchemaName = &schemaName
}
/* param endpoint: 连接模板信息(Required) */
func (r *GetTableRequest) SetEndpoint(endpoint *dts.TransmissionEndpoint) {
    r.Endpoint = endpoint
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetTableRequest) GetRegionId() string {
    return r.RegionId
}

type GetTableResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetTableResult `json:"result"`
}

type GetTableResult struct {
    Table dts.Table `json:"table"`
}