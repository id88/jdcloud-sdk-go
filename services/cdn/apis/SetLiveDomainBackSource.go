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

type SetLiveDomainBackSourceRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* 回源类型只能为[ips,domain]中的一种 (Optional) */
    SourceType *string `json:"sourceType"`

    /*  (Optional) */
    BackSourceType *string `json:"backSourceType"`

    /* 默认回源host (Optional) */
    DefaultSourceHost *string `json:"defaultSourceHost"`

    /*  (Optional) */
    DomainSource []cdn.DomainSourceInfo `json:"domainSource"`

    /*  (Optional) */
    IpSource []cdn.IpSourceInfo `json:"ipSource"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetLiveDomainBackSourceRequest(
    domain string,
) *SetLiveDomainBackSourceRequest {

	return &SetLiveDomainBackSourceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/liveDomain/{domain}/backSource",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param sourceType: 回源类型只能为[ips,domain]中的一种 (Optional)
 * param backSourceType:  (Optional)
 * param defaultSourceHost: 默认回源host (Optional)
 * param domainSource:  (Optional)
 * param ipSource:  (Optional)
 */
func NewSetLiveDomainBackSourceRequestWithAllParams(
    domain string,
    sourceType *string,
    backSourceType *string,
    defaultSourceHost *string,
    domainSource []cdn.DomainSourceInfo,
    ipSource []cdn.IpSourceInfo,
) *SetLiveDomainBackSourceRequest {

    return &SetLiveDomainBackSourceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveDomain/{domain}/backSource",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        SourceType: sourceType,
        BackSourceType: backSourceType,
        DefaultSourceHost: defaultSourceHost,
        DomainSource: domainSource,
        IpSource: ipSource,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetLiveDomainBackSourceRequestWithoutParam() *SetLiveDomainBackSourceRequest {

    return &SetLiveDomainBackSourceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveDomain/{domain}/backSource",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *SetLiveDomainBackSourceRequest) SetDomain(domain string) {
    r.Domain = domain
}
/* param sourceType: 回源类型只能为[ips,domain]中的一种(Optional) */
func (r *SetLiveDomainBackSourceRequest) SetSourceType(sourceType string) {
    r.SourceType = &sourceType
}
/* param backSourceType: (Optional) */
func (r *SetLiveDomainBackSourceRequest) SetBackSourceType(backSourceType string) {
    r.BackSourceType = &backSourceType
}
/* param defaultSourceHost: 默认回源host(Optional) */
func (r *SetLiveDomainBackSourceRequest) SetDefaultSourceHost(defaultSourceHost string) {
    r.DefaultSourceHost = &defaultSourceHost
}
/* param domainSource: (Optional) */
func (r *SetLiveDomainBackSourceRequest) SetDomainSource(domainSource []cdn.DomainSourceInfo) {
    r.DomainSource = domainSource
}
/* param ipSource: (Optional) */
func (r *SetLiveDomainBackSourceRequest) SetIpSource(ipSource []cdn.IpSourceInfo) {
    r.IpSource = ipSource
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetLiveDomainBackSourceRequest) GetRegionId() string {
    return ""
}

type SetLiveDomainBackSourceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetLiveDomainBackSourceResult `json:"result"`
}

type SetLiveDomainBackSourceResult struct {
}