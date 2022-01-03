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

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// FirewallPolicyParameters defines the desired state of FirewallPolicy
type FirewallPolicyParameters struct {
	// Region is which region the FirewallPolicy will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// A description of the firewall policy.
	Description *string `json:"description,omitempty"`
	// Indicates whether you want Network Firewall to just check the validity of
	// the request, rather than run the request.
	//
	// If set to TRUE, Network Firewall checks whether the request can run successfully,
	// but doesn't actually make the requested changes. The call returns the value
	// that the request would return if you ran it with dry run set to FALSE, but
	// doesn't make additions or changes to your resources. This option allows you
	// to make sure that you have the required permissions to run the request and
	// that your request parameters are valid.
	//
	// If set to FALSE, Network Firewall makes the requested changes to your resources.
	DryRun *bool `json:"dryRun,omitempty"`
	// The rule groups and policy actions to use in the firewall policy.
	// +kubebuilder:validation:Required
	FirewallPolicy *FirewallPolicy_SDK `json:"firewallPolicy"`
	// The descriptive name of the firewall policy. You can't change the name of
	// a firewall policy after you create it.
	// +kubebuilder:validation:Required
	FirewallPolicyName *string `json:"firewallPolicyName"`
	// The key:value pairs to associate with the resource.
	Tags                           []*Tag `json:"tags,omitempty"`
	CustomFirewallPolicyParameters `json:",inline"`
}

// FirewallPolicySpec defines the desired state of FirewallPolicy
type FirewallPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       FirewallPolicyParameters `json:"forProvider"`
}

// FirewallPolicyObservation defines the observed state of FirewallPolicy
type FirewallPolicyObservation struct {
	// The high-level properties of a firewall policy. This, along with the FirewallPolicy,
	// define the policy. You can retrieve all objects for a firewall policy by
	// calling DescribeFirewallPolicy.
	FirewallPolicyResponse *FirewallPolicyResponse `json:"firewallPolicyResponse,omitempty"`
	// A token used for optimistic locking. Network Firewall returns a token to
	// your requests that access the firewall policy. The token marks the state
	// of the policy resource at the time of the request.
	//
	// To make changes to the policy, you provide the token in your request. Network
	// Firewall uses the token to ensure that the policy hasn't changed since you
	// last retrieved it. If it has changed, the operation fails with an InvalidTokenException.
	// If this happens, retrieve the firewall policy again to get a current copy
	// of it with current token. Reapply your changes as needed, then try the operation
	// again using the new token.
	UpdateToken *string `json:"updateToken,omitempty"`
}

// FirewallPolicyStatus defines the observed state of FirewallPolicy.
type FirewallPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          FirewallPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallPolicy is the Schema for the FirewallPolicies API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type FirewallPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallPolicySpec   `json:"spec"`
	Status            FirewallPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallPolicyList contains a list of FirewallPolicies
type FirewallPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallPolicy `json:"items"`
}

// Repository type metadata.
var (
	FirewallPolicyKind             = "FirewallPolicy"
	FirewallPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: FirewallPolicyKind}.String()
	FirewallPolicyKindAPIVersion   = FirewallPolicyKind + "." + GroupVersion.String()
	FirewallPolicyGroupVersionKind = GroupVersion.WithKind(FirewallPolicyKind)
)

func init() {
	SchemeBuilder.Register(&FirewallPolicy{}, &FirewallPolicyList{})
}
