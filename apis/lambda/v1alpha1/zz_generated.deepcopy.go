// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountLimit) DeepCopyInto(out *AccountLimit) {
	*out = *in
	if in.CodeSizeUnzipped != nil {
		in, out := &in.CodeSizeUnzipped, &out.CodeSizeUnzipped
		*out = new(int64)
		**out = **in
	}
	if in.CodeSizeZipped != nil {
		in, out := &in.CodeSizeZipped, &out.CodeSizeZipped
		*out = new(int64)
		**out = **in
	}
	if in.TotalCodeSize != nil {
		in, out := &in.TotalCodeSize, &out.TotalCodeSize
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountLimit.
func (in *AccountLimit) DeepCopy() *AccountLimit {
	if in == nil {
		return nil
	}
	out := new(AccountLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountUsage) DeepCopyInto(out *AccountUsage) {
	*out = *in
	if in.FunctionCount != nil {
		in, out := &in.FunctionCount, &out.FunctionCount
		*out = new(int64)
		**out = **in
	}
	if in.TotalCodeSize != nil {
		in, out := &in.TotalCodeSize, &out.TotalCodeSize
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountUsage.
func (in *AccountUsage) DeepCopy() *AccountUsage {
	if in == nil {
		return nil
	}
	out := new(AccountUsage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CodeSigningConfig) DeepCopyInto(out *CodeSigningConfig) {
	*out = *in
	if in.CodeSigningConfigARN != nil {
		in, out := &in.CodeSigningConfigARN, &out.CodeSigningConfigARN
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.LastModified != nil {
		in, out := &in.LastModified, &out.LastModified
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CodeSigningConfig.
func (in *CodeSigningConfig) DeepCopy() *CodeSigningConfig {
	if in == nil {
		return nil
	}
	out := new(CodeSigningConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomCodeSigningConfigParameters) DeepCopyInto(out *CustomCodeSigningConfigParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomCodeSigningConfigParameters.
func (in *CustomCodeSigningConfigParameters) DeepCopy() *CustomCodeSigningConfigParameters {
	if in == nil {
		return nil
	}
	out := new(CustomCodeSigningConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomFunctionCodeParameters) DeepCopyInto(out *CustomFunctionCodeParameters) {
	*out = *in
	if in.ImageURI != nil {
		in, out := &in.ImageURI, &out.ImageURI
		*out = new(string)
		**out = **in
	}
	if in.S3Key != nil {
		in, out := &in.S3Key, &out.S3Key
		*out = new(string)
		**out = **in
	}
	if in.S3ObjectVersion != nil {
		in, out := &in.S3ObjectVersion, &out.S3ObjectVersion
		*out = new(string)
		**out = **in
	}
	if in.S3Bucket != nil {
		in, out := &in.S3Bucket, &out.S3Bucket
		*out = new(string)
		**out = **in
	}
	if in.S3BucketRef != nil {
		in, out := &in.S3BucketRef, &out.S3BucketRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.S3BucketSelector != nil {
		in, out := &in.S3BucketSelector, &out.S3BucketSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomFunctionCodeParameters.
func (in *CustomFunctionCodeParameters) DeepCopy() *CustomFunctionCodeParameters {
	if in == nil {
		return nil
	}
	out := new(CustomFunctionCodeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomFunctionParameters) DeepCopyInto(out *CustomFunctionParameters) {
	*out = *in
	if in.KMSKeyARNRef != nil {
		in, out := &in.KMSKeyARNRef, &out.KMSKeyARNRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.KMSKeyARNSelector != nil {
		in, out := &in.KMSKeyARNSelector, &out.KMSKeyARNSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.RoleRef != nil {
		in, out := &in.RoleRef, &out.RoleRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.RoleSelector != nil {
		in, out := &in.RoleSelector, &out.RoleSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomFunctionVPCConfigParameters != nil {
		in, out := &in.CustomFunctionVPCConfigParameters, &out.CustomFunctionVPCConfigParameters
		*out = new(CustomFunctionVPCConfigParameters)
		(*in).DeepCopyInto(*out)
	}
	in.CustomFunctionCodeParameters.DeepCopyInto(&out.CustomFunctionCodeParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomFunctionParameters.
func (in *CustomFunctionParameters) DeepCopy() *CustomFunctionParameters {
	if in == nil {
		return nil
	}
	out := new(CustomFunctionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomFunctionVPCConfigParameters) DeepCopyInto(out *CustomFunctionVPCConfigParameters) {
	*out = *in
	if in.SecurityGroupIDs != nil {
		in, out := &in.SecurityGroupIDs, &out.SecurityGroupIDs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SecurityGroupIDRefs != nil {
		in, out := &in.SecurityGroupIDRefs, &out.SecurityGroupIDRefs
		*out = make([]v1.Reference, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupIDSelector != nil {
		in, out := &in.SecurityGroupIDSelector, &out.SecurityGroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIDs != nil {
		in, out := &in.SubnetIDs, &out.SubnetIDs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetIDRefs != nil {
		in, out := &in.SubnetIDRefs, &out.SubnetIDRefs
		*out = make([]v1.Reference, len(*in))
		copy(*out, *in)
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomFunctionVPCConfigParameters.
func (in *CustomFunctionVPCConfigParameters) DeepCopy() *CustomFunctionVPCConfigParameters {
	if in == nil {
		return nil
	}
	out := new(CustomFunctionVPCConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeadLetterConfig) DeepCopyInto(out *DeadLetterConfig) {
	*out = *in
	if in.TargetARN != nil {
		in, out := &in.TargetARN, &out.TargetARN
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeadLetterConfig.
func (in *DeadLetterConfig) DeepCopy() *DeadLetterConfig {
	if in == nil {
		return nil
	}
	out := new(DeadLetterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Environment) DeepCopyInto(out *Environment) {
	*out = *in
	if in.Variables != nil {
		in, out := &in.Variables, &out.Variables
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Environment.
func (in *Environment) DeepCopy() *Environment {
	if in == nil {
		return nil
	}
	out := new(Environment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentError) DeepCopyInto(out *EnvironmentError) {
	*out = *in
	if in.ErrorCode != nil {
		in, out := &in.ErrorCode, &out.ErrorCode
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentError.
func (in *EnvironmentError) DeepCopy() *EnvironmentError {
	if in == nil {
		return nil
	}
	out := new(EnvironmentError)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentResponse) DeepCopyInto(out *EnvironmentResponse) {
	*out = *in
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = new(EnvironmentError)
		(*in).DeepCopyInto(*out)
	}
	if in.Variables != nil {
		in, out := &in.Variables, &out.Variables
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentResponse.
func (in *EnvironmentResponse) DeepCopy() *EnvironmentResponse {
	if in == nil {
		return nil
	}
	out := new(EnvironmentResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemConfig) DeepCopyInto(out *FileSystemConfig) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.LocalMountPath != nil {
		in, out := &in.LocalMountPath, &out.LocalMountPath
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemConfig.
func (in *FileSystemConfig) DeepCopy() *FileSystemConfig {
	if in == nil {
		return nil
	}
	out := new(FileSystemConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Function) DeepCopyInto(out *Function) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
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
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionCode) DeepCopyInto(out *FunctionCode) {
	*out = *in
	if in.ImageURI != nil {
		in, out := &in.ImageURI, &out.ImageURI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionCode.
func (in *FunctionCode) DeepCopy() *FunctionCode {
	if in == nil {
		return nil
	}
	out := new(FunctionCode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionCodeLocation) DeepCopyInto(out *FunctionCodeLocation) {
	*out = *in
	if in.ImageURI != nil {
		in, out := &in.ImageURI, &out.ImageURI
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.RepositoryType != nil {
		in, out := &in.RepositoryType, &out.RepositoryType
		*out = new(string)
		**out = **in
	}
	if in.ResolvedImageURI != nil {
		in, out := &in.ResolvedImageURI, &out.ResolvedImageURI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionCodeLocation.
func (in *FunctionCodeLocation) DeepCopy() *FunctionCodeLocation {
	if in == nil {
		return nil
	}
	out := new(FunctionCodeLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionConfiguration) DeepCopyInto(out *FunctionConfiguration) {
	*out = *in
	if in.CodeSHA256 != nil {
		in, out := &in.CodeSHA256, &out.CodeSHA256
		*out = new(string)
		**out = **in
	}
	if in.CodeSize != nil {
		in, out := &in.CodeSize, &out.CodeSize
		*out = new(int64)
		**out = **in
	}
	if in.DeadLetterConfig != nil {
		in, out := &in.DeadLetterConfig, &out.DeadLetterConfig
		*out = new(DeadLetterConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = new(EnvironmentResponse)
		(*in).DeepCopyInto(*out)
	}
	if in.FileSystemConfigs != nil {
		in, out := &in.FileSystemConfigs, &out.FileSystemConfigs
		*out = make([]*FileSystemConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(FileSystemConfig)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.FunctionARN != nil {
		in, out := &in.FunctionARN, &out.FunctionARN
		*out = new(string)
		**out = **in
	}
	if in.FunctionName != nil {
		in, out := &in.FunctionName, &out.FunctionName
		*out = new(string)
		**out = **in
	}
	if in.Handler != nil {
		in, out := &in.Handler, &out.Handler
		*out = new(string)
		**out = **in
	}
	if in.ImageConfigResponse != nil {
		in, out := &in.ImageConfigResponse, &out.ImageConfigResponse
		*out = new(ImageConfigResponse)
		(*in).DeepCopyInto(*out)
	}
	if in.KMSKeyARN != nil {
		in, out := &in.KMSKeyARN, &out.KMSKeyARN
		*out = new(string)
		**out = **in
	}
	if in.LastModified != nil {
		in, out := &in.LastModified, &out.LastModified
		*out = new(string)
		**out = **in
	}
	if in.LastUpdateStatus != nil {
		in, out := &in.LastUpdateStatus, &out.LastUpdateStatus
		*out = new(string)
		**out = **in
	}
	if in.LastUpdateStatusReason != nil {
		in, out := &in.LastUpdateStatusReason, &out.LastUpdateStatusReason
		*out = new(string)
		**out = **in
	}
	if in.LastUpdateStatusReasonCode != nil {
		in, out := &in.LastUpdateStatusReasonCode, &out.LastUpdateStatusReasonCode
		*out = new(string)
		**out = **in
	}
	if in.Layers != nil {
		in, out := &in.Layers, &out.Layers
		*out = make([]*Layer, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Layer)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.MasterARN != nil {
		in, out := &in.MasterARN, &out.MasterARN
		*out = new(string)
		**out = **in
	}
	if in.MemorySize != nil {
		in, out := &in.MemorySize, &out.MemorySize
		*out = new(int64)
		**out = **in
	}
	if in.PackageType != nil {
		in, out := &in.PackageType, &out.PackageType
		*out = new(string)
		**out = **in
	}
	if in.RevisionID != nil {
		in, out := &in.RevisionID, &out.RevisionID
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.Runtime != nil {
		in, out := &in.Runtime, &out.Runtime
		*out = new(string)
		**out = **in
	}
	if in.SigningJobARN != nil {
		in, out := &in.SigningJobARN, &out.SigningJobARN
		*out = new(string)
		**out = **in
	}
	if in.SigningProfileVersionARN != nil {
		in, out := &in.SigningProfileVersionARN, &out.SigningProfileVersionARN
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.StateReason != nil {
		in, out := &in.StateReason, &out.StateReason
		*out = new(string)
		**out = **in
	}
	if in.StateReasonCode != nil {
		in, out := &in.StateReasonCode, &out.StateReasonCode
		*out = new(string)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int64)
		**out = **in
	}
	if in.TracingConfig != nil {
		in, out := &in.TracingConfig, &out.TracingConfig
		*out = new(TracingConfigResponse)
		(*in).DeepCopyInto(*out)
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.VPCConfig != nil {
		in, out := &in.VPCConfig, &out.VPCConfig
		*out = new(VPCConfigResponse)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionConfiguration.
func (in *FunctionConfiguration) DeepCopy() *FunctionConfiguration {
	if in == nil {
		return nil
	}
	out := new(FunctionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionEventInvokeConfig) DeepCopyInto(out *FunctionEventInvokeConfig) {
	*out = *in
	if in.FunctionARN != nil {
		in, out := &in.FunctionARN, &out.FunctionARN
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionEventInvokeConfig.
func (in *FunctionEventInvokeConfig) DeepCopy() *FunctionEventInvokeConfig {
	if in == nil {
		return nil
	}
	out := new(FunctionEventInvokeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionList) DeepCopyInto(out *FunctionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Function, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
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
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionObservation) DeepCopyInto(out *FunctionObservation) {
	*out = *in
	if in.CodeSHA256 != nil {
		in, out := &in.CodeSHA256, &out.CodeSHA256
		*out = new(string)
		**out = **in
	}
	if in.CodeSize != nil {
		in, out := &in.CodeSize, &out.CodeSize
		*out = new(int64)
		**out = **in
	}
	if in.FunctionARN != nil {
		in, out := &in.FunctionARN, &out.FunctionARN
		*out = new(string)
		**out = **in
	}
	if in.FunctionName != nil {
		in, out := &in.FunctionName, &out.FunctionName
		*out = new(string)
		**out = **in
	}
	if in.ImageConfigResponse != nil {
		in, out := &in.ImageConfigResponse, &out.ImageConfigResponse
		*out = new(ImageConfigResponse)
		(*in).DeepCopyInto(*out)
	}
	if in.LastModified != nil {
		in, out := &in.LastModified, &out.LastModified
		*out = new(string)
		**out = **in
	}
	if in.LastUpdateStatus != nil {
		in, out := &in.LastUpdateStatus, &out.LastUpdateStatus
		*out = new(string)
		**out = **in
	}
	if in.LastUpdateStatusReason != nil {
		in, out := &in.LastUpdateStatusReason, &out.LastUpdateStatusReason
		*out = new(string)
		**out = **in
	}
	if in.LastUpdateStatusReasonCode != nil {
		in, out := &in.LastUpdateStatusReasonCode, &out.LastUpdateStatusReasonCode
		*out = new(string)
		**out = **in
	}
	if in.MasterARN != nil {
		in, out := &in.MasterARN, &out.MasterARN
		*out = new(string)
		**out = **in
	}
	if in.RevisionID != nil {
		in, out := &in.RevisionID, &out.RevisionID
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.SigningJobARN != nil {
		in, out := &in.SigningJobARN, &out.SigningJobARN
		*out = new(string)
		**out = **in
	}
	if in.SigningProfileVersionARN != nil {
		in, out := &in.SigningProfileVersionARN, &out.SigningProfileVersionARN
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.StateReason != nil {
		in, out := &in.StateReason, &out.StateReason
		*out = new(string)
		**out = **in
	}
	if in.StateReasonCode != nil {
		in, out := &in.StateReasonCode, &out.StateReasonCode
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.VPCConfig != nil {
		in, out := &in.VPCConfig, &out.VPCConfig
		*out = new(VPCConfigResponse)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionObservation.
func (in *FunctionObservation) DeepCopy() *FunctionObservation {
	if in == nil {
		return nil
	}
	out := new(FunctionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionParameters) DeepCopyInto(out *FunctionParameters) {
	*out = *in
	if in.CodeSigningConfigARN != nil {
		in, out := &in.CodeSigningConfigARN, &out.CodeSigningConfigARN
		*out = new(string)
		**out = **in
	}
	if in.DeadLetterConfig != nil {
		in, out := &in.DeadLetterConfig, &out.DeadLetterConfig
		*out = new(DeadLetterConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = new(Environment)
		(*in).DeepCopyInto(*out)
	}
	if in.FileSystemConfigs != nil {
		in, out := &in.FileSystemConfigs, &out.FileSystemConfigs
		*out = make([]*FileSystemConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(FileSystemConfig)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Handler != nil {
		in, out := &in.Handler, &out.Handler
		*out = new(string)
		**out = **in
	}
	if in.ImageConfig != nil {
		in, out := &in.ImageConfig, &out.ImageConfig
		*out = new(ImageConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.KMSKeyARN != nil {
		in, out := &in.KMSKeyARN, &out.KMSKeyARN
		*out = new(string)
		**out = **in
	}
	if in.Layers != nil {
		in, out := &in.Layers, &out.Layers
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.MemorySize != nil {
		in, out := &in.MemorySize, &out.MemorySize
		*out = new(int64)
		**out = **in
	}
	if in.PackageType != nil {
		in, out := &in.PackageType, &out.PackageType
		*out = new(string)
		**out = **in
	}
	if in.Publish != nil {
		in, out := &in.Publish, &out.Publish
		*out = new(bool)
		**out = **in
	}
	if in.Runtime != nil {
		in, out := &in.Runtime, &out.Runtime
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int64)
		**out = **in
	}
	if in.TracingConfig != nil {
		in, out := &in.TracingConfig, &out.TracingConfig
		*out = new(TracingConfig)
		(*in).DeepCopyInto(*out)
	}
	in.CustomFunctionParameters.DeepCopyInto(&out.CustomFunctionParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionParameters.
func (in *FunctionParameters) DeepCopy() *FunctionParameters {
	if in == nil {
		return nil
	}
	out := new(FunctionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionSpec) DeepCopyInto(out *FunctionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
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
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
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
func (in *ImageConfig) DeepCopyInto(out *ImageConfig) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.EntryPoint != nil {
		in, out := &in.EntryPoint, &out.EntryPoint
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.WorkingDirectory != nil {
		in, out := &in.WorkingDirectory, &out.WorkingDirectory
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageConfig.
func (in *ImageConfig) DeepCopy() *ImageConfig {
	if in == nil {
		return nil
	}
	out := new(ImageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageConfigError) DeepCopyInto(out *ImageConfigError) {
	*out = *in
	if in.ErrorCode != nil {
		in, out := &in.ErrorCode, &out.ErrorCode
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageConfigError.
func (in *ImageConfigError) DeepCopy() *ImageConfigError {
	if in == nil {
		return nil
	}
	out := new(ImageConfigError)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageConfigResponse) DeepCopyInto(out *ImageConfigResponse) {
	*out = *in
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = new(ImageConfigError)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageConfig != nil {
		in, out := &in.ImageConfig, &out.ImageConfig
		*out = new(ImageConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageConfigResponse.
func (in *ImageConfigResponse) DeepCopy() *ImageConfigResponse {
	if in == nil {
		return nil
	}
	out := new(ImageConfigResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Layer) DeepCopyInto(out *Layer) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.CodeSize != nil {
		in, out := &in.CodeSize, &out.CodeSize
		*out = new(int64)
		**out = **in
	}
	if in.SigningJobARN != nil {
		in, out := &in.SigningJobARN, &out.SigningJobARN
		*out = new(string)
		**out = **in
	}
	if in.SigningProfileVersionARN != nil {
		in, out := &in.SigningProfileVersionARN, &out.SigningProfileVersionARN
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Layer.
func (in *Layer) DeepCopy() *Layer {
	if in == nil {
		return nil
	}
	out := new(Layer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LayerVersionContentOutput) DeepCopyInto(out *LayerVersionContentOutput) {
	*out = *in
	if in.CodeSHA256 != nil {
		in, out := &in.CodeSHA256, &out.CodeSHA256
		*out = new(string)
		**out = **in
	}
	if in.CodeSize != nil {
		in, out := &in.CodeSize, &out.CodeSize
		*out = new(int64)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.SigningJobARN != nil {
		in, out := &in.SigningJobARN, &out.SigningJobARN
		*out = new(string)
		**out = **in
	}
	if in.SigningProfileVersionARN != nil {
		in, out := &in.SigningProfileVersionARN, &out.SigningProfileVersionARN
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LayerVersionContentOutput.
func (in *LayerVersionContentOutput) DeepCopy() *LayerVersionContentOutput {
	if in == nil {
		return nil
	}
	out := new(LayerVersionContentOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LayerVersionsListItem) DeepCopyInto(out *LayerVersionsListItem) {
	*out = *in
	if in.CreatedDate != nil {
		in, out := &in.CreatedDate, &out.CreatedDate
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.LayerVersionARN != nil {
		in, out := &in.LayerVersionARN, &out.LayerVersionARN
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LayerVersionsListItem.
func (in *LayerVersionsListItem) DeepCopy() *LayerVersionsListItem {
	if in == nil {
		return nil
	}
	out := new(LayerVersionsListItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionedConcurrencyConfigListItem) DeepCopyInto(out *ProvisionedConcurrencyConfigListItem) {
	*out = *in
	if in.FunctionARN != nil {
		in, out := &in.FunctionARN, &out.FunctionARN
		*out = new(string)
		**out = **in
	}
	if in.LastModified != nil {
		in, out := &in.LastModified, &out.LastModified
		*out = new(string)
		**out = **in
	}
	if in.StatusReason != nil {
		in, out := &in.StatusReason, &out.StatusReason
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionedConcurrencyConfigListItem.
func (in *ProvisionedConcurrencyConfigListItem) DeepCopy() *ProvisionedConcurrencyConfigListItem {
	if in == nil {
		return nil
	}
	out := new(ProvisionedConcurrencyConfigListItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PutFunctionConcurrencyOutput) DeepCopyInto(out *PutFunctionConcurrencyOutput) {
	*out = *in
	if in.ReservedConcurrentExecutions != nil {
		in, out := &in.ReservedConcurrentExecutions, &out.ReservedConcurrentExecutions
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PutFunctionConcurrencyOutput.
func (in *PutFunctionConcurrencyOutput) DeepCopy() *PutFunctionConcurrencyOutput {
	if in == nil {
		return nil
	}
	out := new(PutFunctionConcurrencyOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracingConfig) DeepCopyInto(out *TracingConfig) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracingConfig.
func (in *TracingConfig) DeepCopy() *TracingConfig {
	if in == nil {
		return nil
	}
	out := new(TracingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracingConfigResponse) DeepCopyInto(out *TracingConfigResponse) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracingConfigResponse.
func (in *TracingConfigResponse) DeepCopy() *TracingConfigResponse {
	if in == nil {
		return nil
	}
	out := new(TracingConfigResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCConfig) DeepCopyInto(out *VPCConfig) {
	*out = *in
	if in.SecurityGroupIDs != nil {
		in, out := &in.SecurityGroupIDs, &out.SecurityGroupIDs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetIDs != nil {
		in, out := &in.SubnetIDs, &out.SubnetIDs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCConfig.
func (in *VPCConfig) DeepCopy() *VPCConfig {
	if in == nil {
		return nil
	}
	out := new(VPCConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCConfigResponse) DeepCopyInto(out *VPCConfigResponse) {
	*out = *in
	if in.SecurityGroupIDs != nil {
		in, out := &in.SecurityGroupIDs, &out.SecurityGroupIDs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetIDs != nil {
		in, out := &in.SubnetIDs, &out.SubnetIDs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCConfigResponse.
func (in *VPCConfigResponse) DeepCopy() *VPCConfigResponse {
	if in == nil {
		return nil
	}
	out := new(VPCConfigResponse)
	in.DeepCopyInto(out)
	return out
}
