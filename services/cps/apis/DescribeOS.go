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
    cps "github.com/jdcloud-api/jdcloud-sdk-go/services/cps/models"
)

type DescribeOSRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 实例类型，可调用接口（describeDeviceTypes）获取指定地域的实例类型，例如：cps.c.normal  */
    DeviceType string `json:"deviceType"`

    /* 操作系统类型，取值范围：CentOS、Ubuntu、Windows (Optional) */
    OsType *string `json:"osType"`
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param deviceType: 实例类型，可调用接口（describeDeviceTypes）获取指定地域的实例类型，例如：cps.c.normal (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeOSRequest(
    regionId string,
    deviceType string,
) *DescribeOSRequest {

	return &DescribeOSRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/os",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DeviceType: deviceType,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param deviceType: 实例类型，可调用接口（describeDeviceTypes）获取指定地域的实例类型，例如：cps.c.normal (Required)
 * param osType: 操作系统类型，取值范围：CentOS、Ubuntu、Windows (Optional)
 */
func NewDescribeOSRequestWithAllParams(
    regionId string,
    deviceType string,
    osType *string,
) *DescribeOSRequest {

    return &DescribeOSRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/os",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DeviceType: deviceType,
        OsType: osType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeOSRequestWithoutParam() *DescribeOSRequest {

    return &DescribeOSRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/os",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域(Required) */
func (r *DescribeOSRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param deviceType: 实例类型，可调用接口（describeDeviceTypes）获取指定地域的实例类型，例如：cps.c.normal(Required) */
func (r *DescribeOSRequest) SetDeviceType(deviceType string) {
    r.DeviceType = deviceType
}

/* param osType: 操作系统类型，取值范围：CentOS、Ubuntu、Windows(Optional) */
func (r *DescribeOSRequest) SetOsType(osType string) {
    r.OsType = &osType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeOSRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeOSResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeOSResult `json:"result"`
}

type DescribeOSResult struct {
    Oss []cps.Os `json:"oss"`
}