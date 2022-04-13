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

package nat_gateway

import (
	"context"
	"errors"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/ec2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/ec2-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.EC2{}
	_ = &svcapitypes.NATGateway{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer exit(err)
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadManyInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newListRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DescribeNatGatewaysOutput
	resp, err = rm.sdkapi.DescribeNatGatewaysWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_MANY", "DescribeNatGateways", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "UNKNOWN" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	found := false
	for _, elem := range resp.NatGateways {
		if elem.ConnectivityType != nil {
			ko.Spec.ConnectivityType = elem.ConnectivityType
		} else {
			ko.Spec.ConnectivityType = nil
		}
		if elem.CreateTime != nil {
			ko.Status.CreateTime = &metav1.Time{*elem.CreateTime}
		} else {
			ko.Status.CreateTime = nil
		}
		if elem.DeleteTime != nil {
			ko.Status.DeleteTime = &metav1.Time{*elem.DeleteTime}
		} else {
			ko.Status.DeleteTime = nil
		}
		if elem.FailureCode != nil {
			ko.Status.FailureCode = elem.FailureCode
		} else {
			ko.Status.FailureCode = nil
		}
		if elem.FailureMessage != nil {
			ko.Status.FailureMessage = elem.FailureMessage
		} else {
			ko.Status.FailureMessage = nil
		}
		if elem.NatGatewayAddresses != nil {
			f5 := []*svcapitypes.NATGatewayAddress{}
			for _, f5iter := range elem.NatGatewayAddresses {
				f5elem := &svcapitypes.NATGatewayAddress{}
				if f5iter.AllocationId != nil {
					f5elem.AllocationID = f5iter.AllocationId
				}
				if f5iter.NetworkInterfaceId != nil {
					f5elem.NetworkInterfaceID = f5iter.NetworkInterfaceId
				}
				if f5iter.PrivateIp != nil {
					f5elem.PrivateIP = f5iter.PrivateIp
				}
				if f5iter.PublicIp != nil {
					f5elem.PublicIP = f5iter.PublicIp
				}
				f5 = append(f5, f5elem)
			}
			ko.Status.NATGatewayAddresses = f5
		} else {
			ko.Status.NATGatewayAddresses = nil
		}
		if elem.NatGatewayId != nil {
			ko.Status.NATGatewayID = elem.NatGatewayId
		} else {
			ko.Status.NATGatewayID = nil
		}
		if elem.ProvisionedBandwidth != nil {
			f7 := &svcapitypes.ProvisionedBandwidth{}
			if elem.ProvisionedBandwidth.ProvisionTime != nil {
				f7.ProvisionTime = &metav1.Time{*elem.ProvisionedBandwidth.ProvisionTime}
			}
			if elem.ProvisionedBandwidth.Provisioned != nil {
				f7.Provisioned = elem.ProvisionedBandwidth.Provisioned
			}
			if elem.ProvisionedBandwidth.RequestTime != nil {
				f7.RequestTime = &metav1.Time{*elem.ProvisionedBandwidth.RequestTime}
			}
			if elem.ProvisionedBandwidth.Requested != nil {
				f7.Requested = elem.ProvisionedBandwidth.Requested
			}
			if elem.ProvisionedBandwidth.Status != nil {
				f7.Status = elem.ProvisionedBandwidth.Status
			}
			ko.Status.ProvisionedBandwidth = f7
		} else {
			ko.Status.ProvisionedBandwidth = nil
		}
		if elem.State != nil {
			ko.Status.State = elem.State
		} else {
			ko.Status.State = nil
		}
		if elem.SubnetId != nil {
			ko.Spec.SubnetID = elem.SubnetId
		} else {
			ko.Spec.SubnetID = nil
		}
		if elem.Tags != nil {
			f10 := []*svcapitypes.Tag{}
			for _, f10iter := range elem.Tags {
				f10elem := &svcapitypes.Tag{}
				if f10iter.Key != nil {
					f10elem.Key = f10iter.Key
				}
				if f10iter.Value != nil {
					f10elem.Value = f10iter.Value
				}
				f10 = append(f10, f10elem)
			}
			ko.Status.Tags = f10
		} else {
			ko.Status.Tags = nil
		}
		if elem.VpcId != nil {
			ko.Status.VPCID = elem.VpcId
		} else {
			ko.Status.VPCID = nil
		}
		found = true
		break
	}
	if !found {
		return nil, ackerr.NotFound
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadManyInput returns true if there are any fields
// for the ReadMany Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadManyInput(
	r *resource,
) bool {
	return r.ko.Status.NATGatewayID == nil

}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.DescribeNatGatewaysInput, error) {
	res := &svcsdk.DescribeNatGatewaysInput{}

	if r.ko.Status.NATGatewayID != nil {
		f3 := []*string{}
		f3 = append(f3, r.ko.Status.NATGatewayID)
		res.SetNatGatewayIds(f3)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer exit(err)
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateNatGatewayOutput
	_ = resp
	resp, err = rm.sdkapi.CreateNatGatewayWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateNatGateway", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.NatGateway.ConnectivityType != nil {
		ko.Spec.ConnectivityType = resp.NatGateway.ConnectivityType
	} else {
		ko.Spec.ConnectivityType = nil
	}
	if resp.NatGateway.CreateTime != nil {
		ko.Status.CreateTime = &metav1.Time{*resp.NatGateway.CreateTime}
	} else {
		ko.Status.CreateTime = nil
	}
	if resp.NatGateway.DeleteTime != nil {
		ko.Status.DeleteTime = &metav1.Time{*resp.NatGateway.DeleteTime}
	} else {
		ko.Status.DeleteTime = nil
	}
	if resp.NatGateway.FailureCode != nil {
		ko.Status.FailureCode = resp.NatGateway.FailureCode
	} else {
		ko.Status.FailureCode = nil
	}
	if resp.NatGateway.FailureMessage != nil {
		ko.Status.FailureMessage = resp.NatGateway.FailureMessage
	} else {
		ko.Status.FailureMessage = nil
	}
	if resp.NatGateway.NatGatewayAddresses != nil {
		f5 := []*svcapitypes.NATGatewayAddress{}
		for _, f5iter := range resp.NatGateway.NatGatewayAddresses {
			f5elem := &svcapitypes.NATGatewayAddress{}
			if f5iter.AllocationId != nil {
				f5elem.AllocationID = f5iter.AllocationId
			}
			if f5iter.NetworkInterfaceId != nil {
				f5elem.NetworkInterfaceID = f5iter.NetworkInterfaceId
			}
			if f5iter.PrivateIp != nil {
				f5elem.PrivateIP = f5iter.PrivateIp
			}
			if f5iter.PublicIp != nil {
				f5elem.PublicIP = f5iter.PublicIp
			}
			f5 = append(f5, f5elem)
		}
		ko.Status.NATGatewayAddresses = f5
	} else {
		ko.Status.NATGatewayAddresses = nil
	}
	if resp.NatGateway.NatGatewayId != nil {
		ko.Status.NATGatewayID = resp.NatGateway.NatGatewayId
	} else {
		ko.Status.NATGatewayID = nil
	}
	if resp.NatGateway.ProvisionedBandwidth != nil {
		f7 := &svcapitypes.ProvisionedBandwidth{}
		if resp.NatGateway.ProvisionedBandwidth.ProvisionTime != nil {
			f7.ProvisionTime = &metav1.Time{*resp.NatGateway.ProvisionedBandwidth.ProvisionTime}
		}
		if resp.NatGateway.ProvisionedBandwidth.Provisioned != nil {
			f7.Provisioned = resp.NatGateway.ProvisionedBandwidth.Provisioned
		}
		if resp.NatGateway.ProvisionedBandwidth.RequestTime != nil {
			f7.RequestTime = &metav1.Time{*resp.NatGateway.ProvisionedBandwidth.RequestTime}
		}
		if resp.NatGateway.ProvisionedBandwidth.Requested != nil {
			f7.Requested = resp.NatGateway.ProvisionedBandwidth.Requested
		}
		if resp.NatGateway.ProvisionedBandwidth.Status != nil {
			f7.Status = resp.NatGateway.ProvisionedBandwidth.Status
		}
		ko.Status.ProvisionedBandwidth = f7
	} else {
		ko.Status.ProvisionedBandwidth = nil
	}
	if resp.NatGateway.State != nil {
		ko.Status.State = resp.NatGateway.State
	} else {
		ko.Status.State = nil
	}
	if resp.NatGateway.SubnetId != nil {
		ko.Spec.SubnetID = resp.NatGateway.SubnetId
	} else {
		ko.Spec.SubnetID = nil
	}
	if resp.NatGateway.Tags != nil {
		f10 := []*svcapitypes.Tag{}
		for _, f10iter := range resp.NatGateway.Tags {
			f10elem := &svcapitypes.Tag{}
			if f10iter.Key != nil {
				f10elem.Key = f10iter.Key
			}
			if f10iter.Value != nil {
				f10elem.Value = f10iter.Value
			}
			f10 = append(f10, f10elem)
		}
		ko.Status.Tags = f10
	} else {
		ko.Status.Tags = nil
	}
	if resp.NatGateway.VpcId != nil {
		ko.Status.VPCID = resp.NatGateway.VpcId
	} else {
		ko.Status.VPCID = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateNatGatewayInput, error) {
	res := &svcsdk.CreateNatGatewayInput{}

	if r.ko.Spec.AllocationID != nil {
		res.SetAllocationId(*r.ko.Spec.AllocationID)
	}
	if r.ko.Spec.ConnectivityType != nil {
		res.SetConnectivityType(*r.ko.Spec.ConnectivityType)
	}
	if r.ko.Spec.SubnetID != nil {
		res.SetSubnetId(*r.ko.Spec.SubnetID)
	}
	if r.ko.Spec.TagSpecifications != nil {
		f3 := []*svcsdk.TagSpecification{}
		for _, f3iter := range r.ko.Spec.TagSpecifications {
			f3elem := &svcsdk.TagSpecification{}
			if f3iter.ResourceType != nil {
				f3elem.SetResourceType(*f3iter.ResourceType)
			}
			if f3iter.Tags != nil {
				f3elemf1 := []*svcsdk.Tag{}
				for _, f3elemf1iter := range f3iter.Tags {
					f3elemf1elem := &svcsdk.Tag{}
					if f3elemf1iter.Key != nil {
						f3elemf1elem.SetKey(*f3elemf1iter.Key)
					}
					if f3elemf1iter.Value != nil {
						f3elemf1elem.SetValue(*f3elemf1iter.Value)
					}
					f3elemf1 = append(f3elemf1, f3elemf1elem)
				}
				f3elem.SetTags(f3elemf1)
			}
			f3 = append(f3, f3elem)
		}
		res.SetTagSpecifications(f3)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {
	// TODO(jaypipes): Figure this out...
	return nil, ackerr.NotImplemented
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer exit(err)
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteNatGatewayOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteNatGatewayWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteNatGateway", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteNatGatewayInput, error) {
	res := &svcsdk.DeleteNatGatewayInput{}

	if r.ko.Status.NATGatewayID != nil {
		res.SetNatGatewayId(*r.ko.Status.NATGatewayID)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.NATGateway,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.Region == nil {
		ko.Status.ACKResourceMetadata.Region = &rm.awsRegion
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}
	var termError *ackerr.TerminalError
	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	// No terminal_errors specified for this resource in generator config
	return false
}
