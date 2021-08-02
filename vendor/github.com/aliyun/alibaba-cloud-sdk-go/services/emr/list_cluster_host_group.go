package emr

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ListClusterHostGroup invokes the emr.ListClusterHostGroup API synchronously
func (client *Client) ListClusterHostGroup(request *ListClusterHostGroupRequest) (response *ListClusterHostGroupResponse, err error) {
	response = CreateListClusterHostGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ListClusterHostGroupWithChan invokes the emr.ListClusterHostGroup API asynchronously
func (client *Client) ListClusterHostGroupWithChan(request *ListClusterHostGroupRequest) (<-chan *ListClusterHostGroupResponse, <-chan error) {
	responseChan := make(chan *ListClusterHostGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListClusterHostGroup(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ListClusterHostGroupWithCallback invokes the emr.ListClusterHostGroup API asynchronously
func (client *Client) ListClusterHostGroupWithCallback(request *ListClusterHostGroupRequest, callback func(response *ListClusterHostGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListClusterHostGroupResponse
		var err error
		defer close(result)
		response, err = client.ListClusterHostGroup(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ListClusterHostGroupRequest is the request struct for api ListClusterHostGroup
type ListClusterHostGroupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StatusList      *[]string        `position:"Query" name:"StatusList"  type:"Repeated"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
	HostGroupName   string           `position:"Query" name:"HostGroupName"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	HostGroupId     string           `position:"Query" name:"HostGroupId"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	HostGroupType   string           `position:"Query" name:"HostGroupType"`
}

// ListClusterHostGroupResponse is the response struct for api ListClusterHostGroup
type ListClusterHostGroupResponse struct {
	*responses.BaseResponse
	RequestId     string                              `json:"RequestId" xml:"RequestId"`
	PageNumber    int                                 `json:"PageNumber" xml:"PageNumber"`
	PageSize      int                                 `json:"PageSize" xml:"PageSize"`
	Total         int                                 `json:"Total" xml:"Total"`
	ClusterId     string                              `json:"ClusterId" xml:"ClusterId"`
	HostGroupList HostGroupListInListClusterHostGroup `json:"HostGroupList" xml:"HostGroupList"`
}

// CreateListClusterHostGroupRequest creates a request to invoke ListClusterHostGroup API
func CreateListClusterHostGroupRequest() (request *ListClusterHostGroupRequest) {
	request = &ListClusterHostGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListClusterHostGroup", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListClusterHostGroupResponse creates a response to parse from ListClusterHostGroup response
func CreateListClusterHostGroupResponse() (response *ListClusterHostGroupResponse) {
	response = &ListClusterHostGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
