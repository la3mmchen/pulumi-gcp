// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Diagflow.Inputs
{

    public sealed class CxPageFormParameterGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The human-readable name of the parameter, unique within the form.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The entity type of the parameter.
        /// Format: projects/-/locations/-/agents/-/entityTypes/&lt;System Entity Type ID&gt; for system entity types (for example, projects/-/locations/-/agents/-/entityTypes/sys.date), or projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;/entityTypes/&lt;Entity Type ID&gt; for developer entity types.
        /// </summary>
        [Input("entityType")]
        public Input<string>? EntityType { get; set; }

        /// <summary>
        /// Defines fill behavior for the parameter.
        /// Structure is documented below.
        /// </summary>
        [Input("fillBehavior")]
        public Input<Inputs.CxPageFormParameterFillBehaviorGetArgs>? FillBehavior { get; set; }

        /// <summary>
        /// Indicates whether the parameter represents a list of values.
        /// </summary>
        [Input("isList")]
        public Input<bool>? IsList { get; set; }

        /// <summary>
        /// Indicates whether the parameter content should be redacted in log.
        /// If redaction is enabled, the parameter content will be replaced by parameter name during logging. Note: the parameter content is subject to redaction if either parameter level redaction or entity type level redaction is enabled.
        /// </summary>
        [Input("redact")]
        public Input<bool>? Redact { get; set; }

        /// <summary>
        /// Indicates whether the parameter is required. Optional parameters will not trigger prompts; however, they are filled if the user specifies them.
        /// Required parameters must be filled before form filling concludes.
        /// </summary>
        [Input("required")]
        public Input<bool>? Required { get; set; }

        public CxPageFormParameterGetArgs()
        {
        }
    }
}