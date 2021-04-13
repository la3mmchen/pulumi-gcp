// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Apigee
{
    /// <summary>
    /// An `Environment Group attachment` in Apigee.
    /// 
    /// To get more information about EnvgroupAttachment, see:
    /// 
    /// * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.envgroups.attachments/create)
    /// * How-to Guides
    ///     * [Creating an environment](https://cloud.google.com/apigee/docs/api-platform/get-started/create-environment)
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// EnvgroupAttachment can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:apigee/envGroupAttachment:EnvGroupAttachment default {{envgroup_id}}/attachments/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:apigee/envGroupAttachment:EnvGroupAttachment default {{envgroup_id}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:apigee/envGroupAttachment:EnvGroupAttachment")]
    public partial class EnvGroupAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// The Apigee environment group associated with the Apigee environment,
        /// in the format `organizations/{{org_name}}/envgroups/{{envgroup_name}}`.
        /// </summary>
        [Output("envgroupId")]
        public Output<string> EnvgroupId { get; private set; } = null!;

        /// <summary>
        /// The resource ID of the environment.
        /// </summary>
        [Output("environment")]
        public Output<string> Environment { get; private set; } = null!;

        /// <summary>
        /// The name of the newly created attachment (output parameter).
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a EnvGroupAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EnvGroupAttachment(string name, EnvGroupAttachmentArgs args, CustomResourceOptions? options = null)
            : base("gcp:apigee/envGroupAttachment:EnvGroupAttachment", name, args ?? new EnvGroupAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EnvGroupAttachment(string name, Input<string> id, EnvGroupAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("gcp:apigee/envGroupAttachment:EnvGroupAttachment", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EnvGroupAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EnvGroupAttachment Get(string name, Input<string> id, EnvGroupAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new EnvGroupAttachment(name, id, state, options);
        }
    }

    public sealed class EnvGroupAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Apigee environment group associated with the Apigee environment,
        /// in the format `organizations/{{org_name}}/envgroups/{{envgroup_name}}`.
        /// </summary>
        [Input("envgroupId", required: true)]
        public Input<string> EnvgroupId { get; set; } = null!;

        /// <summary>
        /// The resource ID of the environment.
        /// </summary>
        [Input("environment", required: true)]
        public Input<string> Environment { get; set; } = null!;

        public EnvGroupAttachmentArgs()
        {
        }
    }

    public sealed class EnvGroupAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Apigee environment group associated with the Apigee environment,
        /// in the format `organizations/{{org_name}}/envgroups/{{envgroup_name}}`.
        /// </summary>
        [Input("envgroupId")]
        public Input<string>? EnvgroupId { get; set; }

        /// <summary>
        /// The resource ID of the environment.
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        /// <summary>
        /// The name of the newly created attachment (output parameter).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public EnvGroupAttachmentState()
        {
        }
    }
}