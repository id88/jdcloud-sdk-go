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

type CreateRoomRequest struct {

    core.JDCloudRequest

    /* 房间名称 (Optional) */
    RoomName *string `json:"roomName"`

    /* 应用ID (Optional) */
    AppId *string `json:"appId"`

    /* JRtc用户ID(创建者ID) (Optional) */
    PeerId *int64 `json:"peerId"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateRoomRequest(
) *CreateRoomRequest {

	return &CreateRoomRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/createRoom",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param roomName: 房间名称 (Optional)
 * param appId: 应用ID (Optional)
 * param peerId: JRtc用户ID(创建者ID) (Optional)
 */
func NewCreateRoomRequestWithAllParams(
    roomName *string,
    appId *string,
    peerId *int64,
) *CreateRoomRequest {

    return &CreateRoomRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/createRoom",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RoomName: roomName,
        AppId: appId,
        PeerId: peerId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateRoomRequestWithoutParam() *CreateRoomRequest {

    return &CreateRoomRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/createRoom",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param roomName: 房间名称(Optional) */
func (r *CreateRoomRequest) SetRoomName(roomName string) {
    r.RoomName = &roomName
}

/* param appId: 应用ID(Optional) */
func (r *CreateRoomRequest) SetAppId(appId string) {
    r.AppId = &appId
}

/* param peerId: JRtc用户ID(创建者ID)(Optional) */
func (r *CreateRoomRequest) SetPeerId(peerId int64) {
    r.PeerId = &peerId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateRoomRequest) GetRegionId() string {
    return ""
}

type CreateRoomResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateRoomResult `json:"result"`
}

type CreateRoomResult struct {
    RoomId int64 `json:"roomId"`
    RoomName string `json:"roomName"`
    AppId string `json:"appId"`
    PeerId int64 `json:"peerId"`
}