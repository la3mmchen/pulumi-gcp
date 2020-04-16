// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container.Outputs
{

    [OutputType]
    public sealed class ClusterAddonsConfigIstioConfig
    {
        /// <summary>
        /// The authentication type between services in Istio. Available options include `AUTH_MUTUAL_TLS`.
        /// </summary>
        public readonly string? Auth;
        /// <summary>
        /// The status of the Istio addon, which makes it easy to set up Istio for services in a
        /// cluster. It is disabled by default. Set `disabled = false` to enable.
        /// </summary>
        public readonly bool Disabled;

        [OutputConstructor]
        private ClusterAddonsConfigIstioConfig(
            string? auth,

            bool disabled)
        {
            Auth = auth;
            Disabled = disabled;
        }
    }
}