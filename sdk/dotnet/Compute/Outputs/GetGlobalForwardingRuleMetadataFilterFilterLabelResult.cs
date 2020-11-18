// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Outputs
{

    [OutputType]
    public sealed class GetGlobalForwardingRuleMetadataFilterFilterLabelResult
    {
        /// <summary>
        /// The name of the global forwarding rule.
        /// </summary>
        public readonly string Name;
        public readonly string Value;

        [OutputConstructor]
        private GetGlobalForwardingRuleMetadataFilterFilterLabelResult(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}