// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpcaccess

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ConnectorSubnet struct {
	Name      *string `pulumi:"name"`
	ProjectId *string `pulumi:"projectId"`
}

// ConnectorSubnetInput is an input type that accepts ConnectorSubnetArgs and ConnectorSubnetOutput values.
// You can construct a concrete instance of `ConnectorSubnetInput` via:
//
//          ConnectorSubnetArgs{...}
type ConnectorSubnetInput interface {
	pulumi.Input

	ToConnectorSubnetOutput() ConnectorSubnetOutput
	ToConnectorSubnetOutputWithContext(context.Context) ConnectorSubnetOutput
}

type ConnectorSubnetArgs struct {
	Name      pulumi.StringPtrInput `pulumi:"name"`
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
}

func (ConnectorSubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorSubnet)(nil)).Elem()
}

func (i ConnectorSubnetArgs) ToConnectorSubnetOutput() ConnectorSubnetOutput {
	return i.ToConnectorSubnetOutputWithContext(context.Background())
}

func (i ConnectorSubnetArgs) ToConnectorSubnetOutputWithContext(ctx context.Context) ConnectorSubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorSubnetOutput)
}

func (i ConnectorSubnetArgs) ToConnectorSubnetPtrOutput() ConnectorSubnetPtrOutput {
	return i.ToConnectorSubnetPtrOutputWithContext(context.Background())
}

func (i ConnectorSubnetArgs) ToConnectorSubnetPtrOutputWithContext(ctx context.Context) ConnectorSubnetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorSubnetOutput).ToConnectorSubnetPtrOutputWithContext(ctx)
}

// ConnectorSubnetPtrInput is an input type that accepts ConnectorSubnetArgs, ConnectorSubnetPtr and ConnectorSubnetPtrOutput values.
// You can construct a concrete instance of `ConnectorSubnetPtrInput` via:
//
//          ConnectorSubnetArgs{...}
//
//  or:
//
//          nil
type ConnectorSubnetPtrInput interface {
	pulumi.Input

	ToConnectorSubnetPtrOutput() ConnectorSubnetPtrOutput
	ToConnectorSubnetPtrOutputWithContext(context.Context) ConnectorSubnetPtrOutput
}

type connectorSubnetPtrType ConnectorSubnetArgs

func ConnectorSubnetPtr(v *ConnectorSubnetArgs) ConnectorSubnetPtrInput {
	return (*connectorSubnetPtrType)(v)
}

func (*connectorSubnetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorSubnet)(nil)).Elem()
}

func (i *connectorSubnetPtrType) ToConnectorSubnetPtrOutput() ConnectorSubnetPtrOutput {
	return i.ToConnectorSubnetPtrOutputWithContext(context.Background())
}

func (i *connectorSubnetPtrType) ToConnectorSubnetPtrOutputWithContext(ctx context.Context) ConnectorSubnetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorSubnetPtrOutput)
}

type ConnectorSubnetOutput struct{ *pulumi.OutputState }

func (ConnectorSubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorSubnet)(nil)).Elem()
}

func (o ConnectorSubnetOutput) ToConnectorSubnetOutput() ConnectorSubnetOutput {
	return o
}

func (o ConnectorSubnetOutput) ToConnectorSubnetOutputWithContext(ctx context.Context) ConnectorSubnetOutput {
	return o
}

func (o ConnectorSubnetOutput) ToConnectorSubnetPtrOutput() ConnectorSubnetPtrOutput {
	return o.ToConnectorSubnetPtrOutputWithContext(context.Background())
}

func (o ConnectorSubnetOutput) ToConnectorSubnetPtrOutputWithContext(ctx context.Context) ConnectorSubnetPtrOutput {
	return o.ApplyT(func(v ConnectorSubnet) *ConnectorSubnet {
		return &v
	}).(ConnectorSubnetPtrOutput)
}
func (o ConnectorSubnetOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectorSubnet) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ConnectorSubnetOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectorSubnet) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

type ConnectorSubnetPtrOutput struct{ *pulumi.OutputState }

func (ConnectorSubnetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorSubnet)(nil)).Elem()
}

func (o ConnectorSubnetPtrOutput) ToConnectorSubnetPtrOutput() ConnectorSubnetPtrOutput {
	return o
}

func (o ConnectorSubnetPtrOutput) ToConnectorSubnetPtrOutputWithContext(ctx context.Context) ConnectorSubnetPtrOutput {
	return o
}

func (o ConnectorSubnetPtrOutput) Elem() ConnectorSubnetOutput {
	return o.ApplyT(func(v *ConnectorSubnet) ConnectorSubnet { return *v }).(ConnectorSubnetOutput)
}

func (o ConnectorSubnetPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorSubnet) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ConnectorSubnetPtrOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorSubnet) *string {
		if v == nil {
			return nil
		}
		return v.ProjectId
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectorSubnetOutput{})
	pulumi.RegisterOutputType(ConnectorSubnetPtrOutput{})
}