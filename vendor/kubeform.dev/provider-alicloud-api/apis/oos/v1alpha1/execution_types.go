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

type Execution struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExecutionSpec   `json:"spec,omitempty"`
	Status            ExecutionStatus `json:"status,omitempty"`
}

type ExecutionSpec struct {
	State *ExecutionSpecResource `json:"state,omitempty" tf:"-"`

	Resource ExecutionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ExecutionSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Counters *string `json:"counters,omitempty" tf:"counters"`
	// +optional
	CreateDate *string `json:"createDate,omitempty" tf:"create_date"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	EndDate *string `json:"endDate,omitempty" tf:"end_date"`
	// +optional
	ExecutedBy *string `json:"executedBy,omitempty" tf:"executed_by"`
	// +optional
	IsParent *bool `json:"isParent,omitempty" tf:"is_parent"`
	// +optional
	LoopMode *string `json:"loopMode,omitempty" tf:"loop_mode"`
	// +optional
	Mode *string `json:"mode,omitempty" tf:"mode"`
	// +optional
	Outputs *string `json:"outputs,omitempty" tf:"outputs"`
	// +optional
	Parameters *string `json:"parameters,omitempty" tf:"parameters"`
	// +optional
	ParentExecutionID *string `json:"parentExecutionID,omitempty" tf:"parent_execution_id"`
	// +optional
	RamRole *string `json:"ramRole,omitempty" tf:"ram_role"`
	// +optional
	SafetyCheck *string `json:"safetyCheck,omitempty" tf:"safety_check"`
	// +optional
	StartDate *string `json:"startDate,omitempty" tf:"start_date"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	StatusMessage *string `json:"statusMessage,omitempty" tf:"status_message"`
	// +optional
	TemplateContent *string `json:"templateContent,omitempty" tf:"template_content"`
	// +optional
	TemplateID   *string `json:"templateID,omitempty" tf:"template_id"`
	TemplateName *string `json:"templateName" tf:"template_name"`
	// +optional
	TemplateVersion *string `json:"templateVersion,omitempty" tf:"template_version"`
	// +optional
	UpdateDate *string `json:"updateDate,omitempty" tf:"update_date"`
}

type ExecutionStatus struct {
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

// ExecutionList is a list of Executions
type ExecutionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Execution CRD objects
	Items []Execution `json:"items,omitempty"`
}