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

package client

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    nativecontainer "github.com/jdcloud-api/jdcloud-sdk-go/services/nativecontainer/apis"
    "encoding/json"
    "errors"
)

type NativecontainerClient struct {
    core.JDCloudClient
}

func NewNativecontainerClient(credential *core.Credential) *NativecontainerClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("nativecontainer.jdcloud-api.com")

    return &NativecontainerClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "nativecontainer",
            Revision:    "1.0.1",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *NativecontainerClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *NativecontainerClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 创建一个 secret，用于存放镜像仓库机密相关信息。
 */
func (c *NativecontainerClient) CreateSecret(request *nativecontainer.CreateSecretRequest) (*nativecontainer.CreateSecretResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nativecontainer.CreateSecretResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除单个 secret
 */
func (c *NativecontainerClient) DeleteSecret(request *nativecontainer.DeleteSecretRequest) (*nativecontainer.DeleteSecretResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nativecontainer.DeleteSecretResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询一台原生容器的详细信息
 */
func (c *NativecontainerClient) DescribeContainer(request *nativecontainer.DescribeContainerRequest) (*nativecontainer.DescribeContainerResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nativecontainer.DescribeContainerResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 容器状态必须为 stopped、running 或 error状态。 <br>
按量付费的实例，如不主动删除将一直运行，不再使用的实例，可通过本接口主动停用。<br>
只能支持主动删除按量计费类型的实例。包年包月过期的容器也可以删除，其它的情况还请发工单系统。计费状态异常的容器无法删除。
 */
func (c *NativecontainerClient) DeleteContainer(request *nativecontainer.DeleteContainerRequest) (*nativecontainer.DeleteContainerResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nativecontainer.DeleteContainerResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询资源的配额，支持：原生容器 pod 和 secret.
 */
func (c *NativecontainerClient) DescribeQuota(request *nativecontainer.DescribeQuotaRequest) (*nativecontainer.DescribeQuotaResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nativecontainer.DescribeQuotaResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询单个 secret 详情
 */
func (c *NativecontainerClient) DescribeSecret(request *nativecontainer.DescribeSecretRequest) (*nativecontainer.DescribeSecretResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nativecontainer.DescribeSecretResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 停止处于运行状态的单个实例，处于任务执行中的容器无法启动。
 */
func (c *NativecontainerClient) StopContainer(request *nativecontainer.StopContainerRequest) (*nativecontainer.StopContainerResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nativecontainer.StopContainerResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 容器解绑公网 IP，解绑的是主网卡、主内网 IP 对应的弹性 IP.
 */
func (c *NativecontainerClient) DisassociateElasticIp(request *nativecontainer.DisassociateElasticIpRequest) (*nativecontainer.DisassociateElasticIpResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nativecontainer.DisassociateElasticIpResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询单个容器日志
 */
func (c *NativecontainerClient) GetLogs(request *nativecontainer.GetLogsRequest) (*nativecontainer.GetLogsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nativecontainer.GetLogsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 重置原生容器，对已有原生容器使用新的镜像重置。
原容器 id 不变，不涉及计费变动，暂不支持修改实例类型，不会改变原生容器所在的物理节点，也不支持修改已经使用的系统盘和数据盘以及网络相关参数。
- 镜像
    - 容器的镜像通过镜像名称来确定
    - nginx:tag 或 mysql/mysql-server:tag 这样命名的镜像表示 docker hub 官方镜像
    - container-registry/image:tag 这样命名的镜像表示私有仓储的镜像
    - 私有仓储必须兼容 docker registry 认证机制，并通过 secret 来保存机密信息
- 其他
    - rebuild 之前容器必须处于关闭状态
    - rebuild 完成后，容器仍为关闭状态
 */
func (c *NativecontainerClient) RebuildContainer(request *nativecontainer.RebuildContainerRequest) (*nativecontainer.RebuildContainerResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nativecontainer.RebuildContainerResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 容器绑定弹性公网 IP，绑定的是主网卡、主内网IP对应的弹性IP. <br>
一台云主机只能绑定一个弹性公网 IP(主网卡)，若主网卡已存在弹性公网IP，会返回错误。<br>
如果是黑名单中的用户，会返回错误。
 */
func (c *NativecontainerClient) AssociateElasticIp(request *nativecontainer.AssociateElasticIpRequest) (*nativecontainer.AssociateElasticIpResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nativecontainer.AssociateElasticIpResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建exec
 */
func (c *NativecontainerClient) ExecCreate(request *nativecontainer.ExecCreateRequest) (*nativecontainer.ExecCreateResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nativecontainer.ExecCreateResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取exec退出码
 */
func (c *NativecontainerClient) ExecGetExitCode(request *nativecontainer.ExecGetExitCodeRequest) (*nativecontainer.ExecGetExitCodeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nativecontainer.ExecGetExitCodeResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建一台或多台指定配置容器。
- 创建容器需要通过实名认证
- 镜像
    - 容器的镜像通过镜像名称来确定
    - nginx:tag 或 mysql/mysql-server:tag 这样命名的镜像表示 docker hub 官方镜像
    - container-registry/image:tag 这样命名的镜像表示私有仓储的镜像
    - 私有仓储必须兼容 docker registry 认证机制，并通过 secret 来保存机密信息
- hostname 规范
    - 支持两种方式：以标签方式书写或以完整主机名方式书写
    - 标签规范
        - 0-9，a-z(不分大小写)和 -（减号），其他的都是无效的字符串
        - 不能以减号开始，也不能以减号结尾
        - 最小1个字符，最大63个字符
    - 完整的主机名由一系列标签与点连接组成
        - 标签与标签之间使用“.”(点)进行连接
        - 不能以“.”(点)开始，也不能以“.”(点)结尾
        - 整个主机名（包括标签以及分隔点“.”）最多有63个ASCII字符
- 网络配置
    - 指定主网卡配置信息
        - 必须指定一个子网
        - 一台云主机创建时必须指定一个安全组，至多指定 5 个安全组
        - 可以指定 elasticIp 规格来约束创建的弹性 IP，带宽取值范围 [1-200]Mbps，步进 1Mbps
        - 可以指定网卡的主 IP(primaryIpAddress)，该 IP 需要在子网 IP 范围内且未被占用，指定子网 IP 时 maxCount 只能为1
        - 安全组 securityGroup 需与子网 Subnet 在同一个私有网络 VPC 内
        - 主网卡 deviceIndex 设置为 1
- 存储
    - volume 分为 root volume 和 data volume，root volume 的挂载目录是 /，data volume 的挂载目录可以随意指定
    - volume 的底层存储介质当前只支持 cloud 类别，也就是云硬盘
    - 系统盘
        - 云硬盘类型可以选择 ssd、premium-hdd
        - 磁盘大小
            - ssd：范围 [10, 100]GB，步长为 10G
            - premium-hdd：范围 [20, 1000]GB，步长为 10G
        - 自动删除
            - 云盘默认跟随容器实例自动删除，如果是包年包月的数据盘或共享型数据盘，此参数不生效
        - 可以选择已存在的云硬盘
    - 数据盘
        - 云硬盘类型可以选择 ssd、premium-hdd
        - 磁盘大小
            - ssd：范围[20,1000]GB，步长为10G
            - premium-hdd：范围[20,3000]GB，步长为10G
        - 自动删除
            - 默认自动删除
        - 可以选择已存在的云硬盘
        - 单个容器最多可以挂载 7 个 data volume
- 计费
  - 弹性IP的计费模式，如果选择按用量类型可以单独设置，其它计费模式都以主机为准
  - 云硬盘的计费模式以主机为准
- 容器日志
    - 默认在本地分配10MB的存储空间，自动 rotate
- 其他
    - 创建完成后，容器状态为running
    - maxCount 为最大努力，不保证一定能达到 maxCount
 */
func (c *NativecontainerClient) CreateContainers(request *nativecontainer.CreateContainersRequest) (*nativecontainer.CreateContainersResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nativecontainer.CreateContainersResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 启动处于关闭状态的单个容器，处在任务执行中的容器无法启动。<br>
容器实例或其绑定的云盘已欠费时，容器将无法正常启动。<br>
 */
func (c *NativecontainerClient) StartContainer(request *nativecontainer.StartContainerRequest) (*nativecontainer.StartContainerResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nativecontainer.StartContainerResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询 secret 列表。<br> 
此接口支持分页查询，默认每页20条。
 */
func (c *NativecontainerClient) DescribeSecrets(request *nativecontainer.DescribeSecretsRequest) (*nativecontainer.DescribeSecretsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nativecontainer.DescribeSecretsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改容器的 名称 和 描述。
 */
func (c *NativecontainerClient) ModifyContainerAttribute(request *nativecontainer.ModifyContainerAttributeRequest) (*nativecontainer.ModifyContainerAttributeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nativecontainer.ModifyContainerAttributeResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 调整TTY大小
 */
func (c *NativecontainerClient) ResizeTTY(request *nativecontainer.ResizeTTYRequest) (*nativecontainer.ResizeTTYResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nativecontainer.ResizeTTYResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 批量查询原生容器的详细信息<br>
此接口支持分页查询，默认每页20条。
 */
func (c *NativecontainerClient) DescribeContainers(request *nativecontainer.DescribeContainersRequest) (*nativecontainer.DescribeContainersResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nativecontainer.DescribeContainersResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

