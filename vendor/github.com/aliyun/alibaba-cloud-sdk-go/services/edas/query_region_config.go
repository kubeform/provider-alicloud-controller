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

// QueryRegionConfig invokes the edas.QueryRegionConfig API synchronously
func (client *Client) QueryRegionConfig(request *QueryRegionConfigRequest) (response *QueryRegionConfigResponse, err error) {
	response = CreateQueryRegionConfigResponse()
	err = client.DoAction(request, response)
	return
}

// QueryRegionConfigWithChan invokes the edas.QueryRegionConfig API asynchronously
func (client *Client) QueryRegionConfigWithChan(request *QueryRegionConfigRequest) (<-chan *QueryRegionConfigResponse, <-chan error) {
	responseChan := make(chan *QueryRegionConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryRegionConfig(request)
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

// QueryRegionConfigWithCallback invokes the edas.QueryRegionConfig API asynchronously
func (client *Client) QueryRegionConfigWithCallback(request *QueryRegionConfigRequest, callback func(response *QueryRegionConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryRegionConfigResponse
		var err error
		defer close(result)
		response, err = client.QueryRegionConfig(request)
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

// QueryRegionConfigRequest is the request struct for api QueryRegionConfig
type QueryRegionConfigRequest struct {
	*requests.RoaRequest
}

// QueryRegionConfigResponse is the response struct for api QueryRegionConfig
type QueryRegionConfigResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	Code         int          `json:"Code" xml:"Code"`
	Message      string       `json:"Message" xml:"Message"`
	RegionConfig RegionConfig `json:"RegionConfig" xml:"RegionConfig"`
}

// CreateQueryRegionConfigRequest creates a request to invoke QueryRegionConfig API
func CreateQueryRegionConfigRequest() (request *QueryRegionConfigRequest) {
	request = &QueryRegionConfigRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "QueryRegionConfig", "/pop/v5/region_config", "Edas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateQueryRegionConfigResponse creates a response to parse from QueryRegionConfig response
func CreateQueryRegionConfigResponse() (response *QueryRegionConfigResponse) {
	response = &QueryRegionConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
