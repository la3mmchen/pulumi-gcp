// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigee

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An `Environment Group attachment` in Apigee.
//
// To get more information about EnvgroupAttachment, see:
//
// * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.envgroups.attachments/create)
// * How-to Guides
//     * [Creating an environment](https://cloud.google.com/apigee/docs/api-platform/get-started/create-environment)
//
// ## Example Usage
//
// ## Import
//
// EnvgroupAttachment can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:apigee/envGroupAttachment:EnvGroupAttachment default {{envgroup_id}}/attachments/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:apigee/envGroupAttachment:EnvGroupAttachment default {{envgroup_id}}/{{name}}
// ```
type EnvGroupAttachment struct {
	pulumi.CustomResourceState

	// The Apigee environment group associated with the Apigee environment,
	// in the format `organizations/{{org_name}}/envgroups/{{envgroup_name}}`.
	EnvgroupId pulumi.StringOutput `pulumi:"envgroupId"`
	// The resource ID of the environment.
	Environment pulumi.StringOutput `pulumi:"environment"`
	// The name of the newly created attachment (output parameter).
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewEnvGroupAttachment registers a new resource with the given unique name, arguments, and options.
func NewEnvGroupAttachment(ctx *pulumi.Context,
	name string, args *EnvGroupAttachmentArgs, opts ...pulumi.ResourceOption) (*EnvGroupAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvgroupId == nil {
		return nil, errors.New("invalid value for required argument 'EnvgroupId'")
	}
	if args.Environment == nil {
		return nil, errors.New("invalid value for required argument 'Environment'")
	}
	var resource EnvGroupAttachment
	err := ctx.RegisterResource("gcp:apigee/envGroupAttachment:EnvGroupAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvGroupAttachment gets an existing EnvGroupAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvGroupAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvGroupAttachmentState, opts ...pulumi.ResourceOption) (*EnvGroupAttachment, error) {
	var resource EnvGroupAttachment
	err := ctx.ReadResource("gcp:apigee/envGroupAttachment:EnvGroupAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnvGroupAttachment resources.
type envGroupAttachmentState struct {
	// The Apigee environment group associated with the Apigee environment,
	// in the format `organizations/{{org_name}}/envgroups/{{envgroup_name}}`.
	EnvgroupId *string `pulumi:"envgroupId"`
	// The resource ID of the environment.
	Environment *string `pulumi:"environment"`
	// The name of the newly created attachment (output parameter).
	Name *string `pulumi:"name"`
}

type EnvGroupAttachmentState struct {
	// The Apigee environment group associated with the Apigee environment,
	// in the format `organizations/{{org_name}}/envgroups/{{envgroup_name}}`.
	EnvgroupId pulumi.StringPtrInput
	// The resource ID of the environment.
	Environment pulumi.StringPtrInput
	// The name of the newly created attachment (output parameter).
	Name pulumi.StringPtrInput
}

func (EnvGroupAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*envGroupAttachmentState)(nil)).Elem()
}

type envGroupAttachmentArgs struct {
	// The Apigee environment group associated with the Apigee environment,
	// in the format `organizations/{{org_name}}/envgroups/{{envgroup_name}}`.
	EnvgroupId string `pulumi:"envgroupId"`
	// The resource ID of the environment.
	Environment string `pulumi:"environment"`
}

// The set of arguments for constructing a EnvGroupAttachment resource.
type EnvGroupAttachmentArgs struct {
	// The Apigee environment group associated with the Apigee environment,
	// in the format `organizations/{{org_name}}/envgroups/{{envgroup_name}}`.
	EnvgroupId pulumi.StringInput
	// The resource ID of the environment.
	Environment pulumi.StringInput
}

func (EnvGroupAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*envGroupAttachmentArgs)(nil)).Elem()
}

