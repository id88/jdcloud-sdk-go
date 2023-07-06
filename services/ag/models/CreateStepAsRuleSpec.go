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


type CreateStepAsRuleSpec struct {

    /* 告警任务ID，步进规则必须绑定一个告警任务  */
    AsAlarmId string `json:"asAlarmId"`

    /* 伸缩调整方式，取值范围：[`Number`,`Percentage`,`Total`]
- `Number`：增加或减少指定数量的实例
- `Percentage`：增加或减少指定百分比的实例
- `Total`：将当前伸缩组的实例数量调整到指定数量
  */
    AdjustmentType string `json:"adjustmentType"`

    /* 步进调整策略数组  */
    StepAdjustments []CreateStepAdjustmentSpec `json:"stepAdjustments"`
}
