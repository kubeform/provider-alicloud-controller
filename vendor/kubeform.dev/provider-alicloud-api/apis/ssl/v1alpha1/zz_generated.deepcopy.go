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
func (in *CertificatesServiceCertificate) DeepCopyInto(out *CertificatesServiceCertificate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesServiceCertificate.
func (in *CertificatesServiceCertificate) DeepCopy() *CertificatesServiceCertificate {
	if in == nil {
		return nil
	}
	out := new(CertificatesServiceCertificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificatesServiceCertificate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesServiceCertificateList) DeepCopyInto(out *CertificatesServiceCertificateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CertificatesServiceCertificate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesServiceCertificateList.
func (in *CertificatesServiceCertificateList) DeepCopy() *CertificatesServiceCertificateList {
	if in == nil {
		return nil
	}
	out := new(CertificatesServiceCertificateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificatesServiceCertificateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesServiceCertificateSpec) DeepCopyInto(out *CertificatesServiceCertificateSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(CertificatesServiceCertificateSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesServiceCertificateSpec.
func (in *CertificatesServiceCertificateSpec) DeepCopy() *CertificatesServiceCertificateSpec {
	if in == nil {
		return nil
	}
	out := new(CertificatesServiceCertificateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesServiceCertificateSpecResource) DeepCopyInto(out *CertificatesServiceCertificateSpecResource) {
	*out = *in
	if in.Cert != nil {
		in, out := &in.Cert, &out.Cert
		*out = new(string)
		**out = **in
	}
	if in.CertificateName != nil {
		in, out := &in.CertificateName, &out.CertificateName
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Lang != nil {
		in, out := &in.Lang, &out.Lang
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesServiceCertificateSpecResource.
func (in *CertificatesServiceCertificateSpecResource) DeepCopy() *CertificatesServiceCertificateSpecResource {
	if in == nil {
		return nil
	}
	out := new(CertificatesServiceCertificateSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesServiceCertificateStatus) DeepCopyInto(out *CertificatesServiceCertificateStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesServiceCertificateStatus.
func (in *CertificatesServiceCertificateStatus) DeepCopy() *CertificatesServiceCertificateStatus {
	if in == nil {
		return nil
	}
	out := new(CertificatesServiceCertificateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnClientCert) DeepCopyInto(out *VpnClientCert) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnClientCert.
func (in *VpnClientCert) DeepCopy() *VpnClientCert {
	if in == nil {
		return nil
	}
	out := new(VpnClientCert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpnClientCert) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnClientCertList) DeepCopyInto(out *VpnClientCertList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpnClientCert, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnClientCertList.
func (in *VpnClientCertList) DeepCopy() *VpnClientCertList {
	if in == nil {
		return nil
	}
	out := new(VpnClientCertList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpnClientCertList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnClientCertSpec) DeepCopyInto(out *VpnClientCertSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(VpnClientCertSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnClientCertSpec.
func (in *VpnClientCertSpec) DeepCopy() *VpnClientCertSpec {
	if in == nil {
		return nil
	}
	out := new(VpnClientCertSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnClientCertSpecResource) DeepCopyInto(out *VpnClientCertSpecResource) {
	*out = *in
	if in.CaCert != nil {
		in, out := &in.CaCert, &out.CaCert
		*out = new(string)
		**out = **in
	}
	if in.ClientCert != nil {
		in, out := &in.ClientCert, &out.ClientCert
		*out = new(string)
		**out = **in
	}
	if in.ClientConfig != nil {
		in, out := &in.ClientConfig, &out.ClientConfig
		*out = new(string)
		**out = **in
	}
	if in.ClientKey != nil {
		in, out := &in.ClientKey, &out.ClientKey
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.SslVPNServerID != nil {
		in, out := &in.SslVPNServerID, &out.SslVPNServerID
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnClientCertSpecResource.
func (in *VpnClientCertSpecResource) DeepCopy() *VpnClientCertSpecResource {
	if in == nil {
		return nil
	}
	out := new(VpnClientCertSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnClientCertStatus) DeepCopyInto(out *VpnClientCertStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnClientCertStatus.
func (in *VpnClientCertStatus) DeepCopy() *VpnClientCertStatus {
	if in == nil {
		return nil
	}
	out := new(VpnClientCertStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnServer) DeepCopyInto(out *VpnServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnServer.
func (in *VpnServer) DeepCopy() *VpnServer {
	if in == nil {
		return nil
	}
	out := new(VpnServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpnServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnServerList) DeepCopyInto(out *VpnServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpnServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnServerList.
func (in *VpnServerList) DeepCopy() *VpnServerList {
	if in == nil {
		return nil
	}
	out := new(VpnServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpnServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnServerSpec) DeepCopyInto(out *VpnServerSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(VpnServerSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnServerSpec.
func (in *VpnServerSpec) DeepCopy() *VpnServerSpec {
	if in == nil {
		return nil
	}
	out := new(VpnServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnServerSpecResource) DeepCopyInto(out *VpnServerSpecResource) {
	*out = *in
	if in.Cipher != nil {
		in, out := &in.Cipher, &out.Cipher
		*out = new(string)
		**out = **in
	}
	if in.ClientIPPool != nil {
		in, out := &in.ClientIPPool, &out.ClientIPPool
		*out = new(string)
		**out = **in
	}
	if in.Compress != nil {
		in, out := &in.Compress, &out.Compress
		*out = new(bool)
		**out = **in
	}
	if in.Connections != nil {
		in, out := &in.Connections, &out.Connections
		*out = new(int64)
		**out = **in
	}
	if in.InternetIP != nil {
		in, out := &in.InternetIP, &out.InternetIP
		*out = new(string)
		**out = **in
	}
	if in.LocalSubnet != nil {
		in, out := &in.LocalSubnet, &out.LocalSubnet
		*out = new(string)
		**out = **in
	}
	if in.MaxConnections != nil {
		in, out := &in.MaxConnections, &out.MaxConnections
		*out = new(int64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnServerSpecResource.
func (in *VpnServerSpecResource) DeepCopy() *VpnServerSpecResource {
	if in == nil {
		return nil
	}
	out := new(VpnServerSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnServerStatus) DeepCopyInto(out *VpnServerStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnServerStatus.
func (in *VpnServerStatus) DeepCopy() *VpnServerStatus {
	if in == nil {
		return nil
	}
	out := new(VpnServerStatus)
	in.DeepCopyInto(out)
	return out
}
