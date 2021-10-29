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
func (in *ProviderAccess) DeepCopyInto(out *ProviderAccess) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderAccess.
func (in *ProviderAccess) DeepCopy() *ProviderAccess {
	if in == nil {
		return nil
	}
	out := new(ProviderAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProviderAccess) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderAccessAuthorization) DeepCopyInto(out *ProviderAccessAuthorization) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderAccessAuthorization.
func (in *ProviderAccessAuthorization) DeepCopy() *ProviderAccessAuthorization {
	if in == nil {
		return nil
	}
	out := new(ProviderAccessAuthorization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProviderAccessAuthorization) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderAccessAuthorizationList) DeepCopyInto(out *ProviderAccessAuthorizationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProviderAccessAuthorization, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderAccessAuthorizationList.
func (in *ProviderAccessAuthorizationList) DeepCopy() *ProviderAccessAuthorizationList {
	if in == nil {
		return nil
	}
	out := new(ProviderAccessAuthorizationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProviderAccessAuthorizationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderAccessAuthorizationSpec) DeepCopyInto(out *ProviderAccessAuthorizationSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ProviderAccessAuthorizationSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderAccessAuthorizationSpec.
func (in *ProviderAccessAuthorizationSpec) DeepCopy() *ProviderAccessAuthorizationSpec {
	if in == nil {
		return nil
	}
	out := new(ProviderAccessAuthorizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderAccessAuthorizationSpecAws) DeepCopyInto(out *ProviderAccessAuthorizationSpecAws) {
	*out = *in
	if in.IamAssumedRoleArn != nil {
		in, out := &in.IamAssumedRoleArn, &out.IamAssumedRoleArn
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderAccessAuthorizationSpecAws.
func (in *ProviderAccessAuthorizationSpecAws) DeepCopy() *ProviderAccessAuthorizationSpecAws {
	if in == nil {
		return nil
	}
	out := new(ProviderAccessAuthorizationSpecAws)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderAccessAuthorizationSpecFeatureUsages) DeepCopyInto(out *ProviderAccessAuthorizationSpecFeatureUsages) {
	*out = *in
	if in.FeatureID != nil {
		in, out := &in.FeatureID, &out.FeatureID
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.FeatureType != nil {
		in, out := &in.FeatureType, &out.FeatureType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderAccessAuthorizationSpecFeatureUsages.
func (in *ProviderAccessAuthorizationSpecFeatureUsages) DeepCopy() *ProviderAccessAuthorizationSpecFeatureUsages {
	if in == nil {
		return nil
	}
	out := new(ProviderAccessAuthorizationSpecFeatureUsages)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderAccessAuthorizationSpecResource) DeepCopyInto(out *ProviderAccessAuthorizationSpecResource) {
	*out = *in
	if in.AuthorizedDate != nil {
		in, out := &in.AuthorizedDate, &out.AuthorizedDate
		*out = new(string)
		**out = **in
	}
	if in.Aws != nil {
		in, out := &in.Aws, &out.Aws
		*out = new(ProviderAccessAuthorizationSpecAws)
		(*in).DeepCopyInto(*out)
	}
	if in.FeatureUsages != nil {
		in, out := &in.FeatureUsages, &out.FeatureUsages
		*out = make([]ProviderAccessAuthorizationSpecFeatureUsages, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.RoleID != nil {
		in, out := &in.RoleID, &out.RoleID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderAccessAuthorizationSpecResource.
func (in *ProviderAccessAuthorizationSpecResource) DeepCopy() *ProviderAccessAuthorizationSpecResource {
	if in == nil {
		return nil
	}
	out := new(ProviderAccessAuthorizationSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderAccessAuthorizationStatus) DeepCopyInto(out *ProviderAccessAuthorizationStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderAccessAuthorizationStatus.
func (in *ProviderAccessAuthorizationStatus) DeepCopy() *ProviderAccessAuthorizationStatus {
	if in == nil {
		return nil
	}
	out := new(ProviderAccessAuthorizationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderAccessList) DeepCopyInto(out *ProviderAccessList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProviderAccess, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderAccessList.
func (in *ProviderAccessList) DeepCopy() *ProviderAccessList {
	if in == nil {
		return nil
	}
	out := new(ProviderAccessList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProviderAccessList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderAccessSetup) DeepCopyInto(out *ProviderAccessSetup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderAccessSetup.
func (in *ProviderAccessSetup) DeepCopy() *ProviderAccessSetup {
	if in == nil {
		return nil
	}
	out := new(ProviderAccessSetup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProviderAccessSetup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderAccessSetupList) DeepCopyInto(out *ProviderAccessSetupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProviderAccessSetup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderAccessSetupList.
func (in *ProviderAccessSetupList) DeepCopy() *ProviderAccessSetupList {
	if in == nil {
		return nil
	}
	out := new(ProviderAccessSetupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProviderAccessSetupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderAccessSetupSpec) DeepCopyInto(out *ProviderAccessSetupSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ProviderAccessSetupSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderAccessSetupSpec.
func (in *ProviderAccessSetupSpec) DeepCopy() *ProviderAccessSetupSpec {
	if in == nil {
		return nil
	}
	out := new(ProviderAccessSetupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderAccessSetupSpecAwsConfig) DeepCopyInto(out *ProviderAccessSetupSpecAwsConfig) {
	*out = *in
	if in.AtlasAssumedRoleExternalID != nil {
		in, out := &in.AtlasAssumedRoleExternalID, &out.AtlasAssumedRoleExternalID
		*out = new(string)
		**out = **in
	}
	if in.AtlasAwsAccountArn != nil {
		in, out := &in.AtlasAwsAccountArn, &out.AtlasAwsAccountArn
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderAccessSetupSpecAwsConfig.
func (in *ProviderAccessSetupSpecAwsConfig) DeepCopy() *ProviderAccessSetupSpecAwsConfig {
	if in == nil {
		return nil
	}
	out := new(ProviderAccessSetupSpecAwsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderAccessSetupSpecResource) DeepCopyInto(out *ProviderAccessSetupSpecResource) {
	*out = *in
	if in.Aws != nil {
		in, out := &in.Aws, &out.Aws
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.AwsConfig != nil {
		in, out := &in.AwsConfig, &out.AwsConfig
		*out = make([]ProviderAccessSetupSpecAwsConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CreatedDate != nil {
		in, out := &in.CreatedDate, &out.CreatedDate
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.ProviderName != nil {
		in, out := &in.ProviderName, &out.ProviderName
		*out = new(string)
		**out = **in
	}
	if in.RoleID != nil {
		in, out := &in.RoleID, &out.RoleID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderAccessSetupSpecResource.
func (in *ProviderAccessSetupSpecResource) DeepCopy() *ProviderAccessSetupSpecResource {
	if in == nil {
		return nil
	}
	out := new(ProviderAccessSetupSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderAccessSetupStatus) DeepCopyInto(out *ProviderAccessSetupStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderAccessSetupStatus.
func (in *ProviderAccessSetupStatus) DeepCopy() *ProviderAccessSetupStatus {
	if in == nil {
		return nil
	}
	out := new(ProviderAccessSetupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderAccessSpec) DeepCopyInto(out *ProviderAccessSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ProviderAccessSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderAccessSpec.
func (in *ProviderAccessSpec) DeepCopy() *ProviderAccessSpec {
	if in == nil {
		return nil
	}
	out := new(ProviderAccessSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderAccessSpecFeatureUsages) DeepCopyInto(out *ProviderAccessSpecFeatureUsages) {
	*out = *in
	if in.FeatureID != nil {
		in, out := &in.FeatureID, &out.FeatureID
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.FeatureType != nil {
		in, out := &in.FeatureType, &out.FeatureType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderAccessSpecFeatureUsages.
func (in *ProviderAccessSpecFeatureUsages) DeepCopy() *ProviderAccessSpecFeatureUsages {
	if in == nil {
		return nil
	}
	out := new(ProviderAccessSpecFeatureUsages)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderAccessSpecResource) DeepCopyInto(out *ProviderAccessSpecResource) {
	*out = *in
	if in.AtlasAssumedRoleExternalID != nil {
		in, out := &in.AtlasAssumedRoleExternalID, &out.AtlasAssumedRoleExternalID
		*out = new(string)
		**out = **in
	}
	if in.AtlasAwsAccountArn != nil {
		in, out := &in.AtlasAwsAccountArn, &out.AtlasAwsAccountArn
		*out = new(string)
		**out = **in
	}
	if in.AuthorizedDate != nil {
		in, out := &in.AuthorizedDate, &out.AuthorizedDate
		*out = new(string)
		**out = **in
	}
	if in.CreatedDate != nil {
		in, out := &in.CreatedDate, &out.CreatedDate
		*out = new(string)
		**out = **in
	}
	if in.FeatureUsages != nil {
		in, out := &in.FeatureUsages, &out.FeatureUsages
		*out = make([]ProviderAccessSpecFeatureUsages, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IamAssumedRoleArn != nil {
		in, out := &in.IamAssumedRoleArn, &out.IamAssumedRoleArn
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.ProviderName != nil {
		in, out := &in.ProviderName, &out.ProviderName
		*out = new(string)
		**out = **in
	}
	if in.RoleID != nil {
		in, out := &in.RoleID, &out.RoleID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderAccessSpecResource.
func (in *ProviderAccessSpecResource) DeepCopy() *ProviderAccessSpecResource {
	if in == nil {
		return nil
	}
	out := new(ProviderAccessSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderAccessStatus) DeepCopyInto(out *ProviderAccessStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderAccessStatus.
func (in *ProviderAccessStatus) DeepCopy() *ProviderAccessStatus {
	if in == nil {
		return nil
	}
	out := new(ProviderAccessStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSnapshot) DeepCopyInto(out *ProviderSnapshot) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSnapshot.
func (in *ProviderSnapshot) DeepCopy() *ProviderSnapshot {
	if in == nil {
		return nil
	}
	out := new(ProviderSnapshot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProviderSnapshot) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSnapshotBackupPolicy) DeepCopyInto(out *ProviderSnapshotBackupPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSnapshotBackupPolicy.
func (in *ProviderSnapshotBackupPolicy) DeepCopy() *ProviderSnapshotBackupPolicy {
	if in == nil {
		return nil
	}
	out := new(ProviderSnapshotBackupPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProviderSnapshotBackupPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSnapshotBackupPolicyList) DeepCopyInto(out *ProviderSnapshotBackupPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProviderSnapshotBackupPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSnapshotBackupPolicyList.
func (in *ProviderSnapshotBackupPolicyList) DeepCopy() *ProviderSnapshotBackupPolicyList {
	if in == nil {
		return nil
	}
	out := new(ProviderSnapshotBackupPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProviderSnapshotBackupPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSnapshotBackupPolicySpec) DeepCopyInto(out *ProviderSnapshotBackupPolicySpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ProviderSnapshotBackupPolicySpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSnapshotBackupPolicySpec.
func (in *ProviderSnapshotBackupPolicySpec) DeepCopy() *ProviderSnapshotBackupPolicySpec {
	if in == nil {
		return nil
	}
	out := new(ProviderSnapshotBackupPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSnapshotBackupPolicySpecPolicies) DeepCopyInto(out *ProviderSnapshotBackupPolicySpecPolicies) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.PolicyItem != nil {
		in, out := &in.PolicyItem, &out.PolicyItem
		*out = make([]ProviderSnapshotBackupPolicySpecPoliciesPolicyItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSnapshotBackupPolicySpecPolicies.
func (in *ProviderSnapshotBackupPolicySpecPolicies) DeepCopy() *ProviderSnapshotBackupPolicySpecPolicies {
	if in == nil {
		return nil
	}
	out := new(ProviderSnapshotBackupPolicySpecPolicies)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSnapshotBackupPolicySpecPoliciesPolicyItem) DeepCopyInto(out *ProviderSnapshotBackupPolicySpecPoliciesPolicyItem) {
	*out = *in
	if in.FrequencyInterval != nil {
		in, out := &in.FrequencyInterval, &out.FrequencyInterval
		*out = new(int64)
		**out = **in
	}
	if in.FrequencyType != nil {
		in, out := &in.FrequencyType, &out.FrequencyType
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RetentionUnit != nil {
		in, out := &in.RetentionUnit, &out.RetentionUnit
		*out = new(string)
		**out = **in
	}
	if in.RetentionValue != nil {
		in, out := &in.RetentionValue, &out.RetentionValue
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSnapshotBackupPolicySpecPoliciesPolicyItem.
func (in *ProviderSnapshotBackupPolicySpecPoliciesPolicyItem) DeepCopy() *ProviderSnapshotBackupPolicySpecPoliciesPolicyItem {
	if in == nil {
		return nil
	}
	out := new(ProviderSnapshotBackupPolicySpecPoliciesPolicyItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSnapshotBackupPolicySpecResource) DeepCopyInto(out *ProviderSnapshotBackupPolicySpecResource) {
	*out = *in
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.NextSnapshot != nil {
		in, out := &in.NextSnapshot, &out.NextSnapshot
		*out = new(string)
		**out = **in
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]ProviderSnapshotBackupPolicySpecPolicies, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.ReferenceHourOfDay != nil {
		in, out := &in.ReferenceHourOfDay, &out.ReferenceHourOfDay
		*out = new(int64)
		**out = **in
	}
	if in.ReferenceMinuteOfHour != nil {
		in, out := &in.ReferenceMinuteOfHour, &out.ReferenceMinuteOfHour
		*out = new(int64)
		**out = **in
	}
	if in.RestoreWindowDays != nil {
		in, out := &in.RestoreWindowDays, &out.RestoreWindowDays
		*out = new(int64)
		**out = **in
	}
	if in.UpdateSnapshots != nil {
		in, out := &in.UpdateSnapshots, &out.UpdateSnapshots
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSnapshotBackupPolicySpecResource.
func (in *ProviderSnapshotBackupPolicySpecResource) DeepCopy() *ProviderSnapshotBackupPolicySpecResource {
	if in == nil {
		return nil
	}
	out := new(ProviderSnapshotBackupPolicySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSnapshotBackupPolicyStatus) DeepCopyInto(out *ProviderSnapshotBackupPolicyStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSnapshotBackupPolicyStatus.
func (in *ProviderSnapshotBackupPolicyStatus) DeepCopy() *ProviderSnapshotBackupPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(ProviderSnapshotBackupPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSnapshotList) DeepCopyInto(out *ProviderSnapshotList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProviderSnapshot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSnapshotList.
func (in *ProviderSnapshotList) DeepCopy() *ProviderSnapshotList {
	if in == nil {
		return nil
	}
	out := new(ProviderSnapshotList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProviderSnapshotList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSnapshotRestoreJob) DeepCopyInto(out *ProviderSnapshotRestoreJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSnapshotRestoreJob.
func (in *ProviderSnapshotRestoreJob) DeepCopy() *ProviderSnapshotRestoreJob {
	if in == nil {
		return nil
	}
	out := new(ProviderSnapshotRestoreJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProviderSnapshotRestoreJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSnapshotRestoreJobList) DeepCopyInto(out *ProviderSnapshotRestoreJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProviderSnapshotRestoreJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSnapshotRestoreJobList.
func (in *ProviderSnapshotRestoreJobList) DeepCopy() *ProviderSnapshotRestoreJobList {
	if in == nil {
		return nil
	}
	out := new(ProviderSnapshotRestoreJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProviderSnapshotRestoreJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSnapshotRestoreJobSpec) DeepCopyInto(out *ProviderSnapshotRestoreJobSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ProviderSnapshotRestoreJobSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSnapshotRestoreJobSpec.
func (in *ProviderSnapshotRestoreJobSpec) DeepCopy() *ProviderSnapshotRestoreJobSpec {
	if in == nil {
		return nil
	}
	out := new(ProviderSnapshotRestoreJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSnapshotRestoreJobSpecDeliveryTypeConfig) DeepCopyInto(out *ProviderSnapshotRestoreJobSpecDeliveryTypeConfig) {
	*out = *in
	if in.Automated != nil {
		in, out := &in.Automated, &out.Automated
		*out = new(bool)
		**out = **in
	}
	if in.Download != nil {
		in, out := &in.Download, &out.Download
		*out = new(bool)
		**out = **in
	}
	if in.OplogInc != nil {
		in, out := &in.OplogInc, &out.OplogInc
		*out = new(int64)
		**out = **in
	}
	if in.OplogTs != nil {
		in, out := &in.OplogTs, &out.OplogTs
		*out = new(int64)
		**out = **in
	}
	if in.PointInTime != nil {
		in, out := &in.PointInTime, &out.PointInTime
		*out = new(bool)
		**out = **in
	}
	if in.PointInTimeUtcSeconds != nil {
		in, out := &in.PointInTimeUtcSeconds, &out.PointInTimeUtcSeconds
		*out = new(int64)
		**out = **in
	}
	if in.TargetClusterName != nil {
		in, out := &in.TargetClusterName, &out.TargetClusterName
		*out = new(string)
		**out = **in
	}
	if in.TargetProjectID != nil {
		in, out := &in.TargetProjectID, &out.TargetProjectID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSnapshotRestoreJobSpecDeliveryTypeConfig.
func (in *ProviderSnapshotRestoreJobSpecDeliveryTypeConfig) DeepCopy() *ProviderSnapshotRestoreJobSpecDeliveryTypeConfig {
	if in == nil {
		return nil
	}
	out := new(ProviderSnapshotRestoreJobSpecDeliveryTypeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSnapshotRestoreJobSpecResource) DeepCopyInto(out *ProviderSnapshotRestoreJobSpecResource) {
	*out = *in
	if in.Cancelled != nil {
		in, out := &in.Cancelled, &out.Cancelled
		*out = new(bool)
		**out = **in
	}
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.DeliveryType != nil {
		in, out := &in.DeliveryType, &out.DeliveryType
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.DeliveryTypeConfig != nil {
		in, out := &in.DeliveryTypeConfig, &out.DeliveryTypeConfig
		*out = new(ProviderSnapshotRestoreJobSpecDeliveryTypeConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.DeliveryURL != nil {
		in, out := &in.DeliveryURL, &out.DeliveryURL
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Expired != nil {
		in, out := &in.Expired, &out.Expired
		*out = new(bool)
		**out = **in
	}
	if in.ExpiresAt != nil {
		in, out := &in.ExpiresAt, &out.ExpiresAt
		*out = new(string)
		**out = **in
	}
	if in.FinishedAt != nil {
		in, out := &in.FinishedAt, &out.FinishedAt
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.SnapshotID != nil {
		in, out := &in.SnapshotID, &out.SnapshotID
		*out = new(string)
		**out = **in
	}
	if in.SnapshotRestoreJobID != nil {
		in, out := &in.SnapshotRestoreJobID, &out.SnapshotRestoreJobID
		*out = new(string)
		**out = **in
	}
	if in.Timestamp != nil {
		in, out := &in.Timestamp, &out.Timestamp
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSnapshotRestoreJobSpecResource.
func (in *ProviderSnapshotRestoreJobSpecResource) DeepCopy() *ProviderSnapshotRestoreJobSpecResource {
	if in == nil {
		return nil
	}
	out := new(ProviderSnapshotRestoreJobSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSnapshotRestoreJobStatus) DeepCopyInto(out *ProviderSnapshotRestoreJobStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSnapshotRestoreJobStatus.
func (in *ProviderSnapshotRestoreJobStatus) DeepCopy() *ProviderSnapshotRestoreJobStatus {
	if in == nil {
		return nil
	}
	out := new(ProviderSnapshotRestoreJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSnapshotSpec) DeepCopyInto(out *ProviderSnapshotSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ProviderSnapshotSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSnapshotSpec.
func (in *ProviderSnapshotSpec) DeepCopy() *ProviderSnapshotSpec {
	if in == nil {
		return nil
	}
	out := new(ProviderSnapshotSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSnapshotSpecResource) DeepCopyInto(out *ProviderSnapshotSpecResource) {
	*out = *in
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ExpiresAt != nil {
		in, out := &in.ExpiresAt, &out.ExpiresAt
		*out = new(string)
		**out = **in
	}
	if in.MasterKeyUUID != nil {
		in, out := &in.MasterKeyUUID, &out.MasterKeyUUID
		*out = new(string)
		**out = **in
	}
	if in.MongodVersion != nil {
		in, out := &in.MongodVersion, &out.MongodVersion
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.RetentionInDays != nil {
		in, out := &in.RetentionInDays, &out.RetentionInDays
		*out = new(int64)
		**out = **in
	}
	if in.SnapshotID != nil {
		in, out := &in.SnapshotID, &out.SnapshotID
		*out = new(string)
		**out = **in
	}
	if in.SnapshotType != nil {
		in, out := &in.SnapshotType, &out.SnapshotType
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.StorageSizeBytes != nil {
		in, out := &in.StorageSizeBytes, &out.StorageSizeBytes
		*out = new(int64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSnapshotSpecResource.
func (in *ProviderSnapshotSpecResource) DeepCopy() *ProviderSnapshotSpecResource {
	if in == nil {
		return nil
	}
	out := new(ProviderSnapshotSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSnapshotStatus) DeepCopyInto(out *ProviderSnapshotStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSnapshotStatus.
func (in *ProviderSnapshotStatus) DeepCopy() *ProviderSnapshotStatus {
	if in == nil {
		return nil
	}
	out := new(ProviderSnapshotStatus)
	in.DeepCopyInto(out)
	return out
}
