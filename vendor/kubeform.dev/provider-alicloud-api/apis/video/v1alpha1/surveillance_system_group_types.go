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

type SurveillanceSystemGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SurveillanceSystemGroupSpec   `json:"spec,omitempty"`
	Status            SurveillanceSystemGroupStatus `json:"status,omitempty"`
}

type SurveillanceSystemGroupSpec struct {
	State *SurveillanceSystemGroupSpecResource `json:"state,omitempty" tf:"-"`

	Resource SurveillanceSystemGroupSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SurveillanceSystemGroupSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Callback *string `json:"callback,omitempty" tf:"callback"`
	// +optional
	CaptureImage *int64 `json:"captureImage,omitempty" tf:"capture_image"`
	// +optional
	CaptureInterval *int64 `json:"captureInterval,omitempty" tf:"capture_interval"`
	// +optional
	CaptureOssBucket *string `json:"captureOssBucket,omitempty" tf:"capture_oss_bucket"`
	// +optional
	CaptureOssPath *string `json:"captureOssPath,omitempty" tf:"capture_oss_path"`
	// +optional
	CaptureVideo *int64 `json:"captureVideo,omitempty" tf:"capture_video"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	Enabled    *bool   `json:"enabled,omitempty" tf:"enabled"`
	GroupName  *string `json:"groupName" tf:"group_name"`
	InProtocol *string `json:"inProtocol" tf:"in_protocol"`
	// +optional
	LazyPull    *bool   `json:"lazyPull,omitempty" tf:"lazy_pull"`
	OutProtocol *string `json:"outProtocol" tf:"out_protocol"`
	PlayDomain  *string `json:"playDomain" tf:"play_domain"`
	PushDomain  *string `json:"pushDomain" tf:"push_domain"`
	// +optional
	Status *bool `json:"status,omitempty" tf:"status"`
}

type SurveillanceSystemGroupStatus struct {
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

// SurveillanceSystemGroupList is a list of SurveillanceSystemGroups
type SurveillanceSystemGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SurveillanceSystemGroup CRD objects
	Items []SurveillanceSystemGroup `json:"items,omitempty"`
}