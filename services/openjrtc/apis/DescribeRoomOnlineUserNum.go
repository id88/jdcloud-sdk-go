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

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
)

type DescribeRoomOnlineUserNumRequest struct {

    core.JDCloudRequest

    /* 房间ID  */
    RoomId int `json:"roomId"`
}

/*
 * param roomId: 房间ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeRoomOnlineUserNumRequest(
    roomId int,
) *DescribeRoomOnlineUserNumRequest {

	return &DescribeRoomOnlineUserNumRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/describeRoomOnlineUserNum/{roomId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RoomId: roomId,
	}
}

/*
 * param roomId: 房间ID (Required)
 */
func NewDescribeRoomOnlineUserNumRequestWithAllParams(
    roomId int,
) *DescribeRoomOnlineUserNumRequest {

    return &DescribeRoomOnlineUserNumRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeRoomOnlineUserNum/{roomId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RoomId: roomId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeRoomOnlineUserNumRequestWithoutParam() *DescribeRoomOnlineUserNumRequest {

    return &DescribeRoomOnlineUserNumRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeRoomOnlineUserNum/{roomId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param roomId: 房间ID(Required) */
func (r *DescribeRoomOnlineUserNumRequest) SetRoomId(roomId int) {
    r.RoomId = roomId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeRoomOnlineUserNumRequest) GetRegionId() string {
    return ""
}

type DescribeRoomOnlineUserNumResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeRoomOnlineUserNumResult `json:"result"`
}

type DescribeRoomOnlineUserNumResult struct {
    AppId string `json:"appId"`
    RoomId int64 `json:"roomId"`
    Number int `json:"number"`
    CreateTime string `json:"createTime"`
}