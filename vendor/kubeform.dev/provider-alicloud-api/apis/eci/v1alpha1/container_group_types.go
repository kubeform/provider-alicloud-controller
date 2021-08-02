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

type ContainerGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerGroupSpec   `json:"spec,omitempty"`
	Status            ContainerGroupStatus `json:"status,omitempty"`
}

type ContainerGroupSpecContainersEnvironmentVars struct {
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type ContainerGroupSpecContainersPorts struct {
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// +optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`
}

type ContainerGroupSpecContainersVolumeMounts struct {
	// +optional
	MountPath *string `json:"mountPath,omitempty" tf:"mount_path"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	ReadOnly *bool `json:"readOnly,omitempty" tf:"read_only"`
}

type ContainerGroupSpecContainers struct {
	// +optional
	Args []string `json:"args,omitempty" tf:"args"`
	// +optional
	Commands []string `json:"commands,omitempty" tf:"commands"`
	// +optional
	Cpu *float64 `json:"cpu,omitempty" tf:"cpu"`
	// +optional
	EnvironmentVars []ContainerGroupSpecContainersEnvironmentVars `json:"environmentVars,omitempty" tf:"environment_vars"`
	// +optional
	Gpu   *int64  `json:"gpu,omitempty" tf:"gpu"`
	Image *string `json:"image" tf:"image"`
	// +optional
	ImagePullPolicy *string `json:"imagePullPolicy,omitempty" tf:"image_pull_policy"`
	// +optional
	Memory *float64 `json:"memory,omitempty" tf:"memory"`
	Name   *string  `json:"name" tf:"name"`
	// +optional
	Ports []ContainerGroupSpecContainersPorts `json:"ports,omitempty" tf:"ports"`
	// +optional
	Ready *bool `json:"ready,omitempty" tf:"ready"`
	// +optional
	RestartCount *int64 `json:"restartCount,omitempty" tf:"restart_count"`
	// +optional
	VolumeMounts []ContainerGroupSpecContainersVolumeMounts `json:"volumeMounts,omitempty" tf:"volume_mounts"`
	// +optional
	WorkingDir *string `json:"workingDir,omitempty" tf:"working_dir"`
}

type ContainerGroupSpecDnsConfigOptions struct {
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type ContainerGroupSpecDnsConfig struct {
	// +optional
	NameServers []string `json:"nameServers,omitempty" tf:"name_servers"`
	// +optional
	Options []ContainerGroupSpecDnsConfigOptions `json:"options,omitempty" tf:"options"`
	// +optional
	Searches []string `json:"searches,omitempty" tf:"searches"`
}

type ContainerGroupSpecEciSecurityContextSysctls struct {
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type ContainerGroupSpecEciSecurityContext struct {
	// +optional
	Sysctls []ContainerGroupSpecEciSecurityContextSysctls `json:"sysctls,omitempty" tf:"sysctls"`
}

type ContainerGroupSpecHostAliases struct {
	// +optional
	Hostnames []string `json:"hostnames,omitempty" tf:"hostnames"`
	// +optional
	Ip *string `json:"ip,omitempty" tf:"ip"`
}

type ContainerGroupSpecInitContainersEnvironmentVars struct {
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type ContainerGroupSpecInitContainersPorts struct {
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// +optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`
}

type ContainerGroupSpecInitContainersVolumeMounts struct {
	// +optional
	MountPath *string `json:"mountPath,omitempty" tf:"mount_path"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	ReadOnly *bool `json:"readOnly,omitempty" tf:"read_only"`
}

type ContainerGroupSpecInitContainers struct {
	// +optional
	Args []string `json:"args,omitempty" tf:"args"`
	// +optional
	Commands []string `json:"commands,omitempty" tf:"commands"`
	// +optional
	Cpu *float64 `json:"cpu,omitempty" tf:"cpu"`
	// +optional
	EnvironmentVars []ContainerGroupSpecInitContainersEnvironmentVars `json:"environmentVars,omitempty" tf:"environment_vars"`
	// +optional
	Gpu *int64 `json:"gpu,omitempty" tf:"gpu"`
	// +optional
	Image *string `json:"image,omitempty" tf:"image"`
	// +optional
	ImagePullPolicy *string `json:"imagePullPolicy,omitempty" tf:"image_pull_policy"`
	// +optional
	Memory *float64 `json:"memory,omitempty" tf:"memory"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Ports []ContainerGroupSpecInitContainersPorts `json:"ports,omitempty" tf:"ports"`
	// +optional
	Ready *bool `json:"ready,omitempty" tf:"ready"`
	// +optional
	RestartCount *int64 `json:"restartCount,omitempty" tf:"restart_count"`
	// +optional
	VolumeMounts []ContainerGroupSpecInitContainersVolumeMounts `json:"volumeMounts,omitempty" tf:"volume_mounts"`
	// +optional
	WorkingDir *string `json:"workingDir,omitempty" tf:"working_dir"`
}

type ContainerGroupSpecVolumesConfigFileVolumeConfigFileToPaths struct {
	// +optional
	Content *string `json:"content,omitempty" tf:"content"`
	// +optional
	Path *string `json:"path,omitempty" tf:"path"`
}

type ContainerGroupSpecVolumes struct {
	// +optional
	ConfigFileVolumeConfigFileToPaths []ContainerGroupSpecVolumesConfigFileVolumeConfigFileToPaths `json:"configFileVolumeConfigFileToPaths,omitempty" tf:"config_file_volume_config_file_to_paths"`
	// +optional
	DiskVolumeDiskID *string `json:"diskVolumeDiskID,omitempty" tf:"disk_volume_disk_id"`
	// +optional
	DiskVolumeFsType *string `json:"diskVolumeFsType,omitempty" tf:"disk_volume_fs_type"`
	// +optional
	FlexVolumeDriver *string `json:"flexVolumeDriver,omitempty" tf:"flex_volume_driver"`
	// +optional
	FlexVolumeFsType *string `json:"flexVolumeFsType,omitempty" tf:"flex_volume_fs_type"`
	// +optional
	FlexVolumeOptions *string `json:"flexVolumeOptions,omitempty" tf:"flex_volume_options"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	NfsVolumePath *string `json:"nfsVolumePath,omitempty" tf:"nfs_volume_path"`
	// +optional
	NfsVolumeReadOnly *bool `json:"nfsVolumeReadOnly,omitempty" tf:"nfs_volume_read_only"`
	// +optional
	NfsVolumeServer *string `json:"nfsVolumeServer,omitempty" tf:"nfs_volume_server"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type ContainerGroupSpec struct {
	State *ContainerGroupSpecResource `json:"state,omitempty" tf:"-"`

	Resource ContainerGroupSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ContainerGroupSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ContainerGroupName *string                        `json:"containerGroupName" tf:"container_group_name"`
	Containers         []ContainerGroupSpecContainers `json:"containers" tf:"containers"`
	// +optional
	Cpu *float64 `json:"cpu,omitempty" tf:"cpu"`
	// +optional
	DnsConfig *ContainerGroupSpecDnsConfig `json:"dnsConfig,omitempty" tf:"dns_config"`
	// +optional
	EciSecurityContext *ContainerGroupSpecEciSecurityContext `json:"eciSecurityContext,omitempty" tf:"eci_security_context"`
	// +optional
	HostAliases []ContainerGroupSpecHostAliases `json:"hostAliases,omitempty" tf:"host_aliases"`
	// +optional
	InitContainers []ContainerGroupSpecInitContainers `json:"initContainers,omitempty" tf:"init_containers"`
	// +optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type"`
	// +optional
	Memory *float64 `json:"memory,omitempty" tf:"memory"`
	// +optional
	RamRoleName *string `json:"ramRoleName,omitempty" tf:"ram_role_name"`
	// +optional
	ResourceGroupID *string `json:"resourceGroupID,omitempty" tf:"resource_group_id"`
	// +optional
	RestartPolicy   *string `json:"restartPolicy,omitempty" tf:"restart_policy"`
	SecurityGroupID *string `json:"securityGroupID" tf:"security_group_id"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	Volumes   []ContainerGroupSpecVolumes `json:"volumes,omitempty" tf:"volumes"`
	VswitchID *string                     `json:"vswitchID" tf:"vswitch_id"`
	// +optional
	ZoneID *string `json:"zoneID,omitempty" tf:"zone_id"`
}

type ContainerGroupStatus struct {
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

// ContainerGroupList is a list of ContainerGroups
type ContainerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ContainerGroup CRD objects
	Items []ContainerGroup `json:"items,omitempty"`
}
