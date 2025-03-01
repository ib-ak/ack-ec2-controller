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

package vpc_endpoint

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	acktypes "github.com/aws-controllers-k8s/runtime/pkg/types"

	svcapitypes "github.com/aws-controllers-k8s/ec2-controller/apis/v1alpha1"
)

// ResolveReferences finds if there are any Reference field(s) present
// inside AWSResource passed in the parameter and attempts to resolve
// those reference field(s) into target field(s).
// It returns an AWSResource with resolved reference(s), and an error if the
// passed AWSResource's reference field(s) cannot be resolved.
// This method also adds/updates the ConditionTypeReferencesResolved for the
// AWSResource.
func (rm *resourceManager) ResolveReferences(
	ctx context.Context,
	apiReader client.Reader,
	res acktypes.AWSResource,
) (acktypes.AWSResource, error) {
	namespace := res.MetaObject().GetNamespace()
	ko := rm.concreteResource(res).ko.DeepCopy()
	err := validateReferenceFields(ko)
	if err == nil {
		err = resolveReferenceForRouteTableIDs(ctx, apiReader, namespace, ko)
	}
	if err == nil {
		err = resolveReferenceForSecurityGroupIDs(ctx, apiReader, namespace, ko)
	}
	if err == nil {
		err = resolveReferenceForSubnetIDs(ctx, apiReader, namespace, ko)
	}
	if err == nil {
		err = resolveReferenceForVPCID(ctx, apiReader, namespace, ko)
	}

	// If there was an error while resolving any reference, reset all the
	// resolved values so that they do not get persisted inside etcd
	if err != nil {
		ko = rm.concreteResource(res).ko.DeepCopy()
	}
	if hasNonNilReferences(ko) {
		return ackcondition.WithReferencesResolvedCondition(&resource{ko}, err)
	}
	return &resource{ko}, err
}

// validateReferenceFields validates the reference field and corresponding
// identifier field.
func validateReferenceFields(ko *svcapitypes.VPCEndpoint) error {
	if ko.Spec.RouteTableRefs != nil && ko.Spec.RouteTableIDs != nil {
		return ackerr.ResourceReferenceAndIDNotSupportedFor("RouteTableIDs", "RouteTableRefs")
	}
	if ko.Spec.SecurityGroupRefs != nil && ko.Spec.SecurityGroupIDs != nil {
		return ackerr.ResourceReferenceAndIDNotSupportedFor("SecurityGroupIDs", "SecurityGroupRefs")
	}
	if ko.Spec.SubnetRefs != nil && ko.Spec.SubnetIDs != nil {
		return ackerr.ResourceReferenceAndIDNotSupportedFor("SubnetIDs", "SubnetRefs")
	}
	if ko.Spec.VPCRef != nil && ko.Spec.VPCID != nil {
		return ackerr.ResourceReferenceAndIDNotSupportedFor("VPCID", "VPCRef")
	}
	if ko.Spec.VPCRef == nil && ko.Spec.VPCID == nil {
		return ackerr.ResourceReferenceOrIDRequiredFor("VPCID", "VPCRef")
	}
	return nil
}

// hasNonNilReferences returns true if resource contains a reference to another
// resource
func hasNonNilReferences(ko *svcapitypes.VPCEndpoint) bool {
	return false || (ko.Spec.RouteTableRefs != nil) || (ko.Spec.SecurityGroupRefs != nil) || (ko.Spec.SubnetRefs != nil) || (ko.Spec.VPCRef != nil)
}

// resolveReferenceForRouteTableIDs reads the resource referenced
// from RouteTableRefs field and sets the RouteTableIDs
// from referenced resource
func resolveReferenceForRouteTableIDs(
	ctx context.Context,
	apiReader client.Reader,
	namespace string,
	ko *svcapitypes.VPCEndpoint,
) error {
	if len(ko.Spec.RouteTableRefs) > 0 {
		resolved0 := []*string{}
		for _, iter0 := range ko.Spec.RouteTableRefs {
			arr := iter0.From
			if arr == nil || arr.Name == nil || *arr.Name == "" {
				return fmt.Errorf("provided resource reference is nil or empty: RouteTableRefs")
			}
			obj := &svcapitypes.RouteTable{}
			if err := getReferencedResourceState_RouteTable(ctx, apiReader, obj, *arr.Name, namespace); err != nil {
				return err
			}
			resolved0 = append(resolved0, (*string)(obj.Status.RouteTableID))
		}
		ko.Spec.RouteTableIDs = resolved0
	}

	return nil
}

