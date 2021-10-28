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
	apiv1alpha1 "kubeform.dev/apimachinery/api/v1alpha1"

	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

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
func (in *GatewaySpecBandwidthPackages) DeepCopyInto(out *GatewaySpecBandwidthPackages) {
	*out = *in
	if in.Bandwidth != nil {
		in, out := &in.Bandwidth, &out.Bandwidth
		*out = new(int64)
		**out = **in
	}
	if in.IpCount != nil {
		in, out := &in.IpCount, &out.IpCount
		*out = new(int64)
		**out = **in
	}
	if in.PublicIPAddresses != nil {
		in, out := &in.PublicIPAddresses, &out.PublicIPAddresses
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewaySpecBandwidthPackages.
func (in *GatewaySpecBandwidthPackages) DeepCopy() *GatewaySpecBandwidthPackages {
	if in == nil {
		return nil
	}
	out := new(GatewaySpecBandwidthPackages)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewaySpecResource) DeepCopyInto(out *GatewaySpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.BandwidthPackageIDS != nil {
		in, out := &in.BandwidthPackageIDS, &out.BandwidthPackageIDS
		*out = new(string)
		**out = **in
	}
	if in.BandwidthPackages != nil {
		in, out := &in.BandwidthPackages, &out.BandwidthPackages
		*out = make([]GatewaySpecBandwidthPackages, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeletionProtection != nil {
		in, out := &in.DeletionProtection, &out.DeletionProtection
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DryRun != nil {
		in, out := &in.DryRun, &out.DryRun
		*out = new(bool)
		**out = **in
	}
	if in.Force != nil {
		in, out := &in.Force, &out.Force
		*out = new(bool)
		**out = **in
	}
	if in.ForwardTableIDS != nil {
		in, out := &in.ForwardTableIDS, &out.ForwardTableIDS
		*out = new(string)
		**out = **in
	}
	if in.InstanceChargeType != nil {
		in, out := &in.InstanceChargeType, &out.InstanceChargeType
		*out = new(string)
		**out = **in
	}
	if in.InternetChargeType != nil {
		in, out := &in.InternetChargeType, &out.InternetChargeType
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NatGatewayName != nil {
		in, out := &in.NatGatewayName, &out.NatGatewayName
		*out = new(string)
		**out = **in
	}
	if in.NatType != nil {
		in, out := &in.NatType, &out.NatType
		*out = new(string)
		**out = **in
	}
	if in.PaymentType != nil {
		in, out := &in.PaymentType, &out.PaymentType
		*out = new(string)
		**out = **in
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(int64)
		**out = **in
	}
	if in.SnatTableIDS != nil {
		in, out := &in.SnatTableIDS, &out.SnatTableIDS
		*out = new(string)
		**out = **in
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(string)
		**out = **in
	}
	if in.Specification != nil {
		in, out := &in.Specification, &out.Specification
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
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
