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

type InstanceSpecParameters struct {
	Name  *string `json:"name" tf:"name"`
	Value *string `json:"value" tf:"value"`
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
	AutoRenew *bool `json:"autoRenew,omitempty" tf:"auto_renew"`
	// +optional
	AutoRenewPeriod *int64 `json:"autoRenewPeriod,omitempty" tf:"auto_renew_period"`
	// +optional
	AutoUseCoupon *bool `json:"autoUseCoupon,omitempty" tf:"auto_use_coupon"`
	// +optional
	// Deprecated
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`
	// +optional
	BackupID *string `json:"backupID,omitempty" tf:"backup_id"`
	// +optional
	BackupPeriod []string `json:"backupPeriod,omitempty" tf:"backup_period"`
	// +optional
	BackupTime *string `json:"backupTime,omitempty" tf:"backup_time"`
	// +optional
	Bandwidth *int64 `json:"bandwidth,omitempty" tf:"bandwidth"`
	// +optional
	BusinessInfo *string `json:"businessInfo,omitempty" tf:"business_info"`
	// +optional
	Capacity *int64 `json:"capacity,omitempty" tf:"capacity"`
	// +optional
	Config map[string]string `json:"config,omitempty" tf:"config"`
	// +optional
	ConnectionDomain *string `json:"connectionDomain,omitempty" tf:"connection_domain"`
	// +optional
	// Deprecated
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string"`
	// +optional
	// Deprecated
	ConnectionStringPrefix *string `json:"connectionStringPrefix,omitempty" tf:"connection_string_prefix"`
	// +optional
	CouponNo *string `json:"couponNo,omitempty" tf:"coupon_no"`
	// +optional
	DbInstanceName *string `json:"dbInstanceName,omitempty" tf:"db_instance_name"`
	// +optional
	DedicatedHostGroupID *string `json:"dedicatedHostGroupID,omitempty" tf:"dedicated_host_group_id"`
	// +optional
	DryRun *bool `json:"dryRun,omitempty" tf:"dry_run"`
	// +optional
	EnableBackupLog *int64 `json:"enableBackupLog,omitempty" tf:"enable_backup_log"`
	// +optional
	// Deprecated
	EnablePublic *bool `json:"enablePublic,omitempty" tf:"enable_public"`
	// +optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time"`
	// +optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version"`
	// +optional
	ForceUpgrade *bool `json:"forceUpgrade,omitempty" tf:"force_upgrade"`
	// +optional
	GlobalInstance *bool `json:"globalInstance,omitempty" tf:"global_instance"`
	// +optional
	GlobalInstanceID *string `json:"globalInstanceID,omitempty" tf:"global_instance_id"`
	// +optional
	// Deprecated
	InstanceChargeType *string `json:"instanceChargeType,omitempty" tf:"instance_charge_type"`
	// +optional
	InstanceClass *string `json:"instanceClass,omitempty" tf:"instance_class"`
	// +optional
	// Deprecated
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name"`
	// +optional
	InstanceReleaseProtection *bool `json:"instanceReleaseProtection,omitempty" tf:"instance_release_protection"`
	// +optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type"`
	// +optional
	KmsEncryptedPassword *string `json:"kmsEncryptedPassword,omitempty" tf:"kms_encrypted_password"`
	// +optional
	KmsEncryptionContext map[string]string `json:"kmsEncryptionContext,omitempty" tf:"kms_encryption_context"`
	// +optional
	MaintainEndTime *string `json:"maintainEndTime,omitempty" tf:"maintain_end_time"`
	// +optional
	MaintainStartTime *string `json:"maintainStartTime,omitempty" tf:"maintain_start_time"`
	// +optional
	ModifyMode *int64 `json:"modifyMode,omitempty" tf:"modify_mode"`
	// +optional
	// Deprecated
	NodeType *string `json:"nodeType,omitempty" tf:"node_type"`
	// +optional
	OrderType *string `json:"orderType,omitempty" tf:"order_type"`
	// +optional
	// Deprecated
	Parameters []InstanceSpecParameters `json:"parameters,omitempty" tf:"parameters"`
	// +optional
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	PaymentType *string `json:"paymentType,omitempty" tf:"payment_type"`
	// +optional
	Period *string `json:"period,omitempty" tf:"period"`
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// +optional
	PrivateConnectionPort *string `json:"privateConnectionPort,omitempty" tf:"private_connection_port"`
	// +optional
	PrivateConnectionPrefix *string `json:"privateConnectionPrefix,omitempty" tf:"private_connection_prefix"`
	// +optional
	PrivateIP *string `json:"privateIP,omitempty" tf:"private_ip"`
	// +optional
	Qps *int64 `json:"qps,omitempty" tf:"qps"`
	// +optional
	ResourceGroupID *string `json:"resourceGroupID,omitempty" tf:"resource_group_id"`
	// +optional
	RestoreTime *string `json:"restoreTime,omitempty" tf:"restore_time"`
	// +optional
	SecondaryZoneID *string `json:"secondaryZoneID,omitempty" tf:"secondary_zone_id"`
	// +optional
	SecurityGroupID *string `json:"securityGroupID,omitempty" tf:"security_group_id"`
	// +optional
	SecurityIPGroupAttribute *string `json:"securityIPGroupAttribute,omitempty" tf:"security_ip_group_attribute"`
	// +optional
	SecurityIPGroupName *string `json:"securityIPGroupName,omitempty" tf:"security_ip_group_name"`
	// +optional
	SecurityIPS []string `json:"securityIPS,omitempty" tf:"security_ips"`
	// +optional
	SrcdbInstanceID *string `json:"srcdbInstanceID,omitempty" tf:"srcdb_instance_id"`
	// +optional
	SslEnable *string `json:"sslEnable,omitempty" tf:"ssl_enable"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	VpcAuthMode *string `json:"vpcAuthMode,omitempty" tf:"vpc_auth_mode"`
	// +optional
	VswitchID *string `json:"vswitchID,omitempty" tf:"vswitch_id"`
	// +optional
	ZoneID *string `json:"zoneID,omitempty" tf:"zone_id"`
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
