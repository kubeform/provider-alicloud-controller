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

type BandwidthPackage struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BandwidthPackageSpec   `json:"spec,omitempty"`
	Status            BandwidthPackageStatus `json:"status,omitempty"`
}

type BandwidthPackageSpec struct {
	State *BandwidthPackageSpecResource `json:"state,omitempty" tf:"-"`

	Resource BandwidthPackageSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type BandwidthPackageSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Bandwidth *string `json:"bandwidth" tf:"bandwidth"`
	// +optional
	BandwidthPackageName *string `json:"bandwidthPackageName,omitempty" tf:"bandwidth_package_name"`
	// +optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	Force *string `json:"force,omitempty" tf:"force"`
	// +optional
	InternetChargeType *string `json:"internetChargeType,omitempty" tf:"internet_charge_type"`
	// +optional
	Isp *string `json:"isp,omitempty" tf:"isp"`
	// +optional
	// Deprecated
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Ratio *int64 `json:"ratio,omitempty" tf:"ratio"`
	// +optional
	ResourceGroupID *string `json:"resourceGroupID,omitempty" tf:"resource_group_id"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	Zone *string `json:"zone,omitempty" tf:"zone"`
}

type BandwidthPackageStatus struct {
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

// BandwidthPackageList is a list of BandwidthPackages
type BandwidthPackageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BandwidthPackage CRD objects
	Items []BandwidthPackage `json:"items,omitempty"`
}