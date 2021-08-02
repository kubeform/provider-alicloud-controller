package drds

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

// CreateShardTask invokes the drds.CreateShardTask API synchronously
func (client *Client) CreateShardTask(request *CreateShardTaskRequest) (response *CreateShardTaskResponse, err error) {
	response = CreateCreateShardTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateShardTaskWithChan invokes the drds.CreateShardTask API asynchronously
func (client *Client) CreateShardTaskWithChan(request *CreateShardTaskRequest) (<-chan *CreateShardTaskResponse, <-chan error) {
	responseChan := make(chan *CreateShardTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateShardTask(request)
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

// CreateShardTaskWithCallback invokes the drds.CreateShardTask API asynchronously
func (client *Client) CreateShardTaskWithCallback(request *CreateShardTaskRequest, callback func(response *CreateShardTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateShardTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateShardTask(request)
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

// CreateShardTaskRequest is the request struct for api CreateShardTask
type CreateShardTaskRequest struct {
	*requests.RpcRequest
	TaskType        string `position:"Query" name:"TaskType"`
	TargetTableName string `position:"Query" name:"TargetTableName"`
	DrdsInstanceId  string `position:"Query" name:"DrdsInstanceId"`
	DbName          string `position:"Query" name:"DbName"`
	SourceTableName string `position:"Query" name:"SourceTableName"`
}

// CreateShardTaskResponse is the response struct for api CreateShardTask
type CreateShardTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateCreateShardTaskRequest creates a request to invoke CreateShardTask API
func CreateCreateShardTaskRequest() (request *CreateShardTaskRequest) {
	request = &CreateShardTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "CreateShardTask", "drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateShardTaskResponse creates a response to parse from CreateShardTask response
func CreateCreateShardTaskResponse() (response *CreateShardTaskResponse) {
	response = &CreateShardTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
