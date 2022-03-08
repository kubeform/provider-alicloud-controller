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

type ManagedKubernetes struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedKubernetesSpec   `json:"spec,omitempty"`
	Status            ManagedKubernetesStatus `json:"status,omitempty"`
}

type ManagedKubernetesSpecAddons struct {
	// +optional
	Config *string `json:"config,omitempty" tf:"config"`
	// +optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type ManagedKubernetesSpecCertificateAuthority struct {
	// +optional
	ClientCert *string `json:"clientCert,omitempty" tf:"client_cert"`
	// +optional
	ClientKey *string `json:"clientKey,omitempty" tf:"client_key"`
	// +optional
	ClusterCert *string `json:"clusterCert,omitempty" tf:"cluster_cert"`
}

type ManagedKubernetesSpecConnections struct {
	// +optional
	ApiServerInternet *string `json:"apiServerInternet,omitempty" tf:"api_server_internet"`
	// +optional
	ApiServerIntranet *string `json:"apiServerIntranet,omitempty" tf:"api_server_intranet"`
	// +optional
	MasterPublicIP *string `json:"masterPublicIP,omitempty" tf:"master_public_ip"`
	// +optional
	ServiceDomain *string `json:"serviceDomain,omitempty" tf:"service_domain"`
}

type ManagedKubernetesSpecLogConfig struct {
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	Type    *string `json:"type" tf:"type"`
}

type ManagedKubernetesSpecMaintenanceWindow struct {
	Duration        *string `json:"duration" tf:"duration"`
	Enable          *bool   `json:"enable" tf:"enable"`
	MaintenanceTime *string `json:"maintenanceTime" tf:"maintenance_time"`
	WeeklyPeriod    *string `json:"weeklyPeriod" tf:"weekly_period"`
}

type ManagedKubernetesSpecRuntime struct {
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
}

