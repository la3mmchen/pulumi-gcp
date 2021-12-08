// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Vertex.Outputs
{

    [OutputType]
    public sealed class AiMetadataStoreEncryptionSpec
    {
        /// <summary>
        /// Required. The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource.
        /// Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the resource is created.
        /// </summary>
        public readonly string? KmsKeyName;

        [OutputConstructor]
        private AiMetadataStoreEncryptionSpec(string? kmsKeyName)
        {
            KmsKeyName = kmsKeyName;
        }
    }
}