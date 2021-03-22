// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventarc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type TriggerDestination struct {
	// Cloud Run fully-managed service that receives the events. The service should be running in the same project as the trigger.
	// The `matchingCriteria` block supports:
	CloudRunService *TriggerDestinationCloudRunService `pulumi:"cloudRunService"`
}

// TriggerDestinationInput is an input type that accepts TriggerDestinationArgs and TriggerDestinationOutput values.
// You can construct a concrete instance of `TriggerDestinationInput` via:
//
//          TriggerDestinationArgs{...}
type TriggerDestinationInput interface {
	pulumi.Input

	ToTriggerDestinationOutput() TriggerDestinationOutput
	ToTriggerDestinationOutputWithContext(context.Context) TriggerDestinationOutput
}

type TriggerDestinationArgs struct {
	// Cloud Run fully-managed service that receives the events. The service should be running in the same project as the trigger.
	// The `matchingCriteria` block supports:
	CloudRunService TriggerDestinationCloudRunServicePtrInput `pulumi:"cloudRunService"`
}

func (TriggerDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerDestination)(nil)).Elem()
}

func (i TriggerDestinationArgs) ToTriggerDestinationOutput() TriggerDestinationOutput {
	return i.ToTriggerDestinationOutputWithContext(context.Background())
}

func (i TriggerDestinationArgs) ToTriggerDestinationOutputWithContext(ctx context.Context) TriggerDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerDestinationOutput)
}

func (i TriggerDestinationArgs) ToTriggerDestinationPtrOutput() TriggerDestinationPtrOutput {
	return i.ToTriggerDestinationPtrOutputWithContext(context.Background())
}

func (i TriggerDestinationArgs) ToTriggerDestinationPtrOutputWithContext(ctx context.Context) TriggerDestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerDestinationOutput).ToTriggerDestinationPtrOutputWithContext(ctx)
}

// TriggerDestinationPtrInput is an input type that accepts TriggerDestinationArgs, TriggerDestinationPtr and TriggerDestinationPtrOutput values.
// You can construct a concrete instance of `TriggerDestinationPtrInput` via:
//
//          TriggerDestinationArgs{...}
//
//  or:
//
//          nil
type TriggerDestinationPtrInput interface {
	pulumi.Input

	ToTriggerDestinationPtrOutput() TriggerDestinationPtrOutput
	ToTriggerDestinationPtrOutputWithContext(context.Context) TriggerDestinationPtrOutput
}

type triggerDestinationPtrType TriggerDestinationArgs

func TriggerDestinationPtr(v *TriggerDestinationArgs) TriggerDestinationPtrInput {
	return (*triggerDestinationPtrType)(v)
}

func (*triggerDestinationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerDestination)(nil)).Elem()
}

func (i *triggerDestinationPtrType) ToTriggerDestinationPtrOutput() TriggerDestinationPtrOutput {
	return i.ToTriggerDestinationPtrOutputWithContext(context.Background())
}

func (i *triggerDestinationPtrType) ToTriggerDestinationPtrOutputWithContext(ctx context.Context) TriggerDestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerDestinationPtrOutput)
}

type TriggerDestinationOutput struct{ *pulumi.OutputState }

func (TriggerDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerDestination)(nil)).Elem()
}

func (o TriggerDestinationOutput) ToTriggerDestinationOutput() TriggerDestinationOutput {
	return o
}

func (o TriggerDestinationOutput) ToTriggerDestinationOutputWithContext(ctx context.Context) TriggerDestinationOutput {
	return o
}

func (o TriggerDestinationOutput) ToTriggerDestinationPtrOutput() TriggerDestinationPtrOutput {
	return o.ToTriggerDestinationPtrOutputWithContext(context.Background())
}

func (o TriggerDestinationOutput) ToTriggerDestinationPtrOutputWithContext(ctx context.Context) TriggerDestinationPtrOutput {
	return o.ApplyT(func(v TriggerDestination) *TriggerDestination {
		return &v
	}).(TriggerDestinationPtrOutput)
}

