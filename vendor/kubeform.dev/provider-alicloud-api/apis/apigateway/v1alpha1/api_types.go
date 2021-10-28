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

type Api struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiSpec   `json:"spec,omitempty"`
	Status            ApiStatus `json:"status,omitempty"`
}

type ApiSpecConstantParameters struct {
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	In          *string `json:"in" tf:"in"`
	Name        *string `json:"name" tf:"name"`
	Value       *string `json:"value" tf:"value"`
}

type ApiSpecFcServiceConfig struct {
	// +optional
	ArnRole      *string `json:"arnRole,omitempty" tf:"arn_role"`
	FunctionName *string `json:"functionName" tf:"function_name"`
	Region       *string `json:"region" tf:"region"`
	ServiceName  *string `json:"serviceName" tf:"service_name"`
	Timeout      *int64  `json:"timeout" tf:"timeout"`
}

type ApiSpecHttpServiceConfig struct {
	Address *string `json:"address" tf:"address"`
	// +optional
	AoneName *string `json:"aoneName,omitempty" tf:"aone_name"`
	Method   *string `json:"method" tf:"method"`
	Path     *string `json:"path" tf:"path"`
	Timeout  *int64  `json:"timeout" tf:"timeout"`
}

type ApiSpecHttpVpcServiceConfig struct {
	// +optional
	AoneName *string `json:"aoneName,omitempty" tf:"aone_name"`
	Method   *string `json:"method" tf:"method"`
	Name     *string `json:"name" tf:"name"`
	Path     *string `json:"path" tf:"path"`
	Timeout  *int64  `json:"timeout" tf:"timeout"`
}

type ApiSpecMockServiceConfig struct {
	// +optional
	AoneName *string `json:"aoneName,omitempty" tf:"aone_name"`
	Result   *string `json:"result" tf:"result"`
}

type ApiSpecRequestConfig struct {
	// +optional
	BodyFormat *string `json:"bodyFormat,omitempty" tf:"body_format"`
	Method     *string `json:"method" tf:"method"`
	Mode       *string `json:"mode" tf:"mode"`
	Path       *string `json:"path" tf:"path"`
	Protocol   *string `json:"protocol" tf:"protocol"`
}

type ApiSpecRequestParameters struct {
	// +optional
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	In          *string `json:"in" tf:"in"`
	InService   *string `json:"inService" tf:"in_service"`
	Name        *string `json:"name" tf:"name"`
	NameService *string `json:"nameService" tf:"name_service"`
	Required    *string `json:"required" tf:"required"`
	Type        *string `json:"type" tf:"type"`
}

type ApiSpecSystemParameters struct {
	In          *string `json:"in" tf:"in"`
	Name        *string `json:"name" tf:"name"`
	NameService *string `json:"nameService" tf:"name_service"`
}

type ApiSpec struct {
	State *ApiSpecResource `json:"state,omitempty" tf:"-"`

	Resource ApiSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ApiSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ApiID    *string `json:"apiID,omitempty" tf:"api_id"`
	AuthType *string `json:"authType" tf:"auth_type"`
	// +optional
	ConstantParameters []ApiSpecConstantParameters `json:"constantParameters,omitempty" tf:"constant_parameters"`
	Description        *string                     `json:"description" tf:"description"`
	// +optional
	FcServiceConfig *ApiSpecFcServiceConfig `json:"fcServiceConfig,omitempty" tf:"fc_service_config"`
	GroupID         *string                 `json:"groupID" tf:"group_id"`
	// +optional
	HttpServiceConfig *ApiSpecHttpServiceConfig `json:"httpServiceConfig,omitempty" tf:"http_service_config"`
	// +optional
	HttpVpcServiceConfig *ApiSpecHttpVpcServiceConfig `json:"httpVpcServiceConfig,omitempty" tf:"http_vpc_service_config"`
	// +optional
	MockServiceConfig *ApiSpecMockServiceConfig `json:"mockServiceConfig,omitempty" tf:"mock_service_config"`
	Name              *string                   `json:"name" tf:"name"`
	RequestConfig     *ApiSpecRequestConfig     `json:"requestConfig" tf:"request_config"`
	// +optional
	RequestParameters []ApiSpecRequestParameters `json:"requestParameters,omitempty" tf:"request_parameters"`
	ServiceType       *string                    `json:"serviceType" tf:"service_type"`
	// +optional
	StageNames []string `json:"stageNames,omitempty" tf:"stage_names"`
	// +optional
	SystemParameters []ApiSpecSystemParameters `json:"systemParameters,omitempty" tf:"system_parameters"`
}

type ApiStatus struct {
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

// ApiList is a list of Apis
type ApiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Api CRD objects
	Items []Api `json:"items,omitempty"`
}
