# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class SecurityScanConfig(pulumi.CustomResource):
    authentication: pulumi.Output[dict]
    blacklist_patterns: pulumi.Output[list]
    display_name: pulumi.Output[str]
    export_to_security_command_center: pulumi.Output[str]
    max_qps: pulumi.Output[float]
    name: pulumi.Output[str]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    schedule: pulumi.Output[dict]
    starting_urls: pulumi.Output[list]
    target_platforms: pulumi.Output[list]
    user_agent: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, authentication=None, blacklist_patterns=None, display_name=None, export_to_security_command_center=None, max_qps=None, project=None, schedule=None, starting_urls=None, target_platforms=None, user_agent=None, __name__=None, __opts__=None):
        """
        A ScanConfig resource contains the configurations to launch a scan.
        
        > **Warning:** This resource is in beta, and should be used with the terraform-provider-google-beta provider.
        See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta resources.
        
        To get more information about ScanConfig, see:
        
        * [API documentation](https://cloud.google.com/security-scanner/docs/reference/rest/v1beta/projects.scanConfigs)
        * How-to Guides
            * [Using Cloud Security Scanner](https://cloud.google.com/security-scanner/docs/scanning)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
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

        __props__['authentication'] = authentication

        __props__['blacklist_patterns'] = blacklist_patterns

        if display_name is None:
            raise TypeError("Missing required property 'display_name'")
        __props__['display_name'] = display_name

        __props__['export_to_security_command_center'] = export_to_security_command_center

        __props__['max_qps'] = max_qps

        __props__['project'] = project

        __props__['schedule'] = schedule

        if starting_urls is None:
            raise TypeError("Missing required property 'starting_urls'")
        __props__['starting_urls'] = starting_urls

        __props__['target_platforms'] = target_platforms

        __props__['user_agent'] = user_agent

        __props__['name'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(SecurityScanConfig, __self__).__init__(
            'gcp:compute/securityScanConfig:SecurityScanConfig',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
