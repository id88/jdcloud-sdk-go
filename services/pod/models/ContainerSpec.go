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


type ContainerSpec struct {

    /* 容器名称，符合DNS-1123 label规范，在一个Pod内不可重复、不支持修改  */
    Name string `json:"name"`

    /* 容器执行命令，如果不指定默认是docker镜像的ENTRYPOINT。总长度256个字符。 (Optional) */
    Command []string `json:"command"`

    /* 容器执行命令的参数，如果不指定默认是docker镜像的CMD。总长度2048个字符。 (Optional) */
    Args []string `json:"args"`

    /* 容器执行的环境变量；如果和镜像中的环境变量Key相同，会覆盖镜像中的值。数组范围：[0-100] (Optional) */
    Env []EnvSpec `json:"env"`

    /* 镜像名称 </br>
容器镜像名字。 nginx:latest。长度范围：[1-639]
1. Docker Hub官方镜像通过类似nginx, mysql/mysql-server的名字指定 </br> 
2. repository长度最大256个字符，tag最大128个字符，registry最大255个字符 </br> 
  */
    Image string `json:"image"`

    /* 镜像仓库认证信息。如果目前不传，默认选择dockerHub镜像 (Optional) */
    Secret *string `json:"secret"`

    /* 容器是否分配tty。默认不分配 (Optional) */
    Tty *bool `json:"tty"`

    /* 容器的工作目录。如果不指定，默认是根目录（/）；必须是绝对路径；长度范围：[0-1024] (Optional) */
    WorkingDir *string `json:"workingDir"`

    /* 容器存活探针配置 (Optional) */
    LivenessProbe *ProbeSpec `json:"livenessProbe"`

    /* 容器服务就绪探针配置 (Optional) */
    ReadinessProbe *ProbeSpec `json:"readinessProbe"`

    /* 容器计算资源配置 (Optional) */
    Resources *ResourceRequestsSpec `json:"resources"`

    /* 容器计算资源配置  */
    SystemDisk *CloudDiskSpec `json:"systemDisk"`

    /* 云盘挂载信息 (Optional) */
    VolumeMounts []VolumeMountSpec `json:"volumeMounts"`
}
