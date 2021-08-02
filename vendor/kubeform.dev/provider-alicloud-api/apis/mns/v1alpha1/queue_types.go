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

type Queue struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QueueSpec   `json:"spec,omitempty"`
	Status            QueueStatus `json:"status,omitempty"`
}

type QueueSpec struct {
	State *QueueSpecResource `json:"state,omitempty" tf:"-"`

	Resource QueueSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type QueueSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	DelaySeconds *int64 `json:"delaySeconds,omitempty" tf:"delay_seconds"`
	// +optional
	MaximumMessageSize *int64 `json:"maximumMessageSize,omitempty" tf:"maximum_message_size"`
	// +optional
	MessageRetentionPeriod *int64  `json:"messageRetentionPeriod,omitempty" tf:"message_retention_period"`
	Name                   *string `json:"name" tf:"name"`
	// +optional
	PollingWaitSeconds *int64 `json:"pollingWaitSeconds,omitempty" tf:"polling_wait_seconds"`
	// +optional
	VisibilityTimeout *int64 `json:"visibilityTimeout,omitempty" tf:"visibility_timeout"`
}

type QueueStatus struct {
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

// QueueList is a list of Queues
type QueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Queue CRD objects
	Items []Queue `json:"items,omitempty"`
}