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

// EnableSmartAGDpiMonitor invokes the smartag.EnableSmartAGDpiMonitor API synchronously
func (client *Client) EnableSmartAGDpiMonitor(request *EnableSmartAGDpiMonitorRequest) (response *EnableSmartAGDpiMonitorResponse, err error) {
	response = CreateEnableSmartAGDpiMonitorResponse()
	err = client.DoAction(request, response)
	return
}

// EnableSmartAGDpiMonitorWithChan invokes the smartag.EnableSmartAGDpiMonitor API asynchronously
func (client *Client) EnableSmartAGDpiMonitorWithChan(request *EnableSmartAGDpiMonitorRequest) (<-chan *EnableSmartAGDpiMonitorResponse, <-chan error) {
	responseChan := make(chan *EnableSmartAGDpiMonitorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableSmartAGDpiMonitor(request)
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

// EnableSmartAGDpiMonitorWithCallback invokes the smartag.EnableSmartAGDpiMonitor API asynchronously
func (client *Client) EnableSmartAGDpiMonitorWithCallback(request *EnableSmartAGDpiMonitorRequest, callback func(response *EnableSmartAGDpiMonitorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableSmartAGDpiMonitorResponse
		var err error
		defer close(result)
		response, err = client.EnableSmartAGDpiMonitor(request)
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

// EnableSmartAGDpiMonitorRequest is the request struct for api EnableSmartAGDpiMonitor
type EnableSmartAGDpiMonitorRequest struct {
	*requests.RpcRequest
	SlsLogStore          string           `position:"Query"`
	ResourceOwnerId      requests.Integer `position:"Query"`
	ClientToken          string           `position:"Query"`
	SlsProjectName       string           `position:"Query"`
	DryRun               requests.Boolean `position:"Query"`
	ResourceOwnerAccount string           `position:"Query"`
	OwnerAccount         string           `position:"Query"`
	OwnerId              requests.Integer `position:"Query"`
	SmartAGId            string           `position:"Query"`
}

// EnableSmartAGDpiMonitorResponse is the response struct for api EnableSmartAGDpiMonitor
type EnableSmartAGDpiMonitorResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEnableSmartAGDpiMonitorRequest creates a request to invoke EnableSmartAGDpiMonitor API
func CreateEnableSmartAGDpiMonitorRequest() (request *EnableSmartAGDpiMonitorRequest) {
	request = &EnableSmartAGDpiMonitorRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "EnableSmartAGDpiMonitor", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnableSmartAGDpiMonitorResponse creates a response to parse from EnableSmartAGDpiMonitor response
func CreateEnableSmartAGDpiMonitorResponse() (response *EnableSmartAGDpiMonitorResponse) {
	response = &EnableSmartAGDpiMonitorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
