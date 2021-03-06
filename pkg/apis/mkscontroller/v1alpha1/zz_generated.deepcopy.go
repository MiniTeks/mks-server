//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2020 The Knative Authors

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
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArrayOrString) DeepCopyInto(out *ArrayOrString) {
	*out = *in
	if in.ArrayVal != nil {
		in, out := &in.ArrayVal, &out.ArrayVal
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArrayOrString.
func (in *ArrayOrString) DeepCopy() *ArrayOrString {
	if in == nil {
		return nil
	}
	out := new(ArrayOrString)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MksConfigMap) DeepCopyInto(out *MksConfigMap) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MksConfigMap.
func (in *MksConfigMap) DeepCopy() *MksConfigMap {
	if in == nil {
		return nil
	}
	out := new(MksConfigMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MksParam) DeepCopyInto(out *MksParam) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MksParam.
func (in *MksParam) DeepCopy() *MksParam {
	if in == nil {
		return nil
	}
	out := new(MksParam)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MksParamSpec) DeepCopyInto(out *MksParamSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MksParamSpec.
func (in *MksParamSpec) DeepCopy() *MksParamSpec {
	if in == nil {
		return nil
	}
	out := new(MksParamSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MksPersistentVolumeClaim) DeepCopyInto(out *MksPersistentVolumeClaim) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MksPersistentVolumeClaim.
func (in *MksPersistentVolumeClaim) DeepCopy() *MksPersistentVolumeClaim {
	if in == nil {
		return nil
	}
	out := new(MksPersistentVolumeClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MksPipeline) DeepCopyInto(out *MksPipeline) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MksPipeline.
func (in *MksPipeline) DeepCopy() *MksPipeline {
	if in == nil {
		return nil
	}
	out := new(MksPipeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MksPipeline) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MksPipelineList) DeepCopyInto(out *MksPipelineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MksPipeline, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MksPipelineList.
func (in *MksPipelineList) DeepCopy() *MksPipelineList {
	if in == nil {
		return nil
	}
	out := new(MksPipelineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MksPipelineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MksPipelineRun) DeepCopyInto(out *MksPipelineRun) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MksPipelineRun.
func (in *MksPipelineRun) DeepCopy() *MksPipelineRun {
	if in == nil {
		return nil
	}
	out := new(MksPipelineRun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MksPipelineRun) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MksPipelineRunList) DeepCopyInto(out *MksPipelineRunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MksPipelineRun, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MksPipelineRunList.
func (in *MksPipelineRunList) DeepCopy() *MksPipelineRunList {
	if in == nil {
		return nil
	}
	out := new(MksPipelineRunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MksPipelineRunList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MksPipelineRunRef) DeepCopyInto(out *MksPipelineRunRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MksPipelineRunRef.
func (in *MksPipelineRunRef) DeepCopy() *MksPipelineRunRef {
	if in == nil {
		return nil
	}
	out := new(MksPipelineRunRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MksPipelineRunSpec) DeepCopyInto(out *MksPipelineRunSpec) {
	*out = *in
	out.PipelineRef = in.PipelineRef
	if in.Workspaces != nil {
		in, out := &in.Workspaces, &out.Workspaces
		*out = make([]MksprWorkspaces, len(*in))
		copy(*out, *in)
	}
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make([]MksParam, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MksPipelineRunSpec.
func (in *MksPipelineRunSpec) DeepCopy() *MksPipelineRunSpec {
	if in == nil {
		return nil
	}
	out := new(MksPipelineRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MksPipelineSpec) DeepCopyInto(out *MksPipelineSpec) {
	*out = *in
	if in.Workspaces != nil {
		in, out := &in.Workspaces, &out.Workspaces
		*out = make([]MksPipelineWorkspace, len(*in))
		copy(*out, *in)
	}
	if in.Param != nil {
		in, out := &in.Param, &out.Param
		*out = make([]ParamSpec, len(*in))
		copy(*out, *in)
	}
	if in.Task != nil {
		in, out := &in.Task, &out.Task
		*out = make([]MksPipelineTask, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MksPipelineSpec.
func (in *MksPipelineSpec) DeepCopy() *MksPipelineSpec {
	if in == nil {
		return nil
	}
	out := new(MksPipelineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MksPipelineTask) DeepCopyInto(out *MksPipelineTask) {
	*out = *in
	out.PipelineTaskRef = in.PipelineTaskRef
	if in.Workspaces != nil {
		in, out := &in.Workspaces, &out.Workspaces
		*out = make([]MksPipelineWorkspace, len(*in))
		copy(*out, *in)
	}
	if in.Param != nil {
		in, out := &in.Param, &out.Param
		*out = make([]Param, len(*in))
		copy(*out, *in)
	}
	if in.RunAfter != nil {
		in, out := &in.RunAfter, &out.RunAfter
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MksPipelineTask.
func (in *MksPipelineTask) DeepCopy() *MksPipelineTask {
	if in == nil {
		return nil
	}
	out := new(MksPipelineTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MksPipelineTaskRef) DeepCopyInto(out *MksPipelineTaskRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MksPipelineTaskRef.
func (in *MksPipelineTaskRef) DeepCopy() *MksPipelineTaskRef {
	if in == nil {
		return nil
	}
	out := new(MksPipelineTaskRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MksPipelineWorkspace) DeepCopyInto(out *MksPipelineWorkspace) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MksPipelineWorkspace.
func (in *MksPipelineWorkspace) DeepCopy() *MksPipelineWorkspace {
	if in == nil {
		return nil
	}
	out := new(MksPipelineWorkspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MksTask) DeepCopyInto(out *MksTask) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MksTask.
func (in *MksTask) DeepCopy() *MksTask {
	if in == nil {
		return nil
	}
	out := new(MksTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MksTask) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MksTaskList) DeepCopyInto(out *MksTaskList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MksTask, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MksTaskList.
func (in *MksTaskList) DeepCopy() *MksTaskList {
	if in == nil {
		return nil
	}
	out := new(MksTaskList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MksTaskList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MksTaskRef) DeepCopyInto(out *MksTaskRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MksTaskRef.
func (in *MksTaskRef) DeepCopy() *MksTaskRef {
	if in == nil {
		return nil
	}
	out := new(MksTaskRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MksTaskRun) DeepCopyInto(out *MksTaskRun) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MksTaskRun.
func (in *MksTaskRun) DeepCopy() *MksTaskRun {
	if in == nil {
		return nil
	}
	out := new(MksTaskRun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MksTaskRun) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MksTaskRunList) DeepCopyInto(out *MksTaskRunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MksTaskRun, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MksTaskRunList.
func (in *MksTaskRunList) DeepCopy() *MksTaskRunList {
	if in == nil {
		return nil
	}
	out := new(MksTaskRunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MksTaskRunList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MksTaskRunSpec) DeepCopyInto(out *MksTaskRunSpec) {
	*out = *in
	out.TaskRef = in.TaskRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MksTaskRunSpec.
func (in *MksTaskRunSpec) DeepCopy() *MksTaskRunSpec {
	if in == nil {
		return nil
	}
	out := new(MksTaskRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MksTaskSpec) DeepCopyInto(out *MksTaskSpec) {
	*out = *in
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]MksTaskSteps, len(*in))
		copy(*out, *in)
	}
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make([]MksParamSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MksTaskSpec.
func (in *MksTaskSpec) DeepCopy() *MksTaskSpec {
	if in == nil {
		return nil
	}
	out := new(MksTaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MksTaskSteps) DeepCopyInto(out *MksTaskSteps) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MksTaskSteps.
func (in *MksTaskSteps) DeepCopy() *MksTaskSteps {
	if in == nil {
		return nil
	}
	out := new(MksTaskSteps)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MksprWorkspaces) DeepCopyInto(out *MksprWorkspaces) {
	*out = *in
	out.PersistentVolumeClaim = in.PersistentVolumeClaim
	out.ConfigMap = in.ConfigMap
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MksprWorkspaces.
func (in *MksprWorkspaces) DeepCopy() *MksprWorkspaces {
	if in == nil {
		return nil
	}
	out := new(MksprWorkspaces)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Param) DeepCopyInto(out *Param) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Param.
func (in *Param) DeepCopy() *Param {
	if in == nil {
		return nil
	}
	out := new(Param)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParamSpec) DeepCopyInto(out *ParamSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParamSpec.
func (in *ParamSpec) DeepCopy() *ParamSpec {
	if in == nil {
		return nil
	}
	out := new(ParamSpec)
	in.DeepCopyInto(out)
	return out
}
