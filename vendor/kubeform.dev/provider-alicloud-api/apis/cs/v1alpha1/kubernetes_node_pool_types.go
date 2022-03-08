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

type KubernetesNodePool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubernetesNodePoolSpec   `json:"spec,omitempty"`
	Status            KubernetesNodePoolStatus `json:"status,omitempty"`
}

type KubernetesNodePoolSpecDataDisks struct {
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
	Size *int64 `json:"size,omitempty" tf:"size"`
	// +optional
	SnapshotID *string `json:"snapshotID,omitempty" tf:"snapshot_id"`
}

type KubernetesNodePoolSpecLabels struct {
	Key *string `json:"key" tf:"key"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type KubernetesNodePoolSpecManagement struct {
	// +optional
	AutoRepair *bool `json:"autoRepair,omitempty" tf:"auto_repair"`
	// +optional
	AutoUpgrade    *bool  `json:"autoUpgrade,omitempty" tf:"auto_upgrade"`
	MaxUnavailable *int64 `json:"maxUnavailable" tf:"max_unavailable"`
	// +optional
	Surge *int64 `json:"surge,omitempty" tf:"surge"`
	// +optional
	SurgePercentage *int64 `json:"surgePercentage,omitempty" tf:"surge_percentage"`
}

type KubernetesNodePoolSpecScalingConfig struct {
	// +optional
	EipBandwidth *int64 `json:"eipBandwidth,omitempty" tf:"eip_bandwidth"`
	// +optional
	EipInternetChargeType *string `json:"eipInternetChargeType,omitempty" tf:"eip_internet_charge_type"`
	// +optional
	IsBondEip *bool  `json:"isBondEip,omitempty" tf:"is_bond_eip"`
	MaxSize   *int64 `json:"maxSize" tf:"max_size"`
	MinSize   *int64 `json:"minSize" tf:"min_size"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type KubernetesNodePoolSpecSpotPriceLimit struct {
	// +optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type"`
	// +optional
	PriceLimit *string `json:"priceLimit,omitempty" tf:"price_limit"`
}

type KubernetesNodePoolSpecTaints struct {
	// +optional
	Effect *string `json:"effect,omitempty" tf:"effect"`
	Key    *string `json:"key" tf:"key"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type KubernetesNodePoolSpec struct {
	State *KubernetesNodePoolSpecResource `json:"state,omitempty" tf:"-"`

	Resource KubernetesNodePoolSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type KubernetesNodePoolSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AutoRenew *bool `json:"autoRenew,omitempty" tf:"auto_renew"`
	// +optional
	AutoRenewPeriod *int64  `json:"autoRenewPeriod,omitempty" tf:"auto_renew_period"`
	ClusterID       *string `json:"clusterID" tf:"cluster_id"`
	// +optional
	DataDisks []KubernetesNodePoolSpecDataDisks `json:"dataDisks,omitempty" tf:"data_disks"`
	// +optional
	DeploymentSetID *string `json:"deploymentSetID,omitempty" tf:"deployment_set_id"`
	// +optional
	DesiredSize *int64 `json:"desiredSize,omitempty" tf:"desired_size"`
	// +optional
	FormatDisk *bool `json:"formatDisk,omitempty" tf:"format_disk"`
	// +optional
	ImageID *string `json:"imageID,omitempty" tf:"image_id"`
	// +optional
	ImageType *string `json:"imageType,omitempty" tf:"image_type"`
	// +optional
	InstallCloudMonitor *bool `json:"installCloudMonitor,omitempty" tf:"install_cloud_monitor"`
	// +optional
	InstanceChargeType *string `json:"instanceChargeType,omitempty" tf:"instance_charge_type"`
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:MinItems=1
	InstanceTypes []string `json:"instanceTypes" tf:"instance_types"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	Instances []string `json:"instances,omitempty" tf:"instances"`
	// +optional
	InternetChargeType *string `json:"internetChargeType,omitempty" tf:"internet_charge_type"`
	// +optional
	InternetMaxBandwidthOut *int64 `json:"internetMaxBandwidthOut,omitempty" tf:"internet_max_bandwidth_out"`
	// +optional
	KeepInstanceName *bool `json:"keepInstanceName,omitempty" tf:"keep_instance_name"`
	// +optional
	KeyName *string `json:"keyName,omitempty" tf:"key_name"`
	// +optional
	KmsEncryptedPassword *string `json:"kmsEncryptedPassword,omitempty" tf:"kms_encrypted_password"`
	// +optional
	Labels []KubernetesNodePoolSpecLabels `json:"labels,omitempty" tf:"labels"`
	// +optional
	Management *KubernetesNodePoolSpecManagement `json:"management,omitempty" tf:"management"`
	Name       *string                           `json:"name" tf:"name"`
	// +optional
	// Deprecated
	NodeCount *int64 `json:"nodeCount,omitempty" tf:"node_count"`
	// +optional
	NodeNameMode *string `json:"nodeNameMode,omitempty" tf:"node_name_mode"`
	// +optional
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	Period *int64 `json:"period,omitempty" tf:"period"`
	// +optional
	PeriodUnit *string `json:"periodUnit,omitempty" tf:"period_unit"`
	// +optional
	// Deprecated
	Platform *string `json:"platform,omitempty" tf:"platform"`
	// +optional
	ResourceGroupID *string `json:"resourceGroupID,omitempty" tf:"resource_group_id"`
	// +optional
	RuntimeName *string `json:"runtimeName,omitempty" tf:"runtime_name"`
	// +optional
	RuntimeVersion *string `json:"runtimeVersion,omitempty" tf:"runtime_version"`
	// +optional
	ScalingConfig *KubernetesNodePoolSpecScalingConfig `json:"scalingConfig,omitempty" tf:"scaling_config"`
	// +optional
	ScalingGroupID *string `json:"scalingGroupID,omitempty" tf:"scaling_group_id"`
	// +optional
	ScalingPolicy *string `json:"scalingPolicy,omitempty" tf:"scaling_policy"`
	// +optional
	// Deprecated
	SecurityGroupID *string `json:"securityGroupID,omitempty" tf:"security_group_id"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	SecurityGroupIDS []string `json:"securityGroupIDS,omitempty" tf:"security_group_ids"`
	// +optional
	SpotPriceLimit []KubernetesNodePoolSpecSpotPriceLimit `json:"spotPriceLimit,omitempty" tf:"spot_price_limit"`
	// +optional
	SpotStrategy *string `json:"spotStrategy,omitempty" tf:"spot_strategy"`
	// +optional
	SystemDiskCategory *string `json:"systemDiskCategory,omitempty" tf:"system_disk_category"`
	// +optional
	SystemDiskPerformanceLevel *string `json:"systemDiskPerformanceLevel,omitempty" tf:"system_disk_performance_level"`
	// +optional
	SystemDiskSize *int64 `json:"systemDiskSize,omitempty" tf:"system_disk_size"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	Taints []KubernetesNodePoolSpecTaints `json:"taints,omitempty" tf:"taints"`
	// +optional
	Unschedulable *bool `json:"unschedulable,omitempty" tf:"unschedulable"`
	// +optional
	UserData *string `json:"userData,omitempty" tf:"user_data"`
	// +optional
	VpcID *string `json:"vpcID,omitempty" tf:"vpc_id"`
	// +kubebuilder:validation:MinItems=1
	VswitchIDS []string `json:"vswitchIDS" tf:"vswitch_ids"`
}

type KubernetesNodePoolStatus struct {
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

// KubernetesNodePoolList is a list of KubernetesNodePools
type KubernetesNodePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubernetesNodePool CRD objects
	Items []KubernetesNodePool `json:"items,omitempty"`
}
