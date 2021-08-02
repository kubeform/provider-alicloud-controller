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

type Bucket struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketSpec   `json:"spec,omitempty"`
	Status            BucketStatus `json:"status,omitempty"`
}

type BucketSpecCorsRule struct {
	// +optional
	AllowedHeaders []string `json:"allowedHeaders,omitempty" tf:"allowed_headers"`
	AllowedMethods []string `json:"allowedMethods" tf:"allowed_methods"`
	AllowedOrigins []string `json:"allowedOrigins" tf:"allowed_origins"`
	// +optional
	ExposeHeaders []string `json:"exposeHeaders,omitempty" tf:"expose_headers"`
	// +optional
	MaxAgeSeconds *int64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds"`
}

type BucketSpecLifecycleRuleAbortMultipartUpload struct {
	// +optional
	CreatedBeforeDate *string `json:"createdBeforeDate,omitempty" tf:"created_before_date"`
	// +optional
	Days *int64 `json:"days,omitempty" tf:"days"`
}

type BucketSpecLifecycleRuleExpiration struct {
	// +optional
	CreatedBeforeDate *string `json:"createdBeforeDate,omitempty" tf:"created_before_date"`
	// +optional
	Date *string `json:"date,omitempty" tf:"date"`
	// +optional
	Days *int64 `json:"days,omitempty" tf:"days"`
	// +optional
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker"`
}

type BucketSpecLifecycleRuleNoncurrentVersionExpiration struct {
	Days *int64 `json:"days" tf:"days"`
}

type BucketSpecLifecycleRuleNoncurrentVersionTransition struct {
	Days         *int64  `json:"days" tf:"days"`
	StorageClass *string `json:"storageClass" tf:"storage_class"`
}

type BucketSpecLifecycleRuleTransitions struct {
	// +optional
	CreatedBeforeDate *string `json:"createdBeforeDate,omitempty" tf:"created_before_date"`
	// +optional
	Days *int64 `json:"days,omitempty" tf:"days"`
	// +optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class"`
}

type BucketSpecLifecycleRule struct {
	// +optional
	AbortMultipartUpload []BucketSpecLifecycleRuleAbortMultipartUpload `json:"abortMultipartUpload,omitempty" tf:"abort_multipart_upload"`
	Enabled              *bool                                         `json:"enabled" tf:"enabled"`
	// +optional
	Expiration []BucketSpecLifecycleRuleExpiration `json:"expiration,omitempty" tf:"expiration"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	NoncurrentVersionExpiration []BucketSpecLifecycleRuleNoncurrentVersionExpiration `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration"`
	// +optional
	NoncurrentVersionTransition []BucketSpecLifecycleRuleNoncurrentVersionTransition `json:"noncurrentVersionTransition,omitempty" tf:"noncurrent_version_transition"`
	// +optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`
	// +optional
	Transitions []BucketSpecLifecycleRuleTransitions `json:"transitions,omitempty" tf:"transitions"`
}

type BucketSpecLogging struct {
	TargetBucket *string `json:"targetBucket" tf:"target_bucket"`
	// +optional
	TargetPrefix *string `json:"targetPrefix,omitempty" tf:"target_prefix"`
}

type BucketSpecRefererConfig struct {
	// +optional
	AllowEmpty *bool    `json:"allowEmpty,omitempty" tf:"allow_empty"`
	Referers   []string `json:"referers" tf:"referers"`
}

type BucketSpecServerSideEncryptionRule struct {
	// +optional
	KmsMasterKeyID *string `json:"kmsMasterKeyID,omitempty" tf:"kms_master_key_id"`
	SseAlgorithm   *string `json:"sseAlgorithm" tf:"sse_algorithm"`
}

type BucketSpecTransferAcceleration struct {
	Enabled *bool `json:"enabled" tf:"enabled"`
}

type BucketSpecVersioning struct {
	Status *string `json:"status" tf:"status"`
}

type BucketSpecWebsite struct {
	// +optional
	ErrorDocument *string `json:"errorDocument,omitempty" tf:"error_document"`
	IndexDocument *string `json:"indexDocument" tf:"index_document"`
}

type BucketSpec struct {
	State *BucketSpecResource `json:"state,omitempty" tf:"-"`

	Resource BucketSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type BucketSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Acl *string `json:"acl,omitempty" tf:"acl"`
	// +optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	CorsRule []BucketSpecCorsRule `json:"corsRule,omitempty" tf:"cors_rule"`
	// +optional
	CreationDate *string `json:"creationDate,omitempty" tf:"creation_date"`
	// +optional
	ExtranetEndpoint *string `json:"extranetEndpoint,omitempty" tf:"extranet_endpoint"`
	// +optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy"`
	// +optional
	IntranetEndpoint *string `json:"intranetEndpoint,omitempty" tf:"intranet_endpoint"`
	// +optional
	// +kubebuilder:validation:MaxItems=1000
	LifecycleRule []BucketSpecLifecycleRule `json:"lifecycleRule,omitempty" tf:"lifecycle_rule"`
	// +optional
	Location *string `json:"location,omitempty" tf:"location"`
	// +optional
	Logging *BucketSpecLogging `json:"logging,omitempty" tf:"logging"`
	// +optional
	// Deprecated
	LoggingIsenable *bool `json:"loggingIsenable,omitempty" tf:"logging_isenable"`
	// +optional
	Owner *string `json:"owner,omitempty" tf:"owner"`
	// +optional
	Policy *string `json:"policy,omitempty" tf:"policy"`
	// +optional
	RedundancyType *string `json:"redundancyType,omitempty" tf:"redundancy_type"`
	// +optional
	RefererConfig *BucketSpecRefererConfig `json:"refererConfig,omitempty" tf:"referer_config"`
	// +optional
	ServerSideEncryptionRule *BucketSpecServerSideEncryptionRule `json:"serverSideEncryptionRule,omitempty" tf:"server_side_encryption_rule"`
	// +optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TransferAcceleration *BucketSpecTransferAcceleration `json:"transferAcceleration,omitempty" tf:"transfer_acceleration"`
	// +optional
	Versioning *BucketSpecVersioning `json:"versioning,omitempty" tf:"versioning"`
	// +optional
	Website *BucketSpecWebsite `json:"website,omitempty" tf:"website"`
}

type BucketStatus struct {
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

// BucketList is a list of Buckets
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Bucket CRD objects
	Items []Bucket `json:"items,omitempty"`
}