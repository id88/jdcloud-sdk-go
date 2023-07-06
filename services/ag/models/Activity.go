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


type Activity struct {

    /* 伸缩活动ID (Optional) */
    ActivityId string `json:"activityId"`

    /* 状态，包括成功：SUCCESS,拒绝：REJECTED,失败：FAILED,执行中：RUNNING,部分成功：WARN (Optional) */
    Status string `json:"status"`

    /* 变化后实例总数 (Optional) */
    TargetTotal int `json:"targetTotal"`

    /* 开始时间 (Optional) */
    StartTime string `json:"startTime"`

    /* 结束时间 (Optional) */
    EndTime string `json:"endTime"`

    /* 描述，展示期望实例数、最大最小数变化至多少，或展示活动应增加移出多少实例 (Optional) */
    Description string `json:"description"`

    /* 活动起因，包括但不限于下述起因：
- 手动添加实例
- 手动移出/删除实例
- 手动修改期望实例数/最大实例数/最小实例数
- 健康检查失败
- 手动执行伸缩规则
- 定时任务触发
- 报警任务触发
 (Optional) */
    Cause string `json:"cause"`

    /* 详细信息，
- 若活动执行成功，展示具体哪些实例（ID）被添加或移出
- 若活动执行失败，则展示失败原因
- 若活动执行部分成功，展示哪些实例（ID）被成功添加或移出，并展示其余资源失败原因
 (Optional) */
    Detail []ActivityDetail `json:"detail"`
}
