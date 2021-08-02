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

type QosPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QosPolicySpec   `json:"spec,omitempty"`
	Status            QosPolicyStatus `json:"status,omitempty"`
}

type QosPolicySpec struct {
	State *QosPolicySpecResource `json:"state,omitempty" tf:"-"`

	Resource QosPolicySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type QosPolicySpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Description   *string `json:"description,omitempty" tf:"description"`
	DestCIDR      *string `json:"destCIDR" tf:"dest_cidr"`
	DestPortRange *string `json:"destPortRange" tf:"dest_port_range"`
	// +optional
	EndTime    *string `json:"endTime,omitempty" tf:"end_time"`
	IpProtocol *string `json:"ipProtocol" tf:"ip_protocol"`
	// +optional
	Name            *string `json:"name,omitempty" tf:"name"`
	Priority        *int64  `json:"priority" tf:"priority"`
	QosID           *string `json:"qosID" tf:"qos_id"`
	SourceCIDR      *string `json:"sourceCIDR" tf:"source_cidr"`
	SourcePortRange *string `json:"sourcePortRange" tf:"source_port_range"`
	// +optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time"`
}

type QosPolicyStatus struct {
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

// QosPolicyList is a list of QosPolicys
type QosPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of QosPolicy CRD objects
	Items []QosPolicy `json:"items,omitempty"`
}