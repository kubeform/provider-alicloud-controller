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

type DedicatedHost struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DedicatedHostSpec   `json:"spec,omitempty"`
	Status            DedicatedHostStatus `json:"status,omitempty"`
}

type DedicatedHostSpecNetworkAttributes struct {
	// +optional
	SlbUDPTimeout *int64 `json:"slbUDPTimeout,omitempty" tf:"slb_udp_timeout"`
	// +optional
	UdpTimeout *int64 `json:"udpTimeout,omitempty" tf:"udp_timeout"`
}

type DedicatedHostSpec struct {
	State *DedicatedHostSpecResource `json:"state,omitempty" tf:"-"`

	Resource DedicatedHostSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type DedicatedHostSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ActionOnMaintenance *string `json:"actionOnMaintenance,omitempty" tf:"action_on_maintenance"`
	// +optional
	AutoPlacement *string `json:"autoPlacement,omitempty" tf:"auto_placement"`
	// +optional
	AutoReleaseTime *string `json:"autoReleaseTime,omitempty" tf:"auto_release_time"`
	// +optional
	AutoRenew *bool `json:"autoRenew,omitempty" tf:"auto_renew"`
	// +optional
	AutoRenewPeriod *int64 `json:"autoRenewPeriod,omitempty" tf:"auto_renew_period"`
	// +optional
	CpuOverCommitRatio *float64 `json:"cpuOverCommitRatio,omitempty" tf:"cpu_over_commit_ratio"`
	// +optional
	DedicatedHostClusterID *string `json:"dedicatedHostClusterID,omitempty" tf:"dedicated_host_cluster_id"`
	// +optional
	DedicatedHostName *string `json:"dedicatedHostName,omitempty" tf:"dedicated_host_name"`
	DedicatedHostType *string `json:"dedicatedHostType" tf:"dedicated_host_type"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DetailFee *bool `json:"detailFee,omitempty" tf:"detail_fee"`
	// +optional
	DryRun *bool `json:"dryRun,omitempty" tf:"dry_run"`
	// +optional
	ExpiredTime *string `json:"expiredTime,omitempty" tf:"expired_time"`
	// +optional
	MinQuantity *int64 `json:"minQuantity,omitempty" tf:"min_quantity"`
	// +optional
	NetworkAttributes *DedicatedHostSpecNetworkAttributes `json:"networkAttributes,omitempty" tf:"network_attributes"`
	// +optional
	PaymentType *string `json:"paymentType,omitempty" tf:"payment_type"`
	// +optional
	ResourceGroupID *string `json:"resourceGroupID,omitempty" tf:"resource_group_id"`
	// +optional
	SaleCycle *string `json:"saleCycle,omitempty" tf:"sale_cycle"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	ZoneID *string `json:"zoneID,omitempty" tf:"zone_id"`
}

type DedicatedHostStatus struct {
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

// DedicatedHostList is a list of DedicatedHosts
type DedicatedHostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DedicatedHost CRD objects
	Items []DedicatedHost `json:"items,omitempty"`
}