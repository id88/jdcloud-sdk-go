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

type CreateViTemplateRequest struct {

    core.JDCloudRequest

    /* 模板名称。长度不超过128个字符。UTF-8编码。
  */
    TemplateName string `json:"templateName"`

    /* 截图间隔 (Optional) */
    ShotInterval *int `json:"shotInterval"`

    /* 间隔类型：time, percent (Optional) */
    ShotIntervalType *string `json:"shotIntervalType"`

    /* 截图格式：png, jpg (Optional) */
    ShotFormat *string `json:"shotFormat"`

    /* 截图帧类型：any, key (Optional) */
    ShotFrameType *string `json:"shotFrameType"`

    /* 截图宽度 (Optional) */
    ShotWidth *int `json:"shotWidth"`

    /* 截图高度 (Optional) */
    ShotHeight *int `json:"shotHeight"`

    /* 截图填充类型：stretch, gauss, black, white (Optional) */
    ShotFillType *string `json:"shotFillType"`

    /* 审查配置，JSON格式  */
    InspectItems string `json:"inspectItems"`
}

/*
 * param templateName: 模板名称。长度不超过128个字符。UTF-8编码。
 (Required)
 * param inspectItems: 审查配置，JSON格式 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateViTemplateRequest(
    templateName string,
    inspectItems string,
) *CreateViTemplateRequest {

	return &CreateViTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/viTemplates",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        TemplateName: templateName,
        InspectItems: inspectItems,
	}
}

/*
 * param templateName: 模板名称。长度不超过128个字符。UTF-8编码。
 (Required)
 * param shotInterval: 截图间隔 (Optional)
 * param shotIntervalType: 间隔类型：time, percent (Optional)
 * param shotFormat: 截图格式：png, jpg (Optional)
 * param shotFrameType: 截图帧类型：any, key (Optional)
 * param shotWidth: 截图宽度 (Optional)
 * param shotHeight: 截图高度 (Optional)
 * param shotFillType: 截图填充类型：stretch, gauss, black, white (Optional)
 * param inspectItems: 审查配置，JSON格式 (Required)
 */
func NewCreateViTemplateRequestWithAllParams(
    templateName string,
    shotInterval *int,
    shotIntervalType *string,
    shotFormat *string,
    shotFrameType *string,
    shotWidth *int,
    shotHeight *int,
    shotFillType *string,
    inspectItems string,
) *CreateViTemplateRequest {

    return &CreateViTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/viTemplates",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        TemplateName: templateName,
        ShotInterval: shotInterval,
        ShotIntervalType: shotIntervalType,
        ShotFormat: shotFormat,
        ShotFrameType: shotFrameType,
        ShotWidth: shotWidth,
        ShotHeight: shotHeight,
        ShotFillType: shotFillType,
        InspectItems: inspectItems,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateViTemplateRequestWithoutParam() *CreateViTemplateRequest {

    return &CreateViTemplateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/viTemplates",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param templateName: 模板名称。长度不超过128个字符。UTF-8编码。
(Required) */
func (r *CreateViTemplateRequest) SetTemplateName(templateName string) {
    r.TemplateName = templateName
}

/* param shotInterval: 截图间隔(Optional) */
func (r *CreateViTemplateRequest) SetShotInterval(shotInterval int) {
    r.ShotInterval = &shotInterval
}

/* param shotIntervalType: 间隔类型：time, percent(Optional) */
func (r *CreateViTemplateRequest) SetShotIntervalType(shotIntervalType string) {
    r.ShotIntervalType = &shotIntervalType
}

/* param shotFormat: 截图格式：png, jpg(Optional) */
func (r *CreateViTemplateRequest) SetShotFormat(shotFormat string) {
    r.ShotFormat = &shotFormat
}

/* param shotFrameType: 截图帧类型：any, key(Optional) */
func (r *CreateViTemplateRequest) SetShotFrameType(shotFrameType string) {
    r.ShotFrameType = &shotFrameType
}

/* param shotWidth: 截图宽度(Optional) */
func (r *CreateViTemplateRequest) SetShotWidth(shotWidth int) {
    r.ShotWidth = &shotWidth
}

/* param shotHeight: 截图高度(Optional) */
func (r *CreateViTemplateRequest) SetShotHeight(shotHeight int) {
    r.ShotHeight = &shotHeight
}

/* param shotFillType: 截图填充类型：stretch, gauss, black, white(Optional) */
func (r *CreateViTemplateRequest) SetShotFillType(shotFillType string) {
    r.ShotFillType = &shotFillType
}

/* param inspectItems: 审查配置，JSON格式(Required) */
func (r *CreateViTemplateRequest) SetInspectItems(inspectItems string) {
    r.InspectItems = inspectItems
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateViTemplateRequest) GetRegionId() string {
    return ""
}

type CreateViTemplateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateViTemplateResult `json:"result"`
}

type CreateViTemplateResult struct {
    TemplateId string `json:"templateId"`
    TemplateName string `json:"templateName"`
    ShotInterval int `json:"shotInterval"`
    ShotIntervalType string `json:"shotIntervalType"`
    ShotFormat string `json:"shotFormat"`
    ShotFrameType string `json:"shotFrameType"`
    ShotWidth int `json:"shotWidth"`
    ShotHeight int `json:"shotHeight"`
    ShotFillType string `json:"shotFillType"`
    InspectItems string `json:"inspectItems"`
    CreateTime string `json:"createTime"`
    UpdateTime string `json:"updateTime"`
}