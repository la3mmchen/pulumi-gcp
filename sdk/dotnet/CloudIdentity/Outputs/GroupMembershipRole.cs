// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudIdentity.Outputs
{

    [OutputType]
    public sealed class GroupMembershipRole
    {
        /// <summary>
        /// The name of the MembershipRole. Must be one of OWNER, MANAGER, MEMBER.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GroupMembershipRole(string name)
        {
            Name = name;
        }
    }
}