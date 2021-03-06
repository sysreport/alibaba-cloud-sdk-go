package imm

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

// CreateFaceJob invokes the imm.CreateFaceJob API synchronously
// api document: https://help.aliyun.com/api/imm/createfacejob.html
func (client *Client) CreateFaceJob(request *CreateFaceJobRequest) (response *CreateFaceJobResponse, err error) {
	response = CreateCreateFaceJobResponse()
	err = client.DoAction(request, response)
	return
}

// CreateFaceJobWithChan invokes the imm.CreateFaceJob API asynchronously
// api document: https://help.aliyun.com/api/imm/createfacejob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateFaceJobWithChan(request *CreateFaceJobRequest) (<-chan *CreateFaceJobResponse, <-chan error) {
	responseChan := make(chan *CreateFaceJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateFaceJob(request)
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

// CreateFaceJobWithCallback invokes the imm.CreateFaceJob API asynchronously
// api document: https://help.aliyun.com/api/imm/createfacejob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateFaceJobWithCallback(request *CreateFaceJobRequest, callback func(response *CreateFaceJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateFaceJobResponse
		var err error
		defer close(result)
		response, err = client.CreateFaceJob(request)
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

// CreateFaceJobRequest is the request struct for api CreateFaceJob
type CreateFaceJobRequest struct {
	*requests.RpcRequest
	Project string `position:"Query" name:"Project"`
	SrcUri  string `position:"Query" name:"SrcUri"`
}

// CreateFaceJobResponse is the response struct for api CreateFaceJob
type CreateFaceJobResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	JobId      string `json:"JobId" xml:"JobId"`
	SetId      string `json:"SetId" xml:"SetId"`
	SrcUri     string `json:"SrcUri" xml:"SrcUri"`
	Percent    int    `json:"Percent" xml:"Percent"`
	CreateTime string `json:"CreateTime" xml:"CreateTime"`
	FinishTime string `json:"FinishTime" xml:"FinishTime"`
	Status     string `json:"Status" xml:"Status"`
}

// CreateCreateFaceJobRequest creates a request to invoke CreateFaceJob API
func CreateCreateFaceJobRequest() (request *CreateFaceJobRequest) {
	request = &CreateFaceJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "CreateFaceJob", "imm", "openAPI")
	return
}

// CreateCreateFaceJobResponse creates a response to parse from CreateFaceJob response
func CreateCreateFaceJobResponse() (response *CreateFaceJobResponse) {
	response = &CreateFaceJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
