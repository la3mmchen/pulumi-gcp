# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'NotificationConfigStreamingConfigArgs',
]

@pulumi.input_type
class NotificationConfigStreamingConfigArgs:
    def __init__(__self__, *,
                 filter: pulumi.Input[str]):
        """
        :param pulumi.Input[str] filter: Expression that defines the filter to apply across create/update
               events of assets or findings as specified by the event type. The
               expression is a list of zero or more restrictions combined via
               logical operators AND and OR. Parentheses are supported, and OR
               has higher precedence than AND.
               Restrictions have the form <field> <operator> <value> and may have
               a - character in front of them to indicate negation. The fields
               map to those defined in the corresponding resource.
               The supported operators are:
               * = for all value types.
               * >, <, >=, <= for integer values.
               * :, meaning substring matching, for strings.
               The supported value types are:
               * string literals in quotes.
               * integer literals without quotes.
               * boolean literals true and false without quotes.
               See
               [Filtering notifications](https://cloud.google.com/security-command-center/docs/how-to-api-filter-notifications)
               for information on how to write a filter.
        """
        pulumi.set(__self__, "filter", filter)

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Input[str]:
        """
        Expression that defines the filter to apply across create/update
        events of assets or findings as specified by the event type. The
        expression is a list of zero or more restrictions combined via
        logical operators AND and OR. Parentheses are supported, and OR
        has higher precedence than AND.
        Restrictions have the form <field> <operator> <value> and may have
        a - character in front of them to indicate negation. The fields
        map to those defined in the corresponding resource.
        The supported operators are:
        * = for all value types.
        * >, <, >=, <= for integer values.
        * :, meaning substring matching, for strings.
        The supported value types are:
        * string literals in quotes.
        * integer literals without quotes.
        * boolean literals true and false without quotes.
        See
        [Filtering notifications](https://cloud.google.com/security-command-center/docs/how-to-api-filter-notifications)
        for information on how to write a filter.
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: pulumi.Input[str]):
        pulumi.set(self, "filter", value)

