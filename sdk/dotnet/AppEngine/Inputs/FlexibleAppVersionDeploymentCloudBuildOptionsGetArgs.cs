// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine.Inputs
{

    public sealed class FlexibleAppVersionDeploymentCloudBuildOptionsGetArgs : Pulumi.ResourceArgs
    {
        [Input("appYamlPath", required: true)]
        public Input<string> AppYamlPath { get; set; } = null!;

        [Input("cloudBuildTimeout")]
        public Input<string>? CloudBuildTimeout { get; set; }

        public FlexibleAppVersionDeploymentCloudBuildOptionsGetArgs()
        {
        }
    }
}