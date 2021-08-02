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

type KubernetesAutoscaler struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubernetesAutoscalerSpec   `json:"spec,omitempty"`
	Status            KubernetesAutoscalerStatus `json:"status,omitempty"`
}

type KubernetesAutoscalerSpecNodepools struct {
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Labels *string `json:"labels,omitempty" tf:"labels"`
	// +optional
	Taints *string `json:"taints,omitempty" tf:"taints"`
}

type KubernetesAutoscalerSpec struct {
	State *KubernetesAutoscalerSpecResource `json:"state,omitempty" tf:"-"`

	Resource KubernetesAutoscalerSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type KubernetesAutoscalerSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ClusterID            *string `json:"clusterID" tf:"cluster_id"`
	CoolDownDuration     *string `json:"coolDownDuration" tf:"cool_down_duration"`
	DeferScaleInDuration *string `json:"deferScaleInDuration" tf:"defer_scale_in_duration"`
	// +optional
	// +kubebuilder:validation:MaxItems=30
	// +kubebuilder:validation:MinItems=1
	Nodepools []KubernetesAutoscalerSpecNodepools `json:"nodepools,omitempty" tf:"nodepools"`
	// +optional
	UseEcsRAMRoleToken *bool   `json:"useEcsRAMRoleToken,omitempty" tf:"use_ecs_ram_role_token"`
	Utilization        *string `json:"utilization" tf:"utilization"`
}

type KubernetesAutoscalerStatus struct {
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

// KubernetesAutoscalerList is a list of KubernetesAutoscalers
type KubernetesAutoscalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubernetesAutoscaler CRD objects
	Items []KubernetesAutoscaler `json:"items,omitempty"`
}
