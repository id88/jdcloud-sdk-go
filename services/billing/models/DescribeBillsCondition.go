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


type DescribeBillsCondition struct {

    /* 业务线 (Optional) */
    AppCode string `json:"appCode"`

    /* 产品 (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* 数据来源：1-统一计费, 2-自计费, 10-物联网, 11-视频云, 12-CDN, 13-PCDN, 14-IDC, 15-通信云 (Optional) */
    DataSource int `json:"dataSource"`

    /* 用户pin  */
    Pin string `json:"pin"`

    /* 开始时间（按出账时间查询，开始结束时间必须在同一个月内）  */
    StartTime string `json:"startTime"`

    /* 结束时间（按出账时间查询，开始结束时间必须在同一个月内）  */
    EndTime string `json:"endTime"`

    /* 计费类型： 1、按配置 2、按用量 3、包年包月 4、按次（一次性） (Optional) */
    BillingType int `json:"billingType"`

    /* 地域 (Optional) */
    Region string `json:"region"`

    /* 资源ID列表 (Optional) */
    ResourceIdList []string `json:"resourceIdList"`

    /* 页码,默认为1 (Optional) */
    PageIndex int `json:"pageIndex"`

    /* 分页大小 (Optional) */
    PageSize int `json:"pageSize"`
}
