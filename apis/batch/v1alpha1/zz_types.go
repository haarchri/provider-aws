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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
)

// +kubebuilder:skipversion
type ArrayProperties struct {
	Size *int64 `json:"size,omitempty"`
}

// +kubebuilder:skipversion
type ArrayPropertiesDetail struct {
	Index *int64 `json:"index,omitempty"`

	Size *int64 `json:"size,omitempty"`
}

// +kubebuilder:skipversion
type ArrayPropertiesSummary struct {
	Index *int64 `json:"index,omitempty"`

	Size *int64 `json:"size,omitempty"`
}

// +kubebuilder:skipversion
type AttemptContainerDetail struct {
	ContainerInstanceARN *string `json:"containerInstanceARN,omitempty"`

	ExitCode *int64 `json:"exitCode,omitempty"`

	LogStreamName *string `json:"logStreamName,omitempty"`

	Reason *string `json:"reason,omitempty"`

	TaskARN *string `json:"taskARN,omitempty"`
}

// +kubebuilder:skipversion
type AttemptDetail struct {
	StatusReason *string `json:"statusReason,omitempty"`
}

// +kubebuilder:skipversion
type ComputeEnvironmentDetail struct {
	ComputeEnvironmentARN *string `json:"computeEnvironmentARN,omitempty"`

	ComputeEnvironmentName *string `json:"computeEnvironmentName,omitempty"`
	// An object representing an Batch compute resource. For more information, see
	// Compute Environments (https://docs.aws.amazon.com/batch/latest/userguide/compute_environments.html)
	// in the Batch User Guide.
	ComputeResources *ComputeResource `json:"computeResources,omitempty"`

	EcsClusterARN *string `json:"ecsClusterARN,omitempty"`

	ServiceRole *string `json:"serviceRole,omitempty"`

	State *string `json:"state,omitempty"`

	Status *string `json:"status,omitempty"`

	StatusReason *string `json:"statusReason,omitempty"`

	Tags map[string]*string `json:"tags,omitempty"`

	Type *string `json:"type_,omitempty"`
}

// +kubebuilder:skipversion
type ComputeEnvironmentOrder struct {
	ComputeEnvironment *string `json:"computeEnvironment,omitempty"`

	Order *int64 `json:"order,omitempty"`
}

// +kubebuilder:skipversion
type ComputeResource struct {
	AllocationStrategy *string `json:"allocationStrategy,omitempty"`

	BidPercentage *int64 `json:"bidPercentage,omitempty"`

	EC2Configuration []*EC2Configuration `json:"ec2Configuration,omitempty"`

	EC2KeyPair *string `json:"ec2KeyPair,omitempty"`

	ImageID *string `json:"imageID,omitempty"`

	InstanceTypes []*string `json:"instanceTypes,omitempty"`
	// An object representing a launch template associated with a compute resource.
	// You must specify either the launch template ID or launch template name in
	// the request, but not both.
	//
	// If security groups are specified using both the securityGroupIds parameter
	// of CreateComputeEnvironment and the launch template, the values in the securityGroupIds
	// parameter of CreateComputeEnvironment will be used.
	//
	// This object isn't applicable to jobs that are running on Fargate resources.
	LaunchTemplate *LaunchTemplateSpecification `json:"launchTemplate,omitempty"`

	MaxvCPUs *int64 `json:"maxvCPUs,omitempty"`

	MinvCPUs *int64 `json:"minvCPUs,omitempty"`

	PlacementGroup *string `json:"placementGroup,omitempty"`

	Tags map[string]*string `json:"tags,omitempty"`

	Type *string `json:"type_,omitempty"`
}

// +kubebuilder:skipversion
type ComputeResourceUpdate struct {
	MaxvCPUs *int64 `json:"maxvCPUs,omitempty"`

	MinvCPUs *int64 `json:"minvCPUs,omitempty"`
}

