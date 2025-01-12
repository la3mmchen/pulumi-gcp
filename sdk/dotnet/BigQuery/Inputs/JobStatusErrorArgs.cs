// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigQuery.Inputs
{

    public sealed class JobStatusErrorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The geographic location of the job. The default value is US.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("message")]
        public Input<string>? Message { get; set; }

        [Input("reason")]
        public Input<string>? Reason { get; set; }

        public JobStatusErrorArgs()
        {
        }
    }
}
