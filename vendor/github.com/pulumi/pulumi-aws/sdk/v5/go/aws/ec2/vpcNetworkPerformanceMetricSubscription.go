// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage an Infrastructure Performance subscription.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.NewVpcNetworkPerformanceMetricSubscription(ctx, "example", &ec2.VpcNetworkPerformanceMetricSubscriptionArgs{
//				Destination: pulumi.String("us-west-1"),
//				Source:      pulumi.String("us-east-1"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type VpcNetworkPerformanceMetricSubscription struct {
	pulumi.CustomResourceState

	// The target Region or Availability Zone that the metric subscription is enabled for. For example, `eu-west-1`.
	Destination pulumi.StringOutput `pulumi:"destination"`
	// The metric used for the enabled subscription. Valid values: `aggregate-latency`. Default: `aggregate-latency`.
	Metric pulumi.StringPtrOutput `pulumi:"metric"`
	// The data aggregation time for the subscription.
	Period pulumi.StringOutput `pulumi:"period"`
	// The source Region or Availability Zone that the metric subscription is enabled for. For example, `us-east-1`.
	Source pulumi.StringOutput `pulumi:"source"`
	// The statistic used for the enabled subscription. Valid values: `p50`. Default: `p50`.
	Statistic pulumi.StringPtrOutput `pulumi:"statistic"`
}

// NewVpcNetworkPerformanceMetricSubscription registers a new resource with the given unique name, arguments, and options.
func NewVpcNetworkPerformanceMetricSubscription(ctx *pulumi.Context,
	name string, args *VpcNetworkPerformanceMetricSubscriptionArgs, opts ...pulumi.ResourceOption) (*VpcNetworkPerformanceMetricSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Destination == nil {
		return nil, errors.New("invalid value for required argument 'Destination'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	var resource VpcNetworkPerformanceMetricSubscription
	err := ctx.RegisterResource("aws:ec2/vpcNetworkPerformanceMetricSubscription:VpcNetworkPerformanceMetricSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcNetworkPerformanceMetricSubscription gets an existing VpcNetworkPerformanceMetricSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcNetworkPerformanceMetricSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcNetworkPerformanceMetricSubscriptionState, opts ...pulumi.ResourceOption) (*VpcNetworkPerformanceMetricSubscription, error) {
	var resource VpcNetworkPerformanceMetricSubscription
	err := ctx.ReadResource("aws:ec2/vpcNetworkPerformanceMetricSubscription:VpcNetworkPerformanceMetricSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcNetworkPerformanceMetricSubscription resources.
type vpcNetworkPerformanceMetricSubscriptionState struct {
	// The target Region or Availability Zone that the metric subscription is enabled for. For example, `eu-west-1`.
	Destination *string `pulumi:"destination"`
	// The metric used for the enabled subscription. Valid values: `aggregate-latency`. Default: `aggregate-latency`.
	Metric *string `pulumi:"metric"`
	// The data aggregation time for the subscription.
	Period *string `pulumi:"period"`
	// The source Region or Availability Zone that the metric subscription is enabled for. For example, `us-east-1`.
	Source *string `pulumi:"source"`
	// The statistic used for the enabled subscription. Valid values: `p50`. Default: `p50`.
	Statistic *string `pulumi:"statistic"`
}

type VpcNetworkPerformanceMetricSubscriptionState struct {
	// The target Region or Availability Zone that the metric subscription is enabled for. For example, `eu-west-1`.
	Destination pulumi.StringPtrInput
	// The metric used for the enabled subscription. Valid values: `aggregate-latency`. Default: `aggregate-latency`.
	Metric pulumi.StringPtrInput
	// The data aggregation time for the subscription.
	Period pulumi.StringPtrInput
	// The source Region or Availability Zone that the metric subscription is enabled for. For example, `us-east-1`.
	Source pulumi.StringPtrInput
	// The statistic used for the enabled subscription. Valid values: `p50`. Default: `p50`.
	Statistic pulumi.StringPtrInput
}

func (VpcNetworkPerformanceMetricSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcNetworkPerformanceMetricSubscriptionState)(nil)).Elem()
}

type vpcNetworkPerformanceMetricSubscriptionArgs struct {
	// The target Region or Availability Zone that the metric subscription is enabled for. For example, `eu-west-1`.
	Destination string `pulumi:"destination"`
	// The metric used for the enabled subscription. Valid values: `aggregate-latency`. Default: `aggregate-latency`.
	Metric *string `pulumi:"metric"`
	// The source Region or Availability Zone that the metric subscription is enabled for. For example, `us-east-1`.
	Source string `pulumi:"source"`
	// The statistic used for the enabled subscription. Valid values: `p50`. Default: `p50`.
	Statistic *string `pulumi:"statistic"`
}

// The set of arguments for constructing a VpcNetworkPerformanceMetricSubscription resource.
type VpcNetworkPerformanceMetricSubscriptionArgs struct {
	// The target Region or Availability Zone that the metric subscription is enabled for. For example, `eu-west-1`.
	Destination pulumi.StringInput
	// The metric used for the enabled subscription. Valid values: `aggregate-latency`. Default: `aggregate-latency`.
	Metric pulumi.StringPtrInput
	// The source Region or Availability Zone that the metric subscription is enabled for. For example, `us-east-1`.
	Source pulumi.StringInput
	// The statistic used for the enabled subscription. Valid values: `p50`. Default: `p50`.
	Statistic pulumi.StringPtrInput
}

func (VpcNetworkPerformanceMetricSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcNetworkPerformanceMetricSubscriptionArgs)(nil)).Elem()
}

type VpcNetworkPerformanceMetricSubscriptionInput interface {
	pulumi.Input

	ToVpcNetworkPerformanceMetricSubscriptionOutput() VpcNetworkPerformanceMetricSubscriptionOutput
	ToVpcNetworkPerformanceMetricSubscriptionOutputWithContext(ctx context.Context) VpcNetworkPerformanceMetricSubscriptionOutput
}

func (*VpcNetworkPerformanceMetricSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcNetworkPerformanceMetricSubscription)(nil)).Elem()
}

