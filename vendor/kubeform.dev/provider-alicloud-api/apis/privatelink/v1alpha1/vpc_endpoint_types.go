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

type VpcEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcEndpointSpec   `json:"spec,omitempty"`
	Status            VpcEndpointStatus `json:"status,omitempty"`
}

type VpcEndpointSpec struct {
	State *VpcEndpointSpecResource `json:"state,omitempty" tf:"-"`

	Resource VpcEndpointSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type VpcEndpointSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Bandwidth *int64 `json:"bandwidth,omitempty" tf:"bandwidth"`
	// +optional
	ConnectionStatus *string `json:"connectionStatus,omitempty" tf:"connection_status"`
	// +optional
	DryRun *bool `json:"dryRun,omitempty" tf:"dry_run"`
	// +optional
	EndpointBusinessStatus *string `json:"endpointBusinessStatus,omitempty" tf:"endpoint_business_status"`
	// +optional
	EndpointDescription *string `json:"endpointDescription,omitempty" tf:"endpoint_description"`
	// +optional
	EndpointDomain   *string  `json:"endpointDomain,omitempty" tf:"endpoint_domain"`
	SecurityGroupIDS []string `json:"securityGroupIDS" tf:"security_group_ids"`
	// +optional
	ServiceID *string `json:"serviceID,omitempty" tf:"service_id"`
	// +optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	VpcEndpointName *string `json:"vpcEndpointName,omitempty" tf:"vpc_endpoint_name"`
	VpcID           *string `json:"vpcID" tf:"vpc_id"`
}

type VpcEndpointStatus struct {
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

// VpcEndpointList is a list of VpcEndpoints
type VpcEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpcEndpoint CRD objects
	Items []VpcEndpoint `json:"items,omitempty"`
}
