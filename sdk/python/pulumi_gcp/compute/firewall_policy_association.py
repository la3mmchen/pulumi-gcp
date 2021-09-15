# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FirewallPolicyAssociationArgs', 'FirewallPolicyAssociation']

@pulumi.input_type
class FirewallPolicyAssociationArgs:
    def __init__(__self__, *,
                 attachment_target: pulumi.Input[str],
                 firewall_policy: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallPolicyAssociation resource.
        :param pulumi.Input[str] attachment_target: The target that the firewall policy is attached to.
        :param pulumi.Input[str] firewall_policy: The firewall policy ID of the association.
        :param pulumi.Input[str] name: The name for an association.
        """
        pulumi.set(__self__, "attachment_target", attachment_target)
        pulumi.set(__self__, "firewall_policy", firewall_policy)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="attachmentTarget")
    def attachment_target(self) -> pulumi.Input[str]:
        """
        The target that the firewall policy is attached to.
        """
        return pulumi.get(self, "attachment_target")

    @attachment_target.setter
    def attachment_target(self, value: pulumi.Input[str]):
        pulumi.set(self, "attachment_target", value)

    @property
    @pulumi.getter(name="firewallPolicy")
    def firewall_policy(self) -> pulumi.Input[str]:
        """
        The firewall policy ID of the association.
        """
        return pulumi.get(self, "firewall_policy")

    @firewall_policy.setter
    def firewall_policy(self, value: pulumi.Input[str]):
        pulumi.set(self, "firewall_policy", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name for an association.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _FirewallPolicyAssociationState:
    def __init__(__self__, *,
                 attachment_target: Optional[pulumi.Input[str]] = None,
                 firewall_policy: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 short_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallPolicyAssociation resources.
        :param pulumi.Input[str] attachment_target: The target that the firewall policy is attached to.
        :param pulumi.Input[str] firewall_policy: The firewall policy ID of the association.
        :param pulumi.Input[str] name: The name for an association.
        :param pulumi.Input[str] short_name: The short name of the firewall policy of the association.
        """
        if attachment_target is not None:
            pulumi.set(__self__, "attachment_target", attachment_target)
        if firewall_policy is not None:
            pulumi.set(__self__, "firewall_policy", firewall_policy)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if short_name is not None:
            pulumi.set(__self__, "short_name", short_name)

    @property
    @pulumi.getter(name="attachmentTarget")
    def attachment_target(self) -> Optional[pulumi.Input[str]]:
        """
        The target that the firewall policy is attached to.
        """
        return pulumi.get(self, "attachment_target")

    @attachment_target.setter
    def attachment_target(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "attachment_target", value)

    @property
    @pulumi.getter(name="firewallPolicy")
    def firewall_policy(self) -> Optional[pulumi.Input[str]]:
        """
        The firewall policy ID of the association.
        """
        return pulumi.get(self, "firewall_policy")

    @firewall_policy.setter
    def firewall_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "firewall_policy", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name for an association.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="shortName")
    def short_name(self) -> Optional[pulumi.Input[str]]:
        """
        The short name of the firewall policy of the association.
        """
        return pulumi.get(self, "short_name")

    @short_name.setter
    def short_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "short_name", value)


class FirewallPolicyAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attachment_target: Optional[pulumi.Input[str]] = None,
                 firewall_policy: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Allows associating hierarchical firewall policies with the target where they are applied. This allows creating policies and rules in a different location than they are applied.

        For more information on applying hierarchical firewall policies see the [official documentation](https://cloud.google.com/vpc/docs/firewall-policies#managing_hierarchical_firewall_policy_resources)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_firewall_policy = gcp.compute.FirewallPolicy("defaultFirewallPolicy",
            parent="organizations/12345",
            short_name="my-policy",
            description="Example Resource")
        default_firewall_policy_association = gcp.compute.FirewallPolicyAssociation("defaultFirewallPolicyAssociation",
            firewall_policy=default_firewall_policy.id,
            attachment_target=google_folder["folder"]["name"])
        ```

        ## Import

        FirewallPolicyAssociation can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/firewallPolicyAssociation:FirewallPolicyAssociation default locations/global/firewallPolicies/{{firewall_policy}}/associations/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/firewallPolicyAssociation:FirewallPolicyAssociation default {{firewall_policy}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attachment_target: The target that the firewall policy is attached to.
        :param pulumi.Input[str] firewall_policy: The firewall policy ID of the association.
        :param pulumi.Input[str] name: The name for an association.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallPolicyAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows associating hierarchical firewall policies with the target where they are applied. This allows creating policies and rules in a different location than they are applied.

        For more information on applying hierarchical firewall policies see the [official documentation](https://cloud.google.com/vpc/docs/firewall-policies#managing_hierarchical_firewall_policy_resources)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_firewall_policy = gcp.compute.FirewallPolicy("defaultFirewallPolicy",
            parent="organizations/12345",
            short_name="my-policy",
            description="Example Resource")
        default_firewall_policy_association = gcp.compute.FirewallPolicyAssociation("defaultFirewallPolicyAssociation",
            firewall_policy=default_firewall_policy.id,
            attachment_target=google_folder["folder"]["name"])
        ```

        ## Import

        FirewallPolicyAssociation can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/firewallPolicyAssociation:FirewallPolicyAssociation default locations/global/firewallPolicies/{{firewall_policy}}/associations/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/firewallPolicyAssociation:FirewallPolicyAssociation default {{firewall_policy}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param FirewallPolicyAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallPolicyAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attachment_target: Optional[pulumi.Input[str]] = None,
                 firewall_policy: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallPolicyAssociationArgs.__new__(FirewallPolicyAssociationArgs)

            if attachment_target is None and not opts.urn:
                raise TypeError("Missing required property 'attachment_target'")
            __props__.__dict__["attachment_target"] = attachment_target
            if firewall_policy is None and not opts.urn:
                raise TypeError("Missing required property 'firewall_policy'")
            __props__.__dict__["firewall_policy"] = firewall_policy
            __props__.__dict__["name"] = name
            __props__.__dict__["short_name"] = None
        super(FirewallPolicyAssociation, __self__).__init__(
            'gcp:compute/firewallPolicyAssociation:FirewallPolicyAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attachment_target: Optional[pulumi.Input[str]] = None,
            firewall_policy: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            short_name: Optional[pulumi.Input[str]] = None) -> 'FirewallPolicyAssociation':
        """
        Get an existing FirewallPolicyAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attachment_target: The target that the firewall policy is attached to.
        :param pulumi.Input[str] firewall_policy: The firewall policy ID of the association.
        :param pulumi.Input[str] name: The name for an association.
        :param pulumi.Input[str] short_name: The short name of the firewall policy of the association.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallPolicyAssociationState.__new__(_FirewallPolicyAssociationState)

        __props__.__dict__["attachment_target"] = attachment_target
        __props__.__dict__["firewall_policy"] = firewall_policy
        __props__.__dict__["name"] = name
        __props__.__dict__["short_name"] = short_name
        return FirewallPolicyAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="attachmentTarget")
    def attachment_target(self) -> pulumi.Output[str]:
        """
        The target that the firewall policy is attached to.
        """
        return pulumi.get(self, "attachment_target")

    @property
    @pulumi.getter(name="firewallPolicy")
    def firewall_policy(self) -> pulumi.Output[str]:
        """
        The firewall policy ID of the association.
        """
        return pulumi.get(self, "firewall_policy")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name for an association.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="shortName")
    def short_name(self) -> pulumi.Output[str]:
        """
        The short name of the firewall policy of the association.
        """
        return pulumi.get(self, "short_name")
