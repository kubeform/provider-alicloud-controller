package edas

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

// ListRecentChangeOrder invokes the edas.ListRecentChangeOrder API synchronously
func (client *Client) ListRecentChangeOrder(request *ListRecentChangeOrderRequest) (response *ListRecentChangeOrderResponse, err error) {
	response = CreateListRecentChangeOrderResponse()
	err = client.DoAction(request, response)
	return
}

// ListRecentChangeOrderWithChan invokes the edas.ListRecentChangeOrder API asynchronously
func (client *Client) ListRecentChangeOrderWithChan(request *ListRecentChangeOrderRequest) (<-chan *ListRecentChangeOrderResponse, <-chan error) {
	responseChan := make(chan *ListRecentChangeOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRecentChangeOrder(request)
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

// ListRecentChangeOrderWithCallback invokes the edas.ListRecentChangeOrder API asynchronously
func (client *Client) ListRecentChangeOrderWithCallback(request *ListRecentChangeOrderRequest, callback func(response *ListRecentChangeOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRecentChangeOrderResponse
		var err error
		defer close(result)
		response, err = client.ListRecentChangeOrder(request)
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

// ListRecentChangeOrderRequest is the request struct for api ListRecentChangeOrder
type ListRecentChangeOrderRequest struct {
	*requests.RoaRequest
	AppId string `position:"Query" name:"AppId"`
}

// ListRecentChangeOrderResponse is the response struct for api ListRecentChangeOrder
type ListRecentChangeOrderResponse struct {
	*responses.BaseResponse
	Code            int             `json:"Code" xml:"Code"`
	Message         string          `json:"Message" xml:"Message"`
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	ChangeOrderList ChangeOrderList `json:"ChangeOrderList" xml:"ChangeOrderList"`
}

// CreateListRecentChangeOrderRequest creates a request to invoke ListRecentChangeOrder API
func CreateListRecentChangeOrderRequest() (request *ListRecentChangeOrderRequest) {
	request = &ListRecentChangeOrderRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "ListRecentChangeOrder", "/pop/v5/changeorder/change_order_list", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListRecentChangeOrderResponse creates a response to parse from ListRecentChangeOrder response
func CreateListRecentChangeOrderResponse() (response *ListRecentChangeOrderResponse) {
	response = &ListRecentChangeOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
