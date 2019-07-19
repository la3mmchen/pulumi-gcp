// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_fhir_store_iam_policy.html.markdown.
type FhirStoreIamPolicy struct {
	s *pulumi.ResourceState
}

// NewFhirStoreIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewFhirStoreIamPolicy(ctx *pulumi.Context,
	name string, args *FhirStoreIamPolicyArgs, opts ...pulumi.ResourceOpt) (*FhirStoreIamPolicy, error) {
	if args == nil || args.FhirStoreId == nil {
		return nil, errors.New("missing required argument 'FhirStoreId'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["fhirStoreId"] = nil
		inputs["policyData"] = nil
	} else {
		inputs["fhirStoreId"] = args.FhirStoreId
		inputs["policyData"] = args.PolicyData
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:healthcare/fhirStoreIamPolicy:FhirStoreIamPolicy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FhirStoreIamPolicy{s: s}, nil
}

// GetFhirStoreIamPolicy gets an existing FhirStoreIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFhirStoreIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *FhirStoreIamPolicyState, opts ...pulumi.ResourceOpt) (*FhirStoreIamPolicy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["etag"] = state.Etag
		inputs["fhirStoreId"] = state.FhirStoreId
		inputs["policyData"] = state.PolicyData
	}
	s, err := ctx.ReadResource("gcp:healthcare/fhirStoreIamPolicy:FhirStoreIamPolicy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FhirStoreIamPolicy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *FhirStoreIamPolicy) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *FhirStoreIamPolicy) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// (Computed) The etag of the FHIR store's IAM policy.
func (r *FhirStoreIamPolicy) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

// The FHIR store ID, in the form
// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
// project setting will be used as a fallback.
func (r *FhirStoreIamPolicy) FhirStoreId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["fhirStoreId"])
}

// The policy data generated by
// a `google_iam_policy` data source.
func (r *FhirStoreIamPolicy) PolicyData() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["policyData"])
}

// Input properties used for looking up and filtering FhirStoreIamPolicy resources.
type FhirStoreIamPolicyState struct {
	// (Computed) The etag of the FHIR store's IAM policy.
	Etag interface{}
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	FhirStoreId interface{}
	// The policy data generated by
	// a `google_iam_policy` data source.
	PolicyData interface{}
}

// The set of arguments for constructing a FhirStoreIamPolicy resource.
type FhirStoreIamPolicyArgs struct {
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	FhirStoreId interface{}
	// The policy data generated by
	// a `google_iam_policy` data source.
	PolicyData interface{}
}