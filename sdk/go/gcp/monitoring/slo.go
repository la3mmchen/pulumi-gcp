// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Service-Level Objective (SLO) describes the level of desired good
// service. It consists of a service-level indicator (SLI), a performance
// goal, and a period over which the objective is to be evaluated against
// that goal. The SLO can use SLIs defined in a number of different manners.
// Typical SLOs might include "99% of requests in each rolling week have
// latency below 200 milliseconds" or "99.5% of requests in each calendar
// month return successfully."
//
//
// To get more information about Slo, see:
//
// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/services.serviceLevelObjectives)
// * How-to Guides
//     * [Service Monitoring](https://cloud.google.com/monitoring/service-monitoring)
//     * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)
type Slo struct {
	pulumi.CustomResourceState

	// Basic Service-Level Indicator (SLI) on a well-known service type. Performance will be computed on the basis of
	// pre-defined metrics. SLIs are used to measure and calculate the quality of the Service's performance with respect to a
	// single aspect of service quality.
	BasicSli SloBasicSliOutput `pulumi:"basicSli"`
	// A calendar period, semantically "since the start of the current <calendarPeriod>".
	CalendarPeriod pulumi.StringPtrOutput `pulumi:"calendarPeriod"`
	// Name used for UI elements listing this SLO.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The fraction of service that must be good in order for this objective to be met. 0 < goal <= 0.999
	Goal pulumi.Float64Output `pulumi:"goal"`
	// The full resource name for this service. The syntax is:
	// projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// A rolling time period, semantically "in the past X days". Must be between 1 to 30 days, inclusive.
	RollingPeriodDays pulumi.IntPtrOutput `pulumi:"rollingPeriodDays"`
	// ID of the service to which this SLO belongs.
	Service pulumi.StringOutput `pulumi:"service"`
	// The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
	SloId pulumi.StringOutput `pulumi:"sloId"`
}

// NewSlo registers a new resource with the given unique name, arguments, and options.
func NewSlo(ctx *pulumi.Context,
	name string, args *SloArgs, opts ...pulumi.ResourceOption) (*Slo, error) {
	if args == nil || args.BasicSli == nil {
		return nil, errors.New("missing required argument 'BasicSli'")
	}
	if args == nil || args.Goal == nil {
		return nil, errors.New("missing required argument 'Goal'")
	}
	if args == nil || args.Service == nil {
		return nil, errors.New("missing required argument 'Service'")
	}
	if args == nil {
		args = &SloArgs{}
	}
	var resource Slo
	err := ctx.RegisterResource("gcp:monitoring/slo:Slo", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSlo gets an existing Slo resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSlo(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SloState, opts ...pulumi.ResourceOption) (*Slo, error) {
	var resource Slo
	err := ctx.ReadResource("gcp:monitoring/slo:Slo", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Slo resources.
type sloState struct {
	// Basic Service-Level Indicator (SLI) on a well-known service type. Performance will be computed on the basis of
	// pre-defined metrics. SLIs are used to measure and calculate the quality of the Service's performance with respect to a
	// single aspect of service quality.
	BasicSli *SloBasicSli `pulumi:"basicSli"`
	// A calendar period, semantically "since the start of the current <calendarPeriod>".
	CalendarPeriod *string `pulumi:"calendarPeriod"`
	// Name used for UI elements listing this SLO.
	DisplayName *string `pulumi:"displayName"`
	// The fraction of service that must be good in order for this objective to be met. 0 < goal <= 0.999
	Goal *float64 `pulumi:"goal"`
	// The full resource name for this service. The syntax is:
	// projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A rolling time period, semantically "in the past X days". Must be between 1 to 30 days, inclusive.
	RollingPeriodDays *int `pulumi:"rollingPeriodDays"`
	// ID of the service to which this SLO belongs.
	Service *string `pulumi:"service"`
	// The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
	SloId *string `pulumi:"sloId"`
}

type SloState struct {
	// Basic Service-Level Indicator (SLI) on a well-known service type. Performance will be computed on the basis of
	// pre-defined metrics. SLIs are used to measure and calculate the quality of the Service's performance with respect to a
	// single aspect of service quality.
	BasicSli SloBasicSliPtrInput
	// A calendar period, semantically "since the start of the current <calendarPeriod>".
	CalendarPeriod pulumi.StringPtrInput
	// Name used for UI elements listing this SLO.
	DisplayName pulumi.StringPtrInput
	// The fraction of service that must be good in order for this objective to be met. 0 < goal <= 0.999
	Goal pulumi.Float64PtrInput
	// The full resource name for this service. The syntax is:
	// projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A rolling time period, semantically "in the past X days". Must be between 1 to 30 days, inclusive.
	RollingPeriodDays pulumi.IntPtrInput
	// ID of the service to which this SLO belongs.
	Service pulumi.StringPtrInput
	// The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
	SloId pulumi.StringPtrInput
}

func (SloState) ElementType() reflect.Type {
	return reflect.TypeOf((*sloState)(nil)).Elem()
}

type sloArgs struct {
	// Basic Service-Level Indicator (SLI) on a well-known service type. Performance will be computed on the basis of
	// pre-defined metrics. SLIs are used to measure and calculate the quality of the Service's performance with respect to a
	// single aspect of service quality.
	BasicSli SloBasicSli `pulumi:"basicSli"`
	// A calendar period, semantically "since the start of the current <calendarPeriod>".
	CalendarPeriod *string `pulumi:"calendarPeriod"`
	// Name used for UI elements listing this SLO.
	DisplayName *string `pulumi:"displayName"`
	// The fraction of service that must be good in order for this objective to be met. 0 < goal <= 0.999
	Goal float64 `pulumi:"goal"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A rolling time period, semantically "in the past X days". Must be between 1 to 30 days, inclusive.
	RollingPeriodDays *int `pulumi:"rollingPeriodDays"`
	// ID of the service to which this SLO belongs.
	Service string `pulumi:"service"`
	// The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
	SloId *string `pulumi:"sloId"`
}

// The set of arguments for constructing a Slo resource.
type SloArgs struct {
	// Basic Service-Level Indicator (SLI) on a well-known service type. Performance will be computed on the basis of
	// pre-defined metrics. SLIs are used to measure and calculate the quality of the Service's performance with respect to a
	// single aspect of service quality.
	BasicSli SloBasicSliInput
	// A calendar period, semantically "since the start of the current <calendarPeriod>".
	CalendarPeriod pulumi.StringPtrInput
	// Name used for UI elements listing this SLO.
	DisplayName pulumi.StringPtrInput
	// The fraction of service that must be good in order for this objective to be met. 0 < goal <= 0.999
	Goal pulumi.Float64Input
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A rolling time period, semantically "in the past X days". Must be between 1 to 30 days, inclusive.
	RollingPeriodDays pulumi.IntPtrInput
	// ID of the service to which this SLO belongs.
	Service pulumi.StringInput
	// The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
	SloId pulumi.StringPtrInput
}

func (SloArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sloArgs)(nil)).Elem()
}