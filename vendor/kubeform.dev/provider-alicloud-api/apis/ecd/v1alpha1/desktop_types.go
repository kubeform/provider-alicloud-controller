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

type Desktop struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DesktopSpec   `json:"spec,omitempty"`
	Status            DesktopStatus `json:"status,omitempty"`
}

type DesktopSpec struct {
	State *DesktopSpecResource `json:"state,omitempty" tf:"-"`

	Resource DesktopSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type DesktopSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Amount *int64 `json:"amount,omitempty" tf:"amount"`
	// +optional
	AutoPay *bool `json:"autoPay,omitempty" tf:"auto_pay"`
	// +optional
	AutoRenew *bool   `json:"autoRenew,omitempty" tf:"auto_renew"`
	BundleID  *string `json:"bundleID" tf:"bundle_id"`
	// +optional
	DesktopName *string `json:"desktopName,omitempty" tf:"desktop_name"`
	// +optional
	DesktopType *string `json:"desktopType,omitempty" tf:"desktop_type"`
	// +optional
	EndUserIDS []string `json:"endUserIDS,omitempty" tf:"end_user_ids"`
	// +optional
	HostName     *string `json:"hostName,omitempty" tf:"host_name"`
	OfficeSiteID *string `json:"officeSiteID" tf:"office_site_id"`
	// +optional
	PaymentType *string `json:"paymentType,omitempty" tf:"payment_type"`
	// +optional
	Period *int64 `json:"period,omitempty" tf:"period"`
	// +optional
	PeriodUnit    *string `json:"periodUnit,omitempty" tf:"period_unit"`
	PolicyGroupID *string `json:"policyGroupID" tf:"policy_group_id"`
	// +optional
	RootDiskSizeGib *int64 `json:"rootDiskSizeGib,omitempty" tf:"root_disk_size_gib"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	StoppedMode *string `json:"stoppedMode,omitempty" tf:"stopped_mode"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	UserAssignMode *string `json:"userAssignMode,omitempty" tf:"user_assign_mode"`
	// +optional
	UserDiskSizeGib *int64 `json:"userDiskSizeGib,omitempty" tf:"user_disk_size_gib"`
}

type DesktopStatus struct {
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

// DesktopList is a list of Desktops
type DesktopList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Desktop CRD objects
	Items []Desktop `json:"items,omitempty"`
}