func (i *VpcNetworkPerformanceMetricSubscription) ToVpcNetworkPerformanceMetricSubscriptionOutput() VpcNetworkPerformanceMetricSubscriptionOutput {
	return i.ToVpcNetworkPerformanceMetricSubscriptionOutputWithContext(context.Background())
}

func (i *VpcNetworkPerformanceMetricSubscription) ToVpcNetworkPerformanceMetricSubscriptionOutputWithContext(ctx context.Context) VpcNetworkPerformanceMetricSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcNetworkPerformanceMetricSubscriptionOutput)
}

// VpcNetworkPerformanceMetricSubscriptionArrayInput is an input type that accepts VpcNetworkPerformanceMetricSubscriptionArray and VpcNetworkPerformanceMetricSubscriptionArrayOutput values.
// You can construct a concrete instance of `VpcNetworkPerformanceMetricSubscriptionArrayInput` via:
//
//	VpcNetworkPerformanceMetricSubscriptionArray{ VpcNetworkPerformanceMetricSubscriptionArgs{...} }
type VpcNetworkPerformanceMetricSubscriptionArrayInput interface {
	pulumi.Input

	ToVpcNetworkPerformanceMetricSubscriptionArrayOutput() VpcNetworkPerformanceMetricSubscriptionArrayOutput
	ToVpcNetworkPerformanceMetricSubscriptionArrayOutputWithContext(context.Context) VpcNetworkPerformanceMetricSubscriptionArrayOutput
}

type VpcNetworkPerformanceMetricSubscriptionArray []VpcNetworkPerformanceMetricSubscriptionInput

func (VpcNetworkPerformanceMetricSubscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcNetworkPerformanceMetricSubscription)(nil)).Elem()
}

func (i VpcNetworkPerformanceMetricSubscriptionArray) ToVpcNetworkPerformanceMetricSubscriptionArrayOutput() VpcNetworkPerformanceMetricSubscriptionArrayOutput {
	return i.ToVpcNetworkPerformanceMetricSubscriptionArrayOutputWithContext(context.Background())
}

func (i VpcNetworkPerformanceMetricSubscriptionArray) ToVpcNetworkPerformanceMetricSubscriptionArrayOutputWithContext(ctx context.Context) VpcNetworkPerformanceMetricSubscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcNetworkPerformanceMetricSubscriptionArrayOutput)
}

// VpcNetworkPerformanceMetricSubscriptionMapInput is an input type that accepts VpcNetworkPerformanceMetricSubscriptionMap and VpcNetworkPerformanceMetricSubscriptionMapOutput values.
// You can construct a concrete instance of `VpcNetworkPerformanceMetricSubscriptionMapInput` via:
//
//	VpcNetworkPerformanceMetricSubscriptionMap{ "key": VpcNetworkPerformanceMetricSubscriptionArgs{...} }
type VpcNetworkPerformanceMetricSubscriptionMapInput interface {
	pulumi.Input

	ToVpcNetworkPerformanceMetricSubscriptionMapOutput() VpcNetworkPerformanceMetricSubscriptionMapOutput
	ToVpcNetworkPerformanceMetricSubscriptionMapOutputWithContext(context.Context) VpcNetworkPerformanceMetricSubscriptionMapOutput
}

type VpcNetworkPerformanceMetricSubscriptionMap map[string]VpcNetworkPerformanceMetricSubscriptionInput

func (VpcNetworkPerformanceMetricSubscriptionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcNetworkPerformanceMetricSubscription)(nil)).Elem()
}

func (i VpcNetworkPerformanceMetricSubscriptionMap) ToVpcNetworkPerformanceMetricSubscriptionMapOutput() VpcNetworkPerformanceMetricSubscriptionMapOutput {
	return i.ToVpcNetworkPerformanceMetricSubscriptionMapOutputWithContext(context.Background())
}

func (i VpcNetworkPerformanceMetricSubscriptionMap) ToVpcNetworkPerformanceMetricSubscriptionMapOutputWithContext(ctx context.Context) VpcNetworkPerformanceMetricSubscriptionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcNetworkPerformanceMetricSubscriptionMapOutput)
}

