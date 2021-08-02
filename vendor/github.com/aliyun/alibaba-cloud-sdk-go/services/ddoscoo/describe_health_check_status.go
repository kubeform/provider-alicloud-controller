package ddoscoo

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

// DescribeHealthCheckStatus invokes the ddoscoo.DescribeHealthCheckStatus API synchronously
func (client *Client) DescribeHealthCheckStatus(request *DescribeHealthCheckStatusRequest) (response *DescribeHealthCheckStatusResponse, err error) {
	response = CreateDescribeHealthCheckStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeHealthCheckStatusWithChan invokes the ddoscoo.DescribeHealthCheckStatus API asynchronously
func (client *Client) DescribeHealthCheckStatusWithChan(request *DescribeHealthCheckStatusRequest) (<-chan *DescribeHealthCheckStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeHealthCheckStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeHealthCheckStatus(request)
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

// DescribeHealthCheckStatusWithCallback invokes the ddoscoo.DescribeHealthCheckStatus API asynchronously
func (client *Client) DescribeHealthCheckStatusWithCallback(request *DescribeHealthCheckStatusRequest, callback func(response *DescribeHealthCheckStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeHealthCheckStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeHealthCheckStatus(request)
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

// DescribeHealthCheckStatusRequest is the request struct for api DescribeHealthCheckStatus
type DescribeHealthCheckStatusRequest struct {
	*requests.RpcRequest
	NetworkRules string `position:"Query" name:"NetworkRules"`
	SourceIp     string `position:"Query" name:"SourceIp"`
}

// DescribeHealthCheckStatusResponse is the response struct for api DescribeHealthCheckStatus
type DescribeHealthCheckStatusResponse struct {
	*responses.BaseResponse
	RequestId         string   `json:"RequestId" xml:"RequestId"`
	HealthCheckStatus []Status `json:"HealthCheckStatus" xml:"HealthCheckStatus"`
}

// CreateDescribeHealthCheckStatusRequest creates a request to invoke DescribeHealthCheckStatus API
func CreateDescribeHealthCheckStatusRequest() (request *DescribeHealthCheckStatusRequest) {
	request = &DescribeHealthCheckStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeHealthCheckStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeHealthCheckStatusResponse creates a response to parse from DescribeHealthCheckStatus response
func CreateDescribeHealthCheckStatusResponse() (response *DescribeHealthCheckStatusResponse) {
	response = &DescribeHealthCheckStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
