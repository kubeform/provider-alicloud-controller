package adb

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

// ModifyDBClusterResourceGroup invokes the adb.ModifyDBClusterResourceGroup API synchronously
func (client *Client) ModifyDBClusterResourceGroup(request *ModifyDBClusterResourceGroupRequest) (response *ModifyDBClusterResourceGroupResponse, err error) {
	response = CreateModifyDBClusterResourceGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBClusterResourceGroupWithChan invokes the adb.ModifyDBClusterResourceGroup API asynchronously
func (client *Client) ModifyDBClusterResourceGroupWithChan(request *ModifyDBClusterResourceGroupRequest) (<-chan *ModifyDBClusterResourceGroupResponse, <-chan error) {
	responseChan := make(chan *ModifyDBClusterResourceGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBClusterResourceGroup(request)
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

// ModifyDBClusterResourceGroupWithCallback invokes the adb.ModifyDBClusterResourceGroup API asynchronously
func (client *Client) ModifyDBClusterResourceGroupWithCallback(request *ModifyDBClusterResourceGroupRequest, callback func(response *ModifyDBClusterResourceGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBClusterResourceGroupResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBClusterResourceGroup(request)
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

// ModifyDBClusterResourceGroupRequest is the request struct for api ModifyDBClusterResourceGroup
type ModifyDBClusterResourceGroupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	NewResourceGroupId   string           `position:"Query" name:"NewResourceGroupId"`
}

// ModifyDBClusterResourceGroupResponse is the response struct for api ModifyDBClusterResourceGroup
type ModifyDBClusterResourceGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBClusterResourceGroupRequest creates a request to invoke ModifyDBClusterResourceGroup API
func CreateModifyDBClusterResourceGroupRequest() (request *ModifyDBClusterResourceGroupRequest) {
	request = &ModifyDBClusterResourceGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("adb", "2019-03-15", "ModifyDBClusterResourceGroup", "ads", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDBClusterResourceGroupResponse creates a response to parse from ModifyDBClusterResourceGroup response
func CreateModifyDBClusterResourceGroupResponse() (response *ModifyDBClusterResourceGroupResponse) {
	response = &ModifyDBClusterResourceGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
