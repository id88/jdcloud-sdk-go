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


type ServiceIpConfigItem struct {

    /* 高防IP (Optional) */
    ServiceIp string `json:"serviceIp"`

    /* 安全状态. <br>- SAFE: 安全<br>- CLEANING: 清洗中<br>- BLOCKING: 封禁中 (Optional) */
    SecurityStatus string `json:"securityStatus"`

    /* 使用状态, 1: 使用中, 0: 备用候选 (Optional) */
    UseStatus int `json:"useStatus"`
}
