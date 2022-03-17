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

type UpgradeDbInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UpgradeDbInstanceSpec   `json:"spec,omitempty"`
	Status            UpgradeDbInstanceStatus `json:"status,omitempty"`
}

type UpgradeDbInstanceSpecParameters struct {
	Name  *string `json:"name" tf:"name"`
	Value *string `json:"value" tf:"value"`
}

type UpgradeDbInstanceSpecPgHbaConf struct {
	Address  *string `json:"address" tf:"address"`
	Database *string `json:"database" tf:"database"`
	// +optional
	Mask   *string `json:"mask,omitempty" tf:"mask"`
	Method *string `json:"method" tf:"method"`
	// +optional
	Option     *string `json:"option,omitempty" tf:"option"`
	PriorityID *int64  `json:"priorityID" tf:"priority_id"`
	Type       *string `json:"type" tf:"type"`
	User       *string `json:"user" tf:"user"`
}

type UpgradeDbInstanceSpec struct {
	State *UpgradeDbInstanceSpecResource `json:"state,omitempty" tf:"-"`

	Resource UpgradeDbInstanceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type UpgradeDbInstanceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Acl *string `json:"acl,omitempty" tf:"acl"`
	// +optional
	AutoUpgradeMinorVersion *string `json:"autoUpgradeMinorVersion,omitempty" tf:"auto_upgrade_minor_version"`
	// +optional
	CaType *string `json:"caType,omitempty" tf:"ca_type"`
	// +optional
	Certificate *string `json:"certificate,omitempty" tf:"certificate"`
	// +optional
	ClientCaCert *string `json:"clientCaCert,omitempty" tf:"client_ca_cert"`
	// +optional
	ClientCaEnabled *int64 `json:"clientCaEnabled,omitempty" tf:"client_ca_enabled"`
	// +optional
	ClientCertRevocationList *string `json:"clientCertRevocationList,omitempty" tf:"client_cert_revocation_list"`
	// +optional
	ClientCrlEnabled *int64  `json:"clientCrlEnabled,omitempty" tf:"client_crl_enabled"`
	CollectStatMode  *string `json:"collectStatMode" tf:"collect_stat_mode"`
	// +optional
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string"`
	// +optional
	ConnectionStringPrefix *string `json:"connectionStringPrefix,omitempty" tf:"connection_string_prefix"`
	DbInstanceClass        *string `json:"dbInstanceClass" tf:"db_instance_class"`
	// +optional
	DbInstanceDescription *string `json:"dbInstanceDescription,omitempty" tf:"db_instance_description"`
	DbInstanceStorage     *int64  `json:"dbInstanceStorage" tf:"db_instance_storage"`
	DbInstanceStorageType *string `json:"dbInstanceStorageType" tf:"db_instance_storage_type"`
	// +optional
	DbName *string `json:"dbName,omitempty" tf:"db_name"`
	// +optional
	DedicatedHostGroupID *string `json:"dedicatedHostGroupID,omitempty" tf:"dedicated_host_group_id"`
	// +optional
	Direction *string `json:"direction,omitempty" tf:"direction"`
	// +optional
	EffectiveTime *string `json:"effectiveTime,omitempty" tf:"effective_time"`
	// +optional
	EncryptionKey *string `json:"encryptionKey,omitempty" tf:"encryption_key"`
	// +optional
	Engine *string `json:"engine,omitempty" tf:"engine"`
	// +optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version"`
	// +optional
	ForceRestart *bool `json:"forceRestart,omitempty" tf:"force_restart"`
	// +optional
	HaMode              *string `json:"haMode,omitempty" tf:"ha_mode"`
	InstanceNetworkType *string `json:"instanceNetworkType" tf:"instance_network_type"`
	// +optional
	MaintainTime *string `json:"maintainTime,omitempty" tf:"maintain_time"`
	// +optional
	Parameters []UpgradeDbInstanceSpecParameters `json:"parameters,omitempty" tf:"parameters"`
	// +optional
	Password    *string `json:"password,omitempty" tf:"password"`
	PaymentType *string `json:"paymentType" tf:"payment_type"`
	// +optional
	PgHbaConf []UpgradeDbInstanceSpecPgHbaConf `json:"pgHbaConf,omitempty" tf:"pg_hba_conf"`
	// +optional
	Port *string `json:"port,omitempty" tf:"port"`
	// +optional
	PrivateIPAddress *string `json:"privateIPAddress,omitempty" tf:"private_ip_address"`
	// +optional
	PrivateKey *string `json:"privateKey,omitempty" tf:"private_key"`
	// +optional
	ReleasedKeepPolicy *string `json:"releasedKeepPolicy,omitempty" tf:"released_keep_policy"`
	// +optional
	ReplicationACL *string `json:"replicationACL,omitempty" tf:"replication_acl"`
	// +optional
	ResourceGroupID *string `json:"resourceGroupID,omitempty" tf:"resource_group_id"`
	// +optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn"`
	// +optional
	SecurityIPS []string `json:"securityIPS,omitempty" tf:"security_ips"`
	// +optional
	ServerCert *string `json:"serverCert,omitempty" tf:"server_cert"`
	// +optional
	ServerKey *string `json:"serverKey,omitempty" tf:"server_key"`
	// +optional
	SourceBiz          *string `json:"sourceBiz,omitempty" tf:"source_biz"`
	SourceDbInstanceID *string `json:"sourceDbInstanceID" tf:"source_db_instance_id"`
	// +optional
	SslEnabled *int64  `json:"sslEnabled,omitempty" tf:"ssl_enabled"`
	SwitchOver *string `json:"switchOver" tf:"switch_over"`
	// +optional
	SwitchTime *string `json:"switchTime,omitempty" tf:"switch_time"`
	// +optional
	SwitchTimeMode *string `json:"switchTimeMode,omitempty" tf:"switch_time_mode"`
	// +optional
	SyncMode           *string `json:"syncMode,omitempty" tf:"sync_mode"`
	TargetMajorVersion *string `json:"targetMajorVersion" tf:"target_major_version"`
	// +optional
	TdeStatus *string `json:"tdeStatus,omitempty" tf:"tde_status"`
	// +optional
	VpcID *string `json:"vpcID,omitempty" tf:"vpc_id"`
	// +optional
	VswitchID *string `json:"vswitchID,omitempty" tf:"vswitch_id"`
	// +optional
	ZoneID *string `json:"zoneID,omitempty" tf:"zone_id"`
	// +optional
	ZoneIDSlave1 *string `json:"zoneIDSlave1,omitempty" tf:"zone_id_slave_1"`
}

type UpgradeDbInstanceStatus struct {
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

// UpgradeDbInstanceList is a list of UpgradeDbInstances
type UpgradeDbInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of UpgradeDbInstance CRD objects
	Items []UpgradeDbInstance `json:"items,omitempty"`
}