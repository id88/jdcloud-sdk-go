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
    pod "github.com/jdcloud-api/jdcloud-sdk-go/services/pod/models"
)

type DescribeContainerRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Pod ID  */
    PodId string `json:"podId"`

    /* container name  */
    ContainerName string `json:"containerName"`
}

/*
 * param regionId: Region ID (Required)
 * param podId: Pod ID (Required)
 * param containerName: container name (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeContainerRequest(
    regionId string,
    podId string,
    containerName string,
) *DescribeContainerRequest {

	return &DescribeContainerRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/pods/{podId}/containers/{containerName}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        PodId: podId,
        ContainerName: containerName,
	}
}

/*
 * param regionId: Region ID (Required)
 * param podId: Pod ID (Required)
 * param containerName: container name (Required)
 */
func NewDescribeContainerRequestWithAllParams(
    regionId string,
    podId string,
    containerName string,
) *DescribeContainerRequest {

    return &DescribeContainerRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pods/{podId}/containers/{containerName}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PodId: podId,
        ContainerName: containerName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeContainerRequestWithoutParam() *DescribeContainerRequest {

    return &DescribeContainerRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pods/{podId}/containers/{containerName}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeContainerRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param podId: Pod ID(Required) */
func (r *DescribeContainerRequest) SetPodId(podId string) {
    r.PodId = podId
}
/* param containerName: container name(Required) */
func (r *DescribeContainerRequest) SetContainerName(containerName string) {
    r.ContainerName = containerName
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeContainerRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeContainerResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeContainerResult `json:"result"`
}

type DescribeContainerResult struct {
    Container pod.Container `json:"container"`
}