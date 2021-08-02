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

type Cluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec,omitempty"`
	Status            ClusterStatus `json:"status,omitempty"`
}

type ClusterSpecNodes struct {
	// +optional
	Eip *string `json:"eip,omitempty" tf:"eip"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	PrivateIP *string `json:"privateIP,omitempty" tf:"private_ip"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
}

type ClusterSpec struct {
	State *ClusterSpecResource `json:"state,omitempty" tf:"-"`

	Resource ClusterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type ClusterSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AgentVersion *string `json:"agentVersion,omitempty" tf:"agent_version"`
	CidrBlock    *string `json:"cidrBlock" tf:"cidr_block"`
	// +optional
	DiskCategory *string `json:"diskCategory,omitempty" tf:"disk_category"`
	// +optional
	DiskSize *int64 `json:"diskSize,omitempty" tf:"disk_size"`
	// +optional
	ImageID      *string `json:"imageID,omitempty" tf:"image_id"`
	InstanceType *string `json:"instanceType" tf:"instance_type"`
	// +optional
	IsOutdated *bool `json:"isOutdated,omitempty" tf:"is_outdated"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`
	// +optional
	NeedSlb *bool `json:"needSlb,omitempty" tf:"need_slb"`
	// +optional
	NodeNumber *int64 `json:"nodeNumber,omitempty" tf:"node_number"`
	// +optional
	Nodes    []ClusterSpecNodes `json:"nodes,omitempty" tf:"nodes"`
	Password *string            `json:"-" sensitive:"true" tf:"password"`
	// +optional
	ReleaseEip *bool `json:"releaseEip,omitempty" tf:"release_eip"`
	// +optional
	SecurityGroupID *string `json:"securityGroupID,omitempty" tf:"security_group_id"`
	// +optional
	// Deprecated
	Size *int64 `json:"size,omitempty" tf:"size"`
	// +optional
	SlbID *string `json:"slbID,omitempty" tf:"slb_id"`
	// +optional
	VpcID     *string `json:"vpcID,omitempty" tf:"vpc_id"`
	VswitchID *string `json:"vswitchID" tf:"vswitch_id"`
}

type ClusterStatus struct {
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

// ClusterList is a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Cluster CRD objects
	Items []Cluster `json:"items,omitempty"`
}
