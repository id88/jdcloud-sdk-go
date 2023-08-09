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


type Topology struct {

    /* 源数据库引擎名称 (Optional) */
    SourceEngine string `json:"sourceEngine"`

    /* 源实例类型：CloudInstance：RDS MySQL(MariaDB,Percona)、RDS SQL Server、RDS PostgreSQL、MongoDB、Kafka、Elasticsearch；ECS：云主机自建数据库；OTHER：有公网IP的自建数据库；EXPRESS：通过专线接入的数据库 (Optional) */
    SourceInstanceType string `json:"sourceInstanceType"`

    /* 目标数据源引擎名称、目标数据源实例类型 (Optional) */
    Destination []DestinationEndpointRule `json:"destination"`
}
