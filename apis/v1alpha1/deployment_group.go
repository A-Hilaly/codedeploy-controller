// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeploymentGroupSpec defines the desired state of DeploymentGroup.
type DeploymentGroupSpec struct {
	// Information to add about Amazon CloudWatch alarms when the deployment group
// is created.
	 AlarmConfiguration *AlarmConfiguration `json:"alarmConfiguration,omitempty"` 
	// The name of an AWS CodeDeploy application associated with the IAM user or
// AWS account.
	 // +kubebuilder:validation:Required
	ApplicationName *string `json:"applicationName"`
	// Configuration information for an automatic rollback that is added when a
// deployment group is created.
	 AutoRollbackConfiguration *AutoRollbackConfiguration `json:"autoRollbackConfiguration,omitempty"` 
	// A list of associated Amazon EC2 Auto Scaling groups.
	 AutoScalingGroups []*string `json:"autoScalingGroups,omitempty"` 
	// Information about blue/green deployment options for a deployment group.
	 BlueGreenDeploymentConfiguration *BlueGreenDeploymentConfiguration `json:"blueGreenDeploymentConfiguration,omitempty"` 
	// If specified, the deployment configuration name can be either one of the
// predefined configurations provided with AWS CodeDeploy or a custom deployment
// configuration that you create by calling the create deployment configuration
// operation.
// 
// CodeDeployDefault.OneAtATime is the default deployment configuration. It
// is used if a configuration isn't specified for the deployment or deployment
// group.
// 
// For more information about the predefined deployment configurations in AWS
// CodeDeploy, see Working with Deployment Configurations in CodeDeploy (https://docs.aws.amazon.com/codedeploy/latest/userguide/deployment-configurations.html)
// in the AWS CodeDeploy User Guide.
	 DeploymentConfigName *string `json:"deploymentConfigName,omitempty"` 
	// The name of a new deployment group for the specified application.
	 // +kubebuilder:validation:Required
	DeploymentGroupName *string `json:"deploymentGroupName"`
	// Information about the type of deployment, in-place or blue/green, that you
// want to run and whether to route deployment traffic behind a load balancer.
	 DeploymentStyle *DeploymentStyle `json:"deploymentStyle,omitempty"` 
	// The Amazon EC2 tags on which to filter. The deployment group includes EC2
// instances with any of the specified tags. Cannot be used in the same call
// as ec2TagSet.
	 EC2TagFilters []*EC2TagFilter `json:"ec2TagFilters,omitempty"` 
	// Information about groups of tags applied to EC2 instances. The deployment
// group includes only EC2 instances identified by all the tag groups. Cannot
// be used in the same call as ec2TagFilters.
	 EC2TagSet *EC2TagSet `json:"ec2TagSet,omitempty"` 
	// The target Amazon ECS services in the deployment group. This applies only
// to deployment groups that use the Amazon ECS compute platform. A target Amazon
// ECS service is specified as an Amazon ECS cluster and service name pair using
// the format <clustername>:<servicename>.
	 EcsServices []*ECSService `json:"ecsServices,omitempty"` 
	// Information about the load balancer used in a deployment.
	 LoadBalancerInfo *LoadBalancerInfo `json:"loadBalancerInfo,omitempty"` 
	// The on-premises instance tags on which to filter. The deployment group includes
// on-premises instances with any of the specified tags. Cannot be used in the
// same call as OnPremisesTagSet.
	 OnPremisesInstanceTagFilters []*TagFilter `json:"onPremisesInstanceTagFilters,omitempty"` 
	// Information about groups of tags applied to on-premises instances. The deployment
// group includes only on-premises instances identified by all of the tag groups.
// Cannot be used in the same call as onPremisesInstanceTagFilters.
	 OnPremisesTagSet *OnPremisesTagSet `json:"onPremisesTagSet,omitempty"` 
	// Indicates what happens when new EC2 instances are launched mid-deployment
// and do not receive the deployed application revision.
// 
// If this option is set to UPDATE or is unspecified, CodeDeploy initiates one
// or more 'auto-update outdated instances' deployments to apply the deployed
// application revision to the new EC2 instances.
// 
// If this option is set to IGNORE, CodeDeploy does not initiate a deployment
// to update the new EC2 instances. This may result in instances having different
// revisions.
	 OutdatedInstancesStrategy *string `json:"outdatedInstancesStrategy,omitempty"` 
	// A service role Amazon Resource Name (ARN) that allows AWS CodeDeploy to act
// on the user's behalf when interacting with AWS services.
	 // +kubebuilder:validation:Required
	ServiceRoleARN *string `json:"serviceRoleARN"`
	// The metadata that you apply to CodeDeploy deployment groups to help you organize
// and categorize them. Each tag consists of a key and an optional value, both
// of which you define.
	 Tags []*Tag `json:"tags,omitempty"` 
	// Information about triggers to create when the deployment group is created.
// For examples, see Create a Trigger for an AWS CodeDeploy Event (https://docs.aws.amazon.com/codedeploy/latest/userguide/how-to-notify-sns.html)
// in the AWS CodeDeploy User Guide.
	 TriggerConfigurations []*TriggerConfig `json:"triggerConfigurations,omitempty"` 
}

// DeploymentGroupStatus defines the observed state of DeploymentGroup
type DeploymentGroupStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// A unique deployment group ID.
	DeploymentGroupID *string `json:"deploymentGroupID,omitempty"`
}

// DeploymentGroup is the Schema for the DeploymentGroups API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type DeploymentGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   DeploymentGroupSpec   `json:"spec,omitempty"`
	Status DeploymentGroupStatus `json:"status,omitempty"`
}

// DeploymentGroupList contains a list of DeploymentGroup
// +kubebuilder:object:root=true
type DeploymentGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items []DeploymentGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DeploymentGroup{}, &DeploymentGroupList{})
}