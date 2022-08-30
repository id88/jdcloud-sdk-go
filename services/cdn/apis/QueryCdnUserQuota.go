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

type QueryCdnUserQuotaRequest struct {

    core.JDCloudRequest
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryCdnUserQuotaRequest(
) *QueryCdnUserQuotaRequest {

	return &QueryCdnUserQuotaRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/user:quota",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 */
func NewQueryCdnUserQuotaRequestWithAllParams(
) *QueryCdnUserQuotaRequest {

    return &QueryCdnUserQuotaRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/user:quota",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryCdnUserQuotaRequestWithoutParam() *QueryCdnUserQuotaRequest {

    return &QueryCdnUserQuotaRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/user:quota",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}



// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryCdnUserQuotaRequest) GetRegionId() string {
    return ""
}

type QueryCdnUserQuotaResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryCdnUserQuotaResult `json:"result"`
}

type QueryCdnUserQuotaResult struct {
    DirAllCount int `json:"dirAllCount"`
    DirUsedCount int `json:"dirUsedCount"`
    DirRemainedCount int `json:"dirRemainedCount"`
    ForbiddenUrlRemainedCount int `json:"forbiddenUrlRemainedCount"`
    ForbiddenUrlUsedCount int `json:"forbiddenUrlUsedCount"`
    ForbiddenUrlAllCount int `json:"forbiddenUrlAllCount"`
    HasForbiddenDomainCount int `json:"hasForbiddenDomainCount"`
    PrefetchRemainedCount int `json:"prefetchRemainedCount"`
    PrefetchAllCount int `json:"prefetchAllCount"`
    PrefetchUsedCount int `json:"prefetchUsedCount"`
    RefreshAllCount int `json:"refreshAllCount"`
    RefreshRemainedCount int `json:"refreshRemainedCount"`
    RefreshUsedCount int `json:"refreshUsedCount"`
    TotalUserDomainQuota int `json:"totalUserDomainQuota"`
    UsedUserDomainQuota int `json:"usedUserDomainQuota"`
    RemainUserDomainQuota int `json:"remainUserDomainQuota"`
}