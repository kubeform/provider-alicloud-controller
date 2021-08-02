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

type Interface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InterfaceSpec   `json:"spec,omitempty"`
	Status            InterfaceStatus `json:"status,omitempty"`
}

type InterfaceSpec struct {
	State *InterfaceSpecResource `json:"state,omitempty" tf:"-"`

	Resource InterfaceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type InterfaceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	Mac *string `json:"mac,omitempty" tf:"mac"`
	// +optional
	// Deprecated
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	NetworkInterfaceName *string `json:"networkInterfaceName,omitempty" tf:"network_interface_name"`
	// +optional
	PrimaryIPAddress *string `json:"primaryIPAddress,omitempty" tf:"primary_ip_address"`
	// +optional
	// Deprecated
	PrivateIP *string `json:"privateIP,omitempty" tf:"private_ip"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	PrivateIPAddresses []string `json:"privateIPAddresses,omitempty" tf:"private_ip_addresses"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	// Deprecated
	PrivateIPS []string `json:"privateIPS,omitempty" tf:"private_ips"`
	// +optional
	// Deprecated
	PrivateIPSCount *int64 `json:"privateIPSCount,omitempty" tf:"private_ips_count"`
	// +optional
	QueueNumber *int64 `json:"queueNumber,omitempty" tf:"queue_number"`
	// +optional
	ResourceGroupID *string `json:"resourceGroupID,omitempty" tf:"resource_group_id"`
	// +optional
	SecondaryPrivateIPAddressCount *int64 `json:"secondaryPrivateIPAddressCount,omitempty" tf:"secondary_private_ip_address_count"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	SecurityGroupIDS []string `json:"securityGroupIDS,omitempty" tf:"security_group_ids"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	// Deprecated
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	Tags      map[string]string `json:"tags,omitempty" tf:"tags"`
	VswitchID *string           `json:"vswitchID" tf:"vswitch_id"`
}

type InterfaceStatus struct {
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

// InterfaceList is a list of Interfaces
type InterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Interface CRD objects
	Items []Interface `json:"items,omitempty"`
}
