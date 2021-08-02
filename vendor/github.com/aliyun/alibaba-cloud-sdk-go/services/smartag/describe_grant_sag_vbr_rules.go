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

// DescribeGrantSagVbrRules invokes the smartag.DescribeGrantSagVbrRules API synchronously
func (client *Client) DescribeGrantSagVbrRules(request *DescribeGrantSagVbrRulesRequest) (response *DescribeGrantSagVbrRulesResponse, err error) {
	response = CreateDescribeGrantSagVbrRulesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGrantSagVbrRulesWithChan invokes the smartag.DescribeGrantSagVbrRules API asynchronously
func (client *Client) DescribeGrantSagVbrRulesWithChan(request *DescribeGrantSagVbrRulesRequest) (<-chan *DescribeGrantSagVbrRulesResponse, <-chan error) {
	responseChan := make(chan *DescribeGrantSagVbrRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGrantSagVbrRules(request)
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

// DescribeGrantSagVbrRulesWithCallback invokes the smartag.DescribeGrantSagVbrRules API asynchronously
func (client *Client) DescribeGrantSagVbrRulesWithCallback(request *DescribeGrantSagVbrRulesRequest, callback func(response *DescribeGrantSagVbrRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGrantSagVbrRulesResponse
		var err error
		defer close(result)
		response, err = client.DescribeGrantSagVbrRules(request)
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

// DescribeGrantSagVbrRulesRequest is the request struct for api DescribeGrantSagVbrRules
type DescribeGrantSagVbrRulesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query"`
	PageNumber           requests.Integer `position:"Query"`
	PageSize             requests.Integer `position:"Query"`
	ResourceOwnerAccount string           `position:"Query"`
	OwnerAccount         string           `position:"Query"`
	OwnerId              requests.Integer `position:"Query"`
	VbrInstanceId        string           `position:"Query"`
	SmartAGId            string           `position:"Query"`
}

// DescribeGrantSagVbrRulesResponse is the response struct for api DescribeGrantSagVbrRules
type DescribeGrantSagVbrRulesResponse struct {
	*responses.BaseResponse
	TotalCount int                                  `json:"TotalCount" xml:"TotalCount"`
	RequestId  string                               `json:"RequestId" xml:"RequestId"`
	PageSize   int                                  `json:"PageSize" xml:"PageSize"`
	PageNumber int                                  `json:"PageNumber" xml:"PageNumber"`
	GrantRules GrantRulesInDescribeGrantSagVbrRules `json:"GrantRules" xml:"GrantRules"`
}

// CreateDescribeGrantSagVbrRulesRequest creates a request to invoke DescribeGrantSagVbrRules API
func CreateDescribeGrantSagVbrRulesRequest() (request *DescribeGrantSagVbrRulesRequest) {
	request = &DescribeGrantSagVbrRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DescribeGrantSagVbrRules", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeGrantSagVbrRulesResponse creates a response to parse from DescribeGrantSagVbrRules response
func CreateDescribeGrantSagVbrRulesResponse() (response *DescribeGrantSagVbrRulesResponse) {
	response = &DescribeGrantSagVbrRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
