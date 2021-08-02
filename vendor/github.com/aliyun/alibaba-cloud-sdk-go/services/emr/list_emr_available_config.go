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

// ListEmrAvailableConfig invokes the emr.ListEmrAvailableConfig API synchronously
func (client *Client) ListEmrAvailableConfig(request *ListEmrAvailableConfigRequest) (response *ListEmrAvailableConfigResponse, err error) {
	response = CreateListEmrAvailableConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ListEmrAvailableConfigWithChan invokes the emr.ListEmrAvailableConfig API asynchronously
func (client *Client) ListEmrAvailableConfigWithChan(request *ListEmrAvailableConfigRequest) (<-chan *ListEmrAvailableConfigResponse, <-chan error) {
	responseChan := make(chan *ListEmrAvailableConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListEmrAvailableConfig(request)
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

// ListEmrAvailableConfigWithCallback invokes the emr.ListEmrAvailableConfig API asynchronously
func (client *Client) ListEmrAvailableConfigWithCallback(request *ListEmrAvailableConfigRequest, callback func(response *ListEmrAvailableConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListEmrAvailableConfigResponse
		var err error
		defer close(result)
		response, err = client.ListEmrAvailableConfig(request)
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

// ListEmrAvailableConfigRequest is the request struct for api ListEmrAvailableConfig
type ListEmrAvailableConfigRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
}

// ListEmrAvailableConfigResponse is the response struct for api ListEmrAvailableConfig
type ListEmrAvailableConfigResponse struct {
	*responses.BaseResponse
	RequestId          string                                     `json:"RequestId" xml:"RequestId"`
	KeyPairNameList    KeyPairNameList                            `json:"KeyPairNameList" xml:"KeyPairNameList"`
	EmrMainVersionList EmrMainVersionListInListEmrAvailableConfig `json:"EmrMainVersionList" xml:"EmrMainVersionList"`
	SecurityGroupList  SecurityGroupListInListEmrAvailableConfig  `json:"SecurityGroupList" xml:"SecurityGroupList"`
	VpcInfoList        VpcInfoList                                `json:"VpcInfoList" xml:"VpcInfoList"`
}

// CreateListEmrAvailableConfigRequest creates a request to invoke ListEmrAvailableConfig API
func CreateListEmrAvailableConfigRequest() (request *ListEmrAvailableConfigRequest) {
	request = &ListEmrAvailableConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListEmrAvailableConfig", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListEmrAvailableConfigResponse creates a response to parse from ListEmrAvailableConfig response
func CreateListEmrAvailableConfigResponse() (response *ListEmrAvailableConfigResponse) {
	response = &ListEmrAvailableConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
