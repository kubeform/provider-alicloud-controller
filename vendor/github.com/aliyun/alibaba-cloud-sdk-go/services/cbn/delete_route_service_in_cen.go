package cbn

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

// DeleteRouteServiceInCen invokes the cbn.DeleteRouteServiceInCen API synchronously
func (client *Client) DeleteRouteServiceInCen(request *DeleteRouteServiceInCenRequest) (response *DeleteRouteServiceInCenResponse, err error) {
	response = CreateDeleteRouteServiceInCenResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteRouteServiceInCenWithChan invokes the cbn.DeleteRouteServiceInCen API asynchronously
func (client *Client) DeleteRouteServiceInCenWithChan(request *DeleteRouteServiceInCenRequest) (<-chan *DeleteRouteServiceInCenResponse, <-chan error) {
	responseChan := make(chan *DeleteRouteServiceInCenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRouteServiceInCen(request)
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

// DeleteRouteServiceInCenWithCallback invokes the cbn.DeleteRouteServiceInCen API asynchronously
func (client *Client) DeleteRouteServiceInCenWithCallback(request *DeleteRouteServiceInCenRequest, callback func(response *DeleteRouteServiceInCenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRouteServiceInCenResponse
		var err error
		defer close(result)
		response, err = client.DeleteRouteServiceInCen(request)
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

// DeleteRouteServiceInCenRequest is the request struct for api DeleteRouteServiceInCen
type DeleteRouteServiceInCenRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CenId                string           `position:"Query" name:"CenId"`
	AccessRegionId       string           `position:"Query" name:"AccessRegionId"`
	Host                 string           `position:"Query" name:"Host"`
	HostRegionId         string           `position:"Query" name:"HostRegionId"`
	HostVpcId            string           `position:"Query" name:"HostVpcId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteRouteServiceInCenResponse is the response struct for api DeleteRouteServiceInCen
type DeleteRouteServiceInCenResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteRouteServiceInCenRequest creates a request to invoke DeleteRouteServiceInCen API
func CreateDeleteRouteServiceInCenRequest() (request *DeleteRouteServiceInCenRequest) {
	request = &DeleteRouteServiceInCenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "DeleteRouteServiceInCen", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteRouteServiceInCenResponse creates a response to parse from DeleteRouteServiceInCen response
func CreateDeleteRouteServiceInCenResponse() (response *DeleteRouteServiceInCenResponse) {
	response = &DeleteRouteServiceInCenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
