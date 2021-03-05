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

package models


type TranscodeJobSummary struct {

    /* 作业ID (Optional) */
    JobId int64 `json:"jobId"`

    /* 视频ID (Optional) */
    VideoId string `json:"videoId"`

    /* 模板ID列表。以转码模板列表方式提交的转码作业，包含此字段。 (Optional) */
    TemplateIds []int64 `json:"templateIds"`

    /* 模板组ID。以转码模板组方式提交的转码作业，包含此字段。 (Optional) */
    TemplateGroupId string `json:"templateGroupId"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 完成时间 (Optional) */
    CompleteTime string `json:"completeTime"`

    /*  (Optional) */
    Tasks []TransTask `json:"tasks"`
}