// Cloud Run fully-managed service that receives the events. The service should be running in the same project as the trigger.
// The `matchingCriteria` block supports:
func (o TriggerDestinationOutput) CloudRunService() TriggerDestinationCloudRunServicePtrOutput {
	return o.ApplyT(func(v TriggerDestination) *TriggerDestinationCloudRunService { return v.CloudRunService }).(TriggerDestinationCloudRunServicePtrOutput)
}

type TriggerDestinationPtrOutput struct{ *pulumi.OutputState }

func (TriggerDestinationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerDestination)(nil)).Elem()
}

func (o TriggerDestinationPtrOutput) ToTriggerDestinationPtrOutput() TriggerDestinationPtrOutput {
	return o
}

func (o TriggerDestinationPtrOutput) ToTriggerDestinationPtrOutputWithContext(ctx context.Context) TriggerDestinationPtrOutput {
	return o
}

func (o TriggerDestinationPtrOutput) Elem() TriggerDestinationOutput {
	return o.ApplyT(func(v *TriggerDestination) TriggerDestination { return *v }).(TriggerDestinationOutput)
}

// Cloud Run fully-managed service that receives the events. The service should be running in the same project as the trigger.
// The `matchingCriteria` block supports:
func (o TriggerDestinationPtrOutput) CloudRunService() TriggerDestinationCloudRunServicePtrOutput {
	return o.ApplyT(func(v *TriggerDestination) *TriggerDestinationCloudRunService {
		if v == nil {
			return nil
		}
		return v.CloudRunService
	}).(TriggerDestinationCloudRunServicePtrOutput)
}

type TriggerDestinationCloudRunService struct {
	// Optional. The relative path on the Cloud Run service the events should be sent to. The value must conform to the definition of URI path segment (section 3.3 of RFC2396). Examples: "/route", "route", "route/subroute".
	Path *string `pulumi:"path"`
	// Required. The region the Cloud Run service is deployed in.
	Region *string `pulumi:"region"`
	// Required. The name of the Cloud run service being addressed (see https://cloud.google.com/run/docs/reference/rest/v1/namespaces.services). Only services located in the same project of the trigger object can be addressed.
	Service string `pulumi:"service"`
}

// TriggerDestinationCloudRunServiceInput is an input type that accepts TriggerDestinationCloudRunServiceArgs and TriggerDestinationCloudRunServiceOutput values.
// You can construct a concrete instance of `TriggerDestinationCloudRunServiceInput` via:
//
//          TriggerDestinationCloudRunServiceArgs{...}
type TriggerDestinationCloudRunServiceInput interface {
	pulumi.Input

	ToTriggerDestinationCloudRunServiceOutput() TriggerDestinationCloudRunServiceOutput
	ToTriggerDestinationCloudRunServiceOutputWithContext(context.Context) TriggerDestinationCloudRunServiceOutput
}

type TriggerDestinationCloudRunServiceArgs struct {
	// Optional. The relative path on the Cloud Run service the events should be sent to. The value must conform to the definition of URI path segment (section 3.3 of RFC2396). Examples: "/route", "route", "route/subroute".
	Path pulumi.StringPtrInput `pulumi:"path"`
	// Required. The region the Cloud Run service is deployed in.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Required. The name of the Cloud run service being addressed (see https://cloud.google.com/run/docs/reference/rest/v1/namespaces.services). Only services located in the same project of the trigger object can be addressed.
	Service pulumi.StringInput `pulumi:"service"`
}

func (TriggerDestinationCloudRunServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerDestinationCloudRunService)(nil)).Elem()
}

func (i TriggerDestinationCloudRunServiceArgs) ToTriggerDestinationCloudRunServiceOutput() TriggerDestinationCloudRunServiceOutput {
	return i.ToTriggerDestinationCloudRunServiceOutputWithContext(context.Background())
}

func (i TriggerDestinationCloudRunServiceArgs) ToTriggerDestinationCloudRunServiceOutputWithContext(ctx context.Context) TriggerDestinationCloudRunServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerDestinationCloudRunServiceOutput)
}

