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

type ConnectVirtualBorderRouter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConnectVirtualBorderRouterSpec   `json:"spec,omitempty"`
	Status            ConnectVirtualBorderRouterStatus `json:"status,omitempty"`
}

type ConnectVirtualBorderRouterSpec struct {
	State *ConnectVirtualBorderRouterSpecResource `json:"state,omitempty" tf:"-"`

	Resource ConnectVirtualBorderRouterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ConnectVirtualBorderRouterSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AssociatedPhysicalConnections *string `json:"associatedPhysicalConnections,omitempty" tf:"associated_physical_connections"`
	// +optional
	Bandwidth *int64 `json:"bandwidth,omitempty" tf:"bandwidth"`
	// +optional
	CircuitCode *string `json:"circuitCode,omitempty" tf:"circuit_code"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DetectMultiplier *int64 `json:"detectMultiplier,omitempty" tf:"detect_multiplier"`
	// +optional
	EnableIpv6     *bool   `json:"enableIpv6,omitempty" tf:"enable_ipv6"`
	LocalGatewayIP *string `json:"localGatewayIP" tf:"local_gateway_ip"`
	// +optional
	LocalIpv6GatewayIP *string `json:"localIpv6GatewayIP,omitempty" tf:"local_ipv6_gateway_ip"`
	// +optional
	MinRxInterval *int64 `json:"minRxInterval,omitempty" tf:"min_rx_interval"`
	// +optional
	MinTxInterval *int64  `json:"minTxInterval,omitempty" tf:"min_tx_interval"`
	PeerGatewayIP *string `json:"peerGatewayIP" tf:"peer_gateway_ip"`
	// +optional
	PeerIpv6GatewayIP *string `json:"peerIpv6GatewayIP,omitempty" tf:"peer_ipv6_gateway_ip"`
	// +optional
	PeeringIpv6SubnetMask *string `json:"peeringIpv6SubnetMask,omitempty" tf:"peering_ipv6_subnet_mask"`
	PeeringSubnetMask     *string `json:"peeringSubnetMask" tf:"peering_subnet_mask"`
	PhysicalConnectionID  *string `json:"physicalConnectionID" tf:"physical_connection_id"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	VbrOwnerID *string `json:"vbrOwnerID,omitempty" tf:"vbr_owner_id"`
	// +optional
	VirtualBorderRouterName *string `json:"virtualBorderRouterName,omitempty" tf:"virtual_border_router_name"`
	VlanID                  *int64  `json:"vlanID" tf:"vlan_id"`
}

type ConnectVirtualBorderRouterStatus struct {
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

// ConnectVirtualBorderRouterList is a list of ConnectVirtualBorderRouters
type ConnectVirtualBorderRouterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ConnectVirtualBorderRouter CRD objects
	Items []ConnectVirtualBorderRouter `json:"items,omitempty"`
}
