// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GameServices.Inputs
{

    public sealed class GameServerClusterConnectionInfoGetArgs : Pulumi.ResourceArgs
    {
        [Input("gkeClusterReference", required: true)]
        public Input<Inputs.GameServerClusterConnectionInfoGkeClusterReferenceGetArgs> GkeClusterReference { get; set; } = null!;

        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        public GameServerClusterConnectionInfoGetArgs()
        {
        }
    }
}