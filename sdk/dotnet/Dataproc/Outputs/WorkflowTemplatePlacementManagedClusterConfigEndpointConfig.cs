// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc.Outputs
{

    [OutputType]
    public sealed class WorkflowTemplatePlacementManagedClusterConfigEndpointConfig
    {
        /// <summary>
        /// Optional. If true, enable http access to specific ports on the cluster from external sources. Defaults to false.
        /// </summary>
        public readonly bool? EnableHttpPortAccess;
        /// <summary>
        /// -
        /// Output only. The map of port descriptions to URLs. Will only be populated if enable_http_port_access is true.
        /// The `gce_cluster_config` block supports:
        /// </summary>
        public readonly ImmutableDictionary<string, string>? HttpPorts;

        [OutputConstructor]
        private WorkflowTemplatePlacementManagedClusterConfigEndpointConfig(
            bool? enableHttpPortAccess,

            ImmutableDictionary<string, string>? httpPorts)
        {
            EnableHttpPortAccess = enableHttpPortAccess;
            HttpPorts = httpPorts;
        }
    }
}