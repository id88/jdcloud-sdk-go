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
    pod "github.com/jdcloud-api/jdcloud-sdk-go/services/pod/apis"
    "encoding/json"
    "errors"
)

type PodClient struct {
    core.JDCloudClient
}

func NewPodClient(credential *core.Credential) *PodClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("pod.jdcloud-api.com")

    return &PodClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "pod",
            Revision:    "1.0.5",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *PodClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *PodClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 查询资源的配额，支持：原生容器 pod 和 secret.
 */
func (c *PodClient) DescribeQuota(request *pod.DescribeQuotaRequest) (*pod.DescribeQuotaResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &pod.DescribeQuotaResponse{}
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
func (c *PodClient) DescribeSecrets(request *pod.DescribeSecretsRequest) (*pod.DescribeSecretsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &pod.DescribeSecretsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建一个 secret，用于存放镜像仓库机密相关信息。
 */
func (c *PodClient) CreateSecret(request *pod.CreateSecretRequest) (*pod.CreateSecretResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &pod.CreateSecretResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询单个 secret 详情
 */
func (c *PodClient) DescribeSecret(request *pod.DescribeSecretRequest) (*pod.DescribeSecretResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &pod.DescribeSecretResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* pod 解绑公网 IP，解绑的是主网卡、主内网 IP 对应的弹性 IP.
 */
func (c *PodClient) DisassociateElasticIp(request *pod.DisassociateElasticIpRequest) (*pod.DisassociateElasticIpResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &pod.DisassociateElasticIpResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询单个容器日志
 */
func (c *PodClient) GetContainerLogs(request *pod.GetContainerLogsRequest) (*pod.GetContainerLogsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &pod.GetContainerLogsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 设置TTY大小 */
func (c *PodClient) ResizeTTY(request *pod.ResizeTTYRequest) (*pod.ResizeTTYResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &pod.ResizeTTYResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询单个容器日志
 */
func (c *PodClient) Attach(request *pod.AttachRequest) (*pod.AttachResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &pod.AttachResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改 pod 的 名称 和 描述。
 */
func (c *PodClient) ModifyPodAttribute(request *pod.ModifyPodAttributeRequest) (*pod.ModifyPodAttributeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &pod.ModifyPodAttributeResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* pod 状态必须为 stopped、running 或 error状态。 <br>
按量付费的实例，如不主动删除将一直运行，不再使用的实例，可通过本接口主动停用。<br>
只能支持主动删除按量计费类型的实例。包年包月过期的 pod 也可以删除，其它的情况还请发工单系统。计费状态异常的容器无法删除。
 [MFA enabled] */
func (c *PodClient) DeletePod(request *pod.DeletePodRequest) (*pod.DeletePodResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &pod.DeletePodResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 对 pod 中的容器使用新的镜像进行重置，pod 需要处于关闭状态。
 */
func (c *PodClient) RebuildPod(request *pod.RebuildPodRequest) (*pod.RebuildPodResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &pod.RebuildPodResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 执行exec，此接口需要升级Http协议到WebSocket */
func (c *PodClient) ExecStart(request *pod.ExecStartRequest) (*pod.ExecStartResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &pod.ExecStartResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* pod 绑定弹性公网 IP，绑定的是主网卡、主内网IP对应的弹性IP. <br>
一个 pod 只能绑定一个弹性公网 IP(主网卡)，若主网卡已存在弹性公网IP，会返回错误。<br>
如果是黑名单中的用户，会返回错误。
 */
func (c *PodClient) AssociateElasticIp(request *pod.AssociateElasticIpRequest) (*pod.AssociateElasticIpResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &pod.AssociateElasticIpResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 启动处于关闭状态的单个 pod ，处在任务执行中的 pod 无法启动。<br>
pod 实例或其绑定的云盘已欠费时，容器将无法正常启动。<br>
 */
func (c *PodClient) StartPod(request *pod.StartPodRequest) (*pod.StartPodResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &pod.StartPodResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 批量查询 pod 的详细信息<br>
此接口支持分页查询，默认每页20条。
 */
func (c *PodClient) DescribePods(request *pod.DescribePodsRequest) (*pod.DescribePodsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &pod.DescribePodsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取exec退出码 */
func (c *PodClient) ExecGetExitCode(request *pod.ExecGetExitCodeRequest) (*pod.ExecGetExitCodeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &pod.ExecGetExitCodeResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建 exec
 */
func (c *PodClient) ExecCreate(request *pod.ExecCreateRequest) (*pod.ExecCreateResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &pod.ExecCreateResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* - 创建pod需要通过实名认证
- hostname规范
    - 支持两种方式：以标签方式书写或以完整主机名方式书写
    - 标签规范
        - 0-9，a-z(不分大小写)和-（减号），其他的都是无效的字符串
        - 不能以减号开始，也不能以减号结尾
        - 最小1个字符，最大63个字符
    - 完整的主机名由一系列标签与点连接组成
        - 标签与标签之间使用“.”(点)进行连接
        - 不能以“.”(点)开始，也不能以“.”(点)结尾
        - 整个主机名（包括标签以及分隔点“.”）最多有63个ASCII字符
    - 正则：`^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9])(.([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9]))*$`
- 网络配置
    - 指定主网卡配置信息
        - 必须指定subnetId
        - 可以指定elasticIp规格来约束创建的弹性IP，带宽取值范围[1-100]Mbps，步进1Mbps
        - 可以指定网卡的主IP(primaryIpAddress)和辅助IP(secondaryIpAddresses)，此时maxCount只能为1
        - 可以设置网卡的自动删除autoDelete属性，指明是否删除实例时自动删除网卡
        - 安全组securityGroup需与子网Subnet在同一个私有网络VPC内
        - 一个 pod 创建时必须指定一个安全组，至多指定5个安全组
        - 主网卡deviceIndex设置为1
- 存储
    - volume分为root volume和data volume，root volume的挂载目录是/，data volume的挂载目录可以随意指定
    - volume的底层存储介质当前只支持cloud类别，也就是云硬盘
    - root volume
        - root volume只能是cloud类别
        - 云硬盘类型可以选择ssd、premium-hdd
        - 磁盘大小
            - ssd：范围[10,100]GB，步长为10G
            - premium-hdd：范围[10,100]GB，步长为10G
        - 自动删除
            - 默认自动删除
        - 可以选择已存在的云硬盘
    - data volume
        - data volume当前只能选择cloud类别
        - 云硬盘类型可以选择ssd、premium-hdd
        - 磁盘大小
            - ssd：范围[20,1000]GB，步长为10G
            - premium-hdd：范围[20,3000]GB，步长为10G
        - 自动删除
            - 默认自动删除
        - 可以选择已存在的云硬盘
        - 可以从快照创建磁盘
- pod 容器日志
    - default：默认在本地分配10MB的存储空间，自动rotate
- 其他
    - 创建完成后，pod 状态为running
    - maxCount为最大努力，不保证一定能达到maxCount
 */
func (c *PodClient) CreatePods(request *pod.CreatePodsRequest) (*pod.CreatePodsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &pod.CreatePodsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 停止处于运行状态的单个实例，处于任务执行中的 pod 无法启动。
 */
func (c *PodClient) StopPod(request *pod.StopPodRequest) (*pod.StopPodResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &pod.StopPodResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询一个 pod 的详细信息
 */
func (c *PodClient) DescribePod(request *pod.DescribePodRequest) (*pod.DescribePodResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &pod.DescribePodResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取 pod 中某个容器的详情 */
func (c *PodClient) DecribeContainer(request *pod.DecribeContainerRequest) (*pod.DecribeContainerResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &pod.DecribeContainerResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除单个 secret
 */
func (c *PodClient) DeleteSecret(request *pod.DeleteSecretRequest) (*pod.DeleteSecretResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &pod.DeleteSecretResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* podName 是否符合命名规范，以及查询指定 podName 区域内是否已经存在。
 */
func (c *PodClient) CheckPodName(request *pod.CheckPodNameRequest) (*pod.CheckPodNameResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &pod.CheckPodNameResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

