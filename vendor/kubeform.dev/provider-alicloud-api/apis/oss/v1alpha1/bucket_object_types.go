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

type BucketObject struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketObjectSpec   `json:"spec,omitempty"`
	Status            BucketObjectStatus `json:"status,omitempty"`
}

type BucketObjectSpec struct {
	State *BucketObjectSpecResource `json:"state,omitempty" tf:"-"`

	Resource BucketObjectSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type BucketObjectSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Acl    *string `json:"acl,omitempty" tf:"acl"`
	Bucket *string `json:"bucket" tf:"bucket"`
	// +optional
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control"`
	// +optional
	Content *string `json:"content,omitempty" tf:"content"`
	// +optional
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition"`
	// +optional
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding"`
	// +optional
	ContentLength *string `json:"contentLength,omitempty" tf:"content_length"`
	// +optional
	ContentMd5 *string `json:"contentMd5,omitempty" tf:"content_md5"`
	// +optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type"`
	// +optional
	Etag *string `json:"etag,omitempty" tf:"etag"`
	// +optional
	Expires *string `json:"expires,omitempty" tf:"expires"`
	Key     *string `json:"key" tf:"key"`
	// +optional
	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`
	// +optional
	ServerSideEncryption *string `json:"serverSideEncryption,omitempty" tf:"server_side_encryption"`
	// +optional
	Source *string `json:"source,omitempty" tf:"source"`
	// +optional
	VersionID *string `json:"versionID,omitempty" tf:"version_id"`
}

type BucketObjectStatus struct {
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

// BucketObjectList is a list of BucketObjects
type BucketObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BucketObject CRD objects
	Items []BucketObject `json:"items,omitempty"`
}
