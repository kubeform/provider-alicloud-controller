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

type SchedulerRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SchedulerRuleSpec   `json:"spec,omitempty"`
	Status            SchedulerRuleStatus `json:"status,omitempty"`
}

type SchedulerRuleSpecRules struct {
	// +optional
	Priority *int64 `json:"priority,omitempty" tf:"priority"`
	// +optional
	RegionID *string `json:"regionID,omitempty" tf:"region_id"`
	// +optional
	Status *int64 `json:"status,omitempty" tf:"status"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
	// +optional
	ValueType *int64 `json:"valueType,omitempty" tf:"value_type"`
}

type SchedulerRuleSpec struct {
	State *SchedulerRuleSpecResource `json:"state,omitempty" tf:"-"`

	Resource SchedulerRuleSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type SchedulerRuleSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Cname *string `json:"cname,omitempty" tf:"cname"`
	// +optional
	Param *string `json:"param,omitempty" tf:"param"`
	// +optional
	ResourceGroupID *string                  `json:"resourceGroupID,omitempty" tf:"resource_group_id"`
	RuleName        *string                  `json:"ruleName" tf:"rule_name"`
	RuleType        *int64                   `json:"ruleType" tf:"rule_type"`
	Rules           []SchedulerRuleSpecRules `json:"rules" tf:"rules"`
}

type SchedulerRuleStatus struct {
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

// SchedulerRuleList is a list of SchedulerRules
type SchedulerRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SchedulerRule CRD objects
	Items []SchedulerRule `json:"items,omitempty"`
}