func (i TriggerDestinationCloudRunServiceArgs) ToTriggerDestinationCloudRunServicePtrOutput() TriggerDestinationCloudRunServicePtrOutput {
	return i.ToTriggerDestinationCloudRunServicePtrOutputWithContext(context.Background())
}

func (i TriggerDestinationCloudRunServiceArgs) ToTriggerDestinationCloudRunServicePtrOutputWithContext(ctx context.Context) TriggerDestinationCloudRunServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerDestinationCloudRunServiceOutput).ToTriggerDestinationCloudRunServicePtrOutputWithContext(ctx)
}

// TriggerDestinationCloudRunServicePtrInput is an input type that accepts TriggerDestinationCloudRunServiceArgs, TriggerDestinationCloudRunServicePtr and TriggerDestinationCloudRunServicePtrOutput values.
// You can construct a concrete instance of `TriggerDestinationCloudRunServicePtrInput` via:
//
//          TriggerDestinationCloudRunServiceArgs{...}
//
//  or:
//
//          nil
type TriggerDestinationCloudRunServicePtrInput interface {
	pulumi.Input

	ToTriggerDestinationCloudRunServicePtrOutput() TriggerDestinationCloudRunServicePtrOutput
	ToTriggerDestinationCloudRunServicePtrOutputWithContext(context.Context) TriggerDestinationCloudRunServicePtrOutput
}

type triggerDestinationCloudRunServicePtrType TriggerDestinationCloudRunServiceArgs

func TriggerDestinationCloudRunServicePtr(v *TriggerDestinationCloudRunServiceArgs) TriggerDestinationCloudRunServicePtrInput {
	return (*triggerDestinationCloudRunServicePtrType)(v)
}

func (*triggerDestinationCloudRunServicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerDestinationCloudRunService)(nil)).Elem()
}

func (i *triggerDestinationCloudRunServicePtrType) ToTriggerDestinationCloudRunServicePtrOutput() TriggerDestinationCloudRunServicePtrOutput {
	return i.ToTriggerDestinationCloudRunServicePtrOutputWithContext(context.Background())
}

func (i *triggerDestinationCloudRunServicePtrType) ToTriggerDestinationCloudRunServicePtrOutputWithContext(ctx context.Context) TriggerDestinationCloudRunServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerDestinationCloudRunServicePtrOutput)
}

type TriggerDestinationCloudRunServiceOutput struct{ *pulumi.OutputState }

func (TriggerDestinationCloudRunServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerDestinationCloudRunService)(nil)).Elem()
}

func (o TriggerDestinationCloudRunServiceOutput) ToTriggerDestinationCloudRunServiceOutput() TriggerDestinationCloudRunServiceOutput {
	return o
}

func (o TriggerDestinationCloudRunServiceOutput) ToTriggerDestinationCloudRunServiceOutputWithContext(ctx context.Context) TriggerDestinationCloudRunServiceOutput {
	return o
}

func (o TriggerDestinationCloudRunServiceOutput) ToTriggerDestinationCloudRunServicePtrOutput() TriggerDestinationCloudRunServicePtrOutput {
	return o.ToTriggerDestinationCloudRunServicePtrOutputWithContext(context.Background())
}

func (o TriggerDestinationCloudRunServiceOutput) ToTriggerDestinationCloudRunServicePtrOutputWithContext(ctx context.Context) TriggerDestinationCloudRunServicePtrOutput {
	return o.ApplyT(func(v TriggerDestinationCloudRunService) *TriggerDestinationCloudRunService {
		return &v
	}).(TriggerDestinationCloudRunServicePtrOutput)
}

// Optional. The relative path on the Cloud Run service the events should be sent to. The value must conform to the definition of URI path segment (section 3.3 of RFC2396). Examples: "/route", "route", "route/subroute".
func (o TriggerDestinationCloudRunServiceOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TriggerDestinationCloudRunService) *string { return v.Path }).(pulumi.StringPtrOutput)
}

