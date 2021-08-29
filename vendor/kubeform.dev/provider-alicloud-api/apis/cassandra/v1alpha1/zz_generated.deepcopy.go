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

	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPlan) DeepCopyInto(out *BackupPlan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPlan.
func (in *BackupPlan) DeepCopy() *BackupPlan {
	if in == nil {
		return nil
	}
	out := new(BackupPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupPlan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPlanList) DeepCopyInto(out *BackupPlanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupPlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPlanList.
func (in *BackupPlanList) DeepCopy() *BackupPlanList {
	if in == nil {
		return nil
	}
	out := new(BackupPlanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupPlanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPlanSpec) DeepCopyInto(out *BackupPlanSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(BackupPlanSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPlanSpec.
func (in *BackupPlanSpec) DeepCopy() *BackupPlanSpec {
	if in == nil {
		return nil
	}
	out := new(BackupPlanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPlanSpecResource) DeepCopyInto(out *BackupPlanSpecResource) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = new(bool)
		**out = **in
	}
	if in.BackupPeriod != nil {
		in, out := &in.BackupPeriod, &out.BackupPeriod
		*out = new(string)
		**out = **in
	}
	if in.BackupTime != nil {
		in, out := &in.BackupTime, &out.BackupTime
		*out = new(string)
		**out = **in
	}
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.DataCenterID != nil {
		in, out := &in.DataCenterID, &out.DataCenterID
		*out = new(string)
		**out = **in
	}
	if in.RetentionPeriod != nil {
		in, out := &in.RetentionPeriod, &out.RetentionPeriod
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPlanSpecResource.
func (in *BackupPlanSpecResource) DeepCopy() *BackupPlanSpecResource {
	if in == nil {
		return nil
	}
	out := new(BackupPlanSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPlanStatus) DeepCopyInto(out *BackupPlanStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPlanStatus.
func (in *BackupPlanStatus) DeepCopy() *BackupPlanStatus {
	if in == nil {
		return nil
	}
	out := new(BackupPlanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ClusterSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecResource) DeepCopyInto(out *ClusterSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AutoRenew != nil {
		in, out := &in.AutoRenew, &out.AutoRenew
		*out = new(bool)
		**out = **in
	}
	if in.AutoRenewPeriod != nil {
		in, out := &in.AutoRenewPeriod, &out.AutoRenewPeriod
		*out = new(int64)
		**out = **in
	}
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.DataCenterName != nil {
		in, out := &in.DataCenterName, &out.DataCenterName
		*out = new(string)
		**out = **in
	}
	if in.DiskSize != nil {
		in, out := &in.DiskSize, &out.DiskSize
		*out = new(int64)
		**out = **in
	}
	if in.DiskType != nil {
		in, out := &in.DiskType, &out.DiskType
		*out = new(string)
		**out = **in
	}
	if in.EnablePublic != nil {
		in, out := &in.EnablePublic, &out.EnablePublic
		*out = new(bool)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(string)
		**out = **in
	}
	if in.IpWhite != nil {
		in, out := &in.IpWhite, &out.IpWhite
		*out = new(string)
		**out = **in
	}
	if in.MaintainEndTime != nil {
		in, out := &in.MaintainEndTime, &out.MaintainEndTime
		*out = new(string)
		**out = **in
	}
	if in.MaintainStartTime != nil {
		in, out := &in.MaintainStartTime, &out.MaintainStartTime
		*out = new(string)
		**out = **in
	}
	if in.MajorVersion != nil {
		in, out := &in.MajorVersion, &out.MajorVersion
		*out = new(string)
		**out = **in
	}
	if in.NodeCount != nil {
		in, out := &in.NodeCount, &out.NodeCount
		*out = new(int64)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.PayType != nil {
		in, out := &in.PayType, &out.PayType
		*out = new(string)
		**out = **in
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(int64)
		**out = **in
	}
	if in.PeriodUnit != nil {
		in, out := &in.PeriodUnit, &out.PeriodUnit
		*out = new(string)
		**out = **in
	}
	if in.PublicPoints != nil {
		in, out := &in.PublicPoints, &out.PublicPoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
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
	if in.VswitchID != nil {
		in, out := &in.VswitchID, &out.VswitchID
		*out = new(string)
		**out = **in
	}
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecResource.
func (in *ClusterSpecResource) DeepCopy() *ClusterSpecResource {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCenter) DeepCopyInto(out *DataCenter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCenter.
func (in *DataCenter) DeepCopy() *DataCenter {
	if in == nil {
		return nil
	}
	out := new(DataCenter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataCenter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCenterList) DeepCopyInto(out *DataCenterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataCenter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCenterList.
func (in *DataCenterList) DeepCopy() *DataCenterList {
	if in == nil {
		return nil
	}
	out := new(DataCenterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataCenterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCenterSpec) DeepCopyInto(out *DataCenterSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(DataCenterSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCenterSpec.
func (in *DataCenterSpec) DeepCopy() *DataCenterSpec {
	if in == nil {
		return nil
	}
	out := new(DataCenterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCenterSpecResource) DeepCopyInto(out *DataCenterSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AutoRenew != nil {
		in, out := &in.AutoRenew, &out.AutoRenew
		*out = new(bool)
		**out = **in
	}
	if in.AutoRenewPeriod != nil {
		in, out := &in.AutoRenewPeriod, &out.AutoRenewPeriod
		*out = new(int64)
		**out = **in
	}
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.DataCenterID != nil {
		in, out := &in.DataCenterID, &out.DataCenterID
		*out = new(string)
		**out = **in
	}
	if in.DataCenterName != nil {
		in, out := &in.DataCenterName, &out.DataCenterName
		*out = new(string)
		**out = **in
	}
	if in.DiskSize != nil {
		in, out := &in.DiskSize, &out.DiskSize
		*out = new(int64)
		**out = **in
	}
	if in.DiskType != nil {
		in, out := &in.DiskType, &out.DiskType
		*out = new(string)
		**out = **in
	}
	if in.EnablePublic != nil {
		in, out := &in.EnablePublic, &out.EnablePublic
		*out = new(bool)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(string)
		**out = **in
	}
	if in.NodeCount != nil {
		in, out := &in.NodeCount, &out.NodeCount
		*out = new(int64)
		**out = **in
	}
	if in.PayType != nil {
		in, out := &in.PayType, &out.PayType
		*out = new(string)
		**out = **in
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(int64)
		**out = **in
	}
	if in.PeriodUnit != nil {
		in, out := &in.PeriodUnit, &out.PeriodUnit
		*out = new(string)
		**out = **in
	}
	if in.PublicPoints != nil {
		in, out := &in.PublicPoints, &out.PublicPoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.VswitchID != nil {
		in, out := &in.VswitchID, &out.VswitchID
		*out = new(string)
		**out = **in
	}
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCenterSpecResource.
func (in *DataCenterSpecResource) DeepCopy() *DataCenterSpecResource {
	if in == nil {
		return nil
	}
	out := new(DataCenterSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCenterStatus) DeepCopyInto(out *DataCenterStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCenterStatus.
func (in *DataCenterStatus) DeepCopy() *DataCenterStatus {
	if in == nil {
		return nil
	}
	out := new(DataCenterStatus)
	in.DeepCopyInto(out)
	return out
}
