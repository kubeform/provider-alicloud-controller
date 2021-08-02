package dds

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

// ReleaseNodePrivateNetworkAddress invokes the dds.ReleaseNodePrivateNetworkAddress API synchronously
func (client *Client) ReleaseNodePrivateNetworkAddress(request *ReleaseNodePrivateNetworkAddressRequest) (response *ReleaseNodePrivateNetworkAddressResponse, err error) {
	response = CreateReleaseNodePrivateNetworkAddressResponse()
	err = client.DoAction(request, response)
	return
}

// ReleaseNodePrivateNetworkAddressWithChan invokes the dds.ReleaseNodePrivateNetworkAddress API asynchronously
func (client *Client) ReleaseNodePrivateNetworkAddressWithChan(request *ReleaseNodePrivateNetworkAddressRequest) (<-chan *ReleaseNodePrivateNetworkAddressResponse, <-chan error) {
	responseChan := make(chan *ReleaseNodePrivateNetworkAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleaseNodePrivateNetworkAddress(request)
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

// ReleaseNodePrivateNetworkAddressWithCallback invokes the dds.ReleaseNodePrivateNetworkAddress API asynchronously
func (client *Client) ReleaseNodePrivateNetworkAddressWithCallback(request *ReleaseNodePrivateNetworkAddressRequest, callback func(response *ReleaseNodePrivateNetworkAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleaseNodePrivateNetworkAddressResponse
		var err error
		defer close(result)
		response, err = client.ReleaseNodePrivateNetworkAddress(request)
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

// ReleaseNodePrivateNetworkAddressRequest is the request struct for api ReleaseNodePrivateNetworkAddress
type ReleaseNodePrivateNetworkAddressRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	NetworkType          string           `position:"Query" name:"NetworkType"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	NodeId               string           `position:"Query" name:"NodeId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ReleaseNodePrivateNetworkAddressResponse is the response struct for api ReleaseNodePrivateNetworkAddress
type ReleaseNodePrivateNetworkAddressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateReleaseNodePrivateNetworkAddressRequest creates a request to invoke ReleaseNodePrivateNetworkAddress API
func CreateReleaseNodePrivateNetworkAddressRequest() (request *ReleaseNodePrivateNetworkAddressRequest) {
	request = &ReleaseNodePrivateNetworkAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "ReleaseNodePrivateNetworkAddress", "", "")
	request.Method = requests.POST
	return
}

// CreateReleaseNodePrivateNetworkAddressResponse creates a response to parse from ReleaseNodePrivateNetworkAddress response
func CreateReleaseNodePrivateNetworkAddressResponse() (response *ReleaseNodePrivateNetworkAddressResponse) {
	response = &ReleaseNodePrivateNetworkAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
