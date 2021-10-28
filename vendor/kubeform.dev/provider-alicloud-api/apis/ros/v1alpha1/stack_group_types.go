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

type StackGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StackGroupSpec   `json:"spec,omitempty"`
	Status            StackGroupStatus `json:"status,omitempty"`
}

type StackGroupSpecParameters struct {
	// +optional
	ParameterKey *string `json:"parameterKey,omitempty" tf:"parameter_key"`
	// +optional
	ParameterValue *string `json:"parameterValue,omitempty" tf:"parameter_value"`
}

type StackGroupSpec struct {
	State *StackGroupSpecResource `json:"state,omitempty" tf:"-"`

	Resource StackGroupSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type StackGroupSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AccountIDS *string `json:"accountIDS,omitempty" tf:"account_ids"`
	// +optional
	AdministrationRoleName *string `json:"administrationRoleName,omitempty" tf:"administration_role_name"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	ExecutionRoleName *string `json:"executionRoleName,omitempty" tf:"execution_role_name"`
	// +optional
	OperationDescription *string `json:"operationDescription,omitempty" tf:"operation_description"`
	// +optional
	OperationPreferences *string `json:"operationPreferences,omitempty" tf:"operation_preferences"`
	// +optional
	Parameters []StackGroupSpecParameters `json:"parameters,omitempty" tf:"parameters"`
	// +optional
	RegionIDS *string `json:"regionIDS,omitempty" tf:"region_ids"`
	// +optional
	StackGroupID   *string `json:"stackGroupID,omitempty" tf:"stack_group_id"`
	StackGroupName *string `json:"stackGroupName" tf:"stack_group_name"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	TemplateBody *string `json:"templateBody,omitempty" tf:"template_body"`
	// +optional
	TemplateURL *string `json:"templateURL,omitempty" tf:"template_url"`
	// +optional
	TemplateVersion *string `json:"templateVersion,omitempty" tf:"template_version"`
}

type StackGroupStatus struct {
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

// StackGroupList is a list of StackGroups
type StackGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StackGroup CRD objects
	Items []StackGroup `json:"items,omitempty"`
}
