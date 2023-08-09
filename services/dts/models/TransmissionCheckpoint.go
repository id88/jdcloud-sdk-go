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


type TransmissionCheckpoint struct {

    /* 数据传输任务起始位点 MySQL GTID，当源库为MySQL，且为增量传输时，可传递此参数（仅限源为MySQL） (Optional) */
    Gtid *string `json:"gtid"`

    /* 数据传输任务起始位点，MASTER_LOG_FILE，当源库为MySQL，且为增量传输时，可选择GTID或者LOG_FILE_POS（仅限源为MySQL） (Optional) */
    LogFile *string `json:"logFile"`

    /* MySQL 数据传输任务开始位点，MASTER_LOG_POS（仅限源为MySQL） (Optional) */
    LogPos *int `json:"logPos"`

    /* 数据传输任务延迟（仅限源为MySQL） (Optional) */
    Delay *int `json:"delay"`

    /* MongoDB checkpoint (Optional) */
    MongoDbCheckpoints []MongoDBCheckpoint `json:"mongoDbCheckpoints"`

    /* TiCDC checkpoint tso (Optional) */
    CheckpointTSO *int `json:"checkpointTSO"`

    /* TiCDC checkpoint timestamp (Optional) */
    CheckpointTime *string `json:"checkpointTime"`
}
