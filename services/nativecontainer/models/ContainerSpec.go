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

import charge "github.com/jdcloud-api/jdcloud-sdk-go/services/charge/models"

type ContainerSpec struct {

    /* 实例类型；参考[文档](https://www.jdcloud.com/help/detail/1992/isCatalog/1)  */
    InstanceType string `json:"instanceType"`

    /* 容器所属可用区，指定agId时非必传<br> 容器、已有云盘的az必须相同，且包含在AG中  */
    Az string `json:"az"`

    /* 容器名称，不可为空，只支持中文、数字、大小写字母、英文下划线“_”及中划线“-”，且不能超过32字符  */
    Name string `json:"name"`

    /* 域名和IP映射的信息；</br> 最大10个alias (Optional) */
    HostAliases []HostAliasSpec `json:"hostAliases"`

    /* 主机名，规范请参考说明文档；默认容器ID (Optional) */
    Hostname *string `json:"hostname"`

    /* 容器执行命令，如果不指定默认是docker镜像的ENTRYPOINT (Optional) */
    Command []string `json:"command"`

    /* 容器执行命令的参数，如果不指定默认是docker镜像的CMD (Optional) */
    Args []string `json:"args"`

    /* 容器执行的环境变量；如果和镜像中的环境变量Key相同，会覆盖镜像中的值；</br> 最大100对 (Optional) */
    Envs []EnvVar `json:"envs"`

    /* 镜像名称 </br> 1. Docker Hub官方镜像通过类似nginx, mysql/mysql-server的名字指定 </br> </br> repository长度最大256个字符，tag最大128个字符，registry最大255个字符  */
    Image string `json:"image"`

    /* 镜像仓库认证信息；使用Docker Hub和京东云CR的镜像不需要secret (Optional) */
    Secret *string `json:"secret"`

    /* 容器是否分配tty。默认不分配 (Optional) */
    Tty *bool `json:"tty"`

    /* 容器的工作目录。如果不指定，默认是根目录（/）；必须是绝对路径 (Optional) */
    WorkingDir *string `json:"workingDir"`

    /* 根Volume信息  */
    RootVolume *VolumeMountSpec `json:"rootVolume"`

    /* 挂载的数据Volume信息；最多7个 (Optional) */
    DataVolumes []VolumeMountSpec `json:"dataVolumes"`

    /* 主网卡主IP关联的弹性IP规格 (Optional) */
    ElasticIp *ElasticIpSpec `json:"elasticIp"`

    /* 主网卡配置信息  */
    PrimaryNetworkInterface *ContainerNetworkInterfaceAttachmentSpec `json:"primaryNetworkInterface"`

    /* 容器日志配置信息；默认会在本地分配10MB的存储空间 (Optional) */
    LogConfiguration *LogConfiguration `json:"logConfiguration"`

    /* 容器描述 (Optional) */
    Description *string `json:"description"`

    /* 计费配置；如不指定，默认计费类型是后付费-按使用时常付费 (Optional) */
    Charge *charge.ChargeSpec `json:"charge"`

    /* 用户普通标签集合 (Optional) */
    UserTags []Tag `json:"userTags"`

    /* 资源组ID (Optional) */
    ResourceGroupId *string `json:"resourceGroupId"`
}
