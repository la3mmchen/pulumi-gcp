// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine.Outputs
{

    [OutputType]
    public sealed class StandardAppVersionAutomaticScaling
    {
        /// <summary>
        /// Number of concurrent requests an automatic scaling instance can accept before the scheduler spawns a new instance.
        /// Defaults to a runtime-specific value.
        /// </summary>
        public readonly int? MaxConcurrentRequests;
        /// <summary>
        /// Maximum number of idle instances that should be maintained for this version.
        /// </summary>
        public readonly int? MaxIdleInstances;
        /// <summary>
        /// Maximum amount of time that a request should wait in the pending queue before starting a new instance to handle it.
        /// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        /// </summary>
        public readonly string? MaxPendingLatency;
        /// <summary>
        /// Minimum number of idle instances that should be maintained for this version. Only applicable for the default version of a service.
        /// </summary>
        public readonly int? MinIdleInstances;
        /// <summary>
        /// Minimum amount of time a request should wait in the pending queue before starting a new instance to handle it.
        /// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        /// </summary>
        public readonly string? MinPendingLatency;
        /// <summary>
        /// Scheduler settings for standard environment.  Structure is documented below.
        /// </summary>
        public readonly Outputs.StandardAppVersionAutomaticScalingStandardSchedulerSettings? StandardSchedulerSettings;

        [OutputConstructor]
        private StandardAppVersionAutomaticScaling(
            int? maxConcurrentRequests,

            int? maxIdleInstances,

            string? maxPendingLatency,

            int? minIdleInstances,

            string? minPendingLatency,

            Outputs.StandardAppVersionAutomaticScalingStandardSchedulerSettings? standardSchedulerSettings)
        {
            MaxConcurrentRequests = maxConcurrentRequests;
            MaxIdleInstances = maxIdleInstances;
            MaxPendingLatency = maxPendingLatency;
            MinIdleInstances = minIdleInstances;
            MinPendingLatency = minPendingLatency;
            StandardSchedulerSettings = standardSchedulerSettings;
        }
    }
}