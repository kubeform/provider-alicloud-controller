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

// UpgradeDBClusterVersion invokes the polardb.UpgradeDBClusterVersion API synchronously
func (client *Client) UpgradeDBClusterVersion(request *UpgradeDBClusterVersionRequest) (response *UpgradeDBClusterVersionResponse, err error) {
	response = CreateUpgradeDBClusterVersionResponse()
	err = client.DoAction(request, response)
	return
}

// UpgradeDBClusterVersionWithChan invokes the polardb.UpgradeDBClusterVersion API asynchronously
func (client *Client) UpgradeDBClusterVersionWithChan(request *UpgradeDBClusterVersionRequest) (<-chan *UpgradeDBClusterVersionResponse, <-chan error) {
	responseChan := make(chan *UpgradeDBClusterVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpgradeDBClusterVersion(request)
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

// UpgradeDBClusterVersionWithCallback invokes the polardb.UpgradeDBClusterVersion API asynchronously
func (client *Client) UpgradeDBClusterVersionWithCallback(request *UpgradeDBClusterVersionRequest, callback func(response *UpgradeDBClusterVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpgradeDBClusterVersionResponse
		var err error
		defer close(result)
		response, err = client.UpgradeDBClusterVersion(request)
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

// UpgradeDBClusterVersionRequest is the request struct for api UpgradeDBClusterVersion
type UpgradeDBClusterVersionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	UpgradeType          string           `position:"Query" name:"UpgradeType"`
	PlannedEndTime       string           `position:"Query" name:"PlannedEndTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PlannedStartTime     string           `position:"Query" name:"PlannedStartTime"`
	FromTimeService      requests.Boolean `position:"Query" name:"FromTimeService"`
}

// UpgradeDBClusterVersionResponse is the response struct for api UpgradeDBClusterVersion
type UpgradeDBClusterVersionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpgradeDBClusterVersionRequest creates a request to invoke UpgradeDBClusterVersion API
func CreateUpgradeDBClusterVersionRequest() (request *UpgradeDBClusterVersionRequest) {
	request = &UpgradeDBClusterVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "UpgradeDBClusterVersion", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpgradeDBClusterVersionResponse creates a response to parse from UpgradeDBClusterVersion response
func CreateUpgradeDBClusterVersionResponse() (response *UpgradeDBClusterVersionResponse) {
	response = &UpgradeDBClusterVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
