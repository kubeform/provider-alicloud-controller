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

// RemoveBackupsSet invokes the drds.RemoveBackupsSet API synchronously
func (client *Client) RemoveBackupsSet(request *RemoveBackupsSetRequest) (response *RemoveBackupsSetResponse, err error) {
	response = CreateRemoveBackupsSetResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveBackupsSetWithChan invokes the drds.RemoveBackupsSet API asynchronously
func (client *Client) RemoveBackupsSetWithChan(request *RemoveBackupsSetRequest) (<-chan *RemoveBackupsSetResponse, <-chan error) {
	responseChan := make(chan *RemoveBackupsSetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveBackupsSet(request)
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

// RemoveBackupsSetWithCallback invokes the drds.RemoveBackupsSet API asynchronously
func (client *Client) RemoveBackupsSetWithCallback(request *RemoveBackupsSetRequest, callback func(response *RemoveBackupsSetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveBackupsSetResponse
		var err error
		defer close(result)
		response, err = client.RemoveBackupsSet(request)
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

// RemoveBackupsSetRequest is the request struct for api RemoveBackupsSet
type RemoveBackupsSetRequest struct {
	*requests.RpcRequest
	BackupId       string `position:"Query" name:"BackupId"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

// RemoveBackupsSetResponse is the response struct for api RemoveBackupsSet
type RemoveBackupsSetResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    string `json:"Result" xml:"Result"`
}

// CreateRemoveBackupsSetRequest creates a request to invoke RemoveBackupsSet API
func CreateRemoveBackupsSetRequest() (request *RemoveBackupsSetRequest) {
	request = &RemoveBackupsSetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "RemoveBackupsSet", "drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRemoveBackupsSetResponse creates a response to parse from RemoveBackupsSet response
func CreateRemoveBackupsSetResponse() (response *RemoveBackupsSetResponse) {
	response = &RemoveBackupsSetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
