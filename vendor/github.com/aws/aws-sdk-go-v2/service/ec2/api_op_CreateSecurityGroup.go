// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a security group.
//
// A security group acts as a virtual firewall for your instance to control
// inbound and outbound traffic. For more information, see [Amazon EC2 security groups]in the Amazon Elastic
// Compute Cloud User Guide and [Security groups for your VPC]in the Amazon Virtual Private Cloud User Guide.
//
// When you create a security group, you specify a friendly name of your choice.
// You can't have two security groups for the same VPC with the same name.
//
// You have a default security group for use in your VPC. If you don't specify a
// security group when you launch an instance, the instance is launched into the
// appropriate default security group. A default security group includes a default
// rule that grants instances unrestricted network access to each other.
//
// You can add or remove rules from your security groups using AuthorizeSecurityGroupIngress, AuthorizeSecurityGroupEgress, RevokeSecurityGroupIngress, and RevokeSecurityGroupEgress.
//
// For more information about VPC security group limits, see [Amazon VPC Limits].
//
// [Amazon VPC Limits]: https://docs.aws.amazon.com/vpc/latest/userguide/amazon-vpc-limits.html
// [Amazon EC2 security groups]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-network-security.html
// [Security groups for your VPC]: https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_SecurityGroups.html
func (c *Client) CreateSecurityGroup(ctx context.Context, params *CreateSecurityGroupInput, optFns ...func(*Options)) (*CreateSecurityGroupOutput, error) {
	if params == nil {
		params = &CreateSecurityGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSecurityGroup", params, optFns, c.addOperationCreateSecurityGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSecurityGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSecurityGroupInput struct {

	// A description for the security group.
	//
	// Constraints: Up to 255 characters in length
	//
	// Valid characters: a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=&;{}!$*
	//
	// This member is required.
	Description *string

	// The name of the security group.
	//
	// Constraints: Up to 255 characters in length. Cannot start with sg- .
	//
	// Valid characters: a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=&;{}!$*
	//
	// This member is required.
	GroupName *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The tags to assign to the security group.
	TagSpecifications []types.TagSpecification

	// The ID of the VPC. Required for a nondefault VPC.
	VpcId *string

	noSmithyDocumentSerde
}

type CreateSecurityGroupOutput struct {

	// The ID of the security group.
	GroupId *string

	// The tags assigned to the security group.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSecurityGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateSecurityGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateSecurityGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateSecurityGroup"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpCreateSecurityGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSecurityGroup(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateSecurityGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateSecurityGroup",
	}
}