type ManagedKubernetesSpecTaints struct {
	// +optional
	Effect *string `json:"effect,omitempty" tf:"effect"`
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type ManagedKubernetesSpecWorkerDataDisks struct {
	// +optional
	AutoSnapshotPolicyID *string `json:"autoSnapshotPolicyID,omitempty" tf:"auto_snapshot_policy_id"`
	// +optional
	Category *string `json:"category,omitempty" tf:"category"`
	// +optional
	Device *string `json:"device,omitempty" tf:"device"`
	// +optional
	Encrypted *string `json:"encrypted,omitempty" tf:"encrypted"`
	// +optional
	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	PerformanceLevel *string `json:"performanceLevel,omitempty" tf:"performance_level"`
	// +optional
	Size *string `json:"size,omitempty" tf:"size"`
	// +optional
	SnapshotID *string `json:"snapshotID,omitempty" tf:"snapshot_id"`
}

type ManagedKubernetesSpecWorkerNodes struct {
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	PrivateIP *string `json:"privateIP,omitempty" tf:"private_ip"`
}

type ManagedKubernetesSpec struct {
	State *ManagedKubernetesSpecResource `json:"state,omitempty" tf:"-"`

	Resource ManagedKubernetesSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ManagedKubernetesSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Addons []ManagedKubernetesSpecAddons `json:"addons,omitempty" tf:"addons"`
	// +optional
	ApiAudiences []string `json:"apiAudiences,omitempty" tf:"api_audiences"`
	// +optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`
	// +optional
	CertificateAuthority map[string]ManagedKubernetesSpecCertificateAuthority `json:"certificateAuthority,omitempty" tf:"certificate_authority"`
	// +optional
	ClientCert *string `json:"clientCert,omitempty" tf:"client_cert"`
	// +optional
	ClientKey *string `json:"clientKey,omitempty" tf:"client_key"`
	// +optional
	ClusterCaCert *string `json:"clusterCaCert,omitempty" tf:"cluster_ca_cert"`
	// cluster local domain
	// +optional
	ClusterDomain *string `json:"clusterDomain,omitempty" tf:"cluster_domain"`
	// +optional
	ClusterNetworkType *string `json:"clusterNetworkType,omitempty" tf:"cluster_network_type"`
	// +optional
	ClusterSpec *string `json:"clusterSpec,omitempty" tf:"cluster_spec"`
	// +optional
	Connections map[string]ManagedKubernetesSpecConnections `json:"connections,omitempty" tf:"connections"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	ControlPlaneLogComponents []string `json:"controlPlaneLogComponents,omitempty" tf:"control_plane_log_components"`
	// +optional
	ControlPlaneLogProject *string `json:"controlPlaneLogProject,omitempty" tf:"control_plane_log_project"`
	// +optional
	ControlPlaneLogTtl *string `json:"controlPlaneLogTtl,omitempty" tf:"control_plane_log_ttl"`
	// +optional
	CpuPolicy *string `json:"cpuPolicy,omitempty" tf:"cpu_policy"`
	// +optional
	CustomSan *string `json:"customSan,omitempty" tf:"custom_san"`
	// +optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection"`
	// +optional
	EnableSSH *bool `json:"enableSSH,omitempty" tf:"enable_ssh"`
	// disk encryption key, only in ack-pro
	// +optional
	EncryptionProviderKey *string `json:"encryptionProviderKey,omitempty" tf:"encryption_provider_key"`
	// +optional
	ExcludeAutoscalerNodes *bool `json:"excludeAutoscalerNodes,omitempty" tf:"exclude_autoscaler_nodes"`
	// +optional
	ForceUpdate *bool `json:"forceUpdate,omitempty" tf:"force_update"`
	// +optional
	ImageID *string `json:"imageID,omitempty" tf:"image_id"`
	// +optional
	InstallCloudMonitor *bool `json:"installCloudMonitor,omitempty" tf:"install_cloud_monitor"`
	// +optional
	IsEnterpriseSecurityGroup *bool `json:"isEnterpriseSecurityGroup,omitempty" tf:"is_enterprise_security_group"`
	// +optional
	KeyName *string `json:"keyName,omitempty" tf:"key_name"`
	// +optional
	KmsEncryptedPassword *string `json:"kmsEncryptedPassword,omitempty" tf:"kms_encrypted_password"`
	// +optional
	KmsEncryptionContext map[string]string `json:"kmsEncryptionContext,omitempty" tf:"kms_encryption_context"`
	// +optional
	KubeConfig *string `json:"kubeConfig,omitempty" tf:"kube_config"`
	// +optional
	LoadBalancerSpec *string `json:"loadBalancerSpec,omitempty" tf:"load_balancer_spec"`
	// +optional
	LogConfig *ManagedKubernetesSpecLogConfig `json:"logConfig,omitempty" tf:"log_config"`
	// +optional
	MaintenanceWindow *ManagedKubernetesSpecMaintenanceWindow `json:"maintenanceWindow,omitempty" tf:"maintenance_window"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`
	// +optional
	NatGatewayID *string `json:"natGatewayID,omitempty" tf:"nat_gateway_id"`
	// +optional
	NewNATGateway *bool `json:"newNATGateway,omitempty" tf:"new_nat_gateway"`
	// +optional
	NodeCIDRMask *int64 `json:"nodeCIDRMask,omitempty" tf:"node_cidr_mask"`
	// +optional
	NodeNameMode *string `json:"nodeNameMode,omitempty" tf:"node_name_mode"`
	// +optional
	NodePortRange *string `json:"nodePortRange,omitempty" tf:"node_port_range"`
	// +optional
	OsType *string `json:"osType,omitempty" tf:"os_type"`
	// +optional
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	Platform *string `json:"platform,omitempty" tf:"platform"`
	// +optional
	PodCIDR *string `json:"podCIDR,omitempty" tf:"pod_cidr"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	PodVswitchIDS []string `json:"podVswitchIDS,omitempty" tf:"pod_vswitch_ids"`
	// +optional
	ProxyMode *string `json:"proxyMode,omitempty" tf:"proxy_mode"`
	// +optional
	RdsInstances []string `json:"rdsInstances,omitempty" tf:"rds_instances"`
	// +optional
	ResourceGroupID *string `json:"resourceGroupID,omitempty" tf:"resource_group_id"`
	// +optional
	RetainResources []string `json:"retainResources,omitempty" tf:"retain_resources"`
	// +optional
	Runtime map[string]ManagedKubernetesSpecRuntime `json:"runtime,omitempty" tf:"runtime"`
	// +optional
	SecurityGroupID *string `json:"securityGroupID,omitempty" tf:"security_group_id"`
	// +optional
	ServiceAccountIssuer *string `json:"serviceAccountIssuer,omitempty" tf:"service_account_issuer"`
	// +optional
	ServiceCIDR *string `json:"serviceCIDR,omitempty" tf:"service_cidr"`
	// +optional
	// Deprecated
	SlbID *string `json:"slbID,omitempty" tf:"slb_id"`
	// +optional
	SlbInternet *string `json:"slbInternet,omitempty" tf:"slb_internet"`
	// +optional
	SlbInternetEnabled *bool `json:"slbInternetEnabled,omitempty" tf:"slb_internet_enabled"`
	// +optional
	SlbIntranet *string `json:"slbIntranet,omitempty" tf:"slb_intranet"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	Taints []ManagedKubernetesSpecTaints `json:"taints,omitempty" tf:"taints"`
	// +optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone"`
	// +optional
	UserCa *string `json:"userCa,omitempty" tf:"user_ca"`
	// +optional
	UserData *string `json:"userData,omitempty" tf:"user_data"`
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
	// +optional
	VpcID *string `json:"vpcID,omitempty" tf:"vpc_id"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	// +kubebuilder:validation:MinItems=3
	VswitchIDS []string `json:"vswitchIDS,omitempty" tf:"vswitch_ids"`
	// +optional
	WorkerAutoRenew *bool `json:"workerAutoRenew,omitempty" tf:"worker_auto_renew"`
	// +optional
	WorkerAutoRenewPeriod *int64 `json:"workerAutoRenewPeriod,omitempty" tf:"worker_auto_renew_period"`
	// +optional
	WorkerDataDiskCategory *string `json:"workerDataDiskCategory,omitempty" tf:"worker_data_disk_category"`
	// +optional
	WorkerDataDiskSize *int64 `json:"workerDataDiskSize,omitempty" tf:"worker_data_disk_size"`
	// +optional
	WorkerDataDisks []ManagedKubernetesSpecWorkerDataDisks `json:"workerDataDisks,omitempty" tf:"worker_data_disks"`
	// +optional
	WorkerDiskCategory *string `json:"workerDiskCategory,omitempty" tf:"worker_disk_category"`
	// +optional
	WorkerDiskPerformanceLevel *string `json:"workerDiskPerformanceLevel,omitempty" tf:"worker_disk_performance_level"`
	// +optional
	WorkerDiskSize *int64 `json:"workerDiskSize,omitempty" tf:"worker_disk_size"`
	// +optional
	WorkerDiskSnapshotPolicyID *string `json:"workerDiskSnapshotPolicyID,omitempty" tf:"worker_disk_snapshot_policy_id"`
	// +optional
	WorkerInstanceChargeType *string `json:"workerInstanceChargeType,omitempty" tf:"worker_instance_charge_type"`
	// +optional
	WorkerInstanceType *string `json:"workerInstanceType,omitempty" tf:"worker_instance_type"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	WorkerInstanceTypes []string `json:"workerInstanceTypes,omitempty" tf:"worker_instance_types"`
	// +optional
	WorkerNodes []ManagedKubernetesSpecWorkerNodes `json:"workerNodes,omitempty" tf:"worker_nodes"`
	// +optional
	WorkerNumber *int64 `json:"workerNumber,omitempty" tf:"worker_number"`
	// +optional
	// +kubebuilder:validation:MaxItems=3
	// +kubebuilder:validation:MinItems=1
	WorkerNumbers []int64 `json:"workerNumbers,omitempty" tf:"worker_numbers"`
	// +optional
	WorkerPeriod *int64 `json:"workerPeriod,omitempty" tf:"worker_period"`
	// +optional
	WorkerPeriodUnit *string `json:"workerPeriodUnit,omitempty" tf:"worker_period_unit"`
	// +optional
	WorkerRAMRoleName *string `json:"workerRAMRoleName,omitempty" tf:"worker_ram_role_name"`
	// +kubebuilder:validation:MinItems=1
	WorkerVswitchIDS []string `json:"workerVswitchIDS" tf:"worker_vswitch_ids"`
}

type ManagedKubernetesStatus struct {
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

// ManagedKubernetesList is a list of ManagedKubernetess
type ManagedKubernetesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ManagedKubernetes CRD objects
	Items []ManagedKubernetes `json:"items,omitempty"`
}
