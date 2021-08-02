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

// DescribeUnBlackholeCount invokes the ddoscoo.DescribeUnBlackholeCount API synchronously
func (client *Client) DescribeUnBlackholeCount(request *DescribeUnBlackholeCountRequest) (response *DescribeUnBlackholeCountResponse, err error) {
	response = CreateDescribeUnBlackholeCountResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUnBlackholeCountWithChan invokes the ddoscoo.DescribeUnBlackholeCount API asynchronously
func (client *Client) DescribeUnBlackholeCountWithChan(request *DescribeUnBlackholeCountRequest) (<-chan *DescribeUnBlackholeCountResponse, <-chan error) {
	responseChan := make(chan *DescribeUnBlackholeCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUnBlackholeCount(request)
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

// DescribeUnBlackholeCountWithCallback invokes the ddoscoo.DescribeUnBlackholeCount API asynchronously
func (client *Client) DescribeUnBlackholeCountWithCallback(request *DescribeUnBlackholeCountRequest, callback func(response *DescribeUnBlackholeCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUnBlackholeCountResponse
		var err error
		defer close(result)
		response, err = client.DescribeUnBlackholeCount(request)
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

// DescribeUnBlackholeCountRequest is the request struct for api DescribeUnBlackholeCount
type DescribeUnBlackholeCountRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
}

// DescribeUnBlackholeCountResponse is the response struct for api DescribeUnBlackholeCount
type DescribeUnBlackholeCountResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	TotalCount  int    `json:"TotalCount" xml:"TotalCount"`
	RemainCount int    `json:"RemainCount" xml:"RemainCount"`
}

// CreateDescribeUnBlackholeCountRequest creates a request to invoke DescribeUnBlackholeCount API
func CreateDescribeUnBlackholeCountRequest() (request *DescribeUnBlackholeCountRequest) {
	request = &DescribeUnBlackholeCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeUnBlackholeCount", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeUnBlackholeCountResponse creates a response to parse from DescribeUnBlackholeCount response
func CreateDescribeUnBlackholeCountResponse() (response *DescribeUnBlackholeCountResponse) {
	response = &DescribeUnBlackholeCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
