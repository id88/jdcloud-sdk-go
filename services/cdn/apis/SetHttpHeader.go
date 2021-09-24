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

type SetHttpHeaderRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* 0表示header在边缘生效，1表示header回源生效，2表示在边缘和回源都生效，该字段不传时默认header在边缘和回源都生效 (Optional) */
    EdgeType *int `json:"edgeType"`

    /* header类型[resp,req],resp：配置响应头，req：配置请求头 (Optional) */
    HeaderType *string `json:"headerType"`

    /* header名，例如：Content-Disposition，可自定义，长度不能超过256个字符，不能包含中文字符，不能包含$和_，不支持设置如下头名：["Content-Length","Date","Host","Content-Encoding","If-Modified-Since","If-Range","Content-Type","Transfer-Encoding","Cache-Control","Last-Modified","Connection", "Content-Range","ETag","Age","Authentication-Info","Proxy-Authenticate","Retry-After","Set-Cookie","Vary","Content-Location","Meter","Allow","Error","X-Trace", "Proxy-Connection"] (Optional) */
    HeaderName *string `json:"headerName"`

    /* header值，不能包含($,_,#)，不能超过256个字符 (Optional) */
    HeaderValue *string `json:"headerValue"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetHttpHeaderRequest(
    domain string,
) *SetHttpHeaderRequest {

	return &SetHttpHeaderRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain/{domain}/httpHeader",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param edgeType: 0表示header在边缘生效，1表示header回源生效，2表示在边缘和回源都生效，该字段不传时默认header在边缘和回源都生效 (Optional)
 * param headerType: header类型[resp,req],resp：配置响应头，req：配置请求头 (Optional)
 * param headerName: header名，例如：Content-Disposition，可自定义，长度不能超过256个字符，不能包含中文字符，不能包含$和_，不支持设置如下头名：["Content-Length","Date","Host","Content-Encoding","If-Modified-Since","If-Range","Content-Type","Transfer-Encoding","Cache-Control","Last-Modified","Connection", "Content-Range","ETag","Age","Authentication-Info","Proxy-Authenticate","Retry-After","Set-Cookie","Vary","Content-Location","Meter","Allow","Error","X-Trace", "Proxy-Connection"] (Optional)
 * param headerValue: header值，不能包含($,_,#)，不能超过256个字符 (Optional)
 */
func NewSetHttpHeaderRequestWithAllParams(
    domain string,
    edgeType *int,
    headerType *string,
    headerName *string,
    headerValue *string,
) *SetHttpHeaderRequest {

    return &SetHttpHeaderRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/httpHeader",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        EdgeType: edgeType,
        HeaderType: headerType,
        HeaderName: headerName,
        HeaderValue: headerValue,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetHttpHeaderRequestWithoutParam() *SetHttpHeaderRequest {

    return &SetHttpHeaderRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/httpHeader",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *SetHttpHeaderRequest) SetDomain(domain string) {
    r.Domain = domain
}

/* param edgeType: 0表示header在边缘生效，1表示header回源生效，2表示在边缘和回源都生效，该字段不传时默认header在边缘和回源都生效(Optional) */
func (r *SetHttpHeaderRequest) SetEdgeType(edgeType int) {
    r.EdgeType = &edgeType
}

/* param headerType: header类型[resp,req],resp：配置响应头，req：配置请求头(Optional) */
func (r *SetHttpHeaderRequest) SetHeaderType(headerType string) {
    r.HeaderType = &headerType
}

/* param headerName: header名，例如：Content-Disposition，可自定义，长度不能超过256个字符，不能包含中文字符，不能包含$和_，不支持设置如下头名：["Content-Length","Date","Host","Content-Encoding","If-Modified-Since","If-Range","Content-Type","Transfer-Encoding","Cache-Control","Last-Modified","Connection", "Content-Range","ETag","Age","Authentication-Info","Proxy-Authenticate","Retry-After","Set-Cookie","Vary","Content-Location","Meter","Allow","Error","X-Trace", "Proxy-Connection"](Optional) */
func (r *SetHttpHeaderRequest) SetHeaderName(headerName string) {
    r.HeaderName = &headerName
}

/* param headerValue: header值，不能包含($,_,#)，不能超过256个字符(Optional) */
func (r *SetHttpHeaderRequest) SetHeaderValue(headerValue string) {
    r.HeaderValue = &headerValue
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetHttpHeaderRequest) GetRegionId() string {
    return ""
}

type SetHttpHeaderResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetHttpHeaderResult `json:"result"`
}

type SetHttpHeaderResult struct {
}