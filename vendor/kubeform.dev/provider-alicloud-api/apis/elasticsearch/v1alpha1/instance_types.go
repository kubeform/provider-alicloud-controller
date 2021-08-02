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

type Instance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec,omitempty"`
	Status            InstanceStatus `json:"status,omitempty"`
}

type InstanceSpec struct {
	State *InstanceSpecResource `json:"state,omitempty" tf:"-"`

	Resource InstanceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type InstanceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ClientNodeAmount *int64 `json:"clientNodeAmount,omitempty" tf:"client_node_amount"`
	// +optional
	ClientNodeSpec *string `json:"clientNodeSpec,omitempty" tf:"client_node_spec"`
	DataNodeAmount *int64  `json:"dataNodeAmount" tf:"data_node_amount"`
	// +optional
	DataNodeDiskEncrypted *bool   `json:"dataNodeDiskEncrypted,omitempty" tf:"data_node_disk_encrypted"`
	DataNodeDiskSize      *int64  `json:"dataNodeDiskSize" tf:"data_node_disk_size"`
	DataNodeDiskType      *string `json:"dataNodeDiskType" tf:"data_node_disk_type"`
	DataNodeSpec          *string `json:"dataNodeSpec" tf:"data_node_spec"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	Domain *string `json:"domain,omitempty" tf:"domain"`
	// +optional
	EnableKibanaPrivateNetwork *bool `json:"enableKibanaPrivateNetwork,omitempty" tf:"enable_kibana_private_network"`
	// +optional
	EnableKibanaPublicNetwork *bool `json:"enableKibanaPublicNetwork,omitempty" tf:"enable_kibana_public_network"`
	// +optional
	EnablePublic *bool `json:"enablePublic,omitempty" tf:"enable_public"`
	// +optional
	InstanceChargeType *string `json:"instanceChargeType,omitempty" tf:"instance_charge_type"`
	// +optional
	KibanaDomain *string `json:"kibanaDomain,omitempty" tf:"kibana_domain"`
	// +optional
	KibanaPort *int64 `json:"kibanaPort,omitempty" tf:"kibana_port"`
	// +optional
	KibanaPrivateWhitelist []string `json:"kibanaPrivateWhitelist,omitempty" tf:"kibana_private_whitelist"`
	// +optional
	KibanaWhitelist []string `json:"kibanaWhitelist,omitempty" tf:"kibana_whitelist"`
	// +optional
	KmsEncryptedPassword *string `json:"kmsEncryptedPassword,omitempty" tf:"kms_encrypted_password"`
	// +optional
	KmsEncryptionContext map[string]string `json:"kmsEncryptionContext,omitempty" tf:"kms_encryption_context"`
	// +optional
	MasterNodeSpec *string `json:"masterNodeSpec,omitempty" tf:"master_node_spec"`
	// +optional
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	Period *int64 `json:"period,omitempty" tf:"period"`
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// +optional
	PrivateWhitelist []string `json:"privateWhitelist,omitempty" tf:"private_whitelist"`
	// +optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`
	// +optional
	PublicWhitelist []string `json:"publicWhitelist,omitempty" tf:"public_whitelist"`
	// +optional
	ResourceGroupID *string `json:"resourceGroupID,omitempty" tf:"resource_group_id"`
	// +optional
	SettingConfig map[string]string `json:"settingConfig,omitempty" tf:"setting_config"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	Tags      map[string]string `json:"tags,omitempty" tf:"tags"`
	Version   *string           `json:"version" tf:"version"`
	VswitchID *string           `json:"vswitchID" tf:"vswitch_id"`
	// +optional
	ZoneCount *int64 `json:"zoneCount,omitempty" tf:"zone_count"`
}

type InstanceStatus struct {
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

// InstanceList is a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Instance CRD objects
	Items []Instance `json:"items,omitempty"`
}