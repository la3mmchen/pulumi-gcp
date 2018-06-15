# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class SharedVPCServiceProject(pulumi.CustomResource):
    """
    Enables the Google Compute Engine
    [Shared VPC](https://cloud.google.com/compute/docs/shared-vpc)
    feature for a project, assigning it as a Shared VPC service project associated
    with a given host project.
    
    For more information, see,
    [the Project API documentation](https://cloud.google.com/compute/docs/reference/latest/projects),
    where the Shared VPC feature is referred to by its former name "XPN".
    """
    def __init__(__self__, __name__, __opts__=None, host_project=None, service_project=None):
        """Create a SharedVPCServiceProject resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not host_project:
            raise TypeError('Missing required property host_project')
        elif not isinstance(host_project, basestring):
            raise TypeError('Expected property host_project to be a basestring')
        __self__.host_project = host_project
        """
        The ID of a host project to associate.
        """
        __props__['hostProject'] = host_project

        if not service_project:
            raise TypeError('Missing required property service_project')
        elif not isinstance(service_project, basestring):
            raise TypeError('Expected property service_project to be a basestring')
        __self__.service_project = service_project
        """
        The ID of the project that will serve as a Shared VPC service project.
        """
        __props__['serviceProject'] = service_project

        super(SharedVPCServiceProject, __self__).__init__(
            'gcp:compute/sharedVPCServiceProject:SharedVPCServiceProject',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'hostProject' in outs:
            self.host_project = outs['hostProject']
        if 'serviceProject' in outs:
            self.service_project = outs['serviceProject']