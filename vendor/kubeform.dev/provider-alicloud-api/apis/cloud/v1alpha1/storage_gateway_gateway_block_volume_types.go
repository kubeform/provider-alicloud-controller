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

type StorageGatewayGatewayBlockVolume struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageGatewayGatewayBlockVolumeSpec   `json:"spec,omitempty"`
	Status            StorageGatewayGatewayBlockVolumeStatus `json:"status,omitempty"`
}

type StorageGatewayGatewayBlockVolumeSpec struct {
	State *StorageGatewayGatewayBlockVolumeSpecResource `json:"state,omitempty" tf:"-"`

	Resource StorageGatewayGatewayBlockVolumeSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type StorageGatewayGatewayBlockVolumeSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CacheMode *string `json:"cacheMode,omitempty" tf:"cache_mode"`
	// +optional
	ChapEnabled *bool `json:"chapEnabled,omitempty" tf:"chap_enabled"`
	// +optional
	ChapInPassword *string `json:"chapInPassword,omitempty" tf:"chap_in_password"`
	// +optional
	ChapInUser *string `json:"chapInUser,omitempty" tf:"chap_in_user"`
	// +optional
	ChunkSize              *int64  `json:"chunkSize,omitempty" tf:"chunk_size"`
	GatewayBlockVolumeName *string `json:"gatewayBlockVolumeName" tf:"gateway_block_volume_name"`
	GatewayID              *string `json:"gatewayID" tf:"gateway_id"`
	// +optional
	IndexID *string `json:"indexID,omitempty" tf:"index_id"`
	// +optional
	IsSourceDeletion *bool `json:"isSourceDeletion,omitempty" tf:"is_source_deletion"`
	// +optional
	LocalPath     *string `json:"localPath,omitempty" tf:"local_path"`
	OssBucketName *string `json:"ossBucketName" tf:"oss_bucket_name"`
	// +optional
	OssBucketSsl *bool   `json:"ossBucketSsl,omitempty" tf:"oss_bucket_ssl"`
	OssEndpoint  *string `json:"ossEndpoint" tf:"oss_endpoint"`
	Protocol     *string `json:"protocol" tf:"protocol"`
	// +optional
	Recovery *bool `json:"recovery,omitempty" tf:"recovery"`
	// +optional
	Size *int64 `json:"size,omitempty" tf:"size"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
}

type StorageGatewayGatewayBlockVolumeStatus struct {
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

// StorageGatewayGatewayBlockVolumeList is a list of StorageGatewayGatewayBlockVolumes
type StorageGatewayGatewayBlockVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageGatewayGatewayBlockVolume CRD objects
	Items []StorageGatewayGatewayBlockVolume `json:"items,omitempty"`
}