// Required. The region the Cloud Run service is deployed in.
func (o TriggerDestinationCloudRunServiceOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TriggerDestinationCloudRunService) *string { return v.Region }).(pulumi.StringPtrOutput)
}

// Required. The name of the Cloud run service being addressed (see https://cloud.google.com/run/docs/reference/rest/v1/namespaces.services). Only services located in the same project of the trigger object can be addressed.
func (o TriggerDestinationCloudRunServiceOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v TriggerDestinationCloudRunService) string { return v.Service }).(pulumi.StringOutput)
}

type TriggerDestinationCloudRunServicePtrOutput struct{ *pulumi.OutputState }

func (TriggerDestinationCloudRunServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerDestinationCloudRunService)(nil)).Elem()
}

func (o TriggerDestinationCloudRunServicePtrOutput) ToTriggerDestinationCloudRunServicePtrOutput() TriggerDestinationCloudRunServicePtrOutput {
	return o
}

func (o TriggerDestinationCloudRunServicePtrOutput) ToTriggerDestinationCloudRunServicePtrOutputWithContext(ctx context.Context) TriggerDestinationCloudRunServicePtrOutput {
	return o
}

func (o TriggerDestinationCloudRunServicePtrOutput) Elem() TriggerDestinationCloudRunServiceOutput {
	return o.ApplyT(func(v *TriggerDestinationCloudRunService) TriggerDestinationCloudRunService { return *v }).(TriggerDestinationCloudRunServiceOutput)
}

// Optional. The relative path on the Cloud Run service the events should be sent to. The value must conform to the definition of URI path segment (section 3.3 of RFC2396). Examples: "/route", "route", "route/subroute".
func (o TriggerDestinationCloudRunServicePtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TriggerDestinationCloudRunService) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

// Required. The region the Cloud Run service is deployed in.
func (o TriggerDestinationCloudRunServicePtrOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TriggerDestinationCloudRunService) *string {
		if v == nil {
			return nil
		}
		return v.Region
	}).(pulumi.StringPtrOutput)
}

// Required. The name of the Cloud run service being addressed (see https://cloud.google.com/run/docs/reference/rest/v1/namespaces.services). Only services located in the same project of the trigger object can be addressed.
func (o TriggerDestinationCloudRunServicePtrOutput) Service() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TriggerDestinationCloudRunService) *string {
		if v == nil {
			return nil
		}
		return &v.Service
	}).(pulumi.StringPtrOutput)
}

type TriggerMatchingCriteria struct {
	// Required. The name of a CloudEvents attribute. Currently, only a subset of attributes can be specified. All triggers MUST provide a matching criteria for the 'type' attribute.
	Attribute string `pulumi:"attribute"`
	// Required. The value for the attribute.
	Value string `pulumi:"value"`
}

// TriggerMatchingCriteriaInput is an input type that accepts TriggerMatchingCriteriaArgs and TriggerMatchingCriteriaOutput values.
// You can construct a concrete instance of `TriggerMatchingCriteriaInput` via:
//
//          TriggerMatchingCriteriaArgs{...}
type TriggerMatchingCriteriaInput interface {
	pulumi.Input

	ToTriggerMatchingCriteriaOutput() TriggerMatchingCriteriaOutput
	ToTriggerMatchingCriteriaOutputWithContext(context.Context) TriggerMatchingCriteriaOutput
}

type TriggerMatchingCriteriaArgs struct {
	// Required. The name of a CloudEvents attribute. Currently, only a subset of attributes can be specified. All triggers MUST provide a matching criteria for the 'type' attribute.
	Attribute pulumi.StringInput `pulumi:"attribute"`
	// Required. The value for the attribute.
	Value pulumi.StringInput `pulumi:"value"`
}

func (TriggerMatchingCriteriaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerMatchingCriteria)(nil)).Elem()
}

func (i TriggerMatchingCriteriaArgs) ToTriggerMatchingCriteriaOutput() TriggerMatchingCriteriaOutput {
	return i.ToTriggerMatchingCriteriaOutputWithContext(context.Background())
}

