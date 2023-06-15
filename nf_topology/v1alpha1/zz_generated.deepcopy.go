//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 Nephio.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	nf_deploymentsv1alpha1 "github.com/nephio-project/api/nf_deployments/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFClass) DeepCopyInto(out *NFClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFClass.
func (in *NFClass) DeepCopy() *NFClass {
	if in == nil {
		return nil
	}
	out := new(NFClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NFClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFClassList) DeepCopyInto(out *NFClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NFClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFClassList.
func (in *NFClassList) DeepCopy() *NFClassList {
	if in == nil {
		return nil
	}
	out := new(NFClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NFClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFClassSpec) DeepCopyInto(out *NFClassSpec) {
	*out = *in
	out.PackageRef = in.PackageRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFClassSpec.
func (in *NFClassSpec) DeepCopy() *NFClassSpec {
	if in == nil {
		return nil
	}
	out := new(NFClassSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFClassStatus) DeepCopyInto(out *NFClassStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFClassStatus.
func (in *NFClassStatus) DeepCopy() *NFClassStatus {
	if in == nil {
		return nil
	}
	out := new(NFClassStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFConnectivity) DeepCopyInto(out *NFConnectivity) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFConnectivity.
func (in *NFConnectivity) DeepCopy() *NFConnectivity {
	if in == nil {
		return nil
	}
	out := new(NFConnectivity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFInstance) DeepCopyInto(out *NFInstance) {
	*out = *in
	in.ClusterSelector.DeepCopyInto(&out.ClusterSelector)
	in.NFTemplate.DeepCopyInto(&out.NFTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFInstance.
func (in *NFInstance) DeepCopy() *NFInstance {
	if in == nil {
		return nil
	}
	out := new(NFInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFInterface) DeepCopyInto(out *NFInterface) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFInterface.
func (in *NFInterface) DeepCopy() *NFInterface {
	if in == nil {
		return nil
	}
	out := new(NFInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFTemplate) DeepCopyInto(out *NFTemplate) {
	*out = *in
	in.Capacity.DeepCopyInto(&out.Capacity)
	if in.NFInterfaces != nil {
		in, out := &in.NFInterfaces, &out.NFInterfaces
		*out = make([]NFInterface, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFTemplate.
func (in *NFTemplate) DeepCopy() *NFTemplate {
	if in == nil {
		return nil
	}
	out := new(NFTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFTopoTracker) DeepCopyInto(out *NFTopoTracker) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFTopoTracker.
func (in *NFTopoTracker) DeepCopy() *NFTopoTracker {
	if in == nil {
		return nil
	}
	out := new(NFTopoTracker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NFTopoTracker) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFTopoTrackerInstance) DeepCopyInto(out *NFTopoTrackerInstance) {
	*out = *in
	if in.Connectivities != nil {
		in, out := &in.Connectivities, &out.Connectivities
		*out = make([]NFConnectivity, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFTopoTrackerInstance.
func (in *NFTopoTrackerInstance) DeepCopy() *NFTopoTrackerInstance {
	if in == nil {
		return nil
	}
	out := new(NFTopoTrackerInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFTopoTrackerList) DeepCopyInto(out *NFTopoTrackerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NFTopoTracker, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFTopoTrackerList.
func (in *NFTopoTrackerList) DeepCopy() *NFTopoTrackerList {
	if in == nil {
		return nil
	}
	out := new(NFTopoTrackerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NFTopoTrackerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFTopoTrackerSpec) DeepCopyInto(out *NFTopoTrackerSpec) {
	*out = *in
	if in.NFInstances != nil {
		in, out := &in.NFInstances, &out.NFInstances
		*out = make([]NFTopoTrackerInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFTopoTrackerSpec.
func (in *NFTopoTrackerSpec) DeepCopy() *NFTopoTrackerSpec {
	if in == nil {
		return nil
	}
	out := new(NFTopoTrackerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFTopoTrackerStatus) DeepCopyInto(out *NFTopoTrackerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]nf_deploymentsv1alpha1.NFDeploymentConditionType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFTopoTrackerStatus.
func (in *NFTopoTrackerStatus) DeepCopy() *NFTopoTrackerStatus {
	if in == nil {
		return nil
	}
	out := new(NFTopoTrackerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFTopology) DeepCopyInto(out *NFTopology) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFTopology.
func (in *NFTopology) DeepCopy() *NFTopology {
	if in == nil {
		return nil
	}
	out := new(NFTopology)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NFTopology) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFTopologyList) DeepCopyInto(out *NFTopologyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NFTopology, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFTopologyList.
func (in *NFTopologyList) DeepCopy() *NFTopologyList {
	if in == nil {
		return nil
	}
	out := new(NFTopologyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NFTopologyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFTopologySpec) DeepCopyInto(out *NFTopologySpec) {
	*out = *in
	if in.NFInstances != nil {
		in, out := &in.NFInstances, &out.NFInstances
		*out = make([]NFInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFTopologySpec.
func (in *NFTopologySpec) DeepCopy() *NFTopologySpec {
	if in == nil {
		return nil
	}
	out := new(NFTopologySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFTopologyStatus) DeepCopyInto(out *NFTopologyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFTopologyStatus.
func (in *NFTopologyStatus) DeepCopy() *NFTopologyStatus {
	if in == nil {
		return nil
	}
	out := new(NFTopologyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageRevisionReference) DeepCopyInto(out *PackageRevisionReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageRevisionReference.
func (in *PackageRevisionReference) DeepCopy() *PackageRevisionReference {
	if in == nil {
		return nil
	}
	out := new(PackageRevisionReference)
	in.DeepCopyInto(out)
	return out
}