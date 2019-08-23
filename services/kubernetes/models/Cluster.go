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


type Cluster struct {

    /* 集群id (Optional) */
    ClusterId string `json:"clusterId"`

    /* 名称 (Optional) */
    Name string `json:"name"`

    /* 描述 (Optional) */
    Description string `json:"description"`

    /* kubernetes的版本 (Optional) */
    Version string `json:"version"`

    /* 集群所在的az (Optional) */
    Azs []string `json:"azs"`

    /* 节点组列表 (Optional) */
    NodeGroups []NodeGroup `json:"nodeGroups"`

    /* k8s的cluster的cidr (Optional) */
    ClusterCidr string `json:"clusterCidr"`

    /* 认证信息 (Optional) */
    MasterAuth MasterAuth `json:"masterAuth"`

    /* 状态  [pending,running,reconciling（升级时的状态）, deleting, deleted, error] (Optional) */
    ClusterState string `json:"clusterState"`

    /* 状态变更原因 (Optional) */
    StateMessage string `json:"stateMessage"`

    /* 更新时间 (Optional) */
    UpdateTime string `json:"updateTime"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 用户的AccessKey，插件调用open-api时的认证凭证 (Optional) */
    AccessKey string `json:"accessKey"`

    /*  (Optional) */
    BasicAuth bool `json:"basicAuth"`

    /*  (Optional) */
    ClientCertificate bool `json:"clientCertificate"`

    /* 用户访问的ip (Optional) */
    Endpoint string `json:"endpoint"`

    /* endpoint的port (Optional) */
    EndpointPort string `json:"endpointPort"`

    /* endpoint的dashboard port (Optional) */
    DashboardPort string `json:"dashboardPort"`

    /* deprecated 优先以addonsConfig中的配置为准 <br>用户是否启用集群自定义监控，true 表示开启用，false 表示未开启用 (Optional) */
    UserMetrics bool `json:"userMetrics"`

    /* 集群组件配置信息 (Optional) */
    AddonsConfig []AddonConfig `json:"addonsConfig"`

    /* 是否开启集群自动升级，true 表示开启，false 表示未开启 (Optional) */
    AutoUpgrade bool `json:"autoUpgrade"`

    /* 配置集群维护策略 (Optional) */
    MaintenanceWindow MaintenanceWindow `json:"maintenanceWindow"`

    /* 集群升级计划信息, 仅展示最新一条升级计划信息 (Optional) */
    UpgradePlan UpgradePlan `json:"upgradePlan"`

    /* 控制节点操作进度 (Optional) */
    MasterProgress MaintenanceWindow `json:"masterProgress"`
}
