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

type MigrationInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MigrationInstanceSpec   `json:"spec,omitempty"`
	Status            MigrationInstanceStatus `json:"status,omitempty"`
}

type MigrationInstanceSpec struct {
	State *MigrationInstanceSpecResource `json:"state,omitempty" tf:"-"`

	Resource MigrationInstanceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type MigrationInstanceSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ComputeUnit *int64 `json:"computeUnit,omitempty" tf:"compute_unit"`
	// +optional
	DatabaseCount                 *int64  `json:"databaseCount,omitempty" tf:"database_count"`
	DestinationEndpointEngineName *string `json:"destinationEndpointEngineName" tf:"destination_endpoint_engine_name"`
	DestinationEndpointRegion     *string `json:"destinationEndpointRegion" tf:"destination_endpoint_region"`
	// +optional
	DtsInstanceID *string `json:"dtsInstanceID,omitempty" tf:"dts_instance_id"`
	// +optional
	InstanceClass            *string `json:"instanceClass,omitempty" tf:"instance_class"`
	PaymentType              *string `json:"paymentType" tf:"payment_type"`
	SourceEndpointEngineName *string `json:"sourceEndpointEngineName" tf:"source_endpoint_engine_name"`
	SourceEndpointRegion     *string `json:"sourceEndpointRegion" tf:"source_endpoint_region"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	SyncArchitecture *string `json:"syncArchitecture,omitempty" tf:"sync_architecture"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

type MigrationInstanceStatus struct {
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

// MigrationInstanceList is a list of MigrationInstances
type MigrationInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MigrationInstance CRD objects
	Items []MigrationInstance `json:"items,omitempty"`
}