// +kubebuilder:skipversion
type ContainerDetail struct {
	Command []*string `json:"command,omitempty"`

	ContainerInstanceARN *string `json:"containerInstanceARN,omitempty"`

	ExecutionRoleARN *string `json:"executionRoleARN,omitempty"`

	ExitCode *int64 `json:"exitCode,omitempty"`

	Image *string `json:"image,omitempty"`

	InstanceType *string `json:"instanceType,omitempty"`

	JobRoleARN *string `json:"jobRoleARN,omitempty"`

	LogStreamName *string `json:"logStreamName,omitempty"`

	Memory *int64 `json:"memory,omitempty"`

	Reason *string `json:"reason,omitempty"`

	TaskARN *string `json:"taskARN,omitempty"`

	User *string `json:"user,omitempty"`

	Vcpus *int64 `json:"vcpus,omitempty"`
}

// +kubebuilder:skipversion
type ContainerOverrides struct {
	Command []*string `json:"command,omitempty"`

	InstanceType *string `json:"instanceType,omitempty"`

	Memory *int64 `json:"memory,omitempty"`

	Vcpus *int64 `json:"vcpus,omitempty"`
}

// +kubebuilder:skipversion
type ContainerProperties struct {
	Command []*string `json:"command,omitempty"`

	ExecutionRoleARN *string `json:"executionRoleARN,omitempty"`

	Image *string `json:"image,omitempty"`

	InstanceType *string `json:"instanceType,omitempty"`

	JobRoleARN *string `json:"jobRoleARN,omitempty"`

	Memory *int64 `json:"memory,omitempty"`

	User *string `json:"user,omitempty"`

	Vcpus *int64 `json:"vcpus,omitempty"`
}

// +kubebuilder:skipversion
type ContainerSummary struct {
	ExitCode *int64 `json:"exitCode,omitempty"`

	Reason *string `json:"reason,omitempty"`
}

// +kubebuilder:skipversion
type Device struct {
	ContainerPath *string `json:"containerPath,omitempty"`

	HostPath *string `json:"hostPath,omitempty"`
}

// +kubebuilder:skipversion
type EC2Configuration struct {
	ImageIDOverride *string `json:"imageIDOverride,omitempty"`

	ImageType *string `json:"imageType,omitempty"`
}

// +kubebuilder:skipversion
type EFSAuthorizationConfig struct {
	AccessPointID *string `json:"accessPointID,omitempty"`
}

// +kubebuilder:skipversion
type EFSVolumeConfiguration struct {
	FileSystemID *string `json:"fileSystemID,omitempty"`

	RootDirectory *string `json:"rootDirectory,omitempty"`

	TransitEncryptionPort *int64 `json:"transitEncryptionPort,omitempty"`
}

// +kubebuilder:skipversion
type EvaluateOnExit struct {
	OnExitCode *string `json:"onExitCode,omitempty"`

	OnReason *string `json:"onReason,omitempty"`

	OnStatusReason *string `json:"onStatusReason,omitempty"`
}

// +kubebuilder:skipversion
type FargatePlatformConfiguration struct {
	PlatformVersion *string `json:"platformVersion,omitempty"`
}

// +kubebuilder:skipversion
type Host struct {
	SourcePath *string `json:"sourcePath,omitempty"`
}

// +kubebuilder:skipversion
type JobDependency struct {
	JobID *string `json:"jobID,omitempty"`
}

// +kubebuilder:skipversion
type JobDetail struct {
	JobARN *string `json:"jobARN,omitempty"`

	JobDefinition *string `json:"jobDefinition,omitempty"`

	JobID *string `json:"jobID,omitempty"`

	JobName *string `json:"jobName,omitempty"`

	JobQueue *string `json:"jobQueue,omitempty"`

	StatusReason *string `json:"statusReason,omitempty"`

	Tags map[string]*string `json:"tags,omitempty"`
}

// +kubebuilder:skipversion
type JobQueueDetail struct {
	ComputeEnvironmentOrder []*ComputeEnvironmentOrder `json:"computeEnvironmentOrder,omitempty"`

	JobQueueARN *string `json:"jobQueueARN,omitempty"`

	JobQueueName *string `json:"jobQueueName,omitempty"`

	Priority *int64 `json:"priority,omitempty"`

	State *string `json:"state,omitempty"`

	Status *string `json:"status,omitempty"`

	StatusReason *string `json:"statusReason,omitempty"`

	Tags map[string]*string `json:"tags,omitempty"`
}

