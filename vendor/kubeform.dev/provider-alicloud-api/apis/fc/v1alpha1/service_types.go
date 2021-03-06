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

type Service struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceSpec   `json:"spec,omitempty"`
	Status            ServiceStatus `json:"status,omitempty"`
}

type ServiceSpecLogConfig struct {
	Logstore *string `json:"logstore" tf:"logstore"`
	Project  *string `json:"project" tf:"project"`
}

type ServiceSpecNasConfigMountPoints struct {
	MountDir   *string `json:"mountDir" tf:"mount_dir"`
	ServerAddr *string `json:"serverAddr" tf:"server_addr"`
}

type ServiceSpecNasConfig struct {
	GroupID     *int64                            `json:"groupID" tf:"group_id"`
	MountPoints []ServiceSpecNasConfigMountPoints `json:"mountPoints" tf:"mount_points"`
	UserID      *int64                            `json:"userID" tf:"user_id"`
}

type ServiceSpecVpcConfig struct {
	SecurityGroupID *string `json:"securityGroupID" tf:"security_group_id"`
	// +optional
	VpcID      *string  `json:"vpcID,omitempty" tf:"vpc_id"`
	VswitchIDS []string `json:"vswitchIDS" tf:"vswitch_ids"`
}

type ServiceSpec struct {
	State *ServiceSpecResource `json:"state,omitempty" tf:"-"`

	Resource ServiceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ServiceSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	InternetAccess *bool `json:"internetAccess,omitempty" tf:"internet_access"`
	// +optional
	LastModified *string `json:"lastModified,omitempty" tf:"last_modified"`
	// +optional
	LogConfig *ServiceSpecLogConfig `json:"logConfig,omitempty" tf:"log_config"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`
	// +optional
	NasConfig *ServiceSpecNasConfig `json:"nasConfig,omitempty" tf:"nas_config"`
	// +optional
	Publish *bool `json:"publish,omitempty" tf:"publish"`
	// +optional
	Role *string `json:"role,omitempty" tf:"role"`
	// +optional
	ServiceID *string `json:"serviceID,omitempty" tf:"service_id"`
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
	// +optional
	VpcConfig *ServiceSpecVpcConfig `json:"vpcConfig,omitempty" tf:"vpc_config"`
}

type ServiceStatus struct {
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

// ServiceList is a list of Services
type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Service CRD objects
	Items []Service `json:"items,omitempty"`
}
