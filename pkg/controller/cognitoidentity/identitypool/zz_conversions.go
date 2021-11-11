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

package identitypool

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/cognitoidentity"

	svcapitypes "github.com/crossplane/provider-aws/apis/cognitoidentity/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateDescribeIdentityPoolInput returns input for read
// operation.
func GenerateDescribeIdentityPoolInput(cr *svcapitypes.IDentityPool) *svcsdk.DescribeIdentityPoolInput {
	res := &svcsdk.DescribeIdentityPoolInput{}

	if cr.Status.AtProvider.IDentityPoolID != nil {
		res.SetIdentityPoolId(*cr.Status.AtProvider.IDentityPoolID)
	}

	return res
}

// GenerateIDentityPool returns the current state in the form of *svcapitypes.IDentityPool.
func GenerateIDentityPool(resp *svcsdk.IdentityPool) *svcapitypes.IDentityPool {
	cr := &svcapitypes.IDentityPool{}

	if resp.IdentityPoolId != nil {
		cr.Status.AtProvider.IDentityPoolID = resp.IdentityPoolId
	} else {
		cr.Status.AtProvider.IDentityPoolID = nil
	}

return cr
}

// GenerateCreateIdentityPoolInput returns a create input.
func GenerateCreateIdentityPoolInput(cr *svcapitypes.IDentityPool) *svcsdk.CreateIdentityPoolInput {
	res := &svcsdk.CreateIdentityPoolInput{}

	if cr.Spec.ForProvider.AllowClassicFlow != nil {
		res.SetAllowClassicFlow(*cr.Spec.ForProvider.AllowClassicFlow)
	}
	if cr.Spec.ForProvider.AllowUnauthenticatedIDentities != nil {
		res.SetAllowUnauthenticatedIdentities(*cr.Spec.ForProvider.AllowUnauthenticatedIDentities)
	}
	if cr.Spec.ForProvider.CognitoIDentityProviders != nil {
		f2 := []*svcsdk.Provider{}
		for _, f2iter := range cr.Spec.ForProvider.CognitoIDentityProviders {
			f2elem := &svcsdk.Provider{}
			if f2iter.ClientID != nil {
				f2elem.SetClientId(*f2iter.ClientID)
			}
			if f2iter.ProviderName != nil {
				f2elem.SetProviderName(*f2iter.ProviderName)
			}
			if f2iter.ServerSideTokenCheck != nil {
				f2elem.SetServerSideTokenCheck(*f2iter.ServerSideTokenCheck)
			}
			f2 = append(f2, f2elem)
		}
		res.SetCognitoIdentityProviders(f2)
	}
	if cr.Spec.ForProvider.DeveloperProviderName != nil {
		res.SetDeveloperProviderName(*cr.Spec.ForProvider.DeveloperProviderName)
	}
	if cr.Spec.ForProvider.IDentityPoolName != nil {
		res.SetIdentityPoolName(*cr.Spec.ForProvider.IDentityPoolName)
	}
	if cr.Spec.ForProvider.IDentityPoolTags != nil {
		f5 := map[string]*string{}
		for f5key, f5valiter := range cr.Spec.ForProvider.IDentityPoolTags {
			var f5val string
			f5val = *f5valiter
			f5[f5key] = &f5val
		}
		res.SetIdentityPoolTags(f5)
	}
	if cr.Spec.ForProvider.OpenIDConnectProviderARNs != nil {
		f6 := []*string{}
		for _, f6iter := range cr.Spec.ForProvider.OpenIDConnectProviderARNs {
			var f6elem string
			f6elem = *f6iter
			f6 = append(f6, &f6elem)
		}
		res.SetOpenIdConnectProviderARNs(f6)
	}
	if cr.Spec.ForProvider.SamlProviderARNs != nil {
		f7 := []*string{}
		for _, f7iter := range cr.Spec.ForProvider.SamlProviderARNs {
			var f7elem string
			f7elem = *f7iter
			f7 = append(f7, &f7elem)
		}
		res.SetSamlProviderARNs(f7)
	}
	if cr.Spec.ForProvider.SupportedLoginProviders != nil {
		f8 := map[string]*string{}
		for f8key, f8valiter := range cr.Spec.ForProvider.SupportedLoginProviders {
			var f8val string
			f8val = *f8valiter
			f8[f8key] = &f8val
		}
		res.SetSupportedLoginProviders(f8)
	}

	return res
}
// GenerateIdentityPool returns an update input.
func GenerateIdentityPool(cr *svcapitypes.IDentityPool) *svcsdk.IdentityPool {
	res := &svcsdk.IdentityPool{}

	if cr.Spec.ForProvider.AllowClassicFlow != nil {
		res.SetAllowClassicFlow(*cr.Spec.ForProvider.AllowClassicFlow)
	}
	if cr.Spec.ForProvider.AllowUnauthenticatedIDentities != nil {
		res.SetAllowUnauthenticatedIdentities(*cr.Spec.ForProvider.AllowUnauthenticatedIDentities)
	}
	if cr.Spec.ForProvider.CognitoIDentityProviders != nil {
		f2 := []*svcsdk.Provider{}
		for _, f2iter := range cr.Spec.ForProvider.CognitoIDentityProviders {
			f2elem := &svcsdk.Provider{}
			if f2iter.ClientID != nil {
				f2elem.SetClientId(*f2iter.ClientID)
			}
			if f2iter.ProviderName != nil {
				f2elem.SetProviderName(*f2iter.ProviderName)
			}
			if f2iter.ServerSideTokenCheck != nil {
				f2elem.SetServerSideTokenCheck(*f2iter.ServerSideTokenCheck)
			}
			f2 = append(f2, f2elem)
		}
		res.SetCognitoIdentityProviders(f2)
	}
	if cr.Spec.ForProvider.DeveloperProviderName != nil {
		res.SetDeveloperProviderName(*cr.Spec.ForProvider.DeveloperProviderName)
	}
	if cr.Status.AtProvider.IDentityPoolID != nil {
		res.SetIdentityPoolId(*cr.Status.AtProvider.IDentityPoolID)
	}
	if cr.Spec.ForProvider.IDentityPoolName != nil {
		res.SetIdentityPoolName(*cr.Spec.ForProvider.IDentityPoolName)
	}
	if cr.Spec.ForProvider.IDentityPoolTags != nil {
		f6 := map[string]*string{}
		for f6key, f6valiter := range cr.Spec.ForProvider.IDentityPoolTags {
			var f6val string
			f6val = *f6valiter
			f6[f6key] = &f6val
		}
		res.SetIdentityPoolTags(f6)
	}
	if cr.Spec.ForProvider.OpenIDConnectProviderARNs != nil {
		f7 := []*string{}
		for _, f7iter := range cr.Spec.ForProvider.OpenIDConnectProviderARNs {
			var f7elem string
			f7elem = *f7iter
			f7 = append(f7, &f7elem)
		}
		res.SetOpenIdConnectProviderARNs(f7)
	}
	if cr.Spec.ForProvider.SamlProviderARNs != nil {
		f8 := []*string{}
		for _, f8iter := range cr.Spec.ForProvider.SamlProviderARNs {
			var f8elem string
			f8elem = *f8iter
			f8 = append(f8, &f8elem)
		}
		res.SetSamlProviderARNs(f8)
	}
	if cr.Spec.ForProvider.SupportedLoginProviders != nil {
		f9 := map[string]*string{}
		for f9key, f9valiter := range cr.Spec.ForProvider.SupportedLoginProviders {
			var f9val string
			f9val = *f9valiter
			f9[f9key] = &f9val
		}
		res.SetSupportedLoginProviders(f9)
	}

	return res
}

// GenerateDeleteIdentityPoolInput returns a deletion input.
func GenerateDeleteIdentityPoolInput(cr *svcapitypes.IDentityPool) *svcsdk.DeleteIdentityPoolInput {
	res := &svcsdk.DeleteIdentityPoolInput{}

	if cr.Status.AtProvider.IDentityPoolID != nil {
		res.SetIdentityPoolId(*cr.Status.AtProvider.IDentityPoolID)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "UNKNOWN" 
}