// getReferencedResourceState_RouteTable looks up whether a referenced resource
// exists and is in a ACK.ResourceSynced=True state. If the referenced resource does exist and is
// in a Synced state, returns nil, otherwise returns `ackerr.ResourceReferenceTerminalFor` or
// `ResourceReferenceNotSyncedFor` depending on if the resource is in a Terminal state.
func getReferencedResourceState_RouteTable(
	ctx context.Context,
	apiReader client.Reader,
	obj *svcapitypes.RouteTable,
	name string, // the Kubernetes name of the referenced resource
	namespace string, // the Kubernetes namespace of the referenced resource
) error {
	namespacedName := types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
	err := apiReader.Get(ctx, namespacedName, obj)
	if err != nil {
		return err
	}
	var refResourceSynced, refResourceTerminal bool
	for _, cond := range obj.Status.Conditions {
		if cond.Type == ackv1alpha1.ConditionTypeResourceSynced &&
			cond.Status == corev1.ConditionTrue {
			refResourceSynced = true
		}
		if cond.Type == ackv1alpha1.ConditionTypeTerminal &&
			cond.Status == corev1.ConditionTrue {
			return ackerr.ResourceReferenceTerminalFor(
				"RouteTable",
				namespace, name)
		}
	}
	if refResourceTerminal {
		return ackerr.ResourceReferenceTerminalFor(
			"RouteTable",
			namespace, name)
	}
	if !refResourceSynced {
		return ackerr.ResourceReferenceNotSyncedFor(
			"RouteTable",
			namespace, name)
	}
	if obj.Status.RouteTableID == nil {
		return ackerr.ResourceReferenceMissingTargetFieldFor(
			"RouteTable",
			namespace, name,
			"Status.RouteTableID")
	}
	return nil
}

// resolveReferenceForSecurityGroupIDs reads the resource referenced
// from SecurityGroupRefs field and sets the SecurityGroupIDs
// from referenced resource
func resolveReferenceForSecurityGroupIDs(
	ctx context.Context,
	apiReader client.Reader,
	namespace string,
	ko *svcapitypes.VPCEndpoint,
) error {
	if len(ko.Spec.SecurityGroupRefs) > 0 {
		resolved0 := []*string{}
		for _, iter0 := range ko.Spec.SecurityGroupRefs {
			arr := iter0.From
			if arr == nil || arr.Name == nil || *arr.Name == "" {
				return fmt.Errorf("provided resource reference is nil or empty: SecurityGroupRefs")
			}
			obj := &svcapitypes.SecurityGroup{}
			if err := getReferencedResourceState_SecurityGroup(ctx, apiReader, obj, *arr.Name, namespace); err != nil {
				return err
			}
			resolved0 = append(resolved0, (*string)(obj.Status.ID))
		}
		ko.Spec.SecurityGroupIDs = resolved0
	}

	return nil
}

// getReferencedResourceState_SecurityGroup looks up whether a referenced resource
// exists and is in a ACK.ResourceSynced=True state. If the referenced resource does exist and is
// in a Synced state, returns nil, otherwise returns `ackerr.ResourceReferenceTerminalFor` or
// `ResourceReferenceNotSyncedFor` depending on if the resource is in a Terminal state.
func getReferencedResourceState_SecurityGroup(
	ctx context.Context,
	apiReader client.Reader,
	obj *svcapitypes.SecurityGroup,
	name string, // the Kubernetes name of the referenced resource
	namespace string, // the Kubernetes namespace of the referenced resource
) error {
	namespacedName := types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
	err := apiReader.Get(ctx, namespacedName, obj)
	if err != nil {
		return err
	}
	var refResourceSynced, refResourceTerminal bool
	for _, cond := range obj.Status.Conditions {
		if cond.Type == ackv1alpha1.ConditionTypeResourceSynced &&
			cond.Status == corev1.ConditionTrue {
			refResourceSynced = true
		}
		if cond.Type == ackv1alpha1.ConditionTypeTerminal &&
			cond.Status == corev1.ConditionTrue {
			return ackerr.ResourceReferenceTerminalFor(
				"SecurityGroup",
				namespace, name)
		}
	}
	if refResourceTerminal {
		return ackerr.ResourceReferenceTerminalFor(
			"SecurityGroup",
			namespace, name)
	}
	if !refResourceSynced {
		return ackerr.ResourceReferenceNotSyncedFor(
			"SecurityGroup",
			namespace, name)
	}
	if obj.Status.ID == nil {
		return ackerr.ResourceReferenceMissingTargetFieldFor(
			"SecurityGroup",
			namespace, name,
			"Status.ID")
	}
	return nil
}

// resolveReferenceForSubnetIDs reads the resource referenced
// from SubnetRefs field and sets the SubnetIDs
// from referenced resource
func resolveReferenceForSubnetIDs(
	ctx context.Context,
	apiReader client.Reader,
	namespace string,
	ko *svcapitypes.VPCEndpoint,
) error {
	if len(ko.Spec.SubnetRefs) > 0 {
		resolved0 := []*string{}
		for _, iter0 := range ko.Spec.SubnetRefs {
			arr := iter0.From
			if arr == nil || arr.Name == nil || *arr.Name == "" {
				return fmt.Errorf("provided resource reference is nil or empty: SubnetRefs")
			}
			obj := &svcapitypes.Subnet{}
			if err := getReferencedResourceState_Subnet(ctx, apiReader, obj, *arr.Name, namespace); err != nil {
				return err
			}
			resolved0 = append(resolved0, (*string)(obj.Status.SubnetID))
		}
		ko.Spec.SubnetIDs = resolved0
	}

	return nil
}

