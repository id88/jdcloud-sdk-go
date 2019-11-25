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

type DeleteUserDomainRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 要删除domain的id集合，以,分隔  */
    DomainIds string `json:"domainIds"`

    /* api分组id (Optional) */
    ApiGroupId *string `json:"apiGroupId"`
}

/*
 * param regionId: 地域ID (Required)
 * param domainIds: 要删除domain的id集合，以,分隔 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteUserDomainRequest(
    regionId string,
    domainIds string,
) *DeleteUserDomainRequest {

	return &DeleteUserDomainRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/userdomain",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DomainIds: domainIds,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param domainIds: 要删除domain的id集合，以,分隔 (Required)
 * param apiGroupId: api分组id (Optional)
 */
func NewDeleteUserDomainRequestWithAllParams(
    regionId string,
    domainIds string,
    apiGroupId *string,
) *DeleteUserDomainRequest {

    return &DeleteUserDomainRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/userdomain",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DomainIds: domainIds,
        ApiGroupId: apiGroupId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteUserDomainRequestWithoutParam() *DeleteUserDomainRequest {

    return &DeleteUserDomainRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/userdomain",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DeleteUserDomainRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param domainIds: 要删除domain的id集合，以,分隔(Required) */
func (r *DeleteUserDomainRequest) SetDomainIds(domainIds string) {
    r.DomainIds = domainIds
}

/* param apiGroupId: api分组id(Optional) */
func (r *DeleteUserDomainRequest) SetApiGroupId(apiGroupId string) {
    r.ApiGroupId = &apiGroupId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteUserDomainRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteUserDomainResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteUserDomainResult `json:"result"`
}

type DeleteUserDomainResult struct {
}