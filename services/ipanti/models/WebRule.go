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


type WebRule struct {

    /* 规则 Id (Optional) */
    Id string `json:"id"`

    /* 实例 Id (Optional) */
    InstanceId string `json:"instanceId"`

    /* 子域名 (Optional) */
    Domain string `json:"domain"`

    /* 规则的 CNAME (Optional) */
    Cname string `json:"cname"`

    /* CNAME 解析状态, 0: 解析异常, 1: 解析正常 (Optional) */
    CnameStatus int `json:"cnameStatus"`

    /* 该规则使用中的高防 IP (Optional) */
    ServiceIp string `json:"serviceIp"`

    /* 已配置的高防 IP 列表 (Optional) */
    ServiceIpConfig ServiceIpConfig `json:"serviceIpConfig"`

    /*  (Optional) */
    Protocol WebRuleProtocol `json:"protocol"`

    /* 是否为自定义端口号, 0: 为默认, 1: 为自定义 (Optional) */
    CustomPortStatus int `json:"customPortStatus"`

    /* HTTP 协议的端口号, 如 80,81 (Optional) */
    Port []int `json:"port"`

    /* HTTPS 协议的端口号, 如 443,8443 (Optional) */
    HttpsPort []int `json:"httpsPort"`

    /* 是否开启 HTTP 回源, 0: 为不开启, 1: 为开启, 当勾选 HTTPS 时可以配置该属性 (Optional) */
    HttpOrigin int `json:"httpOrigin"`

    /* 0: 防御状态, 1: 回源状态 (Optional) */
    Status int `json:"status"`

    /* 回源类型: A 或者 CNAME (Optional) */
    OriginType string `json:"originType"`

    /* 回源域名, originType 为 A 时返回该字段 (Optional) */
    OriginAddr []OriginAddrItem `json:"originAddr"`

    /* 回源域名, originType 为 CNAME 时返回该字段 (Optional) */
    OriginDomain string `json:"originDomain"`

    /* 备用的回源地址列表, 为一个域名或者多个 IP 地址 (Optional) */
    OnlineAddr []string `json:"onlineAddr"`

    /* 证书状态. <br>- 0: 异常<br>- 1: 正常<br>- 2: 证书未上传 (Optional) */
    HttpCertStatus int `json:"httpCertStatus"`

    /* 证书 Id, (废弃, 绑定证书信息通过 certs 字段查看) (Optional) */
    CertId string `json:"certId"`

    /* 证书名称, (废弃, 绑定证书信息通过 certs 字段查看) (Optional) */
    CertName string `json:"certName"`

    /* 证书内容, (废弃, 绑定证书信息通过 certs 字段查看) (Optional) */
    HttpsCertContent string `json:"httpsCertContent"`

    /* 证书私钥, (废弃, 绑定证书信息通过 certs 字段查看) (Optional) */
    HttpsRsaKey string `json:"httpsRsaKey"`

    /* 网站规则绑定证书信息 (Optional) */
    BindCerts []Cert `json:"bindCerts"`

    /* 是否开启 HTTPS 强制跳转, 当 protocol 为 HTTP_HTTPS 时可以配置该属性<br>- 0: 不强跳<br>- 1: 开启强跳 (Optional) */
    ForceJump int `json:"forceJump"`

    /* 转发规则. <br>- wrr: 带权重的轮询<br>- rr:  不带权重的轮询<br>- sh:  源地址hash (Optional) */
    Algorithm string `json:"algorithm"`

    /* CC 状态, 0: CC 关闭, 1: CC 开启 (Optional) */
    CcStatus int `json:"ccStatus"`

    /* webSocket 状态, 0: 关闭, 1: 开启 (Optional) */
    WebSocketStatus int `json:"webSocketStatus"`

    /* 黑名单状态, 0: 关闭, 1: 开启 (Optional) */
    BlackListEnable int `json:"blackListEnable"`

    /* 白名单状态, 0: 关闭, 1: 开启 (Optional) */
    WhiteListEnable int `json:"whiteListEnable"`

    /* 按区域分流回源配置 (Optional) */
    GeoRsRoute []GeoRsRoute `json:"geoRsRoute"`

    /* 是否开启回源长连接, protocol 选项开启 https 时生效, 可取值<br>- on: 开启<br>- off: 关闭 (Optional) */
    EnableKeepalive string `json:"enableKeepalive"`

    /* http 版本, protocol 选项开启 https 时生效, 可取值 http1 或 http2 (Optional) */
    HttpVersion string `json:"httpVersion"`

    /* SSL协议类型, protocol 选项开启 https 时生效, 可取值SSLv2,SSLv3,TLSv1.0,TLSv1.1,TLSv1.2 (Optional) */
    SslProtocols []string `json:"sslProtocols"`

    /* 加密套件等级, protocol 选项开启 https 时生效, 可取值<br>- low: 低级<br>- middle: 中级<br>- high：高级<br>- custom：自定义 (Optional) */
    SuiteLevel string `json:"suiteLevel"`

    /* 自定义加密套件等级, suiteLevel 为 custom 是有效 (Optional) */
    UserSuiteLevel []string `json:"userSuiteLevel"`

    /* 是否允许在 response 中插入 JS, 0: 关闭, 1: 开启 (Optional) */
    JsFingerprintEnable int `json:"jsFingerprintEnable"`

    /* JS 指纹生效范围, 0: 所有页面, 1: 已配置的自定义页面 (Optional) */
    JsFingerprintScope int `json:"jsFingerprintScope"`

    /* CC自定义规则总开关, 0: 关闭, 1: 开启 (Optional) */
    CcCustomStatus int `json:"ccCustomStatus"`

    /* 健康检查开关, 0: 关闭, 1: 开启 (Optional) */
    EnableHealthCheck int `json:"enableHealthCheck"`

    /* 回源连接超时时长, 单位 秒 (Optional) */
    ProxyConnectTimeout int `json:"proxyConnectTimeout"`

    /* 请求头支持下划线, 0: 关闭, 1: 开启 (Optional) */
    EnableUnderscores int `json:"enableUnderscores"`
}
