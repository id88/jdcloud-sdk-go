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


type Flavor struct {

    /* 规格代码,如tidb.s1.xlarge (Optional) */
    InstanceClass string `json:"instanceClass"`

    /* cpu核数 (Optional) */
    Cpu int `json:"cpu"`

    /* 内存大小，单位GB (Optional) */
    MemoryGB int `json:"memoryGB"`

    /* 默认存储规格，单位GB (Optional) */
    DefaultStorageGB int `json:"defaultStorageGB"`

    /* 该规格支持的存储空间，单位GB (Optional) */
    StorageGB []int `json:"storageGB"`
}
