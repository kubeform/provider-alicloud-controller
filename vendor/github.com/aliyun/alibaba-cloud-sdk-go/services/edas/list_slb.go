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

// ListSlb invokes the edas.ListSlb API synchronously
func (client *Client) ListSlb(request *ListSlbRequest) (response *ListSlbResponse, err error) {
	response = CreateListSlbResponse()
	err = client.DoAction(request, response)
	return
}

// ListSlbWithChan invokes the edas.ListSlb API asynchronously
func (client *Client) ListSlbWithChan(request *ListSlbRequest) (<-chan *ListSlbResponse, <-chan error) {
	responseChan := make(chan *ListSlbResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSlb(request)
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

// ListSlbWithCallback invokes the edas.ListSlb API asynchronously
func (client *Client) ListSlbWithCallback(request *ListSlbRequest, callback func(response *ListSlbResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSlbResponse
		var err error
		defer close(result)
		response, err = client.ListSlb(request)
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

// ListSlbRequest is the request struct for api ListSlb
type ListSlbRequest struct {
	*requests.RoaRequest
}

// ListSlbResponse is the response struct for api ListSlb
type ListSlbResponse struct {
	*responses.BaseResponse
	Code      int              `json:"Code" xml:"Code"`
	Message   string           `json:"Message" xml:"Message"`
	RequestId string           `json:"RequestId" xml:"RequestId"`
	SlbList   SlbListInListSlb `json:"SlbList" xml:"SlbList"`
}

// CreateListSlbRequest creates a request to invoke ListSlb API
func CreateListSlbRequest() (request *ListSlbRequest) {
	request = &ListSlbRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "ListSlb", "/pop/v5/slb_list", "Edas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListSlbResponse creates a response to parse from ListSlb response
func CreateListSlbResponse() (response *ListSlbResponse) {
	response = &ListSlbResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
