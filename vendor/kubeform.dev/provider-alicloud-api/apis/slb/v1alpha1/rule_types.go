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

type Rule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RuleSpec   `json:"spec,omitempty"`
	Status            RuleStatus `json:"status,omitempty"`
}

type RuleSpec struct {
	State *RuleSpecResource `json:"state,omitempty" tf:"-"`

	Resource RuleSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type RuleSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Cookie *string `json:"cookie,omitempty" tf:"cookie"`
	// +optional
	CookieTimeout *int64 `json:"cookieTimeout,omitempty" tf:"cookie_timeout"`
	// +optional
	DeleteProtectionValidation *bool `json:"deleteProtectionValidation,omitempty" tf:"delete_protection_validation"`
	// +optional
	Domain       *string `json:"domain,omitempty" tf:"domain"`
	FrontendPort *int64  `json:"frontendPort" tf:"frontend_port"`
	// +optional
	HealthCheck *string `json:"healthCheck,omitempty" tf:"health_check"`
	// +optional
	HealthCheckConnectPort *int64 `json:"healthCheckConnectPort,omitempty" tf:"health_check_connect_port"`
	// +optional
	HealthCheckDomain *string `json:"healthCheckDomain,omitempty" tf:"health_check_domain"`
	// +optional
	HealthCheckHTTPCode *string `json:"healthCheckHTTPCode,omitempty" tf:"health_check_http_code"`
	// +optional
	HealthCheckInterval *int64 `json:"healthCheckInterval,omitempty" tf:"health_check_interval"`
	// +optional
	HealthCheckTimeout *int64 `json:"healthCheckTimeout,omitempty" tf:"health_check_timeout"`
	// +optional
	HealthCheckURI *string `json:"healthCheckURI,omitempty" tf:"health_check_uri"`
	// +optional
	HealthyThreshold *int64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold"`
	// +optional
	ListenerSync   *string `json:"listenerSync,omitempty" tf:"listener_sync"`
	LoadBalancerID *string `json:"loadBalancerID" tf:"load_balancer_id"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Scheduler     *string `json:"scheduler,omitempty" tf:"scheduler"`
	ServerGroupID *string `json:"serverGroupID" tf:"server_group_id"`
	// +optional
	StickySession *string `json:"stickySession,omitempty" tf:"sticky_session"`
	// +optional
	StickySessionType *string `json:"stickySessionType,omitempty" tf:"sticky_session_type"`
	// +optional
	UnhealthyThreshold *int64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold"`
	// +optional
	Url *string `json:"url,omitempty" tf:"url"`
}

type RuleStatus struct {
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

// RuleList is a list of Rules
type RuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Rule CRD objects
	Items []Rule `json:"items,omitempty"`
}