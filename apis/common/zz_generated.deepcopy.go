//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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

package common

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSPrincipal) DeepCopyInto(out *AWSPrincipal) {
	*out = *in
	if in.UserARN != nil {
		in, out := &in.UserARN, &out.UserARN
		*out = new(string)
		**out = **in
	}
	if in.UserARNRef != nil {
		in, out := &in.UserARNRef, &out.UserARNRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.UserARNSelector != nil {
		in, out := &in.UserARNSelector, &out.UserARNSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.AWSAccountID != nil {
		in, out := &in.AWSAccountID, &out.AWSAccountID
		*out = new(string)
		**out = **in
	}
	if in.IAMRoleARN != nil {
		in, out := &in.IAMRoleARN, &out.IAMRoleARN
		*out = new(string)
		**out = **in
	}
	if in.IAMRoleARNRef != nil {
		in, out := &in.IAMRoleARNRef, &out.IAMRoleARNRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.IAMRoleARNSelector != nil {
		in, out := &in.IAMRoleARNSelector, &out.IAMRoleARNSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSPrincipal.
func (in *AWSPrincipal) DeepCopy() *AWSPrincipal {
	if in == nil {
		return nil
	}
	out := new(AWSPrincipal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ConditionPair, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionPair) DeepCopyInto(out *ConditionPair) {
	*out = *in
	if in.ConditionStringValue != nil {
		in, out := &in.ConditionStringValue, &out.ConditionStringValue
		*out = new(string)
		**out = **in
	}
	if in.ConditionDateValue != nil {
		in, out := &in.ConditionDateValue, &out.ConditionDateValue
		*out = (*in).DeepCopy()
	}
	if in.ConditionNumericValue != nil {
		in, out := &in.ConditionNumericValue, &out.ConditionNumericValue
		*out = new(int64)
		**out = **in
	}
	if in.ConditionBooleanValue != nil {
		in, out := &in.ConditionBooleanValue, &out.ConditionBooleanValue
		*out = new(bool)
		**out = **in
	}
	if in.ConditionListValue != nil {
		in, out := &in.ConditionListValue, &out.ConditionListValue
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionPair.
func (in *ConditionPair) DeepCopy() *ConditionPair {
	if in == nil {
		return nil
	}
	out := new(ConditionPair)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcePolicy) DeepCopyInto(out *ResourcePolicy) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Statements != nil {
		in, out := &in.Statements, &out.Statements
		*out = make([]ResourcePolicyStatement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcePolicy.
func (in *ResourcePolicy) DeepCopy() *ResourcePolicy {
	if in == nil {
		return nil
	}
	out := new(ResourcePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcePolicyStatement) DeepCopyInto(out *ResourcePolicyStatement) {
	*out = *in
	if in.SID != nil {
		in, out := &in.SID, &out.SID
		*out = new(string)
		**out = **in
	}
	if in.Principal != nil {
		in, out := &in.Principal, &out.Principal
		*out = new(ResourcePrincipal)
		(*in).DeepCopyInto(*out)
	}
	if in.NotPrincipal != nil {
		in, out := &in.NotPrincipal, &out.NotPrincipal
		*out = new(ResourcePrincipal)
		(*in).DeepCopyInto(*out)
	}
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NotAction != nil {
		in, out := &in.NotAction, &out.NotAction
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NotResource != nil {
		in, out := &in.NotResource, &out.NotResource
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcePolicyStatement.
func (in *ResourcePolicyStatement) DeepCopy() *ResourcePolicyStatement {
	if in == nil {
		return nil
	}
	out := new(ResourcePolicyStatement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcePrincipal) DeepCopyInto(out *ResourcePrincipal) {
	*out = *in
	if in.AWSPrincipals != nil {
		in, out := &in.AWSPrincipals, &out.AWSPrincipals
		*out = make([]AWSPrincipal, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Federated != nil {
		in, out := &in.Federated, &out.Federated
		*out = new(string)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcePrincipal.
func (in *ResourcePrincipal) DeepCopy() *ResourcePrincipal {
	if in == nil {
		return nil
	}
	out := new(ResourcePrincipal)
	in.DeepCopyInto(out)
	return out
}
