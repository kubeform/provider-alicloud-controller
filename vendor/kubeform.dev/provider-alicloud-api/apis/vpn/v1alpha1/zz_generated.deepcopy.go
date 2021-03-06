//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Connection) DeepCopyInto(out *Connection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Connection.
func (in *Connection) DeepCopy() *Connection {
	if in == nil {
		return nil
	}
	out := new(Connection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Connection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionList) DeepCopyInto(out *ConnectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Connection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionList.
func (in *ConnectionList) DeepCopy() *ConnectionList {
	if in == nil {
		return nil
	}
	out := new(ConnectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConnectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionSpec) DeepCopyInto(out *ConnectionSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ConnectionSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionSpec.
func (in *ConnectionSpec) DeepCopy() *ConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(ConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionSpecIkeConfig) DeepCopyInto(out *ConnectionSpecIkeConfig) {
	*out = *in
	if in.IkeAuthAlg != nil {
		in, out := &in.IkeAuthAlg, &out.IkeAuthAlg
		*out = new(string)
		**out = **in
	}
	if in.IkeEncAlg != nil {
		in, out := &in.IkeEncAlg, &out.IkeEncAlg
		*out = new(string)
		**out = **in
	}
	if in.IkeLifetime != nil {
		in, out := &in.IkeLifetime, &out.IkeLifetime
		*out = new(int64)
		**out = **in
	}
	if in.IkeLocalID != nil {
		in, out := &in.IkeLocalID, &out.IkeLocalID
		*out = new(string)
		**out = **in
	}
	if in.IkeMode != nil {
		in, out := &in.IkeMode, &out.IkeMode
		*out = new(string)
		**out = **in
	}
	if in.IkePfs != nil {
		in, out := &in.IkePfs, &out.IkePfs
		*out = new(string)
		**out = **in
	}
	if in.IkeRemoteID != nil {
		in, out := &in.IkeRemoteID, &out.IkeRemoteID
		*out = new(string)
		**out = **in
	}
	if in.IkeVersion != nil {
		in, out := &in.IkeVersion, &out.IkeVersion
		*out = new(string)
		**out = **in
	}
	if in.Psk != nil {
		in, out := &in.Psk, &out.Psk
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionSpecIkeConfig.
func (in *ConnectionSpecIkeConfig) DeepCopy() *ConnectionSpecIkeConfig {
	if in == nil {
		return nil
	}
	out := new(ConnectionSpecIkeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionSpecIpsecConfig) DeepCopyInto(out *ConnectionSpecIpsecConfig) {
	*out = *in
	if in.IpsecAuthAlg != nil {
		in, out := &in.IpsecAuthAlg, &out.IpsecAuthAlg
		*out = new(string)
		**out = **in
	}
	if in.IpsecEncAlg != nil {
		in, out := &in.IpsecEncAlg, &out.IpsecEncAlg
		*out = new(string)
		**out = **in
	}
	if in.IpsecLifetime != nil {
		in, out := &in.IpsecLifetime, &out.IpsecLifetime
		*out = new(int64)
		**out = **in
	}
	if in.IpsecPfs != nil {
		in, out := &in.IpsecPfs, &out.IpsecPfs
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionSpecIpsecConfig.
func (in *ConnectionSpecIpsecConfig) DeepCopy() *ConnectionSpecIpsecConfig {
	if in == nil {
		return nil
	}
	out := new(ConnectionSpecIpsecConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionSpecResource) DeepCopyInto(out *ConnectionSpecResource) {
	*out = *in
	if in.CustomerGatewayID != nil {
		in, out := &in.CustomerGatewayID, &out.CustomerGatewayID
		*out = new(string)
		**out = **in
	}
	if in.EffectImmediately != nil {
		in, out := &in.EffectImmediately, &out.EffectImmediately
		*out = new(bool)
		**out = **in
	}
	if in.IkeConfig != nil {
		in, out := &in.IkeConfig, &out.IkeConfig
		*out = make([]ConnectionSpecIkeConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IpsecConfig != nil {
		in, out := &in.IpsecConfig, &out.IpsecConfig
		*out = make([]ConnectionSpecIpsecConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LocalSubnet != nil {
		in, out := &in.LocalSubnet, &out.LocalSubnet
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RemoteSubnet != nil {
		in, out := &in.RemoteSubnet, &out.RemoteSubnet
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.VpnGatewayID != nil {
		in, out := &in.VpnGatewayID, &out.VpnGatewayID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionSpecResource.
func (in *ConnectionSpecResource) DeepCopy() *ConnectionSpecResource {
	if in == nil {
		return nil
	}
	out := new(ConnectionSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionStatus) DeepCopyInto(out *ConnectionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionStatus.
func (in *ConnectionStatus) DeepCopy() *ConnectionStatus {
	if in == nil {
		return nil
	}
	out := new(ConnectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomerGateway) DeepCopyInto(out *CustomerGateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomerGateway.
func (in *CustomerGateway) DeepCopy() *CustomerGateway {
	if in == nil {
		return nil
	}
	out := new(CustomerGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomerGateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomerGatewayList) DeepCopyInto(out *CustomerGatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CustomerGateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomerGatewayList.
func (in *CustomerGatewayList) DeepCopy() *CustomerGatewayList {
	if in == nil {
		return nil
	}
	out := new(CustomerGatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomerGatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomerGatewaySpec) DeepCopyInto(out *CustomerGatewaySpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(CustomerGatewaySpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomerGatewaySpec.
func (in *CustomerGatewaySpec) DeepCopy() *CustomerGatewaySpec {
	if in == nil {
		return nil
	}
	out := new(CustomerGatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomerGatewaySpecResource) DeepCopyInto(out *CustomerGatewaySpecResource) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.IpAddress != nil {
		in, out := &in.IpAddress, &out.IpAddress
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomerGatewaySpecResource.
func (in *CustomerGatewaySpecResource) DeepCopy() *CustomerGatewaySpecResource {
	if in == nil {
		return nil
	}
	out := new(CustomerGatewaySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomerGatewayStatus) DeepCopyInto(out *CustomerGatewayStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomerGatewayStatus.
func (in *CustomerGatewayStatus) DeepCopy() *CustomerGatewayStatus {
	if in == nil {
		return nil
	}
	out := new(CustomerGatewayStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Gateway) DeepCopyInto(out *Gateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Gateway.
func (in *Gateway) DeepCopy() *Gateway {
	if in == nil {
		return nil
	}
	out := new(Gateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Gateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayList) DeepCopyInto(out *GatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Gateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayList.
func (in *GatewayList) DeepCopy() *GatewayList {
	if in == nil {
		return nil
	}
	out := new(GatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewaySpec) DeepCopyInto(out *GatewaySpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(GatewaySpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewaySpec.
func (in *GatewaySpec) DeepCopy() *GatewaySpec {
	if in == nil {
		return nil
	}
	out := new(GatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewaySpecResource) DeepCopyInto(out *GatewaySpecResource) {
	*out = *in
	if in.Bandwidth != nil {
		in, out := &in.Bandwidth, &out.Bandwidth
		*out = new(int64)
		**out = **in
	}
	if in.BusinessStatus != nil {
		in, out := &in.BusinessStatus, &out.BusinessStatus
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EnableIpsec != nil {
		in, out := &in.EnableIpsec, &out.EnableIpsec
		*out = new(bool)
		**out = **in
	}
	if in.EnableSsl != nil {
		in, out := &in.EnableSsl, &out.EnableSsl
		*out = new(bool)
		**out = **in
	}
	if in.InstanceChargeType != nil {
		in, out := &in.InstanceChargeType, &out.InstanceChargeType
		*out = new(string)
		**out = **in
	}
	if in.InternetIP != nil {
		in, out := &in.InternetIP, &out.InternetIP
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(int64)
		**out = **in
	}
	if in.SslConnections != nil {
		in, out := &in.SslConnections, &out.SslConnections
		*out = new(int64)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.VpcID != nil {
		in, out := &in.VpcID, &out.VpcID
		*out = new(string)
		**out = **in
	}
	if in.VswitchID != nil {
		in, out := &in.VswitchID, &out.VswitchID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewaySpecResource.
func (in *GatewaySpecResource) DeepCopy() *GatewaySpecResource {
	if in == nil {
		return nil
	}
	out := new(GatewaySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayStatus) DeepCopyInto(out *GatewayStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayStatus.
func (in *GatewayStatus) DeepCopy() *GatewayStatus {
	if in == nil {
		return nil
	}
	out := new(GatewayStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteEntry) DeepCopyInto(out *RouteEntry) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteEntry.
func (in *RouteEntry) DeepCopy() *RouteEntry {
	if in == nil {
		return nil
	}
	out := new(RouteEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RouteEntry) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteEntryList) DeepCopyInto(out *RouteEntryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RouteEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteEntryList.
func (in *RouteEntryList) DeepCopy() *RouteEntryList {
	if in == nil {
		return nil
	}
	out := new(RouteEntryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RouteEntryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteEntrySpec) DeepCopyInto(out *RouteEntrySpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(RouteEntrySpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteEntrySpec.
func (in *RouteEntrySpec) DeepCopy() *RouteEntrySpec {
	if in == nil {
		return nil
	}
	out := new(RouteEntrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteEntrySpecResource) DeepCopyInto(out *RouteEntrySpecResource) {
	*out = *in
	if in.NextHop != nil {
		in, out := &in.NextHop, &out.NextHop
		*out = new(string)
		**out = **in
	}
	if in.PublishVpc != nil {
		in, out := &in.PublishVpc, &out.PublishVpc
		*out = new(bool)
		**out = **in
	}
	if in.RouteDest != nil {
		in, out := &in.RouteDest, &out.RouteDest
		*out = new(string)
		**out = **in
	}
	if in.VpnGatewayID != nil {
		in, out := &in.VpnGatewayID, &out.VpnGatewayID
		*out = new(string)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteEntrySpecResource.
func (in *RouteEntrySpecResource) DeepCopy() *RouteEntrySpecResource {
	if in == nil {
		return nil
	}
	out := new(RouteEntrySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteEntryStatus) DeepCopyInto(out *RouteEntryStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteEntryStatus.
func (in *RouteEntryStatus) DeepCopy() *RouteEntryStatus {
	if in == nil {
		return nil
	}
	out := new(RouteEntryStatus)
	in.DeepCopyInto(out)
	return out
}
