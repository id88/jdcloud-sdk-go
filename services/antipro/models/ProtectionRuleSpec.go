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


type ProtectionRuleSpec struct {

    /* 被防护 IP, 缺省时修改防护包实例防护规则 (Optional) */
    Ip *string `json:"ip"`

    /* 防护规则类型, 修改 ip 防护规则时必传, 0: 设置 ip 使用防护包规则, 1: 设置 IP 使用自定义规则 (Optional) */
    Type *int `json:"type"`

    /* 清洗触发值 bps, 修改实例防护规则或自定义 IP 防护规则时必传 (Optional) */
    CleanThresholdBps *int64 `json:"cleanThresholdBps"`

    /* 清洗触发值 pps, 修改实例防护规则或自定义 IP 防护规则时必传 (Optional) */
    CleanThresholdPps *int64 `json:"cleanThresholdPps"`

    /* 虚假源, 0: 关闭, 1: 开启, 修改实例防护规则或自定义 IP 防护规则时必传 (Optional) */
    SpoofIpEnable *int `json:"spoofIpEnable"`

    /* 源新建连接限速, 0: 关闭, 1: 开启, 修改实例防护规则或自定义 IP 防护规则时必传 (Optional) */
    SrcNewConnLimitEnable *int `json:"srcNewConnLimitEnable"`

    /* 源新建连接速率, 修改实例防护规则或自定义 IP 防护规则时必传 (Optional) */
    SrcNewConnLimitValue *int64 `json:"srcNewConnLimitValue"`

    /* 目的新建连接, 0: 关闭, 1: 开启, 修改实例防护规则或自定义 IP 防护规则时必传 (Optional) */
    DstNewConnLimitEnable *int `json:"dstNewConnLimitEnable"`

    /* 目的新建连接速率, 修改实例防护规则或自定义 IP 防护规则时必传 (Optional) */
    DstNewConnLimitValue *int64 `json:"dstNewConnLimitValue"`

    /* 报文最小长度, 取值范围 [1, datagramRangeMax) (Optional) */
    DatagramRangeMin *int64 `json:"datagramRangeMin"`

    /* 报文最大长度, 取值范围 (datagramRangeMin, 1518] (Optional) */
    DatagramRangeMax *int64 `json:"datagramRangeMax"`

    /* geo 拦截地域编码列表. 查询 <a href="http://docs.jdcloud.com/anti-ddos-protection-package/api/describegeoareas">describeGeoAreas</a> 接口获取可设置的地域编码列表 (Optional) */
    GeoBlackList []string `json:"geoBlackList"`
}