type EnvGroupAttachmentInput interface {
	pulumi.Input

	ToEnvGroupAttachmentOutput() EnvGroupAttachmentOutput
	ToEnvGroupAttachmentOutputWithContext(ctx context.Context) EnvGroupAttachmentOutput
}

func (*EnvGroupAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvGroupAttachment)(nil))
}

func (i *EnvGroupAttachment) ToEnvGroupAttachmentOutput() EnvGroupAttachmentOutput {
	return i.ToEnvGroupAttachmentOutputWithContext(context.Background())
}

func (i *EnvGroupAttachment) ToEnvGroupAttachmentOutputWithContext(ctx context.Context) EnvGroupAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvGroupAttachmentOutput)
}

func (i *EnvGroupAttachment) ToEnvGroupAttachmentPtrOutput() EnvGroupAttachmentPtrOutput {
	return i.ToEnvGroupAttachmentPtrOutputWithContext(context.Background())
}

func (i *EnvGroupAttachment) ToEnvGroupAttachmentPtrOutputWithContext(ctx context.Context) EnvGroupAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvGroupAttachmentPtrOutput)
}

type EnvGroupAttachmentPtrInput interface {
	pulumi.Input

	ToEnvGroupAttachmentPtrOutput() EnvGroupAttachmentPtrOutput
	ToEnvGroupAttachmentPtrOutputWithContext(ctx context.Context) EnvGroupAttachmentPtrOutput
}

type envGroupAttachmentPtrType EnvGroupAttachmentArgs

func (*envGroupAttachmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvGroupAttachment)(nil))
}

func (i *envGroupAttachmentPtrType) ToEnvGroupAttachmentPtrOutput() EnvGroupAttachmentPtrOutput {
	return i.ToEnvGroupAttachmentPtrOutputWithContext(context.Background())
}

func (i *envGroupAttachmentPtrType) ToEnvGroupAttachmentPtrOutputWithContext(ctx context.Context) EnvGroupAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvGroupAttachmentPtrOutput)
}

// EnvGroupAttachmentArrayInput is an input type that accepts EnvGroupAttachmentArray and EnvGroupAttachmentArrayOutput values.
// You can construct a concrete instance of `EnvGroupAttachmentArrayInput` via:
//
//          EnvGroupAttachmentArray{ EnvGroupAttachmentArgs{...} }
type EnvGroupAttachmentArrayInput interface {
	pulumi.Input

	ToEnvGroupAttachmentArrayOutput() EnvGroupAttachmentArrayOutput
	ToEnvGroupAttachmentArrayOutputWithContext(context.Context) EnvGroupAttachmentArrayOutput
}

type EnvGroupAttachmentArray []EnvGroupAttachmentInput

func (EnvGroupAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*EnvGroupAttachment)(nil))
}

func (i EnvGroupAttachmentArray) ToEnvGroupAttachmentArrayOutput() EnvGroupAttachmentArrayOutput {
	return i.ToEnvGroupAttachmentArrayOutputWithContext(context.Background())
}

func (i EnvGroupAttachmentArray) ToEnvGroupAttachmentArrayOutputWithContext(ctx context.Context) EnvGroupAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvGroupAttachmentArrayOutput)
}

// EnvGroupAttachmentMapInput is an input type that accepts EnvGroupAttachmentMap and EnvGroupAttachmentMapOutput values.
// You can construct a concrete instance of `EnvGroupAttachmentMapInput` via:
//
//          EnvGroupAttachmentMap{ "key": EnvGroupAttachmentArgs{...} }
type EnvGroupAttachmentMapInput interface {
	pulumi.Input

	ToEnvGroupAttachmentMapOutput() EnvGroupAttachmentMapOutput
	ToEnvGroupAttachmentMapOutputWithContext(context.Context) EnvGroupAttachmentMapOutput
}

type EnvGroupAttachmentMap map[string]EnvGroupAttachmentInput

func (EnvGroupAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*EnvGroupAttachment)(nil))
}