// getReferencedResourceState_Subnet looks up whether a referenced resource
// exists and is in a ACK.ResourceSynced=True state. If the referenced resource does exist and is
// in a Synced state, returns nil, otherwise returns `ackerr.ResourceReferenceTerminalFor` or
// `ResourceReferenceNotSyncedFor` depending on if the resource is in a Terminal state.
func getReferencedResourceState_Subnet(
	ctx context.Context,
	apiReader client.Reader,
	obj *svcapitypes.Subnet,
	name string, // the Kubernetes name of the referenced resource
	namespace string, // the Kubernetes namespace of the referenced resource
) error {
	namespacedName := types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
	err := apiReader.Get(ctx, namespacedName, obj)
	if err != nil {
		return err
	}
	var refResourceSynced, refResourceTerminal bool
	for _, cond := range obj.Status.Conditions {
		if cond.Type == ackv1alpha1.ConditionTypeResourceSynced &&
			cond.Status == corev1.ConditionTrue {
			refResourceSynced = true
		}
		if cond.Type == ackv1alpha1.ConditionTypeTerminal &&
			cond.Status == corev1.ConditionTrue {
			return ackerr.ResourceReferenceTerminalFor(
				"Subnet",
				namespace, name)
		}
	}
	if refResourceTerminal {
		return ackerr.ResourceReferenceTerminalFor(
			"Subnet",
			namespace, name)
	}
	if !refResourceSynced {
		return ackerr.ResourceReferenceNotSyncedFor(
			"Subnet",
			namespace, name)
	}
	if obj.Status.SubnetID == nil {
		return ackerr.ResourceReferenceMissingTargetFieldFor(
			"Subnet",
			namespace, name,
			"Status.SubnetID")
	}
	return nil
}

// resolveReferenceForVPCID reads the resource referenced
// from VPCRef field and sets the VPCID
// from referenced resource
func resolveReferenceForVPCID(
	ctx context.Context,
	apiReader client.Reader,
	namespace string,
	ko *svcapitypes.VPCEndpoint,
) error {
	if ko.Spec.VPCRef != nil && ko.Spec.VPCRef.From != nil {
		arr := ko.Spec.VPCRef.From
		if arr == nil || arr.Name == nil || *arr.Name == "" {
			return fmt.Errorf("provided resource reference is nil or empty: VPCRef")
		}
		obj := &svcapitypes.VPC{}
		if err := getReferencedResourceState_VPC(ctx, apiReader, obj, *arr.Name, namespace); err != nil {
			return err
		}
		ko.Spec.VPCID = (*string)(obj.Status.VPCID)
	}

	return nil
}

// getReferencedResourceState_VPC looks up whether a referenced resource
// exists and is in a ACK.ResourceSynced=True state. If the referenced resource does exist and is
// in a Synced state, returns nil, otherwise returns `ackerr.ResourceReferenceTerminalFor` or
// `ResourceReferenceNotSyncedFor` depending on if the resource is in a Terminal state.
func getReferencedResourceState_VPC(
	ctx context.Context,
	apiReader client.Reader,
	obj *svcapitypes.VPC,
	name string, // the Kubernetes name of the referenced resource
	namespace string, // the Kubernetes namespace of the referenced resource
) error {
	namespacedName := types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
	err := apiReader.Get(ctx, namespacedName, obj)
	if err != nil {
		return err
	}
	var refResourceSynced, refResourceTerminal bool
	for _, cond := range obj.Status.Conditions {
		if cond.Type == ackv1alpha1.ConditionTypeResourceSynced &&
			cond.Status == corev1.ConditionTrue {
			refResourceSynced = true
		}
		if cond.Type == ackv1alpha1.ConditionTypeTerminal &&
			cond.Status == corev1.ConditionTrue {
			return ackerr.ResourceReferenceTerminalFor(
				"VPC",
				namespace, name)
		}
	}
	if refResourceTerminal {
		return ackerr.ResourceReferenceTerminalFor(
			"VPC",
			namespace, name)
	}
	if !refResourceSynced {
		return ackerr.ResourceReferenceNotSyncedFor(
			"VPC",
			namespace, name)
	}
	if obj.Status.VPCID == nil {
		return ackerr.ResourceReferenceMissingTargetFieldFor(
			"VPC",
			namespace, name,
			"Status.VPCID")
	}
	return nil
}
