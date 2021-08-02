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

type Actiontrail struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ActiontrailSpec   `json:"spec,omitempty"`
	Status            ActiontrailStatus `json:"status,omitempty"`
}

type ActiontrailSpec struct {
	State *ActiontrailSpecResource `json:"state,omitempty" tf:"-"`

	Resource ActiontrailSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ActiontrailSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	EventRw *string `json:"eventRw,omitempty" tf:"event_rw"`
	// +optional
	IsOrganizationTrail *bool `json:"isOrganizationTrail,omitempty" tf:"is_organization_trail"`
	// +optional
	// Deprecated
	MnsTopicArn *string `json:"mnsTopicArn,omitempty" tf:"mns_topic_arn"`
	// +optional
	// Deprecated
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	OssBucketName *string `json:"ossBucketName,omitempty" tf:"oss_bucket_name"`
	// +optional
	OssKeyPrefix *string `json:"ossKeyPrefix,omitempty" tf:"oss_key_prefix"`
	// +optional
	OssWriteRoleArn *string `json:"ossWriteRoleArn,omitempty" tf:"oss_write_role_arn"`
	// +optional
	// Deprecated
	RoleName *string `json:"roleName,omitempty" tf:"role_name"`
	// +optional
	SlsProjectArn *string `json:"slsProjectArn,omitempty" tf:"sls_project_arn"`
	// +optional
	SlsWriteRoleArn *string `json:"slsWriteRoleArn,omitempty" tf:"sls_write_role_arn"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	TrailName *string `json:"trailName,omitempty" tf:"trail_name"`
	// +optional
	TrailRegion *string `json:"trailRegion,omitempty" tf:"trail_region"`
}

type ActiontrailStatus struct {
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

// ActiontrailList is a list of Actiontrails
type ActiontrailList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Actiontrail CRD objects
	Items []Actiontrail `json:"items,omitempty"`
}
