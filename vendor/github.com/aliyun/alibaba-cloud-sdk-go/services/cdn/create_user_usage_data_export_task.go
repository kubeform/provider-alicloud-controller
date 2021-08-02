package cdn

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

// CreateUserUsageDataExportTask invokes the cdn.CreateUserUsageDataExportTask API synchronously
func (client *Client) CreateUserUsageDataExportTask(request *CreateUserUsageDataExportTaskRequest) (response *CreateUserUsageDataExportTaskResponse, err error) {
	response = CreateCreateUserUsageDataExportTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateUserUsageDataExportTaskWithChan invokes the cdn.CreateUserUsageDataExportTask API asynchronously
func (client *Client) CreateUserUsageDataExportTaskWithChan(request *CreateUserUsageDataExportTaskRequest) (<-chan *CreateUserUsageDataExportTaskResponse, <-chan error) {
	responseChan := make(chan *CreateUserUsageDataExportTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateUserUsageDataExportTask(request)
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

// CreateUserUsageDataExportTaskWithCallback invokes the cdn.CreateUserUsageDataExportTask API asynchronously
func (client *Client) CreateUserUsageDataExportTaskWithCallback(request *CreateUserUsageDataExportTaskRequest, callback func(response *CreateUserUsageDataExportTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateUserUsageDataExportTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateUserUsageDataExportTask(request)
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

// CreateUserUsageDataExportTaskRequest is the request struct for api CreateUserUsageDataExportTask
type CreateUserUsageDataExportTaskRequest struct {
	*requests.RpcRequest
	TaskName  string           `position:"Query" name:"TaskName"`
	Language  string           `position:"Query" name:"Language"`
	StartTime string           `position:"Query" name:"StartTime"`
	EndTime   string           `position:"Query" name:"EndTime"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
}

// CreateUserUsageDataExportTaskResponse is the response struct for api CreateUserUsageDataExportTask
type CreateUserUsageDataExportTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	StartTime string `json:"StartTime" xml:"StartTime"`
	EndTime   string `json:"EndTime" xml:"EndTime"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
}

// CreateCreateUserUsageDataExportTaskRequest creates a request to invoke CreateUserUsageDataExportTask API
func CreateCreateUserUsageDataExportTaskRequest() (request *CreateUserUsageDataExportTaskRequest) {
	request = &CreateUserUsageDataExportTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "CreateUserUsageDataExportTask", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateUserUsageDataExportTaskResponse creates a response to parse from CreateUserUsageDataExportTask response
func CreateCreateUserUsageDataExportTaskResponse() (response *CreateUserUsageDataExportTaskResponse) {
	response = &CreateUserUsageDataExportTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
