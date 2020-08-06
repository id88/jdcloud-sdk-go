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

type QueryHttp2Request struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryHttp2Request(
    domain string,
) *QueryHttp2Request {

	return &QueryHttp2Request{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain/{domain}/queryHttp2",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 */
func NewQueryHttp2RequestWithAllParams(
    domain string,
) *QueryHttp2Request {

    return &QueryHttp2Request{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/queryHttp2",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryHttp2RequestWithoutParam() *QueryHttp2Request {

    return &QueryHttp2Request{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/queryHttp2",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *QueryHttp2Request) SetDomain(domain string) {
    r.Domain = domain
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryHttp2Request) GetRegionId() string {
    return ""
}

type QueryHttp2Response struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryHttp2Result `json:"result"`
}

type QueryHttp2Result struct {
    Domain string `json:"domain"`
    Status string `json:"status"`
}