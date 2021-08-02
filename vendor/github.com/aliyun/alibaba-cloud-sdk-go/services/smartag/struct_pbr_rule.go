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

// PbrRule is a nested struct in smartag response
type PbrRule struct {
	SrcPort     string `json:"SrcPort" xml:"SrcPort"`
	Description string `json:"Description" xml:"Description"`
	SrcCidr     string `json:"SrcCidr" xml:"SrcCidr"`
	DstPort     string `json:"DstPort" xml:"DstPort"`
	PbrRuleId   string `json:"PbrRuleId" xml:"PbrRuleId"`
	Protocol    string `json:"Protocol" xml:"Protocol"`
	DstCidr     string `json:"DstCidr" xml:"DstCidr"`
	Name        string `json:"Name" xml:"Name"`
}
