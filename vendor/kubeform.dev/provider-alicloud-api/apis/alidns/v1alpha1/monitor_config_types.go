/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type MonitorConfig struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorConfigSpec   `json:"spec,omitempty"`
	Status            MonitorConfigStatus `json:"status,omitempty"`
}

type MonitorConfigSpecIspCityNode struct {
	CityCode *string `json:"cityCode" tf:"city_code"`
	IspCode  *string `json:"ispCode" tf:"isp_code"`
}

type MonitorConfigSpec struct {
	State *MonitorConfigSpecResource `json:"state,omitempty" tf:"-"`

	Resource MonitorConfigSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type MonitorConfigSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AddrPoolID      *string                        `json:"addrPoolID" tf:"addr_pool_id"`
	EvaluationCount *int64                         `json:"evaluationCount" tf:"evaluation_count"`
	Interval        *int64                         `json:"interval" tf:"interval"`
	IspCityNode     []MonitorConfigSpecIspCityNode `json:"ispCityNode" tf:"isp_city_node"`
	// +optional
	Lang              *string `json:"lang,omitempty" tf:"lang"`
	MonitorExtendInfo *string `json:"monitorExtendInfo" tf:"monitor_extend_info"`
	ProtocolType      *string `json:"protocolType" tf:"protocol_type"`
	Timeout           *int64  `json:"timeout" tf:"timeout"`
}

type MonitorConfigStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MonitorConfigList is a list of MonitorConfigs
type MonitorConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MonitorConfig CRD objects
	Items []MonitorConfig `json:"items,omitempty"`
}