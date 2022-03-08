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

type BastionhostInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BastionhostInstanceSpec   `json:"spec,omitempty"`
	Status            BastionhostInstanceStatus `json:"status,omitempty"`
}

type BastionhostInstanceSpec struct {
	State *BastionhostInstanceSpecResource `json:"state,omitempty" tf:"-"`

	Resource BastionhostInstanceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type BastionhostInstanceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Description *string `json:"description" tf:"description"`
	// +optional
	EnablePublicAccess *bool   `json:"enablePublicAccess,omitempty" tf:"enable_public_access"`
	LicenseCode        *string `json:"licenseCode" tf:"license_code"`
	// +optional
	Period *int64 `json:"period,omitempty" tf:"period"`
	// +optional
	ResourceGroupID  *string  `json:"resourceGroupID,omitempty" tf:"resource_group_id"`
	SecurityGroupIDS []string `json:"securityGroupIDS" tf:"security_group_ids"`
	// +optional
	Tags      map[string]string `json:"tags,omitempty" tf:"tags"`
	VswitchID *string           `json:"vswitchID" tf:"vswitch_id"`
}

type BastionhostInstanceStatus struct {
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

// BastionhostInstanceList is a list of BastionhostInstances
type BastionhostInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BastionhostInstance CRD objects
	Items []BastionhostInstance `json:"items,omitempty"`
}
