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
    autotaskpolicy "github.com/jdcloud-api/jdcloud-sdk-go/services/autotaskpolicy/apis"
    "encoding/json"
    "errors"
)

type AutotaskpolicyClient struct {
    core.JDCloudClient
}

func NewAutotaskpolicyClient(credential *core.Credential) *AutotaskpolicyClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("autotaskpolicy.jdcloud-api.com")

    return &AutotaskpolicyClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "autotaskpolicy",
            Revision:    "1.0.1",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *AutotaskpolicyClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *AutotaskpolicyClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

func (c *AutotaskpolicyClient) DisableLogger() {
    c.Logger = core.NewDummyLogger()
}

/* 策略跨区复制。
 */
func (c *AutotaskpolicyClient) CopyPolicy(request *autotaskpolicy.CopyPolicyRequest) (*autotaskpolicy.CopyPolicyResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &autotaskpolicy.CopyPolicyResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询策略详情。
 */
func (c *AutotaskpolicyClient) DescribePolicy(request *autotaskpolicy.DescribePolicyRequest) (*autotaskpolicy.DescribePolicyResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &autotaskpolicy.DescribePolicyResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 开启策略。 */
func (c *AutotaskpolicyClient) EnablePolicy(request *autotaskpolicy.EnablePolicyRequest) (*autotaskpolicy.EnablePolicyResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &autotaskpolicy.EnablePolicyResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 关联资源。
 */
func (c *AutotaskpolicyClient) AssociateExecResource(request *autotaskpolicy.AssociateExecResourceRequest) (*autotaskpolicy.AssociateExecResourceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &autotaskpolicy.AssociateExecResourceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建策略，不同策略类型有不同的规则，具体如下。
- AutoImage: 自动创建镜像策略
    - fireCondition
        - 仅支持"interval"，执行周期仅支持"小时/天/周"，如: "interval 7d"
        - 两次执行时间间隔需要大于12小时
    - fireTime
        - 触发时间，首次创建时距离当前时间必须在30分钟以后至1年以内。
        - 若策略曾经执行过，再次修改该字段时必须满足距离上一次执行时间超过12小时
    - execResource
        - 哪些云主机需要制作镜像，可指定具体云主机ID列表，每个策略最多绑定100个。
        - 仅支持云盘系统盘的虚机
    - execConfig
        - 配置规则例子:
        - [{"key":"includeDeviceName","value":"vdb,vdc"},{"key":"imageLiveDays","10"}]
        - includeDeviceName说明：云主机中的哪些云盘需要制作镜像，可指定具体盘符列表，或指定为"all"。云主机中的vda系统盘不可改变，一定要参与制作镜像。
        - imageLiveDays说明：镜像保留时间，以天为单位，范围1-36500。不指定则永久有效。
 */
func (c *AutotaskpolicyClient) CreatePolicy(request *autotaskpolicy.CreatePolicyRequest) (*autotaskpolicy.CreatePolicyResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &autotaskpolicy.CreatePolicyResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 取消关联资源。
 */
func (c *AutotaskpolicyClient) DisassociateExecResource(request *autotaskpolicy.DisassociateExecResourceRequest) (*autotaskpolicy.DisassociateExecResourceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &autotaskpolicy.DisassociateExecResourceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询策略列表。
 */
func (c *AutotaskpolicyClient) DescribePolicies(request *autotaskpolicy.DescribePoliciesRequest) (*autotaskpolicy.DescribePoliciesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &autotaskpolicy.DescribePoliciesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 更新策略。
 */
func (c *AutotaskpolicyClient) ModifyPolicy(request *autotaskpolicy.ModifyPolicyRequest) (*autotaskpolicy.ModifyPolicyResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &autotaskpolicy.ModifyPolicyResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除策略。
 */
func (c *AutotaskpolicyClient) DeletePolicy(request *autotaskpolicy.DeletePolicyRequest) (*autotaskpolicy.DeletePolicyResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &autotaskpolicy.DeletePolicyResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 关闭策略。 */
func (c *AutotaskpolicyClient) DisablePolicy(request *autotaskpolicy.DisablePolicyRequest) (*autotaskpolicy.DisablePolicyResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &autotaskpolicy.DisablePolicyResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

