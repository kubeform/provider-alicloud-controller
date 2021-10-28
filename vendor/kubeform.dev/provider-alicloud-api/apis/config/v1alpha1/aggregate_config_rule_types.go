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

type AggregateConfigRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AggregateConfigRuleSpec   `json:"spec,omitempty"`
	Status            AggregateConfigRuleStatus `json:"status,omitempty"`
}

type AggregateConfigRuleSpec struct {
	State *AggregateConfigRuleSpecResource `json:"state,omitempty" tf:"-"`

	Resource AggregateConfigRuleSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type AggregateConfigRuleSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AggregateConfigRuleName *string `json:"aggregateConfigRuleName" tf:"aggregate_config_rule_name"`
	AggregatorID            *string `json:"aggregatorID" tf:"aggregator_id"`
	ConfigRuleTriggerTypes  *string `json:"configRuleTriggerTypes" tf:"config_rule_trigger_types"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	ExcludeResourceIDSScope *string `json:"excludeResourceIDSScope,omitempty" tf:"exclude_resource_ids_scope"`
	// +optional
	InputParameters map[string]string `json:"inputParameters,omitempty" tf:"input_parameters"`
	// +optional
	MaximumExecutionFrequency *string `json:"maximumExecutionFrequency,omitempty" tf:"maximum_execution_frequency"`
	// +optional
	RegionIDSScope *string `json:"regionIDSScope,omitempty" tf:"region_ids_scope"`
	// +optional
	ResourceGroupIDSScope *string  `json:"resourceGroupIDSScope,omitempty" tf:"resource_group_ids_scope"`
	ResourceTypesScope    []string `json:"resourceTypesScope" tf:"resource_types_scope"`
	RiskLevel             *int64   `json:"riskLevel" tf:"risk_level"`
	SourceIdentifier      *string  `json:"sourceIdentifier" tf:"source_identifier"`
	SourceOwner           *string  `json:"sourceOwner" tf:"source_owner"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	TagKeyScope *string `json:"tagKeyScope,omitempty" tf:"tag_key_scope"`
	// +optional
	TagValueScope *string `json:"tagValueScope,omitempty" tf:"tag_value_scope"`
}

type AggregateConfigRuleStatus struct {
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

// AggregateConfigRuleList is a list of AggregateConfigRules
type AggregateConfigRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AggregateConfigRule CRD objects
	Items []AggregateConfigRule `json:"items,omitempty"`
}
