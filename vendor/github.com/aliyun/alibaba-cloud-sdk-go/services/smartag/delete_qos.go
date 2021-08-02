package smartag

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

// DeleteQos invokes the smartag.DeleteQos API synchronously
func (client *Client) DeleteQos(request *DeleteQosRequest) (response *DeleteQosResponse, err error) {
	response = CreateDeleteQosResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteQosWithChan invokes the smartag.DeleteQos API asynchronously
func (client *Client) DeleteQosWithChan(request *DeleteQosRequest) (<-chan *DeleteQosResponse, <-chan error) {
	responseChan := make(chan *DeleteQosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteQos(request)
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

// DeleteQosWithCallback invokes the smartag.DeleteQos API asynchronously
func (client *Client) DeleteQosWithCallback(request *DeleteQosRequest, callback func(response *DeleteQosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteQosResponse
		var err error
		defer close(result)
		response, err = client.DeleteQos(request)
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

// DeleteQosRequest is the request struct for api DeleteQos
type DeleteQosRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query"`
	QosId                string           `position:"Query"`
	ResourceOwnerAccount string           `position:"Query"`
	OwnerAccount         string           `position:"Query"`
	OwnerId              requests.Integer `position:"Query"`
}

// DeleteQosResponse is the response struct for api DeleteQos
type DeleteQosResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteQosRequest creates a request to invoke DeleteQos API
func CreateDeleteQosRequest() (request *DeleteQosRequest) {
	request = &DeleteQosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DeleteQos", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteQosResponse creates a response to parse from DeleteQos response
func CreateDeleteQosResponse() (response *DeleteQosResponse) {
	response = &DeleteQosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
