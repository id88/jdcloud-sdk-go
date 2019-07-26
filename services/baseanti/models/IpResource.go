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


type IpResource struct {

    /* 公网 IP 所在区域编码 (Optional) */
    Region string `json:"region"`

    /* 公网 IP 类型或绑定资源类型:
  0: 未知类型,
  1: 弹性公网 IP(IP 为弹性公网 IP, 绑定资源类型未知),
  10: 弹性公网 IP(IP 为弹性公网 IP, 但未绑定资源),
  11: 云主机,
  12: 负载均衡,
  13: 原生容器实例,
  14: 原生容器 Pod,
  2: 云物理服务器,
 (Optional) */
    ResourceType int `json:"resourceType"`

    /* 公网 IP 地址 (Optional) */
    Ip string `json:"ip"`

    /* 带宽上限, 单位 Mbps (Optional) */
    Bandwidth int64 `json:"bandwidth"`

    /* 每秒请求流量 (Optional) */
    CleanThresholdBps int64 `json:"cleanThresholdBps"`

    /* 每秒报文请求数 (Optional) */
    CleanThresholdPps int64 `json:"cleanThresholdPps"`

    /* 黑洞阈值 (Optional) */
    BlackHoleThreshold int64 `json:"blackHoleThreshold"`

    /* 绑定防护包 ID, 为空字符串时表示未绑定防护包 (Optional) */
    InstanceId string `json:"instanceId"`

    /* 绑定防护包名称, 为空字符串时表示未绑定防护包 (Optional) */
    InstanceName string `json:"instanceName"`

    /* 套餐类型, 1: 独享 IP, 2: 共享 IP, 为 0 时未绑定防护包 (Optional) */
    InstanceType int `json:"instanceType"`

    /* 安全状态, 0: 安全, 1: 清洗, 2: 黑洞 (Optional) */
    SafeStatus int `json:"safeStatus"`
}
