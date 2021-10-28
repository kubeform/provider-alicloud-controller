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

type OpenapiImageCache struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpenapiImageCacheSpec   `json:"spec,omitempty"`
	Status            OpenapiImageCacheStatus `json:"status,omitempty"`
}

type OpenapiImageCacheSpecImageRegistryCredential struct {
	// +optional
	Password *string `json:"password,omitempty" tf:"password"`
	// +optional
	Server *string `json:"server,omitempty" tf:"server"`
	// +optional
	UserName *string `json:"userName,omitempty" tf:"user_name"`
}

type OpenapiImageCacheSpec struct {
	State *OpenapiImageCacheSpecResource `json:"state,omitempty" tf:"-"`

	Resource OpenapiImageCacheSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type OpenapiImageCacheSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ContainerGroupID *string `json:"containerGroupID,omitempty" tf:"container_group_id"`
	// +optional
	EipInstanceID  *string `json:"eipInstanceID,omitempty" tf:"eip_instance_id"`
	ImageCacheName *string `json:"imageCacheName" tf:"image_cache_name"`
	// +optional
	ImageCacheSize *int64 `json:"imageCacheSize,omitempty" tf:"image_cache_size"`
	// +optional
	ImageRegistryCredential []OpenapiImageCacheSpecImageRegistryCredential `json:"imageRegistryCredential,omitempty" tf:"image_registry_credential"`
	Images                  []string                                       `json:"images" tf:"images"`
	// +optional
	ResourceGroupID *string `json:"resourceGroupID,omitempty" tf:"resource_group_id"`
	// +optional
	RetentionDays   *int64  `json:"retentionDays,omitempty" tf:"retention_days"`
	SecurityGroupID *string `json:"securityGroupID" tf:"security_group_id"`
	// +optional
	Status    *string `json:"status,omitempty" tf:"status"`
	VswitchID *string `json:"vswitchID" tf:"vswitch_id"`
	// +optional
	ZoneID *string `json:"zoneID,omitempty" tf:"zone_id"`
}

type OpenapiImageCacheStatus struct {
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

// OpenapiImageCacheList is a list of OpenapiImageCaches
type OpenapiImageCacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OpenapiImageCache CRD objects
	Items []OpenapiImageCache `json:"items,omitempty"`
}
