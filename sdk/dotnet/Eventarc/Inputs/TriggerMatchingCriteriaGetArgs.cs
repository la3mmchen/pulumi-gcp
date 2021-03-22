// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Eventarc.Inputs
{

    public sealed class TriggerMatchingCriteriaGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The name of a CloudEvents attribute. Currently, only a subset of attributes can be specified. All triggers MUST provide a matching criteria for the 'type' attribute.
        /// </summary>
        [Input("attribute", required: true)]
        public Input<string> Attribute { get; set; } = null!;

        /// <summary>
        /// Required. The value for the attribute.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public TriggerMatchingCriteriaGetArgs()
        {
        }
    }
}