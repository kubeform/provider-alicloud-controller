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

type Host struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HostSpec   `json:"spec,omitempty"`
	Status            HostStatus `json:"status,omitempty"`
}

type HostSpec struct {
	State *HostSpecResource `json:"state,omitempty" tf:"-"`

	Resource HostSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type HostSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ActiveAddressType *string `json:"activeAddressType" tf:"active_address_type"`
	// +optional
	Comment *string `json:"comment,omitempty" tf:"comment"`
	// +optional
	HostID   *string `json:"hostID,omitempty" tf:"host_id"`
	HostName *string `json:"hostName" tf:"host_name"`
	// +optional
	HostPrivateAddress *string `json:"hostPrivateAddress,omitempty" tf:"host_private_address"`
	// +optional
	HostPublicAddress *string `json:"hostPublicAddress,omitempty" tf:"host_public_address"`
	InstanceID        *string `json:"instanceID" tf:"instance_id"`
	// +optional
	InstanceRegionID *string `json:"instanceRegionID,omitempty" tf:"instance_region_id"`
	OsType           *string `json:"osType" tf:"os_type"`
	Source           *string `json:"source" tf:"source"`
	// +optional
	SourceInstanceID *string `json:"sourceInstanceID,omitempty" tf:"source_instance_id"`
}

type HostStatus struct {
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

// HostList is a list of Hosts
type HostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Host CRD objects
	Items []Host `json:"items,omitempty"`
}
