// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dns.Inputs
{

    public sealed class ManagedZonePrivateVisibilityConfigNetworkArgs : Pulumi.ResourceArgs
    {
        [Input("networkUrl", required: true)]
        public Input<string> NetworkUrl { get; set; } = null!;

        public ManagedZonePrivateVisibilityConfigNetworkArgs()
        {
        }
    }
}