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

type DescribeRdsDatabasesRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* RDS 实例ID，唯一标识一个RDS实例  */
    InstanceId string `json:"instanceId"`

    /* 数据库名。如果不指定数据库名，则返回所有数据库列表<br>- **MySQL：不支持该字段**<br>- **SQL Server：支持该字段** (Optional) */
    DbName *string `json:"dbName"`

    /* 显示数据的页码，默认为1，取值范围：[-1,∞)。pageNumber为-1时，返回所有数据页码；超过总页数时，显示最后一页; (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 每页显示的数据条数，默认为100，取值范围：[10,100]，用于查询列表的接口 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeRdsDatabasesRequest(
    regionId string,
    instanceId string,
) *DescribeRdsDatabasesRequest {

	return &DescribeRdsDatabasesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/ydRdsInstances/{instanceId}/databases",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param dbName: 数据库名。如果不指定数据库名，则返回所有数据库列表<br>- **MySQL：不支持该字段**<br>- **SQL Server：支持该字段** (Optional)
 * param pageNumber: 显示数据的页码，默认为1，取值范围：[-1,∞)。pageNumber为-1时，返回所有数据页码；超过总页数时，显示最后一页; (Optional)
 * param pageSize: 每页显示的数据条数，默认为100，取值范围：[10,100]，用于查询列表的接口 (Optional)
 */
func NewDescribeRdsDatabasesRequestWithAllParams(
    regionId string,
    instanceId string,
    dbName *string,
    pageNumber *int,
    pageSize *int,
) *DescribeRdsDatabasesRequest {

    return &DescribeRdsDatabasesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ydRdsInstances/{instanceId}/databases",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        DbName: dbName,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeRdsDatabasesRequestWithoutParam() *DescribeRdsDatabasesRequest {

    return &DescribeRdsDatabasesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ydRdsInstances/{instanceId}/databases",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *DescribeRdsDatabasesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param instanceId: RDS 实例ID，唯一标识一个RDS实例(Required) */
func (r *DescribeRdsDatabasesRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}
/* param dbName: 数据库名。如果不指定数据库名，则返回所有数据库列表<br>- **MySQL：不支持该字段**<br>- **SQL Server：支持该字段**(Optional) */
func (r *DescribeRdsDatabasesRequest) SetDbName(dbName string) {
    r.DbName = &dbName
}
/* param pageNumber: 显示数据的页码，默认为1，取值范围：[-1,∞)。pageNumber为-1时，返回所有数据页码；超过总页数时，显示最后一页;(Optional) */
func (r *DescribeRdsDatabasesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}
/* param pageSize: 每页显示的数据条数，默认为100，取值范围：[10,100]，用于查询列表的接口(Optional) */
func (r *DescribeRdsDatabasesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeRdsDatabasesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeRdsDatabasesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeRdsDatabasesResult `json:"result"`
}

type DescribeRdsDatabasesResult struct {
    Databases []rds.Database `json:"databases"`
    TotalCount int `json:"totalCount"`
}