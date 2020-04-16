// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class RegionBackendServiceConsistentHashArgs : Pulumi.ResourceArgs
    {
        [Input("httpCookie")]
        public Input<Inputs.RegionBackendServiceConsistentHashHttpCookieArgs>? HttpCookie { get; set; }

        [Input("httpHeaderName")]
        public Input<string>? HttpHeaderName { get; set; }

        [Input("minimumRingSize")]
        public Input<int>? MinimumRingSize { get; set; }

        public RegionBackendServiceConsistentHashArgs()
        {
        }
    }
}