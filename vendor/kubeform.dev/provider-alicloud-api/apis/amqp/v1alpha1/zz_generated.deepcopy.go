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
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Exchange) DeepCopyInto(out *Exchange) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Exchange.
func (in *Exchange) DeepCopy() *Exchange {
	if in == nil {
		return nil
	}
	out := new(Exchange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Exchange) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExchangeList) DeepCopyInto(out *ExchangeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Exchange, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExchangeList.
func (in *ExchangeList) DeepCopy() *ExchangeList {
	if in == nil {
		return nil
	}
	out := new(ExchangeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExchangeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExchangeSpec) DeepCopyInto(out *ExchangeSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ExchangeSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExchangeSpec.
func (in *ExchangeSpec) DeepCopy() *ExchangeSpec {
	if in == nil {
		return nil
	}
	out := new(ExchangeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExchangeSpecResource) DeepCopyInto(out *ExchangeSpecResource) {
	*out = *in
	if in.AlternateExchange != nil {
		in, out := &in.AlternateExchange, &out.AlternateExchange
		*out = new(string)
		**out = **in
	}
	if in.AutoDeleteState != nil {
		in, out := &in.AutoDeleteState, &out.AutoDeleteState
		*out = new(bool)
		**out = **in
	}
	if in.ExchangeName != nil {
		in, out := &in.ExchangeName, &out.ExchangeName
		*out = new(string)
		**out = **in
	}
	if in.ExchangeType != nil {
		in, out := &in.ExchangeType, &out.ExchangeType
		*out = new(string)
		**out = **in
	}
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	if in.Internal != nil {
		in, out := &in.Internal, &out.Internal
		*out = new(bool)
		**out = **in
	}
	if in.VirtualHostName != nil {
		in, out := &in.VirtualHostName, &out.VirtualHostName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExchangeSpecResource.
func (in *ExchangeSpecResource) DeepCopy() *ExchangeSpecResource {
	if in == nil {
		return nil
	}
	out := new(ExchangeSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExchangeStatus) DeepCopyInto(out *ExchangeStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExchangeStatus.
func (in *ExchangeStatus) DeepCopy() *ExchangeStatus {
	if in == nil {
		return nil
	}
	out := new(ExchangeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Queue) DeepCopyInto(out *Queue) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Queue.
func (in *Queue) DeepCopy() *Queue {
	if in == nil {
		return nil
	}
	out := new(Queue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Queue) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueList) DeepCopyInto(out *QueueList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Queue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueList.
func (in *QueueList) DeepCopy() *QueueList {
	if in == nil {
		return nil
	}
	out := new(QueueList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QueueList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueSpec) DeepCopyInto(out *QueueSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(QueueSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueSpec.
func (in *QueueSpec) DeepCopy() *QueueSpec {
	if in == nil {
		return nil
	}
	out := new(QueueSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueSpecResource) DeepCopyInto(out *QueueSpecResource) {
	*out = *in
	if in.AutoDeleteState != nil {
		in, out := &in.AutoDeleteState, &out.AutoDeleteState
		*out = new(bool)
		**out = **in
	}
	if in.AutoExpireState != nil {
		in, out := &in.AutoExpireState, &out.AutoExpireState
		*out = new(string)
		**out = **in
	}
	if in.DeadLetterExchange != nil {
		in, out := &in.DeadLetterExchange, &out.DeadLetterExchange
		*out = new(string)
		**out = **in
	}
	if in.DeadLetterRoutingKey != nil {
		in, out := &in.DeadLetterRoutingKey, &out.DeadLetterRoutingKey
		*out = new(string)
		**out = **in
	}
	if in.ExclusiveState != nil {
		in, out := &in.ExclusiveState, &out.ExclusiveState
		*out = new(bool)
		**out = **in
	}
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	if in.MaxLength != nil {
		in, out := &in.MaxLength, &out.MaxLength
		*out = new(string)
		**out = **in
	}
	if in.MaximumPriority != nil {
		in, out := &in.MaximumPriority, &out.MaximumPriority
		*out = new(int64)
		**out = **in
	}
	if in.MessageTtl != nil {
		in, out := &in.MessageTtl, &out.MessageTtl
		*out = new(string)
		**out = **in
	}
	if in.QueueName != nil {
		in, out := &in.QueueName, &out.QueueName
		*out = new(string)
		**out = **in
	}
	if in.VirtualHostName != nil {
		in, out := &in.VirtualHostName, &out.VirtualHostName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueSpecResource.
func (in *QueueSpecResource) DeepCopy() *QueueSpecResource {
	if in == nil {
		return nil
	}
	out := new(QueueSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueStatus) DeepCopyInto(out *QueueStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueStatus.
func (in *QueueStatus) DeepCopy() *QueueStatus {
	if in == nil {
		return nil
	}
	out := new(QueueStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualHost) DeepCopyInto(out *VirtualHost) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualHost.
func (in *VirtualHost) DeepCopy() *VirtualHost {
	if in == nil {
		return nil
	}
	out := new(VirtualHost)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualHost) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualHostList) DeepCopyInto(out *VirtualHostList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualHost, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualHostList.
func (in *VirtualHostList) DeepCopy() *VirtualHostList {
	if in == nil {
		return nil
	}
	out := new(VirtualHostList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualHostList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualHostSpec) DeepCopyInto(out *VirtualHostSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(VirtualHostSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualHostSpec.
func (in *VirtualHostSpec) DeepCopy() *VirtualHostSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualHostSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualHostSpecResource) DeepCopyInto(out *VirtualHostSpecResource) {
	*out = *in
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	if in.VirtualHostName != nil {
		in, out := &in.VirtualHostName, &out.VirtualHostName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualHostSpecResource.
func (in *VirtualHostSpecResource) DeepCopy() *VirtualHostSpecResource {
	if in == nil {
		return nil
	}
	out := new(VirtualHostSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualHostStatus) DeepCopyInto(out *VirtualHostStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualHostStatus.
func (in *VirtualHostStatus) DeepCopy() *VirtualHostStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualHostStatus)
	in.DeepCopyInto(out)
	return out
}
