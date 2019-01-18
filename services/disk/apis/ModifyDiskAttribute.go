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

type ModifyDiskAttributeRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 云硬盘ID  */
    DiskId string `json:"diskId"`

    /* 云硬盘名称，只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。 (Optional) */
    Name *string `json:"name"`

    /* 云硬盘描述，允许输入UTF-8编码下的全部字符，不超过256字符。 (Optional) */
    Description *string `json:"description"`
}

/*
 * param regionId: 地域ID (Required)
 * param diskId: 云硬盘ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyDiskAttributeRequest(
    regionId string,
    diskId string,
) *ModifyDiskAttributeRequest {

	return &ModifyDiskAttributeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/disks/{diskId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DiskId: diskId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param diskId: 云硬盘ID (Required)
 * param name: 云硬盘名称，只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。 (Optional)
 * param description: 云硬盘描述，允许输入UTF-8编码下的全部字符，不超过256字符。 (Optional)
 */
func NewModifyDiskAttributeRequestWithAllParams(
    regionId string,
    diskId string,
    name *string,
    description *string,
) *ModifyDiskAttributeRequest {

    return &ModifyDiskAttributeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/disks/{diskId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DiskId: diskId,
        Name: name,
        Description: description,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyDiskAttributeRequestWithoutParam() *ModifyDiskAttributeRequest {

    return &ModifyDiskAttributeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/disks/{diskId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *ModifyDiskAttributeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param diskId: 云硬盘ID(Required) */
func (r *ModifyDiskAttributeRequest) SetDiskId(diskId string) {
    r.DiskId = diskId
}

/* param name: 云硬盘名称，只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。(Optional) */
func (r *ModifyDiskAttributeRequest) SetName(name string) {
    r.Name = &name
}

/* param description: 云硬盘描述，允许输入UTF-8编码下的全部字符，不超过256字符。(Optional) */
func (r *ModifyDiskAttributeRequest) SetDescription(description string) {
    r.Description = &description
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyDiskAttributeRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyDiskAttributeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyDiskAttributeResult `json:"result"`
}

type ModifyDiskAttributeResult struct {
}