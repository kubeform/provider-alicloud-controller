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

type Application struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationSpec   `json:"spec,omitempty"`
	Status            ApplicationStatus `json:"status,omitempty"`
}

type ApplicationSpec struct {
	State *ApplicationSpecResource `json:"state,omitempty" tf:"-"`

	Resource ApplicationSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ApplicationSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ApplicationName *string `json:"applicationName" tf:"application_name"`
	// +optional
	BuildPackID *int64  `json:"buildPackID,omitempty" tf:"build_pack_id"`
	ClusterID   *string `json:"clusterID" tf:"cluster_id"`
	// +optional
	Descriotion *string `json:"descriotion,omitempty" tf:"descriotion"`
	// +optional
	EcuInfo []string `json:"ecuInfo,omitempty" tf:"ecu_info"`
	// +optional
	GroupID *string `json:"groupID,omitempty" tf:"group_id"`
	// +optional
	HealthCheckURL *string `json:"healthCheckURL,omitempty" tf:"health_check_url"`
	// +optional
	LogicalRegionID *string `json:"logicalRegionID,omitempty" tf:"logical_region_id"`
	PackageType     *string `json:"packageType" tf:"package_type"`
	// +optional
	PackageVersion *string `json:"packageVersion,omitempty" tf:"package_version"`
	// +optional
	WarURL *string `json:"warURL,omitempty" tf:"war_url"`
}

type ApplicationStatus struct {
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

// ApplicationList is a list of Applications
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Application CRD objects
	Items []Application `json:"items,omitempty"`
}
