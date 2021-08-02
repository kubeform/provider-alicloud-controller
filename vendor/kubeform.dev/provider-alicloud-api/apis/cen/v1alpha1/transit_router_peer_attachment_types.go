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

type TransitRouterPeerAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitRouterPeerAttachmentSpec   `json:"spec,omitempty"`
	Status            TransitRouterPeerAttachmentStatus `json:"status,omitempty"`
}

type TransitRouterPeerAttachmentSpec struct {
	State *TransitRouterPeerAttachmentSpecResource `json:"state,omitempty" tf:"-"`

	Resource TransitRouterPeerAttachmentSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type TransitRouterPeerAttachmentSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AutoPublishRouteEnabled *bool `json:"autoPublishRouteEnabled,omitempty" tf:"auto_publish_route_enabled"`
	// +optional
	Bandwidth *int64 `json:"bandwidth,omitempty" tf:"bandwidth"`
	// +optional
	CenBandwidthPackageID *string `json:"cenBandwidthPackageID,omitempty" tf:"cen_bandwidth_package_id"`
	CenID                 *string `json:"cenID" tf:"cen_id"`
	// +optional
	DryRun                    *bool   `json:"dryRun,omitempty" tf:"dry_run"`
	PeerTransitRouterID       *string `json:"peerTransitRouterID" tf:"peer_transit_router_id"`
	PeerTransitRouterRegionID *string `json:"peerTransitRouterRegionID" tf:"peer_transit_router_region_id"`
	// +optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type"`
	// +optional
	RouteTableAssociationEnabled *bool `json:"routeTableAssociationEnabled,omitempty" tf:"route_table_association_enabled"`
	// +optional
	RouteTablePropagationEnabled *bool `json:"routeTablePropagationEnabled,omitempty" tf:"route_table_propagation_enabled"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	TransitRouterAttachmentDescription *string `json:"transitRouterAttachmentDescription,omitempty" tf:"transit_router_attachment_description"`
	// +optional
	TransitRouterAttachmentID *string `json:"transitRouterAttachmentID,omitempty" tf:"transit_router_attachment_id"`
	// +optional
	TransitRouterAttachmentName *string `json:"transitRouterAttachmentName,omitempty" tf:"transit_router_attachment_name"`
	// +optional
	TransitRouterID *string `json:"transitRouterID,omitempty" tf:"transit_router_id"`
}

type TransitRouterPeerAttachmentStatus struct {
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

// TransitRouterPeerAttachmentList is a list of TransitRouterPeerAttachments
type TransitRouterPeerAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of TransitRouterPeerAttachment CRD objects
	Items []TransitRouterPeerAttachment `json:"items,omitempty"`
}
