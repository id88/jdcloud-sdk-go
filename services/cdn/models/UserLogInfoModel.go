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


type UserLogInfoModel struct {

    /* domain所归属pin (Optional) */
    Pin string `json:"pin"`

    /* 日志上传域名，如：www.a.com (Optional) */
    Domain string `json:"domain"`

    /* 日志上传全路径，如：cdnuserlog/www.a.com/20190412/2019041200-01.gz (Optional) */
    LogFileFullPath string `json:"logFileFullPath"`

    /* 日志粒度：fiveMin(五分钟粒度),hour(一小时粒度),day(一天粒度) (Optional) */
    Interval string `json:"interval"`

    /* 日志类型：gz,log,zip (Optional) */
    LogType string `json:"logType"`

    /* 日志大小，单位：Byte（字节） (Optional) */
    LogSize int64 `json:"logSize"`

    /* MD5值 (Optional) */
    LogMD5 string `json:"logMD5"`

    /* 日志开始时间，格式：yyyy-MM-dd HH:ss，如：2019-04-12 00:00 (Optional) */
    StartTime string `json:"startTime"`

    /* 日志结束时间，格式：yyyy-MM-dd HH:ss 如：2019-04-12 00:05 (Optional) */
    EndTime string `json:"endTime"`
}
