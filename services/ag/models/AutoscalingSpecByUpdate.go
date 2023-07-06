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


type AutoscalingSpecByUpdate struct {

    /* 伸缩组最小实例数，最大值1000。若高可用组分散策略为host或、switch，组内最小实例数及最大实例数不能大于组quota限制 (Optional) */
    MinSize *int `json:"minSize"`

    /* 伸缩组最大实例数，最大值1000。若高可用组分散策略为host或、switch，组内最小实例数及最大实例数不能大于组quota限制 (Optional) */
    MaxSize *int `json:"maxSize"`

    /* 伸缩组期望实例数 (Optional) */
    DesiredCapacity *int `json:"desiredCapacity"`

    /* 伸缩组内实例是否需要健康检查，默认是开启 (Optional) */
    HealthCheck *bool `json:"healthCheck"`

    /* 冷却时间，默认值为300（单位为秒）,范围为0-86400（24小时） (Optional) */
    CoolDownSeconds *int `json:"coolDownSeconds"`

    /* 默认值为均衡分布，Balance (Optional) */
    ScalingPolicy *string `json:"scalingPolicy"`

    /* 实例移出策略，默认值为最早创建的实例,支持：OldestResource（最早创建实例）,NewestResource（最新创建实例） (Optional) */
    RemovalPolicy *string `json:"removalPolicy"`
}
