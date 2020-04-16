// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class URLMapPathMatcherRouteRuleMatchRuleMetadataFilterArgs : Pulumi.ResourceArgs
    {
        [Input("filterLabels", required: true)]
        private InputList<Inputs.URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArgs>? _filterLabels;
        public InputList<Inputs.URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArgs> FilterLabels
        {
            get => _filterLabels ?? (_filterLabels = new InputList<Inputs.URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArgs>());
            set => _filterLabels = value;
        }

        [Input("filterMatchCriteria", required: true)]
        public Input<string> FilterMatchCriteria { get; set; } = null!;

        public URLMapPathMatcherRouteRuleMatchRuleMetadataFilterArgs()
        {
        }
    }
}