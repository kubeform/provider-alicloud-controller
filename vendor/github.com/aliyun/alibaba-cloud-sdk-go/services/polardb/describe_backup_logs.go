package polardb

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

// DescribeBackupLogs invokes the polardb.DescribeBackupLogs API synchronously
func (client *Client) DescribeBackupLogs(request *DescribeBackupLogsRequest) (response *DescribeBackupLogsResponse, err error) {
	response = CreateDescribeBackupLogsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBackupLogsWithChan invokes the polardb.DescribeBackupLogs API asynchronously
func (client *Client) DescribeBackupLogsWithChan(request *DescribeBackupLogsRequest) (<-chan *DescribeBackupLogsResponse, <-chan error) {
	responseChan := make(chan *DescribeBackupLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBackupLogs(request)
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

// DescribeBackupLogsWithCallback invokes the polardb.DescribeBackupLogs API asynchronously
func (client *Client) DescribeBackupLogsWithCallback(request *DescribeBackupLogsRequest, callback func(response *DescribeBackupLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBackupLogsResponse
		var err error
		defer close(result)
		response, err = client.DescribeBackupLogs(request)
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

// DescribeBackupLogsRequest is the request struct for api DescribeBackupLogs
type DescribeBackupLogsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeBackupLogsResponse is the response struct for api DescribeBackupLogs
type DescribeBackupLogsResponse struct {
	*responses.BaseResponse
	RequestId        string                    `json:"RequestId" xml:"RequestId"`
	TotalRecordCount string                    `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumber       string                    `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  string                    `json:"PageRecordCount" xml:"PageRecordCount"`
	Items            ItemsInDescribeBackupLogs `json:"Items" xml:"Items"`
}

// CreateDescribeBackupLogsRequest creates a request to invoke DescribeBackupLogs API
func CreateDescribeBackupLogsRequest() (request *DescribeBackupLogsRequest) {
	request = &DescribeBackupLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DescribeBackupLogs", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeBackupLogsResponse creates a response to parse from DescribeBackupLogs response
func CreateDescribeBackupLogsResponse() (response *DescribeBackupLogsResponse) {
	response = &DescribeBackupLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}