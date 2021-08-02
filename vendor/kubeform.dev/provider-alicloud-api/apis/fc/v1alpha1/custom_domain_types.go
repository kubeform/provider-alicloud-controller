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

type CustomDomain struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CustomDomainSpec   `json:"spec,omitempty"`
	Status            CustomDomainStatus `json:"status,omitempty"`
}

type CustomDomainSpecCertConfig struct {
	CertName    *string `json:"certName" tf:"cert_name"`
	Certificate *string `json:"certificate" tf:"certificate"`
	PrivateKey  *string `json:"-" sensitive:"true" tf:"private_key"`
}

type CustomDomainSpecRouteConfig struct {
	FunctionName *string `json:"functionName" tf:"function_name"`
	// +optional
	Methods []string `json:"methods,omitempty" tf:"methods"`
	Path    *string  `json:"path" tf:"path"`
	// +optional
	Qualifier   *string `json:"qualifier,omitempty" tf:"qualifier"`
	ServiceName *string `json:"serviceName" tf:"service_name"`
}

type CustomDomainSpec struct {
	State *CustomDomainSpecResource `json:"state,omitempty" tf:"-"`

	Resource CustomDomainSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type CustomDomainSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AccountID *string `json:"accountID,omitempty" tf:"account_id"`
	// +optional
	ApiVersion *string `json:"apiVersion,omitempty" tf:"api_version"`
	// +optional
	CertConfig *CustomDomainSpecCertConfig `json:"certConfig,omitempty" tf:"cert_config"`
	// +optional
	CreatedTime *string `json:"createdTime,omitempty" tf:"created_time"`
	DomainName  *string `json:"domainName" tf:"domain_name"`
	// +optional
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" tf:"last_modified_time"`
	Protocol         *string `json:"protocol" tf:"protocol"`
	// +optional
	RouteConfig []CustomDomainSpecRouteConfig `json:"routeConfig,omitempty" tf:"route_config"`
}

type CustomDomainStatus struct {
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

// CustomDomainList is a list of CustomDomains
type CustomDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CustomDomain CRD objects
	Items []CustomDomain `json:"items,omitempty"`
}