type VpcNetworkPerformanceMetricSubscriptionOutput struct{ *pulumi.OutputState }

func (VpcNetworkPerformanceMetricSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcNetworkPerformanceMetricSubscription)(nil)).Elem()
}

func (o VpcNetworkPerformanceMetricSubscriptionOutput) ToVpcNetworkPerformanceMetricSubscriptionOutput() VpcNetworkPerformanceMetricSubscriptionOutput {
	return o
}

func (o VpcNetworkPerformanceMetricSubscriptionOutput) ToVpcNetworkPerformanceMetricSubscriptionOutputWithContext(ctx context.Context) VpcNetworkPerformanceMetricSubscriptionOutput {
	return o
}

// The target Region or Availability Zone that the metric subscription is enabled for. For example, `eu-west-1`.
func (o VpcNetworkPerformanceMetricSubscriptionOutput) Destination() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcNetworkPerformanceMetricSubscription) pulumi.StringOutput { return v.Destination }).(pulumi.StringOutput)
}

// The metric used for the enabled subscription. Valid values: `aggregate-latency`. Default: `aggregate-latency`.
func (o VpcNetworkPerformanceMetricSubscriptionOutput) Metric() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcNetworkPerformanceMetricSubscription) pulumi.StringPtrOutput { return v.Metric }).(pulumi.StringPtrOutput)
}

// The data aggregation time for the subscription.
func (o VpcNetworkPerformanceMetricSubscriptionOutput) Period() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcNetworkPerformanceMetricSubscription) pulumi.StringOutput { return v.Period }).(pulumi.StringOutput)
}

// The source Region or Availability Zone that the metric subscription is enabled for. For example, `us-east-1`.
func (o VpcNetworkPerformanceMetricSubscriptionOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcNetworkPerformanceMetricSubscription) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

// The statistic used for the enabled subscription. Valid values: `p50`. Default: `p50`.
func (o VpcNetworkPerformanceMetricSubscriptionOutput) Statistic() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcNetworkPerformanceMetricSubscription) pulumi.StringPtrOutput { return v.Statistic }).(pulumi.StringPtrOutput)
}

type VpcNetworkPerformanceMetricSubscriptionArrayOutput struct{ *pulumi.OutputState }

func (VpcNetworkPerformanceMetricSubscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcNetworkPerformanceMetricSubscription)(nil)).Elem()
}

func (o VpcNetworkPerformanceMetricSubscriptionArrayOutput) ToVpcNetworkPerformanceMetricSubscriptionArrayOutput() VpcNetworkPerformanceMetricSubscriptionArrayOutput {
	return o
}

func (o VpcNetworkPerformanceMetricSubscriptionArrayOutput) ToVpcNetworkPerformanceMetricSubscriptionArrayOutputWithContext(ctx context.Context) VpcNetworkPerformanceMetricSubscriptionArrayOutput {
	return o
}

func (o VpcNetworkPerformanceMetricSubscriptionArrayOutput) Index(i pulumi.IntInput) VpcNetworkPerformanceMetricSubscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcNetworkPerformanceMetricSubscription {
		return vs[0].([]*VpcNetworkPerformanceMetricSubscription)[vs[1].(int)]
	}).(VpcNetworkPerformanceMetricSubscriptionOutput)
}

type VpcNetworkPerformanceMetricSubscriptionMapOutput struct{ *pulumi.OutputState }

func (VpcNetworkPerformanceMetricSubscriptionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcNetworkPerformanceMetricSubscription)(nil)).Elem()
}

func (o VpcNetworkPerformanceMetricSubscriptionMapOutput) ToVpcNetworkPerformanceMetricSubscriptionMapOutput() VpcNetworkPerformanceMetricSubscriptionMapOutput {
	return o
}

func (o VpcNetworkPerformanceMetricSubscriptionMapOutput) ToVpcNetworkPerformanceMetricSubscriptionMapOutputWithContext(ctx context.Context) VpcNetworkPerformanceMetricSubscriptionMapOutput {
	return o
}

func (o VpcNetworkPerformanceMetricSubscriptionMapOutput) MapIndex(k pulumi.StringInput) VpcNetworkPerformanceMetricSubscriptionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcNetworkPerformanceMetricSubscription {
		return vs[0].(map[string]*VpcNetworkPerformanceMetricSubscription)[vs[1].(string)]
	}).(VpcNetworkPerformanceMetricSubscriptionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcNetworkPerformanceMetricSubscriptionInput)(nil)).Elem(), &VpcNetworkPerformanceMetricSubscription{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcNetworkPerformanceMetricSubscriptionArrayInput)(nil)).Elem(), VpcNetworkPerformanceMetricSubscriptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcNetworkPerformanceMetricSubscriptionMapInput)(nil)).Elem(), VpcNetworkPerformanceMetricSubscriptionMap{})
	pulumi.RegisterOutputType(VpcNetworkPerformanceMetricSubscriptionOutput{})
	pulumi.RegisterOutputType(VpcNetworkPerformanceMetricSubscriptionArrayOutput{})
	pulumi.RegisterOutputType(VpcNetworkPerformanceMetricSubscriptionMapOutput{})
}