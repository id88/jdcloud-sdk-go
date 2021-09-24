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
    cdn "github.com/jdcloud-api/jdcloud-sdk-go/services/cdn/models"
)

type QueryHttpHeaderRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* header生效节点，0边缘，1回源，2两者都 (Optional) */
    EdgeType *int `json:"edgeType"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryHttpHeaderRequest(
    domain string,
) *QueryHttpHeaderRequest {

	return &QueryHttpHeaderRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain/{domain}/httpHeader",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param edgeType: header生效节点，0边缘，1回源，2两者都 (Optional)
 */
func NewQueryHttpHeaderRequestWithAllParams(
    domain string,
    edgeType *int,
) *QueryHttpHeaderRequest {

    return &QueryHttpHeaderRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/httpHeader",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        EdgeType: edgeType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryHttpHeaderRequestWithoutParam() *QueryHttpHeaderRequest {

    return &QueryHttpHeaderRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/httpHeader",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *QueryHttpHeaderRequest) SetDomain(domain string) {
    r.Domain = domain
}

/* param edgeType: header生效节点，0边缘，1回源，2两者都(Optional) */
func (r *QueryHttpHeaderRequest) SetEdgeType(edgeType int) {
    r.EdgeType = &edgeType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryHttpHeaderRequest) GetRegionId() string {
    return ""
}

type QueryHttpHeaderResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryHttpHeaderResult `json:"result"`
}

type QueryHttpHeaderResult struct {
    Domain string `json:"domain"`
    Headers []cdn.QueryHttpHeaderResp `json:"headers"`
}