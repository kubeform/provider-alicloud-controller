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
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Alias) DeepCopyInto(out *Alias) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alias.
func (in *Alias) DeepCopy() *Alias {
	if in == nil {
		return nil
	}
	out := new(Alias)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Alias) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasList) DeepCopyInto(out *AliasList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Alias, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasList.
func (in *AliasList) DeepCopy() *AliasList {
	if in == nil {
		return nil
	}
	out := new(AliasList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AliasList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasSpec) DeepCopyInto(out *AliasSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(AliasSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasSpec.
func (in *AliasSpec) DeepCopy() *AliasSpec {
	if in == nil {
		return nil
	}
	out := new(AliasSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasSpecResource) DeepCopyInto(out *AliasSpecResource) {
	*out = *in
	if in.AliasName != nil {
		in, out := &in.AliasName, &out.AliasName
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.RoutingConfig != nil {
		in, out := &in.RoutingConfig, &out.RoutingConfig
		*out = new(AliasSpecRoutingConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
	if in.ServiceVersion != nil {
		in, out := &in.ServiceVersion, &out.ServiceVersion
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasSpecResource.
func (in *AliasSpecResource) DeepCopy() *AliasSpecResource {
	if in == nil {
		return nil
	}
	out := new(AliasSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasSpecRoutingConfig) DeepCopyInto(out *AliasSpecRoutingConfig) {
	*out = *in
	if in.AdditionalVersionWeights != nil {
		in, out := &in.AdditionalVersionWeights, &out.AdditionalVersionWeights
		*out = new(map[string]float64)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]float64, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasSpecRoutingConfig.
func (in *AliasSpecRoutingConfig) DeepCopy() *AliasSpecRoutingConfig {
	if in == nil {
		return nil
	}
	out := new(AliasSpecRoutingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasStatus) DeepCopyInto(out *AliasStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasStatus.
func (in *AliasStatus) DeepCopy() *AliasStatus {
	if in == nil {
		return nil
	}
	out := new(AliasStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomDomain) DeepCopyInto(out *CustomDomain) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomDomain.
func (in *CustomDomain) DeepCopy() *CustomDomain {
	if in == nil {
		return nil
	}
	out := new(CustomDomain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomDomain) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomDomainList) DeepCopyInto(out *CustomDomainList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CustomDomain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomDomainList.
func (in *CustomDomainList) DeepCopy() *CustomDomainList {
	if in == nil {
		return nil
	}
	out := new(CustomDomainList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomDomainList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomDomainSpec) DeepCopyInto(out *CustomDomainSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(CustomDomainSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomDomainSpec.
func (in *CustomDomainSpec) DeepCopy() *CustomDomainSpec {
	if in == nil {
		return nil
	}
	out := new(CustomDomainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomDomainSpecCertConfig) DeepCopyInto(out *CustomDomainSpecCertConfig) {
	*out = *in
	if in.CertName != nil {
		in, out := &in.CertName, &out.CertName
		*out = new(string)
		**out = **in
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(string)
		**out = **in
	}
	if in.PrivateKey != nil {
		in, out := &in.PrivateKey, &out.PrivateKey
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomDomainSpecCertConfig.
func (in *CustomDomainSpecCertConfig) DeepCopy() *CustomDomainSpecCertConfig {
	if in == nil {
		return nil
	}
	out := new(CustomDomainSpecCertConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomDomainSpecResource) DeepCopyInto(out *CustomDomainSpecResource) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.ApiVersion != nil {
		in, out := &in.ApiVersion, &out.ApiVersion
		*out = new(string)
		**out = **in
	}
	if in.CertConfig != nil {
		in, out := &in.CertConfig, &out.CertConfig
		*out = new(CustomDomainSpecCertConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.CreatedTime != nil {
		in, out := &in.CreatedTime, &out.CreatedTime
		*out = new(string)
		**out = **in
	}
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedTime != nil {
		in, out := &in.LastModifiedTime, &out.LastModifiedTime
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.RouteConfig != nil {
		in, out := &in.RouteConfig, &out.RouteConfig
		*out = make([]CustomDomainSpecRouteConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomDomainSpecResource.
func (in *CustomDomainSpecResource) DeepCopy() *CustomDomainSpecResource {
	if in == nil {
		return nil
	}
	out := new(CustomDomainSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomDomainSpecRouteConfig) DeepCopyInto(out *CustomDomainSpecRouteConfig) {
	*out = *in
	if in.FunctionName != nil {
		in, out := &in.FunctionName, &out.FunctionName
		*out = new(string)
		**out = **in
	}
	if in.Methods != nil {
		in, out := &in.Methods, &out.Methods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Qualifier != nil {
		in, out := &in.Qualifier, &out.Qualifier
		*out = new(string)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomDomainSpecRouteConfig.
func (in *CustomDomainSpecRouteConfig) DeepCopy() *CustomDomainSpecRouteConfig {
	if in == nil {
		return nil
	}
	out := new(CustomDomainSpecRouteConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomDomainStatus) DeepCopyInto(out *CustomDomainStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomDomainStatus.
func (in *CustomDomainStatus) DeepCopy() *CustomDomainStatus {
	if in == nil {
		return nil
	}
	out := new(CustomDomainStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Function) DeepCopyInto(out *Function) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Function.
func (in *Function) DeepCopy() *Function {
	if in == nil {
		return nil
	}
	out := new(Function)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Function) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionAsyncInvokeConfig) DeepCopyInto(out *FunctionAsyncInvokeConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionAsyncInvokeConfig.
func (in *FunctionAsyncInvokeConfig) DeepCopy() *FunctionAsyncInvokeConfig {
	if in == nil {
		return nil
	}
	out := new(FunctionAsyncInvokeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionAsyncInvokeConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionAsyncInvokeConfigList) DeepCopyInto(out *FunctionAsyncInvokeConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FunctionAsyncInvokeConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionAsyncInvokeConfigList.
func (in *FunctionAsyncInvokeConfigList) DeepCopy() *FunctionAsyncInvokeConfigList {
	if in == nil {
		return nil
	}
	out := new(FunctionAsyncInvokeConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionAsyncInvokeConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionAsyncInvokeConfigSpec) DeepCopyInto(out *FunctionAsyncInvokeConfigSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(FunctionAsyncInvokeConfigSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionAsyncInvokeConfigSpec.
func (in *FunctionAsyncInvokeConfigSpec) DeepCopy() *FunctionAsyncInvokeConfigSpec {
	if in == nil {
		return nil
	}
	out := new(FunctionAsyncInvokeConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionAsyncInvokeConfigSpecDestinationConfig) DeepCopyInto(out *FunctionAsyncInvokeConfigSpecDestinationConfig) {
	*out = *in
	if in.OnFailure != nil {
		in, out := &in.OnFailure, &out.OnFailure
		*out = new(FunctionAsyncInvokeConfigSpecDestinationConfigOnFailure)
		(*in).DeepCopyInto(*out)
	}
	if in.OnSuccess != nil {
		in, out := &in.OnSuccess, &out.OnSuccess
		*out = new(FunctionAsyncInvokeConfigSpecDestinationConfigOnSuccess)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionAsyncInvokeConfigSpecDestinationConfig.
func (in *FunctionAsyncInvokeConfigSpecDestinationConfig) DeepCopy() *FunctionAsyncInvokeConfigSpecDestinationConfig {
	if in == nil {
		return nil
	}
	out := new(FunctionAsyncInvokeConfigSpecDestinationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionAsyncInvokeConfigSpecDestinationConfigOnFailure) DeepCopyInto(out *FunctionAsyncInvokeConfigSpecDestinationConfigOnFailure) {
	*out = *in
	if in.Destination != nil {
		in, out := &in.Destination, &out.Destination
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionAsyncInvokeConfigSpecDestinationConfigOnFailure.
func (in *FunctionAsyncInvokeConfigSpecDestinationConfigOnFailure) DeepCopy() *FunctionAsyncInvokeConfigSpecDestinationConfigOnFailure {
	if in == nil {
		return nil
	}
	out := new(FunctionAsyncInvokeConfigSpecDestinationConfigOnFailure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionAsyncInvokeConfigSpecDestinationConfigOnSuccess) DeepCopyInto(out *FunctionAsyncInvokeConfigSpecDestinationConfigOnSuccess) {
	*out = *in
	if in.Destination != nil {
		in, out := &in.Destination, &out.Destination
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionAsyncInvokeConfigSpecDestinationConfigOnSuccess.
func (in *FunctionAsyncInvokeConfigSpecDestinationConfigOnSuccess) DeepCopy() *FunctionAsyncInvokeConfigSpecDestinationConfigOnSuccess {
	if in == nil {
		return nil
	}
	out := new(FunctionAsyncInvokeConfigSpecDestinationConfigOnSuccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionAsyncInvokeConfigSpecResource) DeepCopyInto(out *FunctionAsyncInvokeConfigSpecResource) {
	*out = *in
	if in.CreatedTime != nil {
		in, out := &in.CreatedTime, &out.CreatedTime
		*out = new(string)
		**out = **in
	}
	if in.DestinationConfig != nil {
		in, out := &in.DestinationConfig, &out.DestinationConfig
		*out = new(FunctionAsyncInvokeConfigSpecDestinationConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.FunctionName != nil {
		in, out := &in.FunctionName, &out.FunctionName
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedTime != nil {
		in, out := &in.LastModifiedTime, &out.LastModifiedTime
		*out = new(string)
		**out = **in
	}
	if in.MaximumEventAgeInSeconds != nil {
		in, out := &in.MaximumEventAgeInSeconds, &out.MaximumEventAgeInSeconds
		*out = new(int64)
		**out = **in
	}
	if in.MaximumRetryAttempts != nil {
		in, out := &in.MaximumRetryAttempts, &out.MaximumRetryAttempts
		*out = new(int64)
		**out = **in
	}
	if in.Qualifier != nil {
		in, out := &in.Qualifier, &out.Qualifier
		*out = new(string)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionAsyncInvokeConfigSpecResource.
func (in *FunctionAsyncInvokeConfigSpecResource) DeepCopy() *FunctionAsyncInvokeConfigSpecResource {
	if in == nil {
		return nil
	}
	out := new(FunctionAsyncInvokeConfigSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionAsyncInvokeConfigStatus) DeepCopyInto(out *FunctionAsyncInvokeConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionAsyncInvokeConfigStatus.
func (in *FunctionAsyncInvokeConfigStatus) DeepCopy() *FunctionAsyncInvokeConfigStatus {
	if in == nil {
		return nil
	}
	out := new(FunctionAsyncInvokeConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionList) DeepCopyInto(out *FunctionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Function, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionList.
func (in *FunctionList) DeepCopy() *FunctionList {
	if in == nil {
		return nil
	}
	out := new(FunctionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionSpec) DeepCopyInto(out *FunctionSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(FunctionSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionSpec.
func (in *FunctionSpec) DeepCopy() *FunctionSpec {
	if in == nil {
		return nil
	}
	out := new(FunctionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionSpecCustomContainerConfig) DeepCopyInto(out *FunctionSpecCustomContainerConfig) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = new(string)
		**out = **in
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = new(string)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionSpecCustomContainerConfig.
func (in *FunctionSpecCustomContainerConfig) DeepCopy() *FunctionSpecCustomContainerConfig {
	if in == nil {
		return nil
	}
	out := new(FunctionSpecCustomContainerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionSpecResource) DeepCopyInto(out *FunctionSpecResource) {
	*out = *in
	if in.CaPort != nil {
		in, out := &in.CaPort, &out.CaPort
		*out = new(int64)
		**out = **in
	}
	if in.CodeChecksum != nil {
		in, out := &in.CodeChecksum, &out.CodeChecksum
		*out = new(string)
		**out = **in
	}
	if in.CustomContainerConfig != nil {
		in, out := &in.CustomContainerConfig, &out.CustomContainerConfig
		*out = new(FunctionSpecCustomContainerConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EnvironmentVariables != nil {
		in, out := &in.EnvironmentVariables, &out.EnvironmentVariables
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Filename != nil {
		in, out := &in.Filename, &out.Filename
		*out = new(string)
		**out = **in
	}
	if in.FunctionID != nil {
		in, out := &in.FunctionID, &out.FunctionID
		*out = new(string)
		**out = **in
	}
	if in.Handler != nil {
		in, out := &in.Handler, &out.Handler
		*out = new(string)
		**out = **in
	}
	if in.InitializationTimeout != nil {
		in, out := &in.InitializationTimeout, &out.InitializationTimeout
		*out = new(int64)
		**out = **in
	}
	if in.Initializer != nil {
		in, out := &in.Initializer, &out.Initializer
		*out = new(string)
		**out = **in
	}
	if in.InstanceConcurrency != nil {
		in, out := &in.InstanceConcurrency, &out.InstanceConcurrency
		*out = new(int64)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(string)
		**out = **in
	}
	if in.LastModified != nil {
		in, out := &in.LastModified, &out.LastModified
		*out = new(string)
		**out = **in
	}
	if in.MemorySize != nil {
		in, out := &in.MemorySize, &out.MemorySize
		*out = new(int64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamePrefix != nil {
		in, out := &in.NamePrefix, &out.NamePrefix
		*out = new(string)
		**out = **in
	}
	if in.OssBucket != nil {
		in, out := &in.OssBucket, &out.OssBucket
		*out = new(string)
		**out = **in
	}
	if in.OssKey != nil {
		in, out := &in.OssKey, &out.OssKey
		*out = new(string)
		**out = **in
	}
	if in.Runtime != nil {
		in, out := &in.Runtime, &out.Runtime
		*out = new(string)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionSpecResource.
func (in *FunctionSpecResource) DeepCopy() *FunctionSpecResource {
	if in == nil {
		return nil
	}
	out := new(FunctionSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionStatus) DeepCopyInto(out *FunctionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionStatus.
func (in *FunctionStatus) DeepCopy() *FunctionStatus {
	if in == nil {
		return nil
	}
	out := new(FunctionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Service) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceList) DeepCopyInto(out *ServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Service, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceList.
func (in *ServiceList) DeepCopy() *ServiceList {
	if in == nil {
		return nil
	}
	out := new(ServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpec) DeepCopyInto(out *ServiceSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ServiceSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpec.
func (in *ServiceSpec) DeepCopy() *ServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpecLogConfig) DeepCopyInto(out *ServiceSpecLogConfig) {
	*out = *in
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
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpecLogConfig.
func (in *ServiceSpecLogConfig) DeepCopy() *ServiceSpecLogConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceSpecLogConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpecNasConfig) DeepCopyInto(out *ServiceSpecNasConfig) {
	*out = *in
	if in.GroupID != nil {
		in, out := &in.GroupID, &out.GroupID
		*out = new(int64)
		**out = **in
	}
	if in.MountPoints != nil {
		in, out := &in.MountPoints, &out.MountPoints
		*out = make([]ServiceSpecNasConfigMountPoints, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UserID != nil {
		in, out := &in.UserID, &out.UserID
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpecNasConfig.
func (in *ServiceSpecNasConfig) DeepCopy() *ServiceSpecNasConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceSpecNasConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpecNasConfigMountPoints) DeepCopyInto(out *ServiceSpecNasConfigMountPoints) {
	*out = *in
	if in.MountDir != nil {
		in, out := &in.MountDir, &out.MountDir
		*out = new(string)
		**out = **in
	}
	if in.ServerAddr != nil {
		in, out := &in.ServerAddr, &out.ServerAddr
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpecNasConfigMountPoints.
func (in *ServiceSpecNasConfigMountPoints) DeepCopy() *ServiceSpecNasConfigMountPoints {
	if in == nil {
		return nil
	}
	out := new(ServiceSpecNasConfigMountPoints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpecResource) DeepCopyInto(out *ServiceSpecResource) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.InternetAccess != nil {
		in, out := &in.InternetAccess, &out.InternetAccess
		*out = new(bool)
		**out = **in
	}
	if in.LastModified != nil {
		in, out := &in.LastModified, &out.LastModified
		*out = new(string)
		**out = **in
	}
	if in.LogConfig != nil {
		in, out := &in.LogConfig, &out.LogConfig
		*out = new(ServiceSpecLogConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamePrefix != nil {
		in, out := &in.NamePrefix, &out.NamePrefix
		*out = new(string)
		**out = **in
	}
	if in.NasConfig != nil {
		in, out := &in.NasConfig, &out.NasConfig
		*out = new(ServiceSpecNasConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Publish != nil {
		in, out := &in.Publish, &out.Publish
		*out = new(bool)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.ServiceID != nil {
		in, out := &in.ServiceID, &out.ServiceID
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.VpcConfig != nil {
		in, out := &in.VpcConfig, &out.VpcConfig
		*out = new(ServiceSpecVpcConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpecResource.
func (in *ServiceSpecResource) DeepCopy() *ServiceSpecResource {
	if in == nil {
		return nil
	}
	out := new(ServiceSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpecVpcConfig) DeepCopyInto(out *ServiceSpecVpcConfig) {
	*out = *in
	if in.SecurityGroupID != nil {
		in, out := &in.SecurityGroupID, &out.SecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.VpcID != nil {
		in, out := &in.VpcID, &out.VpcID
		*out = new(string)
		**out = **in
	}
	if in.VswitchIDS != nil {
		in, out := &in.VswitchIDS, &out.VswitchIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpecVpcConfig.
func (in *ServiceSpecVpcConfig) DeepCopy() *ServiceSpecVpcConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceSpecVpcConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceStatus) DeepCopyInto(out *ServiceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceStatus.
func (in *ServiceStatus) DeepCopy() *ServiceStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trigger) DeepCopyInto(out *Trigger) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trigger.
func (in *Trigger) DeepCopy() *Trigger {
	if in == nil {
		return nil
	}
	out := new(Trigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Trigger) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerList) DeepCopyInto(out *TriggerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Trigger, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerList.
func (in *TriggerList) DeepCopy() *TriggerList {
	if in == nil {
		return nil
	}
	out := new(TriggerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TriggerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpec) DeepCopyInto(out *TriggerSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(TriggerSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpec.
func (in *TriggerSpec) DeepCopy() *TriggerSpec {
	if in == nil {
		return nil
	}
	out := new(TriggerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpecResource) DeepCopyInto(out *TriggerSpecResource) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(string)
		**out = **in
	}
	if in.ConfigMns != nil {
		in, out := &in.ConfigMns, &out.ConfigMns
		*out = new(string)
		**out = **in
	}
	if in.Function != nil {
		in, out := &in.Function, &out.Function
		*out = new(string)
		**out = **in
	}
	if in.LastModified != nil {
		in, out := &in.LastModified, &out.LastModified
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamePrefix != nil {
		in, out := &in.NamePrefix, &out.NamePrefix
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(string)
		**out = **in
	}
	if in.SourceArn != nil {
		in, out := &in.SourceArn, &out.SourceArn
		*out = new(string)
		**out = **in
	}
	if in.TriggerID != nil {
		in, out := &in.TriggerID, &out.TriggerID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpecResource.
func (in *TriggerSpecResource) DeepCopy() *TriggerSpecResource {
	if in == nil {
		return nil
	}
	out := new(TriggerSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerStatus) DeepCopyInto(out *TriggerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerStatus.
func (in *TriggerStatus) DeepCopy() *TriggerStatus {
	if in == nil {
		return nil
	}
	out := new(TriggerStatus)
	in.DeepCopyInto(out)
	return out
}
