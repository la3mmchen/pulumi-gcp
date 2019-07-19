# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Hl7StoreIamPolicy(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the HL7v2 store's IAM policy.
    """
    hl7_v2_store_id: pulumi.Output[str]
    """
    The HL7v2 store ID, in the form
    `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
    `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
    project setting will be used as a fallback.
    """
    policy_data: pulumi.Output[str]
    """
    The policy data generated by
    a `google_iam_policy` data source.
    """
    def __init__(__self__, resource_name, opts=None, hl7_v2_store_id=None, policy_data=None, __name__=None, __opts__=None):
        """
        Create a Hl7StoreIamPolicy resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] hl7_v2_store_id: The HL7v2 store ID, in the form
               `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
               `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
               project setting will be used as a fallback.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `google_iam_policy` data source.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_hl7_v2_store_iam_policy.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if hl7_v2_store_id is None:
            raise TypeError("Missing required property 'hl7_v2_store_id'")
        __props__['hl7_v2_store_id'] = hl7_v2_store_id

        if policy_data is None:
            raise TypeError("Missing required property 'policy_data'")
        __props__['policy_data'] = policy_data

        __props__['etag'] = None

        super(Hl7StoreIamPolicy, __self__).__init__(
            'gcp:healthcare/hl7StoreIamPolicy:Hl7StoreIamPolicy',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