func (i TriggerMatchingCriteriaArgs) ToTriggerMatchingCriteriaOutputWithContext(ctx context.Context) TriggerMatchingCriteriaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerMatchingCriteriaOutput)
}

// TriggerMatchingCriteriaArrayInput is an input type that accepts TriggerMatchingCriteriaArray and TriggerMatchingCriteriaArrayOutput values.
// You can construct a concrete instance of `TriggerMatchingCriteriaArrayInput` via:
//
//          TriggerMatchingCriteriaArray{ TriggerMatchingCriteriaArgs{...} }
type TriggerMatchingCriteriaArrayInput interface {
	pulumi.Input

	ToTriggerMatchingCriteriaArrayOutput() TriggerMatchingCriteriaArrayOutput
	ToTriggerMatchingCriteriaArrayOutputWithContext(context.Context) TriggerMatchingCriteriaArrayOutput
}

type TriggerMatchingCriteriaArray []TriggerMatchingCriteriaInput

func (TriggerMatchingCriteriaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TriggerMatchingCriteria)(nil)).Elem()
}

func (i TriggerMatchingCriteriaArray) ToTriggerMatchingCriteriaArrayOutput() TriggerMatchingCriteriaArrayOutput {
	return i.ToTriggerMatchingCriteriaArrayOutputWithContext(context.Background())
}

func (i TriggerMatchingCriteriaArray) ToTriggerMatchingCriteriaArrayOutputWithContext(ctx context.Context) TriggerMatchingCriteriaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerMatchingCriteriaArrayOutput)
}

type TriggerMatchingCriteriaOutput struct{ *pulumi.OutputState }

func (TriggerMatchingCriteriaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerMatchingCriteria)(nil)).Elem()
}

func (o TriggerMatchingCriteriaOutput) ToTriggerMatchingCriteriaOutput() TriggerMatchingCriteriaOutput {
	return o
}

func (o TriggerMatchingCriteriaOutput) ToTriggerMatchingCriteriaOutputWithContext(ctx context.Context) TriggerMatchingCriteriaOutput {
	return o
}

// Required. The name of a CloudEvents attribute. Currently, only a subset of attributes can be specified. All triggers MUST provide a matching criteria for the 'type' attribute.
func (o TriggerMatchingCriteriaOutput) Attribute() pulumi.StringOutput {
	return o.ApplyT(func(v TriggerMatchingCriteria) string { return v.Attribute }).(pulumi.StringOutput)
}

// Required. The value for the attribute.
func (o TriggerMatchingCriteriaOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v TriggerMatchingCriteria) string { return v.Value }).(pulumi.StringOutput)
}

type TriggerMatchingCriteriaArrayOutput struct{ *pulumi.OutputState }

func (TriggerMatchingCriteriaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TriggerMatchingCriteria)(nil)).Elem()
}

func (o TriggerMatchingCriteriaArrayOutput) ToTriggerMatchingCriteriaArrayOutput() TriggerMatchingCriteriaArrayOutput {
	return o
}

func (o TriggerMatchingCriteriaArrayOutput) ToTriggerMatchingCriteriaArrayOutputWithContext(ctx context.Context) TriggerMatchingCriteriaArrayOutput {
	return o
}

func (o TriggerMatchingCriteriaArrayOutput) Index(i pulumi.IntInput) TriggerMatchingCriteriaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TriggerMatchingCriteria {
		return vs[0].([]TriggerMatchingCriteria)[vs[1].(int)]
	}).(TriggerMatchingCriteriaOutput)
}

type TriggerTransport struct {
	Pubsubs []TriggerTransportPubsub `pulumi:"pubsubs"`
}

// TriggerTransportInput is an input type that accepts TriggerTransportArgs and TriggerTransportOutput values.
// You can construct a concrete instance of `TriggerTransportInput` via:
//
//          TriggerTransportArgs{...}
type TriggerTransportInput interface {
	pulumi.Input

	ToTriggerTransportOutput() TriggerTransportOutput
	ToTriggerTransportOutputWithContext(context.Context) TriggerTransportOutput
}

