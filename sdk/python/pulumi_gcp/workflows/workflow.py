# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Workflow']


class Workflow(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 service_account: Optional[pulumi.Input[str]] = None,
                 source_contents: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Workflow program to be executed by Workflows.

        To get more information about Workflow, see:

        * [API documentation](https://cloud.google.com/workflows/docs/reference/rest/v1/projects.locations.workflows)
        * How-to Guides
            * [Managing Workflows](https://cloud.google.com/workflows/docs/creating-updating-workflow)

        ## Example Usage
        ### Workflow Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        test_account = gcp.service_account.Account("testAccount",
            account_id="my-account",
            display_name="Test Service Account")
        example = gcp.workflows.Workflow("example",
            region="us-central1",
            description="Magic",
            service_account=test_account.id,
            source_contents=f\"\"\"# This is a sample workflow, feel free to replace it with your source code
        #
        # This workflow does the following:
        # - reads current time and date information from an external API and stores
        #   the response in CurrentDateTime variable
        # - retrieves a list of Wikipedia articles related to the day of the week
        #   from CurrentDateTime
        # - returns the list of articles as an output of the workflow
        # FYI, In terraform you need to escape the $$ or it will cause errors.

        - getCurrentTime:
            call: http.get
            args:
                url: https://us-central1-workflowsample.cloudfunctions.net/datetime
            result: CurrentDateTime
        - readWikipedia:
            call: http.get
            args:
                url: https://en.wikipedia.org/w/api.php
                query:
                    action: opensearch
                    search: {current_date_time["body"]["dayOfTheWeek"]}
            result: WikiResult
        - returnOutput:
            return: {wiki_result["body"]}
        \"\"\")
        ```

        ## Import

        This resource does not support import.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the workflow provided by the user. Must be at most 1000 unicode characters long.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to this Workflow.
        :param pulumi.Input[str] name: Name of the Workflow.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the
               specified prefix. If this and name are unspecified, a random value is chosen for the name.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region of the workflow.
        :param pulumi.Input[str] service_account: Name of the service account associated with the latest workflow version. This service
               account represents the identity of the workflow and determines what permissions the workflow has.
               Format: projects/{project}/serviceAccounts/{account}.
        :param pulumi.Input[str] source_contents: Workflow code to be executed. The size limit is 32KB.
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

            __props__['description'] = description
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['name_prefix'] = name_prefix
            __props__['project'] = project
            __props__['region'] = region
            __props__['service_account'] = service_account
            __props__['source_contents'] = source_contents
            __props__['create_time'] = None
            __props__['revision_id'] = None
            __props__['state'] = None
            __props__['update_time'] = None
        super(Workflow, __self__).__init__(
            'gcp:workflows/workflow:Workflow',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            name_prefix: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            revision_id: Optional[pulumi.Input[str]] = None,
            service_account: Optional[pulumi.Input[str]] = None,
            source_contents: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[str]] = None) -> 'Workflow':
        """
        Get an existing Workflow resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: The timestamp of when the workflow was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
               fractional digits.
        :param pulumi.Input[str] description: Description of the workflow provided by the user. Must be at most 1000 unicode characters long.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to this Workflow.
        :param pulumi.Input[str] name: Name of the Workflow.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the
               specified prefix. If this and name are unspecified, a random value is chosen for the name.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region of the workflow.
        :param pulumi.Input[str] revision_id: The revision of the workflow. A new one is generated if the service account or source contents is changed.
        :param pulumi.Input[str] service_account: Name of the service account associated with the latest workflow version. This service
               account represents the identity of the workflow and determines what permissions the workflow has.
               Format: projects/{project}/serviceAccounts/{account}.
        :param pulumi.Input[str] source_contents: Workflow code to be executed. The size limit is 32KB.
        :param pulumi.Input[str] state: State of the workflow deployment.
        :param pulumi.Input[str] update_time: The timestamp of when the workflow was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
               nine fractional digits.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["create_time"] = create_time
        __props__["description"] = description
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["name_prefix"] = name_prefix
        __props__["project"] = project
        __props__["region"] = region
        __props__["revision_id"] = revision_id
        __props__["service_account"] = service_account
        __props__["source_contents"] = source_contents
        __props__["state"] = state
        __props__["update_time"] = update_time
        return Workflow(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The timestamp of when the workflow was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
        fractional digits.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description of the workflow provided by the user. Must be at most 1000 unicode characters long.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A set of key/value label pairs to assign to this Workflow.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the Workflow.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> pulumi.Output[str]:
        """
        Creates a unique name beginning with the
        specified prefix. If this and name are unspecified, a random value is chosen for the name.
        """
        return pulumi.get(self, "name_prefix")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[Optional[str]]:
        """
        The region of the workflow.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="revisionId")
    def revision_id(self) -> pulumi.Output[str]:
        """
        The revision of the workflow. A new one is generated if the service account or source contents is changed.
        """
        return pulumi.get(self, "revision_id")

    @property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> pulumi.Output[str]:
        """
        Name of the service account associated with the latest workflow version. This service
        account represents the identity of the workflow and determines what permissions the workflow has.
        Format: projects/{project}/serviceAccounts/{account}.
        """
        return pulumi.get(self, "service_account")

    @property
    @pulumi.getter(name="sourceContents")
    def source_contents(self) -> pulumi.Output[Optional[str]]:
        """
        Workflow code to be executed. The size limit is 32KB.
        """
        return pulumi.get(self, "source_contents")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        State of the workflow deployment.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The timestamp of when the workflow was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
        nine fractional digits.
        """
        return pulumi.get(self, "update_time")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
