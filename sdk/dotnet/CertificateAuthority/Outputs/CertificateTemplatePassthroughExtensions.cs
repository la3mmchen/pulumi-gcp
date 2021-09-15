// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CertificateAuthority.Outputs
{

    [OutputType]
    public sealed class CertificateTemplatePassthroughExtensions
    {
        /// <summary>
        /// Optional. Describes custom X.509 extensions.
        /// </summary>
        public readonly ImmutableArray<Outputs.CertificateTemplatePassthroughExtensionsAdditionalExtension> AdditionalExtensions;
        /// <summary>
        /// Optional. A set of named X.509 extensions. Will be combined with additional_extensions to determine the full set of X.509 extensions.
        /// </summary>
        public readonly ImmutableArray<string> KnownExtensions;

        [OutputConstructor]
        private CertificateTemplatePassthroughExtensions(
            ImmutableArray<Outputs.CertificateTemplatePassthroughExtensionsAdditionalExtension> additionalExtensions,

            ImmutableArray<string> knownExtensions)
        {
            AdditionalExtensions = additionalExtensions;
            KnownExtensions = knownExtensions;
        }
    }
}