type TriggerTransportArgs struct {
	Pubsubs TriggerTransportPubsubArrayInput `pulumi:"pubsubs"`
}

func (TriggerTransportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerTransport)(nil)).Elem()
}

func (i TriggerTransportArgs) ToTriggerTransportOutput() TriggerTransportOutput {
	return i.ToTriggerTransportOutputWithContext(context.Background())
}

func (i TriggerTransportArgs) ToTriggerTransportOutputWithContext(ctx context.Context) TriggerTransportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerTransportOutput)
}

// TriggerTransportArrayInput is an input type that accepts TriggerTransportArray and TriggerTransportArrayOutput values.
// You can construct a concrete instance of `TriggerTransportArrayInput` via:
//
//          TriggerTransportArray{ TriggerTransportArgs{...} }
type TriggerTransportArrayInput interface {
	pulumi.Input

	ToTriggerTransportArrayOutput() TriggerTransportArrayOutput
	ToTriggerTransportArrayOutputWithContext(context.Context) TriggerTransportArrayOutput
}

type TriggerTransportArray []TriggerTransportInput

func (TriggerTransportArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TriggerTransport)(nil)).Elem()
}

func (i TriggerTransportArray) ToTriggerTransportArrayOutput() TriggerTransportArrayOutput {
	return i.ToTriggerTransportArrayOutputWithContext(context.Background())
}

func (i TriggerTransportArray) ToTriggerTransportArrayOutputWithContext(ctx context.Context) TriggerTransportArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerTransportArrayOutput)
}

type TriggerTransportOutput struct{ *pulumi.OutputState }

func (TriggerTransportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerTransport)(nil)).Elem()
}

func (o TriggerTransportOutput) ToTriggerTransportOutput() TriggerTransportOutput {
	return o
}

func (o TriggerTransportOutput) ToTriggerTransportOutputWithContext(ctx context.Context) TriggerTransportOutput {
	return o
}

func (o TriggerTransportOutput) Pubsubs() TriggerTransportPubsubArrayOutput {
	return o.ApplyT(func(v TriggerTransport) []TriggerTransportPubsub { return v.Pubsubs }).(TriggerTransportPubsubArrayOutput)
}

type TriggerTransportArrayOutput struct{ *pulumi.OutputState }

func (TriggerTransportArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TriggerTransport)(nil)).Elem()
}

func (o TriggerTransportArrayOutput) ToTriggerTransportArrayOutput() TriggerTransportArrayOutput {
	return o
}

func (o TriggerTransportArrayOutput) ToTriggerTransportArrayOutputWithContext(ctx context.Context) TriggerTransportArrayOutput {
	return o
}

func (o TriggerTransportArrayOutput) Index(i pulumi.IntInput) TriggerTransportOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TriggerTransport {
		return vs[0].([]TriggerTransport)[vs[1].(int)]
	}).(TriggerTransportOutput)
}

type TriggerTransportPubsub struct {
	Subscription *string `pulumi:"subscription"`
	Topic        *string `pulumi:"topic"`
}

// TriggerTransportPubsubInput is an input type that accepts TriggerTransportPubsubArgs and TriggerTransportPubsubOutput values.
// You can construct a concrete instance of `TriggerTransportPubsubInput` via:
//
//          TriggerTransportPubsubArgs{...}
type TriggerTransportPubsubInput interface {
	pulumi.Input

	ToTriggerTransportPubsubOutput() TriggerTransportPubsubOutput
	ToTriggerTransportPubsubOutputWithContext(context.Context) TriggerTransportPubsubOutput
}

type TriggerTransportPubsubArgs struct {
	Subscription pulumi.StringPtrInput `pulumi:"subscription"`
	Topic        pulumi.StringPtrInput `pulumi:"topic"`
}

func (TriggerTransportPubsubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerTransportPubsub)(nil)).Elem()
}

func (i TriggerTransportPubsubArgs) ToTriggerTransportPubsubOutput() TriggerTransportPubsubOutput {
	return i.ToTriggerTransportPubsubOutputWithContext(context.Background())
}

