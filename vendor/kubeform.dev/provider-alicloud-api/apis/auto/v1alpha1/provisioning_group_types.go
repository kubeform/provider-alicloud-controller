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

type ProvisioningGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProvisioningGroupSpec   `json:"spec,omitempty"`
	Status            ProvisioningGroupStatus `json:"status,omitempty"`
}

type ProvisioningGroupSpecLaunchTemplateConfig struct {
	// +optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type"`
	MaxPrice     *string `json:"maxPrice" tf:"max_price"`
	// +optional
	Priority         *string `json:"priority,omitempty" tf:"priority"`
	VswitchID        *string `json:"vswitchID" tf:"vswitch_id"`
	WeightedCapacity *string `json:"weightedCapacity" tf:"weighted_capacity"`
}

type ProvisioningGroupSpec struct {
	State *ProvisioningGroupSpecResource `json:"state,omitempty" tf:"-"`

	Resource ProvisioningGroupSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ProvisioningGroupSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AutoProvisioningGroupName *string `json:"autoProvisioningGroupName,omitempty" tf:"auto_provisioning_group_name"`
	// +optional
	AutoProvisioningGroupType *string `json:"autoProvisioningGroupType,omitempty" tf:"auto_provisioning_group_type"`
	// +optional
	DefaultTargetCapacityType *string `json:"defaultTargetCapacityType,omitempty" tf:"default_target_capacity_type"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	ExcessCapacityTerminationPolicy *string                                     `json:"excessCapacityTerminationPolicy,omitempty" tf:"excess_capacity_termination_policy"`
	LaunchTemplateConfig            []ProvisioningGroupSpecLaunchTemplateConfig `json:"launchTemplateConfig" tf:"launch_template_config"`
	LaunchTemplateID                *string                                     `json:"launchTemplateID" tf:"launch_template_id"`
	// +optional
	LaunchTemplateVersion *string `json:"launchTemplateVersion,omitempty" tf:"launch_template_version"`
	// +optional
	MaxSpotPrice *float64 `json:"maxSpotPrice,omitempty" tf:"max_spot_price"`
	// +optional
	PayAsYouGoAllocationStrategy *string `json:"payAsYouGoAllocationStrategy,omitempty" tf:"pay_as_you_go_allocation_strategy"`
	// +optional
	PayAsYouGoTargetCapacity *string `json:"payAsYouGoTargetCapacity,omitempty" tf:"pay_as_you_go_target_capacity"`
	// +optional
	SpotAllocationStrategy *string `json:"spotAllocationStrategy,omitempty" tf:"spot_allocation_strategy"`
	// +optional
	SpotInstanceInterruptionBehavior *string `json:"spotInstanceInterruptionBehavior,omitempty" tf:"spot_instance_interruption_behavior"`
	// +optional
	SpotInstancePoolsToUseCount *int64 `json:"spotInstancePoolsToUseCount,omitempty" tf:"spot_instance_pools_to_use_count"`
	// +optional
	SpotTargetCapacity *string `json:"spotTargetCapacity,omitempty" tf:"spot_target_capacity"`
	// +optional
	TerminateInstances *bool `json:"terminateInstances,omitempty" tf:"terminate_instances"`
	// +optional
	TerminateInstancesWithExpiration *bool   `json:"terminateInstancesWithExpiration,omitempty" tf:"terminate_instances_with_expiration"`
	TotalTargetCapacity              *string `json:"totalTargetCapacity" tf:"total_target_capacity"`
	// +optional
	ValidFrom *string `json:"validFrom,omitempty" tf:"valid_from"`
	// +optional
	ValidUntil *string `json:"validUntil,omitempty" tf:"valid_until"`
}

type ProvisioningGroupStatus struct {
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

// ProvisioningGroupList is a list of ProvisioningGroups
type ProvisioningGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ProvisioningGroup CRD objects
	Items []ProvisioningGroup `json:"items,omitempty"`
}
