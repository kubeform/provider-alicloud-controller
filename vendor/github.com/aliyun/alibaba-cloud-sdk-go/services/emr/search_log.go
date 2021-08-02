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

// SearchLog invokes the emr.SearchLog API synchronously
func (client *Client) SearchLog(request *SearchLogRequest) (response *SearchLogResponse, err error) {
	response = CreateSearchLogResponse()
	err = client.DoAction(request, response)
	return
}

// SearchLogWithChan invokes the emr.SearchLog API asynchronously
func (client *Client) SearchLogWithChan(request *SearchLogRequest) (<-chan *SearchLogResponse, <-chan error) {
	responseChan := make(chan *SearchLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchLog(request)
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

// SearchLogWithCallback invokes the emr.SearchLog API asynchronously
func (client *Client) SearchLogWithCallback(request *SearchLogRequest, callback func(response *SearchLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchLogResponse
		var err error
		defer close(result)
		response, err = client.SearchLog(request)
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

// SearchLogRequest is the request struct for api SearchLog
type SearchLogRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Line            requests.Integer `position:"Query" name:"Line"`
	HostName        string           `position:"Query" name:"HostName"`
	LogstoreName    string           `position:"Query" name:"LogstoreName"`
	FromTimestamp   requests.Integer `position:"Query" name:"FromTimestamp"`
	Offset          requests.Integer `position:"Query" name:"Offset"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
	Reverse         requests.Boolean `position:"Query" name:"Reverse"`
	HostInnerIp     string           `position:"Query" name:"HostInnerIp"`
	ToTimestamp     requests.Integer `position:"Query" name:"ToTimestamp"`
	SlsQueryString  string           `position:"Query" name:"SlsQueryString"`
}

// SearchLogResponse is the response struct for api SearchLog
type SearchLogResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	Completed      bool           `json:"Completed" xml:"Completed"`
	SlsLogItemList SlsLogItemList `json:"SlsLogItemList" xml:"SlsLogItemList"`
}

// CreateSearchLogRequest creates a request to invoke SearchLog API
func CreateSearchLogRequest() (request *SearchLogRequest) {
	request = &SearchLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "SearchLog", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSearchLogResponse creates a response to parse from SearchLog response
func CreateSearchLogResponse() (response *SearchLogResponse) {
	response = &SearchLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
