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


type VideoObject struct {

    /* 视频ID (Optional) */
    Id string `json:"id"`

    /* 视频名称 (Optional) */
    Name string `json:"name"`

    /* 视频描述 (Optional) */
    Description string `json:"description"`

    /* 封面图地址 (Optional) */
    CoverUrl string `json:"coverUrl"`

    /* 视频状态。取值范围：
  transcoding - 转码中
  transcode_failed - 转码失败
  normal - 正常
  uploaded - 上传完成（未转码）
 (Optional) */
    Status string `json:"status"`

    /* 文件大小，单位为 Byte (Optional) */
    FileSize int64 `json:"fileSize"`

    /* 文件MD5校验和 (Optional) */
    Checksum string `json:"checksum"`

    /* 视频时长 (Optional) */
    Duration int64 `json:"duration"`

    /* 标签集合 (Optional) */
    Tags []string `json:"tags"`

    /* 分类ID (Optional) */
    CategoryId int64 `json:"categoryId"`

    /* 分类名称 (Optional) */
    CategoryName string `json:"categoryName"`

    /* 转码截图 (Optional) */
    Snapshots []Snapshot `json:"snapshots"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 修改时间 (Optional) */
    UpdateTime string `json:"updateTime"`
}
