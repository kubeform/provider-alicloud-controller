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
func (in *Domain) DeepCopyInto(out *Domain) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Domain.
func (in *Domain) DeepCopy() *Domain {
	if in == nil {
		return nil
	}
	out := new(Domain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Domain) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainConfig) DeepCopyInto(out *DomainConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainConfig.
func (in *DomainConfig) DeepCopy() *DomainConfig {
	if in == nil {
		return nil
	}
	out := new(DomainConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainConfigList) DeepCopyInto(out *DomainConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DomainConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainConfigList.
func (in *DomainConfigList) DeepCopy() *DomainConfigList {
	if in == nil {
		return nil
	}
	out := new(DomainConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainConfigSpec) DeepCopyInto(out *DomainConfigSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(DomainConfigSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainConfigSpec.
func (in *DomainConfigSpec) DeepCopy() *DomainConfigSpec {
	if in == nil {
		return nil
	}
	out := new(DomainConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainConfigSpecFunctionArgs) DeepCopyInto(out *DomainConfigSpecFunctionArgs) {
	*out = *in
	if in.ArgName != nil {
		in, out := &in.ArgName, &out.ArgName
		*out = new(string)
		**out = **in
	}
	if in.ArgValue != nil {
		in, out := &in.ArgValue, &out.ArgValue
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainConfigSpecFunctionArgs.
func (in *DomainConfigSpecFunctionArgs) DeepCopy() *DomainConfigSpecFunctionArgs {
	if in == nil {
		return nil
	}
	out := new(DomainConfigSpecFunctionArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainConfigSpecResource) DeepCopyInto(out *DomainConfigSpecResource) {
	*out = *in
	if in.ConfigID != nil {
		in, out := &in.ConfigID, &out.ConfigID
		*out = new(string)
		**out = **in
	}
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.FunctionArgs != nil {
		in, out := &in.FunctionArgs, &out.FunctionArgs
		*out = make([]DomainConfigSpecFunctionArgs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FunctionName != nil {
		in, out := &in.FunctionName, &out.FunctionName
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainConfigSpecResource.
func (in *DomainConfigSpecResource) DeepCopy() *DomainConfigSpecResource {
	if in == nil {
		return nil
	}
	out := new(DomainConfigSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainConfigStatus) DeepCopyInto(out *DomainConfigStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainConfigStatus.
func (in *DomainConfigStatus) DeepCopy() *DomainConfigStatus {
	if in == nil {
		return nil
	}
	out := new(DomainConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainList) DeepCopyInto(out *DomainList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Domain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainList.
func (in *DomainList) DeepCopy() *DomainList {
	if in == nil {
		return nil
	}
	out := new(DomainList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainNew) DeepCopyInto(out *DomainNew) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainNew.
func (in *DomainNew) DeepCopy() *DomainNew {
	if in == nil {
		return nil
	}
	out := new(DomainNew)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainNew) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainNewList) DeepCopyInto(out *DomainNewList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DomainNew, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainNewList.
func (in *DomainNewList) DeepCopy() *DomainNewList {
	if in == nil {
		return nil
	}
	out := new(DomainNewList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainNewList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainNewSpec) DeepCopyInto(out *DomainNewSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(DomainNewSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainNewSpec.
func (in *DomainNewSpec) DeepCopy() *DomainNewSpec {
	if in == nil {
		return nil
	}
	out := new(DomainNewSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainNewSpecCertificateConfig) DeepCopyInto(out *DomainNewSpecCertificateConfig) {
	*out = *in
	if in.CertName != nil {
		in, out := &in.CertName, &out.CertName
		*out = new(string)
		**out = **in
	}
	if in.CertType != nil {
		in, out := &in.CertType, &out.CertType
		*out = new(string)
		**out = **in
	}
	if in.ForceSet != nil {
		in, out := &in.ForceSet, &out.ForceSet
		*out = new(string)
		**out = **in
	}
	if in.PrivateKey != nil {
		in, out := &in.PrivateKey, &out.PrivateKey
		*out = new(string)
		**out = **in
	}
	if in.ServerCertificate != nil {
		in, out := &in.ServerCertificate, &out.ServerCertificate
		*out = new(string)
		**out = **in
	}
	if in.ServerCertificateStatus != nil {
		in, out := &in.ServerCertificateStatus, &out.ServerCertificateStatus
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainNewSpecCertificateConfig.
func (in *DomainNewSpecCertificateConfig) DeepCopy() *DomainNewSpecCertificateConfig {
	if in == nil {
		return nil
	}
	out := new(DomainNewSpecCertificateConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainNewSpecResource) DeepCopyInto(out *DomainNewSpecResource) {
	*out = *in
	if in.CdnType != nil {
		in, out := &in.CdnType, &out.CdnType
		*out = new(string)
		**out = **in
	}
	if in.CertificateConfig != nil {
		in, out := &in.CertificateConfig, &out.CertificateConfig
		*out = new(DomainNewSpecCertificateConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Cname != nil {
		in, out := &in.Cname, &out.Cname
		*out = new(string)
		**out = **in
	}
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupID != nil {
		in, out := &in.ResourceGroupID, &out.ResourceGroupID
		*out = new(string)
		**out = **in
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]DomainNewSpecSources, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainNewSpecResource.
func (in *DomainNewSpecResource) DeepCopy() *DomainNewSpecResource {
	if in == nil {
		return nil
	}
	out := new(DomainNewSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainNewSpecSources) DeepCopyInto(out *DomainNewSpecSources) {
	*out = *in
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(int64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainNewSpecSources.
func (in *DomainNewSpecSources) DeepCopy() *DomainNewSpecSources {
	if in == nil {
		return nil
	}
	out := new(DomainNewSpecSources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainNewStatus) DeepCopyInto(out *DomainNewStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainNewStatus.
func (in *DomainNewStatus) DeepCopy() *DomainNewStatus {
	if in == nil {
		return nil
	}
	out := new(DomainNewStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSpec) DeepCopyInto(out *DomainSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(DomainSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSpec.
func (in *DomainSpec) DeepCopy() *DomainSpec {
	if in == nil {
		return nil
	}
	out := new(DomainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSpecAuthConfig) DeepCopyInto(out *DomainSpecAuthConfig) {
	*out = *in
	if in.AuthType != nil {
		in, out := &in.AuthType, &out.AuthType
		*out = new(string)
		**out = **in
	}
	if in.MasterKey != nil {
		in, out := &in.MasterKey, &out.MasterKey
		*out = new(string)
		**out = **in
	}
	if in.SlaveKey != nil {
		in, out := &in.SlaveKey, &out.SlaveKey
		*out = new(string)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSpecAuthConfig.
func (in *DomainSpecAuthConfig) DeepCopy() *DomainSpecAuthConfig {
	if in == nil {
		return nil
	}
	out := new(DomainSpecAuthConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSpecCacheConfig) DeepCopyInto(out *DomainSpecCacheConfig) {
	*out = *in
	if in.CacheContent != nil {
		in, out := &in.CacheContent, &out.CacheContent
		*out = new(string)
		**out = **in
	}
	if in.CacheID != nil {
		in, out := &in.CacheID, &out.CacheID
		*out = new(string)
		**out = **in
	}
	if in.CacheType != nil {
		in, out := &in.CacheType, &out.CacheType
		*out = new(string)
		**out = **in
	}
	if in.Ttl != nil {
		in, out := &in.Ttl, &out.Ttl
		*out = new(int64)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSpecCacheConfig.
func (in *DomainSpecCacheConfig) DeepCopy() *DomainSpecCacheConfig {
	if in == nil {
		return nil
	}
	out := new(DomainSpecCacheConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSpecCertificateConfig) DeepCopyInto(out *DomainSpecCertificateConfig) {
	*out = *in
	if in.PrivateKey != nil {
		in, out := &in.PrivateKey, &out.PrivateKey
		*out = new(string)
		**out = **in
	}
	if in.ServerCertificate != nil {
		in, out := &in.ServerCertificate, &out.ServerCertificate
		*out = new(string)
		**out = **in
	}
	if in.ServerCertificateStatus != nil {
		in, out := &in.ServerCertificateStatus, &out.ServerCertificateStatus
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSpecCertificateConfig.
func (in *DomainSpecCertificateConfig) DeepCopy() *DomainSpecCertificateConfig {
	if in == nil {
		return nil
	}
	out := new(DomainSpecCertificateConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSpecHttpHeaderConfig) DeepCopyInto(out *DomainSpecHttpHeaderConfig) {
	*out = *in
	if in.HeaderID != nil {
		in, out := &in.HeaderID, &out.HeaderID
		*out = new(string)
		**out = **in
	}
	if in.HeaderKey != nil {
		in, out := &in.HeaderKey, &out.HeaderKey
		*out = new(string)
		**out = **in
	}
	if in.HeaderValue != nil {
		in, out := &in.HeaderValue, &out.HeaderValue
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSpecHttpHeaderConfig.
func (in *DomainSpecHttpHeaderConfig) DeepCopy() *DomainSpecHttpHeaderConfig {
	if in == nil {
		return nil
	}
	out := new(DomainSpecHttpHeaderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSpecPage404Config) DeepCopyInto(out *DomainSpecPage404Config) {
	*out = *in
	if in.CustomPageURL != nil {
		in, out := &in.CustomPageURL, &out.CustomPageURL
		*out = new(string)
		**out = **in
	}
	if in.ErrorCode != nil {
		in, out := &in.ErrorCode, &out.ErrorCode
		*out = new(string)
		**out = **in
	}
	if in.PageType != nil {
		in, out := &in.PageType, &out.PageType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSpecPage404Config.
func (in *DomainSpecPage404Config) DeepCopy() *DomainSpecPage404Config {
	if in == nil {
		return nil
	}
	out := new(DomainSpecPage404Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSpecParameterFilterConfig) DeepCopyInto(out *DomainSpecParameterFilterConfig) {
	*out = *in
	if in.Enable != nil {
		in, out := &in.Enable, &out.Enable
		*out = new(string)
		**out = **in
	}
	if in.HashKeyArgs != nil {
		in, out := &in.HashKeyArgs, &out.HashKeyArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSpecParameterFilterConfig.
func (in *DomainSpecParameterFilterConfig) DeepCopy() *DomainSpecParameterFilterConfig {
	if in == nil {
		return nil
	}
	out := new(DomainSpecParameterFilterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSpecReferConfig) DeepCopyInto(out *DomainSpecReferConfig) {
	*out = *in
	if in.AllowEmpty != nil {
		in, out := &in.AllowEmpty, &out.AllowEmpty
		*out = new(string)
		**out = **in
	}
	if in.ReferList != nil {
		in, out := &in.ReferList, &out.ReferList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ReferType != nil {
		in, out := &in.ReferType, &out.ReferType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSpecReferConfig.
func (in *DomainSpecReferConfig) DeepCopy() *DomainSpecReferConfig {
	if in == nil {
		return nil
	}
	out := new(DomainSpecReferConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSpecResource) DeepCopyInto(out *DomainSpecResource) {
	*out = *in
	if in.AuthConfig != nil {
		in, out := &in.AuthConfig, &out.AuthConfig
		*out = new(DomainSpecAuthConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.BlockIPS != nil {
		in, out := &in.BlockIPS, &out.BlockIPS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CacheConfig != nil {
		in, out := &in.CacheConfig, &out.CacheConfig
		*out = make([]DomainSpecCacheConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CdnType != nil {
		in, out := &in.CdnType, &out.CdnType
		*out = new(string)
		**out = **in
	}
	if in.CertificateConfig != nil {
		in, out := &in.CertificateConfig, &out.CertificateConfig
		*out = new(DomainSpecCertificateConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.HttpHeaderConfig != nil {
		in, out := &in.HttpHeaderConfig, &out.HttpHeaderConfig
		*out = make([]DomainSpecHttpHeaderConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OptimizeEnable != nil {
		in, out := &in.OptimizeEnable, &out.OptimizeEnable
		*out = new(string)
		**out = **in
	}
	if in.Page404Config != nil {
		in, out := &in.Page404Config, &out.Page404Config
		*out = new(DomainSpecPage404Config)
		(*in).DeepCopyInto(*out)
	}
	if in.PageCompressEnable != nil {
		in, out := &in.PageCompressEnable, &out.PageCompressEnable
		*out = new(string)
		**out = **in
	}
	if in.ParameterFilterConfig != nil {
		in, out := &in.ParameterFilterConfig, &out.ParameterFilterConfig
		*out = new(DomainSpecParameterFilterConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.RangeEnable != nil {
		in, out := &in.RangeEnable, &out.RangeEnable
		*out = new(string)
		**out = **in
	}
	if in.ReferConfig != nil {
		in, out := &in.ReferConfig, &out.ReferConfig
		*out = new(DomainSpecReferConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
	if in.SourcePort != nil {
		in, out := &in.SourcePort, &out.SourcePort
		*out = new(int64)
		**out = **in
	}
	if in.SourceType != nil {
		in, out := &in.SourceType, &out.SourceType
		*out = new(string)
		**out = **in
	}
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VideoSeekEnable != nil {
		in, out := &in.VideoSeekEnable, &out.VideoSeekEnable
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSpecResource.
func (in *DomainSpecResource) DeepCopy() *DomainSpecResource {
	if in == nil {
		return nil
	}
	out := new(DomainSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainStatus) DeepCopyInto(out *DomainStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainStatus.
func (in *DomainStatus) DeepCopy() *DomainStatus {
	if in == nil {
		return nil
	}
	out := new(DomainStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RealTimeLogDelivery) DeepCopyInto(out *RealTimeLogDelivery) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RealTimeLogDelivery.
func (in *RealTimeLogDelivery) DeepCopy() *RealTimeLogDelivery {
	if in == nil {
		return nil
	}
	out := new(RealTimeLogDelivery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RealTimeLogDelivery) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RealTimeLogDeliveryList) DeepCopyInto(out *RealTimeLogDeliveryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RealTimeLogDelivery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RealTimeLogDeliveryList.
func (in *RealTimeLogDeliveryList) DeepCopy() *RealTimeLogDeliveryList {
	if in == nil {
		return nil
	}
	out := new(RealTimeLogDeliveryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RealTimeLogDeliveryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RealTimeLogDeliverySpec) DeepCopyInto(out *RealTimeLogDeliverySpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(RealTimeLogDeliverySpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RealTimeLogDeliverySpec.
func (in *RealTimeLogDeliverySpec) DeepCopy() *RealTimeLogDeliverySpec {
	if in == nil {
		return nil
	}
	out := new(RealTimeLogDeliverySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RealTimeLogDeliverySpecResource) DeepCopyInto(out *RealTimeLogDeliverySpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.Logstore != nil {
		in, out := &in.Logstore, &out.Logstore
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.SlsRegion != nil {
		in, out := &in.SlsRegion, &out.SlsRegion
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RealTimeLogDeliverySpecResource.
func (in *RealTimeLogDeliverySpecResource) DeepCopy() *RealTimeLogDeliverySpecResource {
	if in == nil {
		return nil
	}
	out := new(RealTimeLogDeliverySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RealTimeLogDeliveryStatus) DeepCopyInto(out *RealTimeLogDeliveryStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RealTimeLogDeliveryStatus.
func (in *RealTimeLogDeliveryStatus) DeepCopy() *RealTimeLogDeliveryStatus {
	if in == nil {
		return nil
	}
	out := new(RealTimeLogDeliveryStatus)
	in.DeepCopyInto(out)
	return out
}
