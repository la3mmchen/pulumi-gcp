// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CertificateAuthority.Inputs
{

    public sealed class CertificateTemplatePassthroughExtensionsArgs : Pulumi.ResourceArgs
    {
        [Input("additionalExtensions")]
        private InputList<Inputs.CertificateTemplatePassthroughExtensionsAdditionalExtensionArgs>? _additionalExtensions;

        /// <summary>
        /// Optional. Describes custom X.509 extensions.
        /// </summary>
        public InputList<Inputs.CertificateTemplatePassthroughExtensionsAdditionalExtensionArgs> AdditionalExtensions
        {
            get => _additionalExtensions ?? (_additionalExtensions = new InputList<Inputs.CertificateTemplatePassthroughExtensionsAdditionalExtensionArgs>());
            set => _additionalExtensions = value;
        }

        [Input("knownExtensions")]
        private InputList<string>? _knownExtensions;

        /// <summary>
        /// Optional. A set of named X.509 extensions. Will be combined with additional_extensions to determine the full set of X.509 extensions.
        /// </summary>
        public InputList<string> KnownExtensions
        {
            get => _knownExtensions ?? (_knownExtensions = new InputList<string>());
            set => _knownExtensions = value;
        }

        public CertificateTemplatePassthroughExtensionsArgs()
        {
        }
    }
}