// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class NodeGroupAutoscalingPolicyGetArgs : Pulumi.ResourceArgs
    {
        [Input("maxNodes")]
        public Input<int>? MaxNodes { get; set; }

        [Input("minNodes")]
        public Input<int>? MinNodes { get; set; }

        [Input("mode")]
        public Input<string>? Mode { get; set; }

        public NodeGroupAutoscalingPolicyGetArgs()
        {
        }
    }
}