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

type Snapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnapshotSpec   `json:"spec,omitempty"`
	Status            SnapshotStatus `json:"status,omitempty"`
}

type SnapshotSpec struct {
	State *SnapshotSpecResource `json:"state,omitempty" tf:"-"`

	Resource SnapshotSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SnapshotSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Category *string `json:"category,omitempty" tf:"category"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	DiskID      *string `json:"diskID" tf:"disk_id"`
	// +optional
	Force *bool `json:"force,omitempty" tf:"force"`
	// +optional
	InstantAccess *bool `json:"instantAccess,omitempty" tf:"instant_access"`
	// +optional
	InstantAccessRetentionDays *int64 `json:"instantAccessRetentionDays,omitempty" tf:"instant_access_retention_days"`
	// +optional
	// Deprecated
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	ResourceGroupID *string `json:"resourceGroupID,omitempty" tf:"resource_group_id"`
	// +optional
	RetentionDays *int64 `json:"retentionDays,omitempty" tf:"retention_days"`
	// +optional
	SnapshotName *string `json:"snapshotName,omitempty" tf:"snapshot_name"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

type SnapshotStatus struct {
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

// SnapshotList is a list of Snapshots
type SnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Snapshot CRD objects
	Items []Snapshot `json:"items,omitempty"`
}
