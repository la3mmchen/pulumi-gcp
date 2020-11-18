# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['PreventionDeidentifyTemplate']


class PreventionDeidentifyTemplate(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deidentify_config: Optional[pulumi.Input[pulumi.InputType['PreventionDeidentifyTemplateDeidentifyConfigArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Allows creation of templates to de-idenfity content.

        To get more information about DeidentifyTemplate, see:

        * [API documentation](https://cloud.google.com/dlp/docs/reference/rest/v2/projects.deidentifyTemplates)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/dlp/docs/concepts-templates)

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PreventionDeidentifyTemplateDeidentifyConfigArgs']] deidentify_config: Configuration of the deidentify template
               Structure is documented below.
        :param pulumi.Input[str] description: A description of the template.
        :param pulumi.Input[str] display_name: User set display name of the template.
        :param pulumi.Input[str] parent: The parent of the template in any of the following formats:
               * `projects/{{project}}`
               * `projects/{{project}}/locations/{{location}}`
               * `organizations/{{organization_id}}`
               * `organizations/{{organization_id}}/locations/{{location}}`
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if deidentify_config is None:
                raise TypeError("Missing required property 'deidentify_config'")
            __props__['deidentify_config'] = deidentify_config
            __props__['description'] = description
            __props__['display_name'] = display_name
            if parent is None:
                raise TypeError("Missing required property 'parent'")
            __props__['parent'] = parent
            __props__['name'] = None
        super(PreventionDeidentifyTemplate, __self__).__init__(
            'gcp:dataloss/preventionDeidentifyTemplate:PreventionDeidentifyTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            deidentify_config: Optional[pulumi.Input[pulumi.InputType['PreventionDeidentifyTemplateDeidentifyConfigArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parent: Optional[pulumi.Input[str]] = None) -> 'PreventionDeidentifyTemplate':
        """
        Get an existing PreventionDeidentifyTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PreventionDeidentifyTemplateDeidentifyConfigArgs']] deidentify_config: Configuration of the deidentify template
               Structure is documented below.
        :param pulumi.Input[str] description: A description of the template.
        :param pulumi.Input[str] display_name: User set display name of the template.
        :param pulumi.Input[str] name: Name of the information type.
        :param pulumi.Input[str] parent: The parent of the template in any of the following formats:
               * `projects/{{project}}`
               * `projects/{{project}}/locations/{{location}}`
               * `organizations/{{organization_id}}`
               * `organizations/{{organization_id}}/locations/{{location}}`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["deidentify_config"] = deidentify_config
        __props__["description"] = description
        __props__["display_name"] = display_name
        __props__["name"] = name
        __props__["parent"] = parent
        return PreventionDeidentifyTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deidentifyConfig")
    def deidentify_config(self) -> pulumi.Output['outputs.PreventionDeidentifyTemplateDeidentifyConfig']:
        """
        Configuration of the deidentify template
        Structure is documented below.
        """
        return pulumi.get(self, "deidentify_config")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the template.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        User set display name of the template.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the information type.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parent(self) -> pulumi.Output[str]:
        """
        The parent of the template in any of the following formats:
        * `projects/{{project}}`
        * `projects/{{project}}/locations/{{location}}`
        * `organizations/{{organization_id}}`
        * `organizations/{{organization_id}}/locations/{{location}}`
        """
        return pulumi.get(self, "parent")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
