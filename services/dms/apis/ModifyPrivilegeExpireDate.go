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

type ModifyPrivilegeExpireDateRequest struct {

    core.JDCloudRequest

    /* 授权实例记录对应的主键ID (Optional) */
    IdList []int64 `json:"idList"`

    /* 授权过期时间(yyyy-MM-dd'T'HH:mm:ss.SSS'Z') (Optional) */
    ExpireDate *string `json:"expireDate"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyPrivilegeExpireDateRequest(
) *ModifyPrivilegeExpireDateRequest {

	return &ModifyPrivilegeExpireDateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/management:modifyPrivilegeExpireDate",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param idList: 授权实例记录对应的主键ID (Optional)
 * param expireDate: 授权过期时间(yyyy-MM-dd'T'HH:mm:ss.SSS'Z') (Optional)
 */
func NewModifyPrivilegeExpireDateRequestWithAllParams(
    idList []int64,
    expireDate *string,
) *ModifyPrivilegeExpireDateRequest {

    return &ModifyPrivilegeExpireDateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/management:modifyPrivilegeExpireDate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        IdList: idList,
        ExpireDate: expireDate,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyPrivilegeExpireDateRequestWithoutParam() *ModifyPrivilegeExpireDateRequest {

    return &ModifyPrivilegeExpireDateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/management:modifyPrivilegeExpireDate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param idList: 授权实例记录对应的主键ID(Optional) */
func (r *ModifyPrivilegeExpireDateRequest) SetIdList(idList []int64) {
    r.IdList = idList
}
/* param expireDate: 授权过期时间(yyyy-MM-dd'T'HH:mm:ss.SSS'Z')(Optional) */
func (r *ModifyPrivilegeExpireDateRequest) SetExpireDate(expireDate string) {
    r.ExpireDate = &expireDate
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyPrivilegeExpireDateRequest) GetRegionId() string {
    return ""
}

type ModifyPrivilegeExpireDateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyPrivilegeExpireDateResult `json:"result"`
}

type ModifyPrivilegeExpireDateResult struct {
}