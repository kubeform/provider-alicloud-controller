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

type EdgeKubernetes struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EdgeKubernetesSpec   `json:"spec,omitempty"`
	Status            EdgeKubernetesStatus `json:"status,omitempty"`
}

type EdgeKubernetesSpecAddons struct {
	// +optional
	Config *string `json:"config,omitempty" tf:"config"`
	// +optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type EdgeKubernetesSpecCertificateAuthority struct {
	// +optional
	ClientCert *string `json:"clientCert,omitempty" tf:"client_cert"`
	// +optional
	ClientKey *string `json:"clientKey,omitempty" tf:"client_key"`
	// +optional
	ClusterCert *string `json:"clusterCert,omitempty" tf:"cluster_cert"`
}

type EdgeKubernetesSpecConnections struct {
	// +optional
	ApiServerInternet *string `json:"apiServerInternet,omitempty" tf:"api_server_internet"`
	// +optional
	ApiServerIntranet *string `json:"apiServerIntranet,omitempty" tf:"api_server_intranet"`
	// +optional
	MasterPublicIP *string `json:"masterPublicIP,omitempty" tf:"master_public_ip"`
	// +optional
	ServiceDomain *string `json:"serviceDomain,omitempty" tf:"service_domain"`
}

type EdgeKubernetesSpecLogConfig struct {
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	Type    *string `json:"type" tf:"type"`
}

type EdgeKubernetesSpecWorkerDataDisks struct {
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

type EdgeKubernetesSpecWorkerNodes struct {
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	PrivateIP *string `json:"privateIP,omitempty" tf:"private_ip"`
}

type EdgeKubernetesSpec struct {
	State *EdgeKubernetesSpecResource `json:"state,omitempty" tf:"-"`

	Resource EdgeKubernetesSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type EdgeKubernetesSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Addons []EdgeKubernetesSpecAddons `json:"addons,omitempty" tf:"addons"`
	// +optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`
	// +optional
	CertificateAuthority map[string]EdgeKubernetesSpecCertificateAuthority `json:"certificateAuthority,omitempty" tf:"certificate_authority"`
	// +optional
	ClientCert *string `json:"clientCert,omitempty" tf:"client_cert"`
	// +optional
	ClientKey *string `json:"clientKey,omitempty" tf:"client_key"`
	// +optional
	ClusterCaCert *string `json:"clusterCaCert,omitempty" tf:"cluster_ca_cert"`
	// +optional
	Connections map[string]EdgeKubernetesSpecConnections `json:"connections,omitempty" tf:"connections"`
	// +optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection"`
	// +optional
	ForceUpdate *bool `json:"forceUpdate,omitempty" tf:"force_update"`
	// +optional
	InstallCloudMonitor *bool `json:"installCloudMonitor,omitempty" tf:"install_cloud_monitor"`
	// +optional
	IsEnterpriseSecurityGroup *bool `json:"isEnterpriseSecurityGroup,omitempty" tf:"is_enterprise_security_group"`
	// +optional
	KeyName *string `json:"keyName,omitempty" tf:"key_name"`
	// +optional
	KubeConfig *string `json:"kubeConfig,omitempty" tf:"kube_config"`
	// +optional
	LogConfig *EdgeKubernetesSpecLogConfig `json:"logConfig,omitempty" tf:"log_config"`
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
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	PodCIDR *string `json:"podCIDR,omitempty" tf:"pod_cidr"`
	// +optional
	ProxyMode *string `json:"proxyMode,omitempty" tf:"proxy_mode"`
	// +optional
	RdsInstances []string `json:"rdsInstances,omitempty" tf:"rds_instances"`
	// +optional
	ResourceGroupID *string `json:"resourceGroupID,omitempty" tf:"resource_group_id"`
	// +optional
	SecurityGroupID *string `json:"securityGroupID,omitempty" tf:"security_group_id"`
	// +optional
	ServiceCIDR *string `json:"serviceCIDR,omitempty" tf:"service_cidr"`
	// +optional
	SlbInternet *string `json:"slbInternet,omitempty" tf:"slb_internet"`
	// +optional
	SlbInternetEnabled *bool `json:"slbInternetEnabled,omitempty" tf:"slb_internet_enabled"`
	// +optional
	SlbIntranet *string `json:"slbIntranet,omitempty" tf:"slb_intranet"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	UserData *string `json:"userData,omitempty" tf:"user_data"`
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
	// +optional
	VpcID *string `json:"vpcID,omitempty" tf:"vpc_id"`
	// +optional
	WorkerDataDisks []EdgeKubernetesSpecWorkerDataDisks `json:"workerDataDisks,omitempty" tf:"worker_data_disks"`
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
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:MinItems=1
	WorkerInstanceTypes []string `json:"workerInstanceTypes" tf:"worker_instance_types"`
	// +optional
	WorkerNodes  []EdgeKubernetesSpecWorkerNodes `json:"workerNodes,omitempty" tf:"worker_nodes"`
	WorkerNumber *int64                          `json:"workerNumber" tf:"worker_number"`
	// +kubebuilder:validation:MinItems=1
	WorkerVswitchIDS []string `json:"workerVswitchIDS" tf:"worker_vswitch_ids"`
}

type EdgeKubernetesStatus struct {
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

// EdgeKubernetesList is a list of EdgeKubernetess
type EdgeKubernetesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EdgeKubernetes CRD objects
	Items []EdgeKubernetes `json:"items,omitempty"`
}
