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

	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Actiontrail) DeepCopyInto(out *Actiontrail) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Actiontrail.
func (in *Actiontrail) DeepCopy() *Actiontrail {
	if in == nil {
		return nil
	}
	out := new(Actiontrail)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Actiontrail) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiontrailList) DeepCopyInto(out *ActiontrailList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Actiontrail, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiontrailList.
func (in *ActiontrailList) DeepCopy() *ActiontrailList {
	if in == nil {
		return nil
	}
	out := new(ActiontrailList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActiontrailList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiontrailSpec) DeepCopyInto(out *ActiontrailSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ActiontrailSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiontrailSpec.
func (in *ActiontrailSpec) DeepCopy() *ActiontrailSpec {
	if in == nil {
		return nil
	}
	out := new(ActiontrailSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiontrailSpecResource) DeepCopyInto(out *ActiontrailSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.EventRw != nil {
		in, out := &in.EventRw, &out.EventRw
		*out = new(string)
		**out = **in
	}
	if in.IsOrganizationTrail != nil {
		in, out := &in.IsOrganizationTrail, &out.IsOrganizationTrail
		*out = new(bool)
		**out = **in
	}
	if in.MnsTopicArn != nil {
		in, out := &in.MnsTopicArn, &out.MnsTopicArn
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OssBucketName != nil {
		in, out := &in.OssBucketName, &out.OssBucketName
		*out = new(string)
		**out = **in
	}
	if in.OssKeyPrefix != nil {
		in, out := &in.OssKeyPrefix, &out.OssKeyPrefix
		*out = new(string)
		**out = **in
	}
	if in.OssWriteRoleArn != nil {
		in, out := &in.OssWriteRoleArn, &out.OssWriteRoleArn
		*out = new(string)
		**out = **in
	}
	if in.RoleName != nil {
		in, out := &in.RoleName, &out.RoleName
		*out = new(string)
		**out = **in
	}
	if in.SlsProjectArn != nil {
		in, out := &in.SlsProjectArn, &out.SlsProjectArn
		*out = new(string)
		**out = **in
	}
	if in.SlsWriteRoleArn != nil {
		in, out := &in.SlsWriteRoleArn, &out.SlsWriteRoleArn
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.TrailName != nil {
		in, out := &in.TrailName, &out.TrailName
		*out = new(string)
		**out = **in
	}
	if in.TrailRegion != nil {
		in, out := &in.TrailRegion, &out.TrailRegion
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiontrailSpecResource.
func (in *ActiontrailSpecResource) DeepCopy() *ActiontrailSpecResource {
	if in == nil {
		return nil
	}
	out := new(ActiontrailSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiontrailStatus) DeepCopyInto(out *ActiontrailStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiontrailStatus.
func (in *ActiontrailStatus) DeepCopy() *ActiontrailStatus {
	if in == nil {
		return nil
	}
	out := new(ActiontrailStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trail) DeepCopyInto(out *Trail) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trail.
func (in *Trail) DeepCopy() *Trail {
	if in == nil {
		return nil
	}
	out := new(Trail)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Trail) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrailList) DeepCopyInto(out *TrailList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Trail, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrailList.
func (in *TrailList) DeepCopy() *TrailList {
	if in == nil {
		return nil
	}
	out := new(TrailList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrailList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrailSpec) DeepCopyInto(out *TrailSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(TrailSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrailSpec.
func (in *TrailSpec) DeepCopy() *TrailSpec {
	if in == nil {
		return nil
	}
	out := new(TrailSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrailSpecResource) DeepCopyInto(out *TrailSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.EventRw != nil {
		in, out := &in.EventRw, &out.EventRw
		*out = new(string)
		**out = **in
	}
	if in.IsOrganizationTrail != nil {
		in, out := &in.IsOrganizationTrail, &out.IsOrganizationTrail
		*out = new(bool)
		**out = **in
	}
	if in.MnsTopicArn != nil {
		in, out := &in.MnsTopicArn, &out.MnsTopicArn
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OssBucketName != nil {
		in, out := &in.OssBucketName, &out.OssBucketName
		*out = new(string)
		**out = **in
	}
	if in.OssKeyPrefix != nil {
		in, out := &in.OssKeyPrefix, &out.OssKeyPrefix
		*out = new(string)
		**out = **in
	}
	if in.OssWriteRoleArn != nil {
		in, out := &in.OssWriteRoleArn, &out.OssWriteRoleArn
		*out = new(string)
		**out = **in
	}
	if in.RoleName != nil {
		in, out := &in.RoleName, &out.RoleName
		*out = new(string)
		**out = **in
	}
	if in.SlsProjectArn != nil {
		in, out := &in.SlsProjectArn, &out.SlsProjectArn
		*out = new(string)
		**out = **in
	}
	if in.SlsWriteRoleArn != nil {
		in, out := &in.SlsWriteRoleArn, &out.SlsWriteRoleArn
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.TrailName != nil {
		in, out := &in.TrailName, &out.TrailName
		*out = new(string)
		**out = **in
	}
	if in.TrailRegion != nil {
		in, out := &in.TrailRegion, &out.TrailRegion
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrailSpecResource.
func (in *TrailSpecResource) DeepCopy() *TrailSpecResource {
	if in == nil {
		return nil
	}
	out := new(TrailSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrailStatus) DeepCopyInto(out *TrailStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrailStatus.
func (in *TrailStatus) DeepCopy() *TrailStatus {
	if in == nil {
		return nil
	}
	out := new(TrailStatus)
	in.DeepCopyInto(out)
	return out
}