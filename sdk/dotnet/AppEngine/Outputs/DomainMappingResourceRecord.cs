// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine.Outputs
{

    [OutputType]
    public sealed class DomainMappingResourceRecord
    {
        public readonly string? Name;
        public readonly string? Rrdata;
        public readonly string? Type;

        [OutputConstructor]
        private DomainMappingResourceRecord(
            string? name,

            string? rrdata,

            string? type)
        {
            Name = name;
            Rrdata = rrdata;
            Type = type;
        }
    }
}