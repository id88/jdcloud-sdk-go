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

type AddCustomLiveStreamWatermarkTemplateRequest struct {

    core.JDCloudRequest

    /* x轴偏移量 单位：像素  */
    OffsetX int `json:"offsetX"`

    /* y轴偏移量 单位：像素  */
    OffsetY int `json:"offsetY"`

    /* 宽  */
    Width int `json:"width"`

    /* 高  */
    Height int `json:"height"`

    /* 录制模板自定义名称  */
    Template string `json:"template"`

    /* 高  */
    Url string `json:"url"`
}

/*
 * param offsetX: x轴偏移量 单位：像素 (Required)
 * param offsetY: y轴偏移量 单位：像素 (Required)
 * param width: 宽 (Required)
 * param height: 高 (Required)
 * param template: 录制模板自定义名称 (Required)
 * param url: 高 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddCustomLiveStreamWatermarkTemplateRequest(
    offsetX int,
    offsetY int,
    width int,
    height int,
    template string,
    url string,
) *AddCustomLiveStreamWatermarkTemplateRequest {

	return &AddCustomLiveStreamWatermarkTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/watermarkCustoms:template",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        OffsetX: offsetX,
        OffsetY: offsetY,
        Width: width,
        Height: height,
        Template: template,
        Url: url,
	}
}

/*
 * param offsetX: x轴偏移量 单位：像素 (Required)
 * param offsetY: y轴偏移量 单位：像素 (Required)
 * param width: 宽 (Required)
 * param height: 高 (Required)
 * param template: 录制模板自定义名称 (Required)
 * param url: 高 (Required)
 */
func NewAddCustomLiveStreamWatermarkTemplateRequestWithAllParams(
    offsetX int,
    offsetY int,
    width int,
    height int,
    template string,
    url string,
) *AddCustomLiveStreamWatermarkTemplateRequest {

    return &AddCustomLiveStreamWatermarkTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/watermarkCustoms:template",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        OffsetX: offsetX,
        OffsetY: offsetY,
        Width: width,
        Height: height,
        Template: template,
        Url: url,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddCustomLiveStreamWatermarkTemplateRequestWithoutParam() *AddCustomLiveStreamWatermarkTemplateRequest {

    return &AddCustomLiveStreamWatermarkTemplateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/watermarkCustoms:template",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param offsetX: x轴偏移量 单位：像素(Required) */
func (r *AddCustomLiveStreamWatermarkTemplateRequest) SetOffsetX(offsetX int) {
    r.OffsetX = offsetX
}

/* param offsetY: y轴偏移量 单位：像素(Required) */
func (r *AddCustomLiveStreamWatermarkTemplateRequest) SetOffsetY(offsetY int) {
    r.OffsetY = offsetY
}

/* param width: 宽(Required) */
func (r *AddCustomLiveStreamWatermarkTemplateRequest) SetWidth(width int) {
    r.Width = width
}

/* param height: 高(Required) */
func (r *AddCustomLiveStreamWatermarkTemplateRequest) SetHeight(height int) {
    r.Height = height
}

/* param template: 录制模板自定义名称(Required) */
func (r *AddCustomLiveStreamWatermarkTemplateRequest) SetTemplate(template string) {
    r.Template = template
}

/* param url: 高(Required) */
func (r *AddCustomLiveStreamWatermarkTemplateRequest) SetUrl(url string) {
    r.Url = url
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddCustomLiveStreamWatermarkTemplateRequest) GetRegionId() string {
    return ""
}

type AddCustomLiveStreamWatermarkTemplateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddCustomLiveStreamWatermarkTemplateResult `json:"result"`
}

type AddCustomLiveStreamWatermarkTemplateResult struct {
}