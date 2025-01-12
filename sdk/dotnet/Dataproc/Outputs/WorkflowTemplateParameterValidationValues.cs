// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc.Outputs
{

    [OutputType]
    public sealed class WorkflowTemplateParameterValidationValues
    {
        /// <summary>
        /// Optional. Corresponds to the label values of reservation resource.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private WorkflowTemplateParameterValidationValues(ImmutableArray<string> values)
        {
            Values = values;
        }
    }
}
