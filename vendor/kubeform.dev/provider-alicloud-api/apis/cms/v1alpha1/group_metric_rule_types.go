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

type GroupMetricRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupMetricRuleSpec   `json:"spec,omitempty"`
	Status            GroupMetricRuleStatus `json:"status,omitempty"`
}

type GroupMetricRuleSpecEscalationsCritical struct {
	// +optional
	ComparisonOperator *string `json:"comparisonOperator,omitempty" tf:"comparison_operator"`
	// +optional
	Statistics *string `json:"statistics,omitempty" tf:"statistics"`
	// +optional
	Threshold *string `json:"threshold,omitempty" tf:"threshold"`
	// +optional
	Times *int64 `json:"times,omitempty" tf:"times"`
}

type GroupMetricRuleSpecEscalationsInfo struct {
	// +optional
	ComparisonOperator *string `json:"comparisonOperator,omitempty" tf:"comparison_operator"`
	// +optional
	Statistics *string `json:"statistics,omitempty" tf:"statistics"`
	// +optional
	Threshold *string `json:"threshold,omitempty" tf:"threshold"`
	// +optional
	Times *int64 `json:"times,omitempty" tf:"times"`
}

type GroupMetricRuleSpecEscalationsWarn struct {
	// +optional
	ComparisonOperator *string `json:"comparisonOperator,omitempty" tf:"comparison_operator"`
	// +optional
	Statistics *string `json:"statistics,omitempty" tf:"statistics"`
	// +optional
	Threshold *string `json:"threshold,omitempty" tf:"threshold"`
	// +optional
	Times *int64 `json:"times,omitempty" tf:"times"`
}

type GroupMetricRuleSpecEscalations struct {
	// +optional
	Critical *GroupMetricRuleSpecEscalationsCritical `json:"critical,omitempty" tf:"critical"`
	// +optional
	Info *GroupMetricRuleSpecEscalationsInfo `json:"info,omitempty" tf:"info"`
	// +optional
	Warn *GroupMetricRuleSpecEscalationsWarn `json:"warn,omitempty" tf:"warn"`
}

type GroupMetricRuleSpec struct {
	State *GroupMetricRuleSpecResource `json:"state,omitempty" tf:"-"`

	Resource GroupMetricRuleSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type GroupMetricRuleSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Category *string `json:"category" tf:"category"`
	// +optional
	ContactGroups *string `json:"contactGroups,omitempty" tf:"contact_groups"`
	// +optional
	Dimensions *string `json:"dimensions,omitempty" tf:"dimensions"`
	// +optional
	EffectiveInterval *string `json:"effectiveInterval,omitempty" tf:"effective_interval"`
	// +optional
	EmailSubject        *string                         `json:"emailSubject,omitempty" tf:"email_subject"`
	Escalations         *GroupMetricRuleSpecEscalations `json:"escalations" tf:"escalations"`
	GroupID             *string                         `json:"groupID" tf:"group_id"`
	GroupMetricRuleName *string                         `json:"groupMetricRuleName" tf:"group_metric_rule_name"`
	// +optional
	Interval   *string `json:"interval,omitempty" tf:"interval"`
	MetricName *string `json:"metricName" tf:"metric_name"`
	Namespace  *string `json:"namespace" tf:"namespace"`
	// +optional
	NoEffectiveInterval *string `json:"noEffectiveInterval,omitempty" tf:"no_effective_interval"`
	// +optional
	Period *int64  `json:"period,omitempty" tf:"period"`
	RuleID *string `json:"ruleID" tf:"rule_id"`
	// +optional
	SilenceTime *int64 `json:"silenceTime,omitempty" tf:"silence_time"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	Webhook *string `json:"webhook,omitempty" tf:"webhook"`
}

type GroupMetricRuleStatus struct {
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

// GroupMetricRuleList is a list of GroupMetricRules
type GroupMetricRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GroupMetricRule CRD objects
	Items []GroupMetricRule `json:"items,omitempty"`
}
