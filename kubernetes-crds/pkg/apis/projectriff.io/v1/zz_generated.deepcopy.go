// +build !ignore_autogenerated

/*
 * Copyright 2018 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Function).DeepCopyInto(out.(*Function))
			return nil
		}, InType: reflect.TypeOf(&Function{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*FunctionList).DeepCopyInto(out.(*FunctionList))
			return nil
		}, InType: reflect.TypeOf(&FunctionList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*FunctionSpec).DeepCopyInto(out.(*FunctionSpec))
			return nil
		}, InType: reflect.TypeOf(&FunctionSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*FunctionStatus).DeepCopyInto(out.(*FunctionStatus))
			return nil
		}, InType: reflect.TypeOf(&FunctionStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Invoker).DeepCopyInto(out.(*Invoker))
			return nil
		}, InType: reflect.TypeOf(&Invoker{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*InvokerFile).DeepCopyInto(out.(*InvokerFile))
			return nil
		}, InType: reflect.TypeOf(&InvokerFile{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*InvokerHandler).DeepCopyInto(out.(*InvokerHandler))
			return nil
		}, InType: reflect.TypeOf(&InvokerHandler{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*InvokerList).DeepCopyInto(out.(*InvokerList))
			return nil
		}, InType: reflect.TypeOf(&InvokerList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*InvokerProperties).DeepCopyInto(out.(*InvokerProperties))
			return nil
		}, InType: reflect.TypeOf(&InvokerProperties{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*InvokerSpec).DeepCopyInto(out.(*InvokerSpec))
			return nil
		}, InType: reflect.TypeOf(&InvokerSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*InvokerStatus).DeepCopyInto(out.(*InvokerStatus))
			return nil
		}, InType: reflect.TypeOf(&InvokerStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Topic).DeepCopyInto(out.(*Topic))
			return nil
		}, InType: reflect.TypeOf(&Topic{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*TopicList).DeepCopyInto(out.(*TopicList))
			return nil
		}, InType: reflect.TypeOf(&TopicList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*TopicSpec).DeepCopyInto(out.(*TopicSpec))
			return nil
		}, InType: reflect.TypeOf(&TopicSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*TopicStatus).DeepCopyInto(out.(*TopicStatus))
			return nil
		}, InType: reflect.TypeOf(&TopicStatus{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Function) DeepCopyInto(out *Function) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		if *in == nil {
			*out = nil
		} else {
			*out = new(FunctionStatus)
			**out = **in
		}
	}
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
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionList) DeepCopyInto(out *FunctionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
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
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionSpec) DeepCopyInto(out *FunctionSpec) {
	*out = *in
	if in.MaxReplicas != nil {
		in, out := &in.MaxReplicas, &out.MaxReplicas
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	if in.IdleTimeoutMs != nil {
		in, out := &in.IdleTimeoutMs, &out.IdleTimeoutMs
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	in.Container.DeepCopyInto(&out.Container)
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
func (in *FunctionStatus) DeepCopyInto(out *FunctionStatus) {
	*out = *in
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
func (in *Invoker) DeepCopyInto(out *Invoker) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		if *in == nil {
			*out = nil
		} else {
			*out = new(InvokerStatus)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Invoker.
func (in *Invoker) DeepCopy() *Invoker {
	if in == nil {
		return nil
	}
	out := new(Invoker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Invoker) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InvokerFile) DeepCopyInto(out *InvokerFile) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvokerFile.
func (in *InvokerFile) DeepCopy() *InvokerFile {
	if in == nil {
		return nil
	}
	out := new(InvokerFile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InvokerHandler) DeepCopyInto(out *InvokerHandler) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvokerHandler.
func (in *InvokerHandler) DeepCopy() *InvokerHandler {
	if in == nil {
		return nil
	}
	out := new(InvokerHandler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InvokerList) DeepCopyInto(out *InvokerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Invoker, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvokerList.
func (in *InvokerList) DeepCopy() *InvokerList {
	if in == nil {
		return nil
	}
	out := new(InvokerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InvokerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InvokerProperties) DeepCopyInto(out *InvokerProperties) {
	*out = *in
	out.Handler = in.Handler
	in.Function.DeepCopyInto(&out.Function)
	in.Topics.DeepCopyInto(&out.Topics)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvokerProperties.
func (in *InvokerProperties) DeepCopy() *InvokerProperties {
	if in == nil {
		return nil
	}
	out := new(InvokerProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InvokerSpec) DeepCopyInto(out *InvokerSpec) {
	*out = *in
	if in.Matchers != nil {
		in, out := &in.Matchers, &out.Matchers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Properties.DeepCopyInto(&out.Properties)
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]InvokerFile, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvokerSpec.
func (in *InvokerSpec) DeepCopy() *InvokerSpec {
	if in == nil {
		return nil
	}
	out := new(InvokerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InvokerStatus) DeepCopyInto(out *InvokerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvokerStatus.
func (in *InvokerStatus) DeepCopy() *InvokerStatus {
	if in == nil {
		return nil
	}
	out := new(InvokerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Topic) DeepCopyInto(out *Topic) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		if *in == nil {
			*out = nil
		} else {
			*out = new(TopicStatus)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Topic.
func (in *Topic) DeepCopy() *Topic {
	if in == nil {
		return nil
	}
	out := new(Topic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Topic) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopicList) DeepCopyInto(out *TopicList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Topic, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopicList.
func (in *TopicList) DeepCopy() *TopicList {
	if in == nil {
		return nil
	}
	out := new(TopicList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TopicList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopicSpec) DeepCopyInto(out *TopicSpec) {
	*out = *in
	if in.Partitions != nil {
		in, out := &in.Partitions, &out.Partitions
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopicSpec.
func (in *TopicSpec) DeepCopy() *TopicSpec {
	if in == nil {
		return nil
	}
	out := new(TopicSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopicStatus) DeepCopyInto(out *TopicStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopicStatus.
func (in *TopicStatus) DeepCopy() *TopicStatus {
	if in == nil {
		return nil
	}
	out := new(TopicStatus)
	in.DeepCopyInto(out)
	return out
}
