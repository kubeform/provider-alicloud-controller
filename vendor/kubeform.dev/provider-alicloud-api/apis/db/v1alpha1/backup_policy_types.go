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

type BackupPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupPolicySpec   `json:"spec,omitempty"`
	Status            BackupPolicyStatus `json:"status,omitempty"`
}

type BackupPolicySpec struct {
	State *BackupPolicySpecResource `json:"state,omitempty" tf:"-"`

	Resource BackupPolicySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type BackupPolicySpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ArchiveBackupKeepCount *int64 `json:"archiveBackupKeepCount,omitempty" tf:"archive_backup_keep_count"`
	// +optional
	ArchiveBackupKeepPolicy *string `json:"archiveBackupKeepPolicy,omitempty" tf:"archive_backup_keep_policy"`
	// +optional
	ArchiveBackupRetentionPeriod *int64 `json:"archiveBackupRetentionPeriod,omitempty" tf:"archive_backup_retention_period"`
	// +optional
	// Deprecated
	BackupPeriod []string `json:"backupPeriod,omitempty" tf:"backup_period"`
	// +optional
	BackupRetentionPeriod *int64 `json:"backupRetentionPeriod,omitempty" tf:"backup_retention_period"`
	// +optional
	// Deprecated
	BackupTime *string `json:"backupTime,omitempty" tf:"backup_time"`
	// +optional
	CompressType *string `json:"compressType,omitempty" tf:"compress_type"`
	// +optional
	EnableBackupLog *bool `json:"enableBackupLog,omitempty" tf:"enable_backup_log"`
	// +optional
	HighSpaceUsageProtection *string `json:"highSpaceUsageProtection,omitempty" tf:"high_space_usage_protection"`
	InstanceID               *string `json:"instanceID" tf:"instance_id"`
	// +optional
	LocalLogRetentionHours *int64 `json:"localLogRetentionHours,omitempty" tf:"local_log_retention_hours"`
	// +optional
	LocalLogRetentionSpace *int64 `json:"localLogRetentionSpace,omitempty" tf:"local_log_retention_space"`
	// +optional
	// Deprecated
	LogBackup *bool `json:"logBackup,omitempty" tf:"log_backup"`
	// +optional
	LogBackupFrequency *string `json:"logBackupFrequency,omitempty" tf:"log_backup_frequency"`
	// +optional
	LogBackupRetentionPeriod *int64 `json:"logBackupRetentionPeriod,omitempty" tf:"log_backup_retention_period"`
	// +optional
	// Deprecated
	LogRetentionPeriod *int64 `json:"logRetentionPeriod,omitempty" tf:"log_retention_period"`
	// +optional
	PreferredBackupPeriod []string `json:"preferredBackupPeriod,omitempty" tf:"preferred_backup_period"`
	// +optional
	PreferredBackupTime *string `json:"preferredBackupTime,omitempty" tf:"preferred_backup_time"`
	// +optional
	// Deprecated
	RetentionPeriod *int64 `json:"retentionPeriod,omitempty" tf:"retention_period"`
}

type BackupPolicyStatus struct {
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

// BackupPolicyList is a list of BackupPolicys
type BackupPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BackupPolicy CRD objects
	Items []BackupPolicy `json:"items,omitempty"`
}