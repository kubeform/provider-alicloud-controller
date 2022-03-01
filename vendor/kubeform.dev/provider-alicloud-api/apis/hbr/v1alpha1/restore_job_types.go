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

type RestoreJob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RestoreJobSpec   `json:"spec,omitempty"`
	Status            RestoreJobStatus `json:"status,omitempty"`
}

type RestoreJobSpec struct {
	State *RestoreJobSpecResource `json:"state,omitempty" tf:"-"`

	Resource RestoreJobSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type RestoreJobSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Exclude *string `json:"exclude,omitempty" tf:"exclude"`
	// +optional
	Include *string `json:"include,omitempty" tf:"include"`
	// +optional
	Options *string `json:"options,omitempty" tf:"options"`
	// +optional
	RestoreJobID *string `json:"restoreJobID,omitempty" tf:"restore_job_id"`
	RestoreType  *string `json:"restoreType" tf:"restore_type"`
	SnapshotHash *string `json:"snapshotHash" tf:"snapshot_hash"`
	SnapshotID   *string `json:"snapshotID" tf:"snapshot_id"`
	SourceType   *string `json:"sourceType" tf:"source_type"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	TargetBucket *string `json:"targetBucket,omitempty" tf:"target_bucket"`
	// +optional
	TargetClientID *string `json:"targetClientID,omitempty" tf:"target_client_id"`
	// +optional
	TargetContainer *string `json:"targetContainer,omitempty" tf:"target_container"`
	// +optional
	TargetContainerClusterID *string `json:"targetContainerClusterID,omitempty" tf:"target_container_cluster_id"`
	// +optional
	TargetCreateTime *string `json:"targetCreateTime,omitempty" tf:"target_create_time"`
	// +optional
	TargetDataSourceID *string `json:"targetDataSourceID,omitempty" tf:"target_data_source_id"`
	// +optional
	TargetFileSystemID *string `json:"targetFileSystemID,omitempty" tf:"target_file_system_id"`
	// +optional
	TargetInstanceID *string `json:"targetInstanceID,omitempty" tf:"target_instance_id"`
	// +optional
	TargetPath *string `json:"targetPath,omitempty" tf:"target_path"`
	// +optional
	TargetPrefix *string `json:"targetPrefix,omitempty" tf:"target_prefix"`
	VaultID      *string `json:"vaultID" tf:"vault_id"`
}

type RestoreJobStatus struct {
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

// RestoreJobList is a list of RestoreJobs
type RestoreJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RestoreJob CRD objects
	Items []RestoreJob `json:"items,omitempty"`
}