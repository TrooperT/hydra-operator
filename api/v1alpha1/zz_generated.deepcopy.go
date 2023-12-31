//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 TrooperT.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIEndpoint) DeepCopyInto(out *APIEndpoint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIEndpoint.
func (in *APIEndpoint) DeepCopy() *APIEndpoint {
	if in == nil {
		return nil
	}
	out := new(APIEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedApiServer) DeepCopyInto(out *ManagedApiServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedApiServer.
func (in *ManagedApiServer) DeepCopy() *ManagedApiServer {
	if in == nil {
		return nil
	}
	out := new(ManagedApiServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedApiServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedApiServerList) DeepCopyInto(out *ManagedApiServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedApiServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedApiServerList.
func (in *ManagedApiServerList) DeepCopy() *ManagedApiServerList {
	if in == nil {
		return nil
	}
	out := new(ManagedApiServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedApiServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedApiServerSpec) DeepCopyInto(out *ManagedApiServerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedApiServerSpec.
func (in *ManagedApiServerSpec) DeepCopy() *ManagedApiServerSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedApiServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedApiServerStatus) DeepCopyInto(out *ManagedApiServerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedApiServerStatus.
func (in *ManagedApiServerStatus) DeepCopy() *ManagedApiServerStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedApiServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedCloudControllerManager) DeepCopyInto(out *ManagedCloudControllerManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedCloudControllerManager.
func (in *ManagedCloudControllerManager) DeepCopy() *ManagedCloudControllerManager {
	if in == nil {
		return nil
	}
	out := new(ManagedCloudControllerManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedCloudControllerManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedCloudControllerManagerList) DeepCopyInto(out *ManagedCloudControllerManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedCloudControllerManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedCloudControllerManagerList.
func (in *ManagedCloudControllerManagerList) DeepCopy() *ManagedCloudControllerManagerList {
	if in == nil {
		return nil
	}
	out := new(ManagedCloudControllerManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedCloudControllerManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedCloudControllerManagerSpec) DeepCopyInto(out *ManagedCloudControllerManagerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedCloudControllerManagerSpec.
func (in *ManagedCloudControllerManagerSpec) DeepCopy() *ManagedCloudControllerManagerSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedCloudControllerManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedCloudControllerManagerStatus) DeepCopyInto(out *ManagedCloudControllerManagerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedCloudControllerManagerStatus.
func (in *ManagedCloudControllerManagerStatus) DeepCopy() *ManagedCloudControllerManagerStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedCloudControllerManagerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedControlPlane) DeepCopyInto(out *ManagedControlPlane) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedControlPlane.
func (in *ManagedControlPlane) DeepCopy() *ManagedControlPlane {
	if in == nil {
		return nil
	}
	out := new(ManagedControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedControlPlane) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedControlPlaneList) DeepCopyInto(out *ManagedControlPlaneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedControlPlane, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedControlPlaneList.
func (in *ManagedControlPlaneList) DeepCopy() *ManagedControlPlaneList {
	if in == nil {
		return nil
	}
	out := new(ManagedControlPlaneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedControlPlaneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedControlPlaneSpec) DeepCopyInto(out *ManagedControlPlaneSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedControlPlaneSpec.
func (in *ManagedControlPlaneSpec) DeepCopy() *ManagedControlPlaneSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedControlPlaneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedControlPlaneStatus) DeepCopyInto(out *ManagedControlPlaneStatus) {
	*out = *in
	if in.APIEndpoints != nil {
		in, out := &in.APIEndpoints, &out.APIEndpoints
		*out = make([]APIEndpoint, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedControlPlaneStatus.
func (in *ManagedControlPlaneStatus) DeepCopy() *ManagedControlPlaneStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedControlPlaneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedControllerManager) DeepCopyInto(out *ManagedControllerManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedControllerManager.
func (in *ManagedControllerManager) DeepCopy() *ManagedControllerManager {
	if in == nil {
		return nil
	}
	out := new(ManagedControllerManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedControllerManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedControllerManagerList) DeepCopyInto(out *ManagedControllerManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedControllerManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedControllerManagerList.
func (in *ManagedControllerManagerList) DeepCopy() *ManagedControllerManagerList {
	if in == nil {
		return nil
	}
	out := new(ManagedControllerManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedControllerManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedControllerManagerSpec) DeepCopyInto(out *ManagedControllerManagerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedControllerManagerSpec.
func (in *ManagedControllerManagerSpec) DeepCopy() *ManagedControllerManagerSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedControllerManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedControllerManagerStatus) DeepCopyInto(out *ManagedControllerManagerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedControllerManagerStatus.
func (in *ManagedControllerManagerStatus) DeepCopy() *ManagedControllerManagerStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedControllerManagerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedEtcd) DeepCopyInto(out *ManagedEtcd) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedEtcd.
func (in *ManagedEtcd) DeepCopy() *ManagedEtcd {
	if in == nil {
		return nil
	}
	out := new(ManagedEtcd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedEtcd) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedEtcdList) DeepCopyInto(out *ManagedEtcdList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedEtcd, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedEtcdList.
func (in *ManagedEtcdList) DeepCopy() *ManagedEtcdList {
	if in == nil {
		return nil
	}
	out := new(ManagedEtcdList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedEtcdList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedEtcdSpec) DeepCopyInto(out *ManagedEtcdSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedEtcdSpec.
func (in *ManagedEtcdSpec) DeepCopy() *ManagedEtcdSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedEtcdSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedEtcdStatus) DeepCopyInto(out *ManagedEtcdStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedEtcdStatus.
func (in *ManagedEtcdStatus) DeepCopy() *ManagedEtcdStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedEtcdStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedScheduler) DeepCopyInto(out *ManagedScheduler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedScheduler.
func (in *ManagedScheduler) DeepCopy() *ManagedScheduler {
	if in == nil {
		return nil
	}
	out := new(ManagedScheduler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedScheduler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedSchedulerList) DeepCopyInto(out *ManagedSchedulerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedScheduler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedSchedulerList.
func (in *ManagedSchedulerList) DeepCopy() *ManagedSchedulerList {
	if in == nil {
		return nil
	}
	out := new(ManagedSchedulerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedSchedulerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedSchedulerSpec) DeepCopyInto(out *ManagedSchedulerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedSchedulerSpec.
func (in *ManagedSchedulerSpec) DeepCopy() *ManagedSchedulerSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedSchedulerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedSchedulerStatus) DeepCopyInto(out *ManagedSchedulerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedSchedulerStatus.
func (in *ManagedSchedulerStatus) DeepCopy() *ManagedSchedulerStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedSchedulerStatus)
	in.DeepCopyInto(out)
	return out
}
