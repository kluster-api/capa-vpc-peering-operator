//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

package v1beta1

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Actions) DeepCopyInto(out *Actions) {
	{
		in := &in
		*out = make(Actions, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Actions.
func (in Actions) DeepCopy() Actions {
	if in == nil {
		return nil
	}
	out := new(Actions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyDocument) DeepCopyInto(out *PolicyDocument) {
	*out = *in
	if in.Statement != nil {
		in, out := &in.Statement, &out.Statement
		*out = make(Statements, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyDocument.
func (in *PolicyDocument) DeepCopy() *PolicyDocument {
	if in == nil {
		return nil
	}
	out := new(PolicyDocument)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PrincipalID) DeepCopyInto(out *PrincipalID) {
	{
		in := &in
		*out = make(PrincipalID, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrincipalID.
func (in PrincipalID) DeepCopy() PrincipalID {
	if in == nil {
		return nil
	}
	out := new(PrincipalID)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Principals) DeepCopyInto(out *Principals) {
	{
		in := &in
		*out = make(Principals, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(PrincipalID, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Principals.
func (in Principals) DeepCopy() Principals {
	if in == nil {
		return nil
	}
	out := new(Principals)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Resources) DeepCopyInto(out *Resources) {
	{
		in := &in
		*out = make(Resources, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources.
func (in Resources) DeepCopy() Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatementEntry) DeepCopyInto(out *StatementEntry) {
	*out = *in
	if in.Principal != nil {
		in, out := &in.Principal, &out.Principal
		*out = make(Principals, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(PrincipalID, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.NotPrincipal != nil {
		in, out := &in.NotPrincipal, &out.NotPrincipal
		*out = make(Principals, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(PrincipalID, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = make(Actions, len(*in))
		copy(*out, *in)
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = make(Resources, len(*in))
		copy(*out, *in)
	}
	out.Condition = in.Condition.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatementEntry.
func (in *StatementEntry) DeepCopy() *StatementEntry {
	if in == nil {
		return nil
	}
	out := new(StatementEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Statements) DeepCopyInto(out *Statements) {
	{
		in := &in
		*out = make(Statements, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Statements.
func (in Statements) DeepCopy() Statements {
	if in == nil {
		return nil
	}
	out := new(Statements)
	in.DeepCopyInto(out)
	return *out
}
