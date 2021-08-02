package slb

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

// LoadBalancer is a nested struct in slb response
type LoadBalancer struct {
	LoadBalancerId               string                      `json:"LoadBalancerId" xml:"LoadBalancerId"`
	LoadBalancerName             string                      `json:"LoadBalancerName" xml:"LoadBalancerName"`
	LoadBalancerStatus           string                      `json:"LoadBalancerStatus" xml:"LoadBalancerStatus"`
	Address                      string                      `json:"Address" xml:"Address"`
	AddressType                  string                      `json:"AddressType" xml:"AddressType"`
	RegionId                     string                      `json:"RegionId" xml:"RegionId"`
	RegionIdAlias                string                      `json:"RegionIdAlias" xml:"RegionIdAlias"`
	VSwitchId                    string                      `json:"VSwitchId" xml:"VSwitchId"`
	VpcId                        string                      `json:"VpcId" xml:"VpcId"`
	NetworkType                  string                      `json:"NetworkType" xml:"NetworkType"`
	MasterZoneId                 string                      `json:"MasterZoneId" xml:"MasterZoneId"`
	SlaveZoneId                  string                      `json:"SlaveZoneId" xml:"SlaveZoneId"`
	InternetChargeType           string                      `json:"InternetChargeType" xml:"InternetChargeType"`
	CreateTime                   string                      `json:"CreateTime" xml:"CreateTime"`
	CreateTimeStamp              int64                       `json:"CreateTimeStamp" xml:"CreateTimeStamp"`
	PayType                      string                      `json:"PayType" xml:"PayType"`
	ResourceGroupId              string                      `json:"ResourceGroupId" xml:"ResourceGroupId"`
	AddressIPVersion             string                      `json:"AddressIPVersion" xml:"AddressIPVersion"`
	BusinessStatus               string                      `json:"BusinessStatus" xml:"BusinessStatus"`
	ModificationProtectionStatus string                      `json:"ModificationProtectionStatus" xml:"ModificationProtectionStatus"`
	ModificationProtectionReason string                      `json:"ModificationProtectionReason" xml:"ModificationProtectionReason"`
	Bandwidth                    int                         `json:"Bandwidth" xml:"Bandwidth"`
	InternetChargeTypeAlias      string                      `json:"InternetChargeTypeAlias" xml:"InternetChargeTypeAlias"`
	LoadBalancerSpec             string                      `json:"LoadBalancerSpec" xml:"LoadBalancerSpec"`
	DeleteProtection             string                      `json:"DeleteProtection" xml:"DeleteProtection"`
	Tags                         TagsInDescribeLoadBalancers `json:"Tags" xml:"Tags"`
}
