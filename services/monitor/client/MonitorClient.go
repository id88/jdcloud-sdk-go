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
    monitor "github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/apis"
    "encoding/json"
    "errors"
)

type MonitorClient struct {
    core.JDCloudClient
}

func NewMonitorClient(credential *core.Credential) *MonitorClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("monitor.jdcloud-api.com")

    return &MonitorClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "monitor",
            Revision:    "1.3.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *MonitorClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *MonitorClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 创建报警规则，可以为某一个实例创建报警规则，也可以为多个实例同时创建报警规则。 */
func (c *MonitorClient) CreateAlarm(request *monitor.CreateAlarmRequest) (*monitor.CreateAlarmResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.CreateAlarmResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查看某资源多个监控项数据，metric介绍1：<a href="https://docs.jdcloud.com/cn/monitoring/metrics">Metrics</a> */
func (c *MonitorClient) DescribeMetricData(request *monitor.DescribeMetricDataRequest) (*monitor.DescribeMetricDataResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.DescribeMetricDataResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询规则的报警联系人 */
func (c *MonitorClient) DescribeAlarmContacts(request *monitor.DescribeAlarmContactsRequest) (*monitor.DescribeAlarmContactsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.DescribeAlarmContactsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询报警历史
检索条件组合优先级从高到低为
1. alarmId
2. serviceCode
2.1 serviceCode + resourceId
2.2 serviceCode + resourceIds
3. serviceCodes
4. 用户所有规则 */
func (c *MonitorClient) DescribeAlarmHistory(request *monitor.DescribeAlarmHistoryRequest) (*monitor.DescribeAlarmHistoryResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.DescribeAlarmHistoryResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除自定义监控规则 */
func (c *MonitorClient) DeleteAlarmsCm(request *monitor.DeleteAlarmsCmRequest) (*monitor.DeleteAlarmsCmResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.DeleteAlarmsCmResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询可用创建监控规则的指标列表,metric介绍：<a href="https://docs.jdcloud.com/cn/monitoring/metrics">Metrics</a> */
func (c *MonitorClient) DescribeMetricsForCreateAlarm(request *monitor.DescribeMetricsForCreateAlarmRequest) (*monitor.DescribeMetricsForCreateAlarmResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.DescribeMetricsForCreateAlarmResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查看某资源的最后一个点,metric介绍：<a href="https://docs.jdcloud.com/cn/monitoring/metrics">Metrics</a> */
func (c *MonitorClient) LastDownsample(request *monitor.LastDownsampleRequest) (*monitor.LastDownsampleResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.LastDownsampleResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询规则详情 */
func (c *MonitorClient) DescribeAlarmsByID(request *monitor.DescribeAlarmsByIDRequest) (*monitor.DescribeAlarmsByIDResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.DescribeAlarmsByIDResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改已创建的报警规则 */
func (c *MonitorClient) UpdateAlarm(request *monitor.UpdateAlarmRequest) (*monitor.UpdateAlarmResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.UpdateAlarmResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询报警历史
检索条件组合优先级从高到低为
1. alarmId
2. serviceCode
2.1 serviceCode + resourceId
2.2 serviceCode + resourceIds
3. serviceCodes
4. 用户所有规则 */
func (c *MonitorClient) DescribeAlarmHistoryAllRegion(request *monitor.DescribeAlarmHistoryAllRegionRequest) (*monitor.DescribeAlarmHistoryAllRegionResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.DescribeAlarmHistoryAllRegionResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 该接口为自定义监控数据上报的接口，方便您将自己采集的时序数据上报到云监控。不同region域名上报不同region的数据，参考：<a href="https://docs.jdcloud.com/cn/monitoring/reporting-monitoring-data">调用说明</a>可上报原始数据和已聚合的统计数据。支持批量上报方式。单次请求最多包含 50 个数据点；数据大小不超过 256k。 */
func (c *MonitorClient) PutMetricData(request *monitor.PutMetricDataRequest) (*monitor.PutMetricDataResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.PutMetricDataResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 禁用报警规则。报警规则禁用后，将停止探测实例的监控项数据。 */
func (c *MonitorClient) DisableAlarm(request *monitor.DisableAlarmRequest) (*monitor.DisableAlarmResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.DisableAlarmResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 启用报警规则，当客户的报警规则处于停止状态时，可以使用此接口启用报警规则。 */
func (c *MonitorClient) EnableAlarm(request *monitor.EnableAlarmRequest) (*monitor.EnableAlarmResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.EnableAlarmResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询规则, 查询参数组合及优先级从高到低为：
1：alarmId不为空
2：serviceCode不为空
2.1：serviceCode + resourceId
2.2: serviceCode + resourceIds
3：serviceCodes不为空
4: 所有规则 */
func (c *MonitorClient) DescribeAlarms(request *monitor.DescribeAlarmsRequest) (*monitor.DescribeAlarmsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.DescribeAlarmsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据产品线查询可用监控项列表,metric介绍：<a href="https://docs.jdcloud.com/cn/monitoring/metrics">Metrics</a> */
func (c *MonitorClient) DescribeMetrics(request *monitor.DescribeMetricsRequest) (*monitor.DescribeMetricsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.DescribeMetricsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

