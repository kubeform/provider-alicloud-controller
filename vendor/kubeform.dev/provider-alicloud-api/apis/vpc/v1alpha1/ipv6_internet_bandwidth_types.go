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

type Ipv6InternetBandwidth struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ipv6InternetBandwidthSpec   `json:"spec,omitempty"`
	Status            Ipv6InternetBandwidthStatus `json:"status,omitempty"`
}

type Ipv6InternetBandwidthSpec struct {
	State *Ipv6InternetBandwidthSpecResource `json:"state,omitempty" tf:"-"`

	Resource Ipv6InternetBandwidthSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type Ipv6InternetBandwidthSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Bandwidth *int64 `json:"bandwidth" tf:"bandwidth"`
	// +optional
	InternetChargeType *string `json:"internetChargeType,omitempty" tf:"internet_charge_type"`
	Ipv6AddressID      *string `json:"ipv6AddressID" tf:"ipv6_address_id"`
	Ipv6GatewayID      *string `json:"ipv6GatewayID" tf:"ipv6_gateway_id"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
}

type Ipv6InternetBandwidthStatus struct {
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

// Ipv6InternetBandwidthList is a list of Ipv6InternetBandwidths
type Ipv6InternetBandwidthList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Ipv6InternetBandwidth CRD objects
	Items []Ipv6InternetBandwidth `json:"items,omitempty"`
}
