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


type CreateContractSpec struct {

    /* 合同名称 (Optional) */
    FileName string `json:"fileName"`

    /* 签署人信息 (Optional) */
    Users []UserInfo `json:"users"`

    /* 签署类型（0 坐标 1 关键字） (Optional) */
    SignPositionType int `json:"signPositionType"`

    /* 签署顺序类型（0 顺序签署，1 无序签署） (Optional) */
    OrderSign int `json:"orderSign"`

    /* 合同文件（base64） (Optional) */
    ContractBase64 string `json:"contractBase64"`

    /* 合同编号 (Optional) */
    ContractNo string `json:"contractNo"`

    /* 合同子编号 (Optional) */
    ContractNoSub string `json:"contractNoSub"`

    /* 备注 (Optional) */
    Note string `json:"note"`

    /* 合同所属用户Id (Optional) */
    OwnerCloudUserId string `json:"ownerCloudUserId"`
}
