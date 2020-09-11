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
    cloudauth "github.com/jdcloud-api/jdcloud-sdk-go/services/cloudauth/models"
)

type CheckLegalPersonRequest struct {

    core.JDCloudRequest

    /*   */
    LegalPersonSpec *cloudauth.LegalPersonSpec `json:"legalPersonSpec"`
}

/*
 * param legalPersonSpec:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCheckLegalPersonRequest(
    legalPersonSpec *cloudauth.LegalPersonSpec,
) *CheckLegalPersonRequest {

	return &CheckLegalPersonRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/company:legalPerson",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        LegalPersonSpec: legalPersonSpec,
	}
}

/*
 * param legalPersonSpec:  (Required)
 */
func NewCheckLegalPersonRequestWithAllParams(
    legalPersonSpec *cloudauth.LegalPersonSpec,
) *CheckLegalPersonRequest {

    return &CheckLegalPersonRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/company:legalPerson",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        LegalPersonSpec: legalPersonSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCheckLegalPersonRequestWithoutParam() *CheckLegalPersonRequest {

    return &CheckLegalPersonRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/company:legalPerson",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param legalPersonSpec: (Required) */
func (r *CheckLegalPersonRequest) SetLegalPersonSpec(legalPersonSpec *cloudauth.LegalPersonSpec) {
    r.LegalPersonSpec = legalPersonSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CheckLegalPersonRequest) GetRegionId() string {
    return ""
}

type CheckLegalPersonResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CheckLegalPersonResult `json:"result"`
}

type CheckLegalPersonResult struct {
    Success bool `json:"success"`
    HasException bool `json:"hasException"`
    Code string `json:"code"`
    Message string `json:"message"`
    Detail string `json:"detail"`
}