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
    yunding "github.com/jdcloud-api/jdcloud-sdk-go/services/yunding/apis"
    "encoding/json"
    "errors"
)

type YundingClient struct {
    core.JDCloudClient
}

func NewYundingClient(credential *core.Credential) *YundingClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("yunding.jdcloud-api.com")

    return &YundingClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "yunding",
            Revision:    "2.0.4",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *YundingClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *YundingClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

func (c *YundingClient) DisableLogger() {
    c.Logger = core.NewDummyLogger()
}

/* 查询RDS实例（MySQL、SQL Server等）的详细信息以及MySQL只读实例详细信息 */
func (c *YundingClient) DescribeRdsInstance(request *yunding.DescribeRdsInstanceRequest) (*yunding.DescribeRdsInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.DescribeRdsInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询云鼎数据库实例列表
 */
func (c *YundingClient) DescribeYdRdsInstances(request *yunding.DescribeYdRdsInstancesRequest) (*yunding.DescribeYdRdsInstancesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.DescribeYdRdsInstancesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除子网 */
func (c *YundingClient) DeleteSubnet(request *yunding.DeleteSubnetRequest) (*yunding.DeleteSubnetResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.DeleteSubnetResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询子网列表 */
func (c *YundingClient) DescribeSubnets(request *yunding.DescribeSubnetsRequest) (*yunding.DescribeSubnetsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.DescribeSubnetsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查看某个RDS实例下所有账号信息，包括账号名称、对各个数据库的访问权限信息等 */
func (c *YundingClient) DescribeRdsAccounts(request *yunding.DescribeRdsAccountsRequest) (*yunding.DescribeRdsAccountsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.DescribeRdsAccountsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建子网 */
func (c *YundingClient) CreateSubnet(request *yunding.CreateSubnetRequest) (*yunding.CreateSubnetResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.CreateSubnetResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 监控数据上报。 */
func (c *YundingClient) PutProductMetricData(request *yunding.PutProductMetricDataRequest) (*yunding.PutProductMetricDataResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.PutProductMetricDataResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 从RDS实例中删除数据库。为便于管理和数据恢复，RDS对用户权限进行了控制，用户仅能通过控制台或本接口删除数据库 [MFA enabled] */
func (c *YundingClient) DeleteRdsDatabase(request *yunding.DeleteRdsDatabaseRequest) (*yunding.DeleteRdsDatabaseResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.DeleteRdsDatabaseResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 云拔测-可用性agent任务查询接口 */
func (c *YundingClient) DescribeTasks(request *yunding.DescribeTasksRequest) (*yunding.DescribeTasksResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.DescribeTasksResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 取消该账号对某个数据库的所有权限。权限取消后，该账号将不能访问此数据库。取消账号对某个数据库的访问权限，不影响该账号对其他数据库的访问权限 */
func (c *YundingClient) RevokePrivilege(request *yunding.RevokePrivilegeRequest) (*yunding.RevokePrivilegeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.RevokePrivilegeResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取当前实例的所有数据库详细信息的列表 */
func (c *YundingClient) DescribeRdsDatabases(request *yunding.DescribeRdsDatabasesRequest) (*yunding.DescribeRdsDatabasesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.DescribeRdsDatabasesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询弹性网卡信息详情 */
func (c *YundingClient) DescribeNetworkInterface(request *yunding.DescribeNetworkInterfaceRequest) (*yunding.DescribeNetworkInterfaceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.DescribeNetworkInterfaceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 监控数据上报。 */
func (c *YundingClient) Put(request *yunding.PutRequest) (*yunding.PutResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.PutResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 给网卡删除secondaryIp接口 */
func (c *YundingClient) UnassignSecondaryIps(request *yunding.UnassignSecondaryIpsRequest) (*yunding.UnassignSecondaryIpsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.UnassignSecondaryIpsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询弹性网卡列表 */
func (c *YundingClient) DescribeNetworkInterfaces(request *yunding.DescribeNetworkInterfacesRequest) (*yunding.DescribeNetworkInterfacesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.DescribeNetworkInterfacesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 给网卡分配secondaryIp接口 */
func (c *YundingClient) AssignSecondaryIps(request *yunding.AssignSecondaryIpsRequest) (*yunding.AssignSecondaryIpsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.AssignSecondaryIpsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改允许访问实例的IP白名单。白名单是允许访问当前实例的IP/IP段列表，缺省情况下，白名单对本VPC开放。如果用户开启了外网访问的功能，还需要对外网的IP配置白名单。 */
func (c *YundingClient) ModifyRdsWhiteList(request *yunding.ModifyRdsWhiteListRequest) (*yunding.ModifyRdsWhiteListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.ModifyRdsWhiteListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建一个数据库。 为了实例的管理和数据恢复，RDS对用户权限进行了限制，用户仅能通过控制台或本接口创建数据库 */
func (c *YundingClient) CreateRdsDatabase(request *yunding.CreateRdsDatabaseRequest) (*yunding.CreateRdsDatabaseResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.CreateRdsDatabaseResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 授予账号的数据库访问权限，即该账号对数据库拥有什么权限。一个账号可以对多个数据库具有访问权限。<br>为便于管理，RDS对权限进行了归类，目前提供以下两种权限<br>- ro：只读权限，用户只能读取数据库中的数据，不能进行创建、插入、删除、更改等操作。<br>- rw：读写权限，用户可以对数据库进行增删改查等操作 */
func (c *YundingClient) GrantRdsPrivilege(request *yunding.GrantRdsPrivilegeRequest) (*yunding.GrantRdsPrivilegeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.GrantRdsPrivilegeResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建网卡接口，只能创建辅助网卡 */
func (c *YundingClient) CreateNetworkInterface(request *yunding.CreateNetworkInterfaceRequest) (*yunding.CreateNetworkInterfaceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.CreateNetworkInterfaceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询子网信息详情 */
func (c *YundingClient) DescribeSubnet(request *yunding.DescribeSubnetRequest) (*yunding.DescribeSubnetResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.DescribeSubnetResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 批量查询云数据库实例列表信息<br>此接口支持分页查询，默认每页20条。 */
func (c *YundingClient) DescribeRdsInstances(request *yunding.DescribeRdsInstancesRequest) (*yunding.DescribeRdsInstancesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.DescribeRdsInstancesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除弹性网卡接口 */
func (c *YundingClient) DeleteNetworkInterface(request *yunding.DeleteNetworkInterfaceRequest) (*yunding.DeleteNetworkInterfaceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.DeleteNetworkInterfaceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询RDS实例（MySQL、SQL Server等）的真实状态 */
func (c *YundingClient) DescribeInstanceInfo(request *yunding.DescribeInstanceInfoRequest) (*yunding.DescribeInstanceInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.DescribeInstanceInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建数据库账号，用户可以使用客户端，应用程序等通过该账号和密码登录RDS数据库实例。<br>为便于管理和恢复，RDS对账号进行了限制，数据库账号只能通过控制台或者OpenAPI进行创建、删除账号以及对账号授权等，用户不能通过SQL语句对账号进行相关操作。 */
func (c *YundingClient) CreateRdsAccount(request *yunding.CreateRdsAccountRequest) (*yunding.CreateRdsAccountResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.CreateRdsAccountResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查看RDS实例当前白名单。白名单是允许访问当前实例的IP/IP段列表，缺省情况下，白名单对本VPC开放。如果用户开启了外网访问的功能，还需要对外网的IP配置白名单。 */
func (c *YundingClient) DescribeRdsWhiteList(request *yunding.DescribeRdsWhiteListRequest) (*yunding.DescribeRdsWhiteListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.DescribeRdsWhiteListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除数据库账号，账号删除后不可恢复，用户无法再使用该账号登录RDS实例 */
func (c *YundingClient) DeleteRdsAccount(request *yunding.DeleteRdsAccountRequest) (*yunding.DeleteRdsAccountResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yunding.DeleteRdsAccountResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