func (i TriggerTransportPubsubArgs) ToTriggerTransportPubsubOutputWithContext(ctx context.Context) TriggerTransportPubsubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerTransportPubsubOutput)
}

// TriggerTransportPubsubArrayInput is an input type that accepts TriggerTransportPubsubArray and TriggerTransportPubsubArrayOutput values.
// You can construct a concrete instance of `TriggerTransportPubsubArrayInput` via:
//
//          TriggerTransportPubsubArray{ TriggerTransportPubsubArgs{...} }
type TriggerTransportPubsubArrayInput interface {
	pulumi.Input

	ToTriggerTransportPubsubArrayOutput() TriggerTransportPubsubArrayOutput
	ToTriggerTransportPubsubArrayOutputWithContext(context.Context) TriggerTransportPubsubArrayOutput
}

type TriggerTransportPubsubArray []TriggerTransportPubsubInput

func (TriggerTransportPubsubArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TriggerTransportPubsub)(nil)).Elem()
}

func (i TriggerTransportPubsubArray) ToTriggerTransportPubsubArrayOutput() TriggerTransportPubsubArrayOutput {
	return i.ToTriggerTransportPubsubArrayOutputWithContext(context.Background())
}

func (i TriggerTransportPubsubArray) ToTriggerTransportPubsubArrayOutputWithContext(ctx context.Context) TriggerTransportPubsubArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerTransportPubsubArrayOutput)
}

type TriggerTransportPubsubOutput struct{ *pulumi.OutputState }

func (TriggerTransportPubsubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerTransportPubsub)(nil)).Elem()
}

func (o TriggerTransportPubsubOutput) ToTriggerTransportPubsubOutput() TriggerTransportPubsubOutput {
	return o
}

func (o TriggerTransportPubsubOutput) ToTriggerTransportPubsubOutputWithContext(ctx context.Context) TriggerTransportPubsubOutput {
	return o
}

func (o TriggerTransportPubsubOutput) Subscription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TriggerTransportPubsub) *string { return v.Subscription }).(pulumi.StringPtrOutput)
}

func (o TriggerTransportPubsubOutput) Topic() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TriggerTransportPubsub) *string { return v.Topic }).(pulumi.StringPtrOutput)
}

type TriggerTransportPubsubArrayOutput struct{ *pulumi.OutputState }

func (TriggerTransportPubsubArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TriggerTransportPubsub)(nil)).Elem()
}

func (o TriggerTransportPubsubArrayOutput) ToTriggerTransportPubsubArrayOutput() TriggerTransportPubsubArrayOutput {
	return o
}

func (o TriggerTransportPubsubArrayOutput) ToTriggerTransportPubsubArrayOutputWithContext(ctx context.Context) TriggerTransportPubsubArrayOutput {
	return o
}

func (o TriggerTransportPubsubArrayOutput) Index(i pulumi.IntInput) TriggerTransportPubsubOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TriggerTransportPubsub {
		return vs[0].([]TriggerTransportPubsub)[vs[1].(int)]
	}).(TriggerTransportPubsubOutput)
}

func init() {
	pulumi.RegisterOutputType(TriggerDestinationOutput{})
	pulumi.RegisterOutputType(TriggerDestinationPtrOutput{})
	pulumi.RegisterOutputType(TriggerDestinationCloudRunServiceOutput{})
	pulumi.RegisterOutputType(TriggerDestinationCloudRunServicePtrOutput{})
	pulumi.RegisterOutputType(TriggerMatchingCriteriaOutput{})
	pulumi.RegisterOutputType(TriggerMatchingCriteriaArrayOutput{})
	pulumi.RegisterOutputType(TriggerTransportOutput{})
	pulumi.RegisterOutputType(TriggerTransportArrayOutput{})
	pulumi.RegisterOutputType(TriggerTransportPubsubOutput{})
	pulumi.RegisterOutputType(TriggerTransportPubsubArrayOutput{})
}