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

type AddLiveRestartDomainRequest struct {

    core.JDCloudRequest

    /* 直播的播放域名
- 回看域名所对应的原播放域名,新建的回看域名将绑定到此播放域名下
  */
    PlayDomain string `json:"playDomain"`

    /* 直播回看域名
- 直播域名必须已经备案完成
  */
    RestartDomain string `json:"restartDomain"`
}

/*
 * param playDomain: 直播的播放域名
- 回看域名所对应的原播放域名,新建的回看域名将绑定到此播放域名下
 (Required)
 * param restartDomain: 直播回看域名
- 直播域名必须已经备案完成
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddLiveRestartDomainRequest(
    playDomain string,
    restartDomain string,
) *AddLiveRestartDomainRequest {

	return &AddLiveRestartDomainRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domains:restart",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        PlayDomain: playDomain,
        RestartDomain: restartDomain,
	}
}

/*
 * param playDomain: 直播的播放域名
- 回看域名所对应的原播放域名,新建的回看域名将绑定到此播放域名下
 (Required)
 * param restartDomain: 直播回看域名
- 直播域名必须已经备案完成
 (Required)
 */
func NewAddLiveRestartDomainRequestWithAllParams(
    playDomain string,
    restartDomain string,
) *AddLiveRestartDomainRequest {

    return &AddLiveRestartDomainRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domains:restart",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        PlayDomain: playDomain,
        RestartDomain: restartDomain,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddLiveRestartDomainRequestWithoutParam() *AddLiveRestartDomainRequest {

    return &AddLiveRestartDomainRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domains:restart",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param playDomain: 直播的播放域名
- 回看域名所对应的原播放域名,新建的回看域名将绑定到此播放域名下
(Required) */
func (r *AddLiveRestartDomainRequest) SetPlayDomain(playDomain string) {
    r.PlayDomain = playDomain
}

/* param restartDomain: 直播回看域名
- 直播域名必须已经备案完成
(Required) */
func (r *AddLiveRestartDomainRequest) SetRestartDomain(restartDomain string) {
    r.RestartDomain = restartDomain
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddLiveRestartDomainRequest) GetRegionId() string {
    return ""
}

type AddLiveRestartDomainResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddLiveRestartDomainResult `json:"result"`
}

type AddLiveRestartDomainResult struct {
}