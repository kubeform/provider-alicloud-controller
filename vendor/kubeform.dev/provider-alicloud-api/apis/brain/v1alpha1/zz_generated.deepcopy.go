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
func (in *IndustrialPidLoop) DeepCopyInto(out *IndustrialPidLoop) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndustrialPidLoop.
func (in *IndustrialPidLoop) DeepCopy() *IndustrialPidLoop {
	if in == nil {
		return nil
	}
	out := new(IndustrialPidLoop)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IndustrialPidLoop) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndustrialPidLoopList) DeepCopyInto(out *IndustrialPidLoopList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IndustrialPidLoop, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndustrialPidLoopList.
func (in *IndustrialPidLoopList) DeepCopy() *IndustrialPidLoopList {
	if in == nil {
		return nil
	}
	out := new(IndustrialPidLoopList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IndustrialPidLoopList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndustrialPidLoopSpec) DeepCopyInto(out *IndustrialPidLoopSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(IndustrialPidLoopSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndustrialPidLoopSpec.
func (in *IndustrialPidLoopSpec) DeepCopy() *IndustrialPidLoopSpec {
	if in == nil {
		return nil
	}
	out := new(IndustrialPidLoopSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndustrialPidLoopSpecResource) DeepCopyInto(out *IndustrialPidLoopSpecResource) {
	*out = *in
	if in.PidLoopConfiguration != nil {
		in, out := &in.PidLoopConfiguration, &out.PidLoopConfiguration
		*out = new(string)
		**out = **in
	}
	if in.PidLoopDcsType != nil {
		in, out := &in.PidLoopDcsType, &out.PidLoopDcsType
		*out = new(string)
		**out = **in
	}
	if in.PidLoopDesc != nil {
		in, out := &in.PidLoopDesc, &out.PidLoopDesc
		*out = new(string)
		**out = **in
	}
	if in.PidLoopIsCrucial != nil {
		in, out := &in.PidLoopIsCrucial, &out.PidLoopIsCrucial
		*out = new(bool)
		**out = **in
	}
	if in.PidLoopName != nil {
		in, out := &in.PidLoopName, &out.PidLoopName
		*out = new(string)
		**out = **in
	}
	if in.PidLoopType != nil {
		in, out := &in.PidLoopType, &out.PidLoopType
		*out = new(string)
		**out = **in
	}
	if in.PidProjectID != nil {
		in, out := &in.PidProjectID, &out.PidProjectID
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndustrialPidLoopSpecResource.
func (in *IndustrialPidLoopSpecResource) DeepCopy() *IndustrialPidLoopSpecResource {
	if in == nil {
		return nil
	}
	out := new(IndustrialPidLoopSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndustrialPidLoopStatus) DeepCopyInto(out *IndustrialPidLoopStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndustrialPidLoopStatus.
func (in *IndustrialPidLoopStatus) DeepCopy() *IndustrialPidLoopStatus {
	if in == nil {
		return nil
	}
	out := new(IndustrialPidLoopStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndustrialPidOrganization) DeepCopyInto(out *IndustrialPidOrganization) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndustrialPidOrganization.
func (in *IndustrialPidOrganization) DeepCopy() *IndustrialPidOrganization {
	if in == nil {
		return nil
	}
	out := new(IndustrialPidOrganization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IndustrialPidOrganization) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndustrialPidOrganizationList) DeepCopyInto(out *IndustrialPidOrganizationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IndustrialPidOrganization, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndustrialPidOrganizationList.
func (in *IndustrialPidOrganizationList) DeepCopy() *IndustrialPidOrganizationList {
	if in == nil {
		return nil
	}
	out := new(IndustrialPidOrganizationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IndustrialPidOrganizationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndustrialPidOrganizationSpec) DeepCopyInto(out *IndustrialPidOrganizationSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(IndustrialPidOrganizationSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndustrialPidOrganizationSpec.
func (in *IndustrialPidOrganizationSpec) DeepCopy() *IndustrialPidOrganizationSpec {
	if in == nil {
		return nil
	}
	out := new(IndustrialPidOrganizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndustrialPidOrganizationSpecResource) DeepCopyInto(out *IndustrialPidOrganizationSpecResource) {
	*out = *in
	if in.ParentPidOrganizationID != nil {
		in, out := &in.ParentPidOrganizationID, &out.ParentPidOrganizationID
		*out = new(string)
		**out = **in
	}
	if in.PidOrganizationName != nil {
		in, out := &in.PidOrganizationName, &out.PidOrganizationName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndustrialPidOrganizationSpecResource.
func (in *IndustrialPidOrganizationSpecResource) DeepCopy() *IndustrialPidOrganizationSpecResource {
	if in == nil {
		return nil
	}
	out := new(IndustrialPidOrganizationSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndustrialPidOrganizationStatus) DeepCopyInto(out *IndustrialPidOrganizationStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndustrialPidOrganizationStatus.
func (in *IndustrialPidOrganizationStatus) DeepCopy() *IndustrialPidOrganizationStatus {
	if in == nil {
		return nil
	}
	out := new(IndustrialPidOrganizationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndustrialPidProject) DeepCopyInto(out *IndustrialPidProject) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndustrialPidProject.
func (in *IndustrialPidProject) DeepCopy() *IndustrialPidProject {
	if in == nil {
		return nil
	}
	out := new(IndustrialPidProject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IndustrialPidProject) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndustrialPidProjectList) DeepCopyInto(out *IndustrialPidProjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IndustrialPidProject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndustrialPidProjectList.
func (in *IndustrialPidProjectList) DeepCopy() *IndustrialPidProjectList {
	if in == nil {
		return nil
	}
	out := new(IndustrialPidProjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IndustrialPidProjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndustrialPidProjectSpec) DeepCopyInto(out *IndustrialPidProjectSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(IndustrialPidProjectSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndustrialPidProjectSpec.
func (in *IndustrialPidProjectSpec) DeepCopy() *IndustrialPidProjectSpec {
	if in == nil {
		return nil
	}
	out := new(IndustrialPidProjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndustrialPidProjectSpecResource) DeepCopyInto(out *IndustrialPidProjectSpecResource) {
	*out = *in
	if in.PidOrganizationID != nil {
		in, out := &in.PidOrganizationID, &out.PidOrganizationID
		*out = new(string)
		**out = **in
	}
	if in.PidProjectDesc != nil {
		in, out := &in.PidProjectDesc, &out.PidProjectDesc
		*out = new(string)
		**out = **in
	}
	if in.PidProjectName != nil {
		in, out := &in.PidProjectName, &out.PidProjectName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndustrialPidProjectSpecResource.
func (in *IndustrialPidProjectSpecResource) DeepCopy() *IndustrialPidProjectSpecResource {
	if in == nil {
		return nil
	}
	out := new(IndustrialPidProjectSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndustrialPidProjectStatus) DeepCopyInto(out *IndustrialPidProjectStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndustrialPidProjectStatus.
func (in *IndustrialPidProjectStatus) DeepCopy() *IndustrialPidProjectStatus {
	if in == nil {
		return nil
	}
	out := new(IndustrialPidProjectStatus)
	in.DeepCopyInto(out)
	return out
}
