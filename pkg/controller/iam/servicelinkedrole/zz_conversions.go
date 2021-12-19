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

// Code generated by ack-generate. DO NOT EDIT.

package servicelinkedrole

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/iam"

	svcapitypes "github.com/crossplane/provider-aws/apis/iam/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateCreateServiceLinkedRoleInput returns a create input.
func GenerateCreateServiceLinkedRoleInput(cr *svcapitypes.ServiceLinkedRole) *svcsdk.CreateServiceLinkedRoleInput {
	res := &svcsdk.CreateServiceLinkedRoleInput{}

	if cr.Spec.ForProvider.AWSServiceName != nil {
		res.SetAWSServiceName(*cr.Spec.ForProvider.AWSServiceName)
	}
	if cr.Spec.ForProvider.CustomSuffix != nil {
		res.SetCustomSuffix(*cr.Spec.ForProvider.CustomSuffix)
	}
	if cr.Spec.ForProvider.Description != nil {
		res.SetDescription(*cr.Spec.ForProvider.Description)
	}

	return res
}

// GenerateDeleteServiceLinkedRoleInput returns a deletion input.
func GenerateDeleteServiceLinkedRoleInput(cr *svcapitypes.ServiceLinkedRole) *svcsdk.DeleteServiceLinkedRoleInput {
	res := &svcsdk.DeleteServiceLinkedRoleInput{}

	if cr.Status.AtProvider.RoleName != nil {
		res.SetRoleName(*cr.Status.AtProvider.RoleName)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "UNKNOWN"
}
