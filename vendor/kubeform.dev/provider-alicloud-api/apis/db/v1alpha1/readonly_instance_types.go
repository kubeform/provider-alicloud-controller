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

type ReadonlyInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReadonlyInstanceSpec   `json:"spec,omitempty"`
	Status            ReadonlyInstanceStatus `json:"status,omitempty"`
}

type ReadonlyInstanceSpecParameters struct {
	Name  *string `json:"name" tf:"name"`
	Value *string `json:"value" tf:"value"`
}

type ReadonlyInstanceSpec struct {
	State *ReadonlyInstanceSpecResource `json:"state,omitempty" tf:"-"`

	Resource ReadonlyInstanceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ReadonlyInstanceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Acl *string `json:"acl,omitempty" tf:"acl"`
	// +optional
	CaType *string `json:"caType,omitempty" tf:"ca_type"`
	// +optional
	ClientCaCert *string `json:"clientCaCert,omitempty" tf:"client_ca_cert"`
	// +optional
	ClientCaEnabled *int64 `json:"clientCaEnabled,omitempty" tf:"client_ca_enabled"`
	// +optional
	ClientCertRevocationList *string `json:"clientCertRevocationList,omitempty" tf:"client_cert_revocation_list"`
	// +optional
	ClientCrlEnabled *int64 `json:"clientCrlEnabled,omitempty" tf:"client_crl_enabled"`
	// +optional
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string"`
	// +optional
	Engine        *string `json:"engine,omitempty" tf:"engine"`
	EngineVersion *string `json:"engineVersion" tf:"engine_version"`
	// +optional
	ForceRestart *bool `json:"forceRestart,omitempty" tf:"force_restart"`
	// +optional
	InstanceName       *string `json:"instanceName,omitempty" tf:"instance_name"`
	InstanceStorage    *int64  `json:"instanceStorage" tf:"instance_storage"`
	InstanceType       *string `json:"instanceType" tf:"instance_type"`
	MasterDbInstanceID *string `json:"masterDbInstanceID" tf:"master_db_instance_id"`
	// +optional
	Parameters []ReadonlyInstanceSpecParameters `json:"parameters,omitempty" tf:"parameters"`
	// +optional
	Port *string `json:"port,omitempty" tf:"port"`
	// +optional
	ReplicationACL *string `json:"replicationACL,omitempty" tf:"replication_acl"`
	// +optional
	ResourceGroupID *string `json:"resourceGroupID,omitempty" tf:"resource_group_id"`
	// +optional
	ServerCert *string `json:"serverCert,omitempty" tf:"server_cert"`
	// +optional
	ServerKey *string `json:"serverKey,omitempty" tf:"server_key"`
	// +optional
	SslEnabled *int64 `json:"sslEnabled,omitempty" tf:"ssl_enabled"`
	// +optional
	SwitchTime *string `json:"switchTime,omitempty" tf:"switch_time"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TargetMinorVersion *string `json:"targetMinorVersion,omitempty" tf:"target_minor_version"`
	// +optional
	UpgradeDbInstanceKernelVersion *bool `json:"upgradeDbInstanceKernelVersion,omitempty" tf:"upgrade_db_instance_kernel_version"`
	// +optional
	UpgradeTime *string `json:"upgradeTime,omitempty" tf:"upgrade_time"`
	// +optional
	VswitchID *string `json:"vswitchID,omitempty" tf:"vswitch_id"`
	// +optional
	ZoneID *string `json:"zoneID,omitempty" tf:"zone_id"`
}

type ReadonlyInstanceStatus struct {
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

// ReadonlyInstanceList is a list of ReadonlyInstances
type ReadonlyInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ReadonlyInstance CRD objects
	Items []ReadonlyInstance `json:"items,omitempty"`
}