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

// UnicomSignConfirm invokes the smartag.UnicomSignConfirm API synchronously
func (client *Client) UnicomSignConfirm(request *UnicomSignConfirmRequest) (response *UnicomSignConfirmResponse, err error) {
	response = CreateUnicomSignConfirmResponse()
	err = client.DoAction(request, response)
	return
}

// UnicomSignConfirmWithChan invokes the smartag.UnicomSignConfirm API asynchronously
func (client *Client) UnicomSignConfirmWithChan(request *UnicomSignConfirmRequest) (<-chan *UnicomSignConfirmResponse, <-chan error) {
	responseChan := make(chan *UnicomSignConfirmResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnicomSignConfirm(request)
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

// UnicomSignConfirmWithCallback invokes the smartag.UnicomSignConfirm API asynchronously
func (client *Client) UnicomSignConfirmWithCallback(request *UnicomSignConfirmRequest, callback func(response *UnicomSignConfirmResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnicomSignConfirmResponse
		var err error
		defer close(result)
		response, err = client.UnicomSignConfirm(request)
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

// UnicomSignConfirmRequest is the request struct for api UnicomSignConfirm
type UnicomSignConfirmRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer             `position:"Query"`
	ResourceOwnerAccount string                       `position:"Query"`
	OwnerAccount         string                       `position:"Query"`
	TmsOrder             *[]UnicomSignConfirmTmsOrder `position:"Query" name:"TmsOrder"  type:"Repeated"`
	OwnerId              requests.Integer             `position:"Query"`
}

// UnicomSignConfirmTmsOrder is a repeated param struct in UnicomSignConfirmRequest
type UnicomSignConfirmTmsOrder struct {
	TmsCode      string `name:"TmsCode"`
	SigningTime  string `name:"SigningTime"`
	TmsOrderCode string `name:"TmsOrderCode"`
	TradeId      string `name:"TradeId"`
}

// UnicomSignConfirmResponse is the response struct for api UnicomSignConfirm
type UnicomSignConfirmResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUnicomSignConfirmRequest creates a request to invoke UnicomSignConfirm API
func CreateUnicomSignConfirmRequest() (request *UnicomSignConfirmRequest) {
	request = &UnicomSignConfirmRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "UnicomSignConfirm", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUnicomSignConfirmResponse creates a response to parse from UnicomSignConfirm response
func CreateUnicomSignConfirmResponse() (response *UnicomSignConfirmResponse) {
	response = &UnicomSignConfirmResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