// +kubebuilder:skipversion
type JobSummary struct {
	JobARN *string `json:"jobARN,omitempty"`

	JobDefinition *string `json:"jobDefinition,omitempty"`

	JobID *string `json:"jobID,omitempty"`

	JobName *string `json:"jobName,omitempty"`

	StatusReason *string `json:"statusReason,omitempty"`
}

// +kubebuilder:skipversion
type JobTimeout struct {
	AttemptDurationSeconds *int64 `json:"attemptDurationSeconds,omitempty"`
}

// +kubebuilder:skipversion
type KeyValuePair struct {
	Name *string `json:"name,omitempty"`

	Value *string `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type KeyValuesPair struct {
	Name *string `json:"name,omitempty"`

	Values []*string `json:"values,omitempty"`
}

// +kubebuilder:skipversion
type LaunchTemplateSpecification struct {
	LaunchTemplateID *string `json:"launchTemplateID,omitempty"`

	LaunchTemplateName *string `json:"launchTemplateName,omitempty"`

	Version *string `json:"version,omitempty"`
}

// +kubebuilder:skipversion
type LinuxParameters struct {
	MaxSwap *int64 `json:"maxSwap,omitempty"`

	SharedMemorySize *int64 `json:"sharedMemorySize,omitempty"`

	Swappiness *int64 `json:"swappiness,omitempty"`
}

// +kubebuilder:skipversion
type MountPoint struct {
	ContainerPath *string `json:"containerPath,omitempty"`

	SourceVolume *string `json:"sourceVolume,omitempty"`
}

// +kubebuilder:skipversion
type NetworkInterface struct {
	AttachmentID *string `json:"attachmentID,omitempty"`

	IPv6Address *string `json:"ipv6Address,omitempty"`

	PrivateIPv4Address *string `json:"privateIPv4Address,omitempty"`
}

// +kubebuilder:skipversion
type NodeDetails struct {
	NodeIndex *int64 `json:"nodeIndex,omitempty"`
}

// +kubebuilder:skipversion
type NodeOverrides struct {
	NumNodes *int64 `json:"numNodes,omitempty"`
}

// +kubebuilder:skipversion
type NodeProperties struct {
	MainNode *int64 `json:"mainNode,omitempty"`

	NumNodes *int64 `json:"numNodes,omitempty"`
}

// +kubebuilder:skipversion
type NodePropertiesSummary struct {
	NodeIndex *int64 `json:"nodeIndex,omitempty"`

	NumNodes *int64 `json:"numNodes,omitempty"`
}

// +kubebuilder:skipversion
type NodePropertyOverride struct {
	TargetNodes *string `json:"targetNodes,omitempty"`
}

// +kubebuilder:skipversion
type NodeRangeProperty struct {
	TargetNodes *string `json:"targetNodes,omitempty"`
}

// +kubebuilder:skipversion
type ResourceRequirement struct {
	Value *string `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type RetryStrategy struct {
	Attempts *int64 `json:"attempts,omitempty"`
}

// +kubebuilder:skipversion
type Secret struct {
	Name *string `json:"name,omitempty"`

	ValueFrom *string `json:"valueFrom,omitempty"`
}

// +kubebuilder:skipversion
type Tmpfs struct {
	ContainerPath *string `json:"containerPath,omitempty"`

	MountOptions []*string `json:"mountOptions,omitempty"`

	Size *int64 `json:"size,omitempty"`
}

// +kubebuilder:skipversion
type Ulimit struct {
	HardLimit *int64 `json:"hardLimit,omitempty"`

	Name *string `json:"name,omitempty"`

	SoftLimit *int64 `json:"softLimit,omitempty"`
}

// +kubebuilder:skipversion
type Volume struct {
	Name *string `json:"name,omitempty"`
}