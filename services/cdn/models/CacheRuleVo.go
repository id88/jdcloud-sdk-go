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


type CacheRuleVo struct {

    /* 此条配置的权重值, 取值范围为1-10,1最大 (Optional) */
    Weight int `json:"weight"`

    /* 缓存时间,单位秒 (Optional) */
    Ttl int64 `json:"ttl"`

    /* 规则内容。其他类型只能以/或者.开头，如/a/b或.jpg (Optional) */
    Content string `json:"content"`

    /* 缓存方式：0、不缓存，1自定义 (Optional) */
    CacheType int `json:"cacheType"`
}
