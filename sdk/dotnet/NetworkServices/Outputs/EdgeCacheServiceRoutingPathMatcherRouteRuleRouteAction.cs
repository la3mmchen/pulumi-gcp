// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.NetworkServices.Outputs
{

    [OutputType]
    public sealed class EdgeCacheServiceRoutingPathMatcherRouteRuleRouteAction
    {
        /// <summary>
        /// The policy to use for defining caching and signed request behaviour for requests that match this route.
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.EdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicy? CdnPolicy;
        /// <summary>
        /// CORSPolicy defines Cross-Origin-Resource-Sharing configuration, including which CORS response headers will be set.
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.EdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCorsPolicy? CorsPolicy;
        /// <summary>
        /// The URL rewrite configuration for requests that match this route.
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.EdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionUrlRewrite? UrlRewrite;

        [OutputConstructor]
        private EdgeCacheServiceRoutingPathMatcherRouteRuleRouteAction(
            Outputs.EdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicy? cdnPolicy,

            Outputs.EdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCorsPolicy? corsPolicy,

            Outputs.EdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionUrlRewrite? urlRewrite)
        {
            CdnPolicy = cdnPolicy;
            CorsPolicy = corsPolicy;
            UrlRewrite = urlRewrite;
        }
    }
}