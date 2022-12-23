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


type RecordListResp struct {

    /* 当前是第几页，page (Optional) */
    Number int `json:"number"`

    /* 是否是最后一页 (Optional) */
    Last bool `json:"last"`

    /* 当前页数量 (Optional) */
    NumberOfElements int `json:"numberOfElements"`

    /* 每页数量 (Optional) */
    Size int `json:"size"`

    /* 总共多少页 (Optional) */
    TotalPages int `json:"totalPages"`

    /* 是否第一页 (Optional) */
    First bool `json:"first"`

    /* 是否为空 (Optional) */
    Empty bool `json:"empty"`

    /* 总数目 (Optional) */
    TotalElements int `json:"totalElements"`

    /*  (Optional) */
    Content []interface{} `json:"content"`
}