func (i EnvGroupAttachmentMap) ToEnvGroupAttachmentMapOutput() EnvGroupAttachmentMapOutput {
	return i.ToEnvGroupAttachmentMapOutputWithContext(context.Background())
}

func (i EnvGroupAttachmentMap) ToEnvGroupAttachmentMapOutputWithContext(ctx context.Context) EnvGroupAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvGroupAttachmentMapOutput)
}

type EnvGroupAttachmentOutput struct {
	*pulumi.OutputState
}

func (EnvGroupAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvGroupAttachment)(nil))
}

func (o EnvGroupAttachmentOutput) ToEnvGroupAttachmentOutput() EnvGroupAttachmentOutput {
	return o
}

func (o EnvGroupAttachmentOutput) ToEnvGroupAttachmentOutputWithContext(ctx context.Context) EnvGroupAttachmentOutput {
	return o
}

func (o EnvGroupAttachmentOutput) ToEnvGroupAttachmentPtrOutput() EnvGroupAttachmentPtrOutput {
	return o.ToEnvGroupAttachmentPtrOutputWithContext(context.Background())
}

func (o EnvGroupAttachmentOutput) ToEnvGroupAttachmentPtrOutputWithContext(ctx context.Context) EnvGroupAttachmentPtrOutput {
	return o.ApplyT(func(v EnvGroupAttachment) *EnvGroupAttachment {
		return &v
	}).(EnvGroupAttachmentPtrOutput)
}

type EnvGroupAttachmentPtrOutput struct {
	*pulumi.OutputState
}

func (EnvGroupAttachmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvGroupAttachment)(nil))
}

func (o EnvGroupAttachmentPtrOutput) ToEnvGroupAttachmentPtrOutput() EnvGroupAttachmentPtrOutput {
	return o
}

func (o EnvGroupAttachmentPtrOutput) ToEnvGroupAttachmentPtrOutputWithContext(ctx context.Context) EnvGroupAttachmentPtrOutput {
	return o
}

type EnvGroupAttachmentArrayOutput struct{ *pulumi.OutputState }

func (EnvGroupAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvGroupAttachment)(nil))
}

func (o EnvGroupAttachmentArrayOutput) ToEnvGroupAttachmentArrayOutput() EnvGroupAttachmentArrayOutput {
	return o
}

func (o EnvGroupAttachmentArrayOutput) ToEnvGroupAttachmentArrayOutputWithContext(ctx context.Context) EnvGroupAttachmentArrayOutput {
	return o
}

func (o EnvGroupAttachmentArrayOutput) Index(i pulumi.IntInput) EnvGroupAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnvGroupAttachment {
		return vs[0].([]EnvGroupAttachment)[vs[1].(int)]
	}).(EnvGroupAttachmentOutput)
}

type EnvGroupAttachmentMapOutput struct{ *pulumi.OutputState }

func (EnvGroupAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EnvGroupAttachment)(nil))
}

func (o EnvGroupAttachmentMapOutput) ToEnvGroupAttachmentMapOutput() EnvGroupAttachmentMapOutput {
	return o
}

func (o EnvGroupAttachmentMapOutput) ToEnvGroupAttachmentMapOutputWithContext(ctx context.Context) EnvGroupAttachmentMapOutput {
	return o
}

func (o EnvGroupAttachmentMapOutput) MapIndex(k pulumi.StringInput) EnvGroupAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EnvGroupAttachment {
		return vs[0].(map[string]EnvGroupAttachment)[vs[1].(string)]
	}).(EnvGroupAttachmentOutput)
}

func init() {
	pulumi.RegisterOutputType(EnvGroupAttachmentOutput{})
	pulumi.RegisterOutputType(EnvGroupAttachmentPtrOutput{})
	pulumi.RegisterOutputType(EnvGroupAttachmentArrayOutput{})
	pulumi.RegisterOutputType(EnvGroupAttachmentMapOutput{})
}