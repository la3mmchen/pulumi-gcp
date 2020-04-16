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
    public sealed class SecurityScanConfigAuthenticationCustomAccount
    {
        public readonly string LoginUrl;
        public readonly string Password;
        public readonly string Username;

        [OutputConstructor]
        private SecurityScanConfigAuthenticationCustomAccount(
            string loginUrl,

            string password,

            string username)
        {
            LoginUrl = loginUrl;
            Password = password;
            Username = username;
        }
    }
}