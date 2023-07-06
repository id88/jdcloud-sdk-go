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


type CreateStepAdjustmentSpec struct {

    /* 指标值上界，取值范围：[`-9.999999e18` ~ `9.999999e18`]  */
    MetricUpperBound float64 `json:"metricUpperBound"`

    /* 指标值下界，取值范围：[`-9.999999e18` ~ `9.999999e18`]  */
    MetricLowerBound float64 `json:"metricLowerBound"`

    /* 伸缩的调整值，负数表示减少，正数表示增加
- `adjustmentType`为`Number`时，取值范围：[`-300` ~ `300`]
- `adjustmentType`为`Percentage`时，单位为百分比，取值范围：[`-100` ~ `10000`]
- `adjustmentType`为`Total`时，取值范围：[`0` ~ `100000`]
  */
    AdjustmentValue int `json:"adjustmentValue"`
}
