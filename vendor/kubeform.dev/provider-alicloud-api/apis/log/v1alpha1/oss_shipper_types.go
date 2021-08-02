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

type OssShipper struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OssShipperSpec   `json:"spec,omitempty"`
	Status            OssShipperStatus `json:"status,omitempty"`
}

type OssShipperSpecParquetConfig struct {
	Name *string `json:"name" tf:"name"`
	Type *string `json:"type" tf:"type"`
}

type OssShipperSpec struct {
	State *OssShipperSpecResource `json:"state,omitempty" tf:"-"`

	Resource OssShipperSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type OssShipperSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	BufferInterval *int64 `json:"bufferInterval" tf:"buffer_interval"`
	BufferSize     *int64 `json:"bufferSize" tf:"buffer_size"`
	// +optional
	CompressType *string `json:"compressType,omitempty" tf:"compress_type"`
	// +optional
	CsvConfigColumns []string `json:"csvConfigColumns,omitempty" tf:"csv_config_columns"`
	// +optional
	CsvConfigDelimiter *string `json:"csvConfigDelimiter,omitempty" tf:"csv_config_delimiter"`
	// +optional
	CsvConfigHeader *bool `json:"csvConfigHeader,omitempty" tf:"csv_config_header"`
	// +optional
	CsvConfigLinefeed *string `json:"csvConfigLinefeed,omitempty" tf:"csv_config_linefeed"`
	// +optional
	CsvConfigNullidentifier *string `json:"csvConfigNullidentifier,omitempty" tf:"csv_config_nullidentifier"`
	// +optional
	CsvConfigQuote *string `json:"csvConfigQuote,omitempty" tf:"csv_config_quote"`
	Format         *string `json:"format" tf:"format"`
	// +optional
	JsonEnableTag *bool   `json:"jsonEnableTag,omitempty" tf:"json_enable_tag"`
	LogstoreName  *string `json:"logstoreName" tf:"logstore_name"`
	OssBucket     *string `json:"ossBucket" tf:"oss_bucket"`
	// +optional
	OssPrefix *string `json:"ossPrefix,omitempty" tf:"oss_prefix"`
	// +optional
	ParquetConfig []OssShipperSpecParquetConfig `json:"parquetConfig,omitempty" tf:"parquet_config"`
	PathFormat    *string                       `json:"pathFormat" tf:"path_format"`
	ProjectName   *string                       `json:"projectName" tf:"project_name"`
	// +optional
	RoleArn     *string `json:"roleArn,omitempty" tf:"role_arn"`
	ShipperName *string `json:"shipperName" tf:"shipper_name"`
}

type OssShipperStatus struct {
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

// OssShipperList is a list of OssShippers
type OssShipperList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OssShipper CRD objects
	Items []OssShipper `json:"items,omitempty"`
}
