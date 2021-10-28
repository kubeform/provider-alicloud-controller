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

type Ingress struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IngressSpec   `json:"spec,omitempty"`
	Status            IngressStatus `json:"status,omitempty"`
}

type IngressSpecDefaultRule struct {
	// +optional
	AppID *string `json:"appID,omitempty" tf:"app_id"`
	// +optional
	AppName *string `json:"appName,omitempty" tf:"app_name"`
	// +optional
	ContainerPort *int64 `json:"containerPort,omitempty" tf:"container_port"`
}

type IngressSpecRules struct {
	AppID         *string `json:"appID" tf:"app_id"`
	AppName       *string `json:"appName" tf:"app_name"`
	ContainerPort *int64  `json:"containerPort" tf:"container_port"`
	Domain        *string `json:"domain" tf:"domain"`
	Path          *string `json:"path" tf:"path"`
}

type IngressSpec struct {
	State *IngressSpecResource `json:"state,omitempty" tf:"-"`

	Resource IngressSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type IngressSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CertID *string `json:"certID,omitempty" tf:"cert_id"`
	// +optional
	DefaultRule *IngressSpecDefaultRule `json:"defaultRule,omitempty" tf:"default_rule"`
	// +optional
	Description  *string            `json:"description,omitempty" tf:"description"`
	ListenerPort *int64             `json:"listenerPort" tf:"listener_port"`
	NamespaceID  *string            `json:"namespaceID" tf:"namespace_id"`
	Rules        []IngressSpecRules `json:"rules" tf:"rules"`
	SlbID        *string            `json:"slbID" tf:"slb_id"`
}

type IngressStatus struct {
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

// IngressList is a list of Ingresss
type IngressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Ingress CRD objects
	Items []Ingress `json:"items,omitempty"`
}
