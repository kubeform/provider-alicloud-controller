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
func (in *ApplicationInfo) DeepCopyInto(out *ApplicationInfo) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationInfo.
func (in *ApplicationInfo) DeepCopy() *ApplicationInfo {
	if in == nil {
		return nil
	}
	out := new(ApplicationInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationInfo) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationInfoList) DeepCopyInto(out *ApplicationInfoList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApplicationInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationInfoList.
func (in *ApplicationInfoList) DeepCopy() *ApplicationInfoList {
	if in == nil {
		return nil
	}
	out := new(ApplicationInfoList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationInfoList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationInfoSpec) DeepCopyInto(out *ApplicationInfoSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ApplicationInfoSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationInfoSpec.
func (in *ApplicationInfoSpec) DeepCopy() *ApplicationInfoSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationInfoSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationInfoSpecDimensions) DeepCopyInto(out *ApplicationInfoSpecDimensions) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationInfoSpecDimensions.
func (in *ApplicationInfoSpecDimensions) DeepCopy() *ApplicationInfoSpecDimensions {
	if in == nil {
		return nil
	}
	out := new(ApplicationInfoSpecDimensions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationInfoSpecResource) DeepCopyInto(out *ApplicationInfoSpecResource) {
	*out = *in
	if in.ApproveValue != nil {
		in, out := &in.ApproveValue, &out.ApproveValue
		*out = new(string)
		**out = **in
	}
	if in.AuditMode != nil {
		in, out := &in.AuditMode, &out.AuditMode
		*out = new(string)
		**out = **in
	}
	if in.AuditReason != nil {
		in, out := &in.AuditReason, &out.AuditReason
		*out = new(string)
		**out = **in
	}
	if in.DesireValue != nil {
		in, out := &in.DesireValue, &out.DesireValue
		*out = new(float64)
		**out = **in
	}
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = make([]ApplicationInfoSpecDimensions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EffectiveTime != nil {
		in, out := &in.EffectiveTime, &out.EffectiveTime
		*out = new(string)
		**out = **in
	}
	if in.ExpireTime != nil {
		in, out := &in.ExpireTime, &out.ExpireTime
		*out = new(string)
		**out = **in
	}
	if in.NoticeType != nil {
		in, out := &in.NoticeType, &out.NoticeType
		*out = new(int64)
		**out = **in
	}
	if in.ProductCode != nil {
		in, out := &in.ProductCode, &out.ProductCode
		*out = new(string)
		**out = **in
	}
	if in.QuotaActionCode != nil {
		in, out := &in.QuotaActionCode, &out.QuotaActionCode
		*out = new(string)
		**out = **in
	}
	if in.QuotaCategory != nil {
		in, out := &in.QuotaCategory, &out.QuotaCategory
		*out = new(string)
		**out = **in
	}
	if in.QuotaDescription != nil {
		in, out := &in.QuotaDescription, &out.QuotaDescription
		*out = new(string)
		**out = **in
	}
	if in.QuotaName != nil {
		in, out := &in.QuotaName, &out.QuotaName
		*out = new(string)
		**out = **in
	}
	if in.QuotaUnit != nil {
		in, out := &in.QuotaUnit, &out.QuotaUnit
		*out = new(string)
		**out = **in
	}
	if in.Reason != nil {
		in, out := &in.Reason, &out.Reason
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationInfoSpecResource.
func (in *ApplicationInfoSpecResource) DeepCopy() *ApplicationInfoSpecResource {
	if in == nil {
		return nil
	}
	out := new(ApplicationInfoSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationInfoStatus) DeepCopyInto(out *ApplicationInfoStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationInfoStatus.
func (in *ApplicationInfoStatus) DeepCopy() *ApplicationInfoStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationInfoStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaAlarm) DeepCopyInto(out *QuotaAlarm) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaAlarm.
func (in *QuotaAlarm) DeepCopy() *QuotaAlarm {
	if in == nil {
		return nil
	}
	out := new(QuotaAlarm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QuotaAlarm) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaAlarmList) DeepCopyInto(out *QuotaAlarmList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]QuotaAlarm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaAlarmList.
func (in *QuotaAlarmList) DeepCopy() *QuotaAlarmList {
	if in == nil {
		return nil
	}
	out := new(QuotaAlarmList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QuotaAlarmList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaAlarmSpec) DeepCopyInto(out *QuotaAlarmSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(QuotaAlarmSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaAlarmSpec.
func (in *QuotaAlarmSpec) DeepCopy() *QuotaAlarmSpec {
	if in == nil {
		return nil
	}
	out := new(QuotaAlarmSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaAlarmSpecQuotaDimensions) DeepCopyInto(out *QuotaAlarmSpecQuotaDimensions) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaAlarmSpecQuotaDimensions.
func (in *QuotaAlarmSpecQuotaDimensions) DeepCopy() *QuotaAlarmSpecQuotaDimensions {
	if in == nil {
		return nil
	}
	out := new(QuotaAlarmSpecQuotaDimensions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaAlarmSpecResource) DeepCopyInto(out *QuotaAlarmSpecResource) {
	*out = *in
	if in.ProductCode != nil {
		in, out := &in.ProductCode, &out.ProductCode
		*out = new(string)
		**out = **in
	}
	if in.QuotaActionCode != nil {
		in, out := &in.QuotaActionCode, &out.QuotaActionCode
		*out = new(string)
		**out = **in
	}
	if in.QuotaAlarmName != nil {
		in, out := &in.QuotaAlarmName, &out.QuotaAlarmName
		*out = new(string)
		**out = **in
	}
	if in.QuotaDimensions != nil {
		in, out := &in.QuotaDimensions, &out.QuotaDimensions
		*out = make([]QuotaAlarmSpecQuotaDimensions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.ThresholdPercent != nil {
		in, out := &in.ThresholdPercent, &out.ThresholdPercent
		*out = new(float64)
		**out = **in
	}
	if in.WebHook != nil {
		in, out := &in.WebHook, &out.WebHook
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaAlarmSpecResource.
func (in *QuotaAlarmSpecResource) DeepCopy() *QuotaAlarmSpecResource {
	if in == nil {
		return nil
	}
	out := new(QuotaAlarmSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaAlarmStatus) DeepCopyInto(out *QuotaAlarmStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaAlarmStatus.
func (in *QuotaAlarmStatus) DeepCopy() *QuotaAlarmStatus {
	if in == nil {
		return nil
	}
	out := new(QuotaAlarmStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaApplication) DeepCopyInto(out *QuotaApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaApplication.
func (in *QuotaApplication) DeepCopy() *QuotaApplication {
	if in == nil {
		return nil
	}
	out := new(QuotaApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QuotaApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaApplicationList) DeepCopyInto(out *QuotaApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]QuotaApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaApplicationList.
func (in *QuotaApplicationList) DeepCopy() *QuotaApplicationList {
	if in == nil {
		return nil
	}
	out := new(QuotaApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QuotaApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaApplicationSpec) DeepCopyInto(out *QuotaApplicationSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(QuotaApplicationSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaApplicationSpec.
func (in *QuotaApplicationSpec) DeepCopy() *QuotaApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(QuotaApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaApplicationSpecDimensions) DeepCopyInto(out *QuotaApplicationSpecDimensions) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaApplicationSpecDimensions.
func (in *QuotaApplicationSpecDimensions) DeepCopy() *QuotaApplicationSpecDimensions {
	if in == nil {
		return nil
	}
	out := new(QuotaApplicationSpecDimensions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaApplicationSpecResource) DeepCopyInto(out *QuotaApplicationSpecResource) {
	*out = *in
	if in.ApproveValue != nil {
		in, out := &in.ApproveValue, &out.ApproveValue
		*out = new(string)
		**out = **in
	}
	if in.AuditMode != nil {
		in, out := &in.AuditMode, &out.AuditMode
		*out = new(string)
		**out = **in
	}
	if in.AuditReason != nil {
		in, out := &in.AuditReason, &out.AuditReason
		*out = new(string)
		**out = **in
	}
	if in.DesireValue != nil {
		in, out := &in.DesireValue, &out.DesireValue
		*out = new(float64)
		**out = **in
	}
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = make([]QuotaApplicationSpecDimensions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EffectiveTime != nil {
		in, out := &in.EffectiveTime, &out.EffectiveTime
		*out = new(string)
		**out = **in
	}
	if in.ExpireTime != nil {
		in, out := &in.ExpireTime, &out.ExpireTime
		*out = new(string)
		**out = **in
	}
	if in.NoticeType != nil {
		in, out := &in.NoticeType, &out.NoticeType
		*out = new(int64)
		**out = **in
	}
	if in.ProductCode != nil {
		in, out := &in.ProductCode, &out.ProductCode
		*out = new(string)
		**out = **in
	}
	if in.QuotaActionCode != nil {
		in, out := &in.QuotaActionCode, &out.QuotaActionCode
		*out = new(string)
		**out = **in
	}
	if in.QuotaCategory != nil {
		in, out := &in.QuotaCategory, &out.QuotaCategory
		*out = new(string)
		**out = **in
	}
	if in.QuotaDescription != nil {
		in, out := &in.QuotaDescription, &out.QuotaDescription
		*out = new(string)
		**out = **in
	}
	if in.QuotaName != nil {
		in, out := &in.QuotaName, &out.QuotaName
		*out = new(string)
		**out = **in
	}
	if in.QuotaUnit != nil {
		in, out := &in.QuotaUnit, &out.QuotaUnit
		*out = new(string)
		**out = **in
	}
	if in.Reason != nil {
		in, out := &in.Reason, &out.Reason
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaApplicationSpecResource.
func (in *QuotaApplicationSpecResource) DeepCopy() *QuotaApplicationSpecResource {
	if in == nil {
		return nil
	}
	out := new(QuotaApplicationSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaApplicationStatus) DeepCopyInto(out *QuotaApplicationStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaApplicationStatus.
func (in *QuotaApplicationStatus) DeepCopy() *QuotaApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(QuotaApplicationStatus)
	in.DeepCopyInto(out)
	return out
}
