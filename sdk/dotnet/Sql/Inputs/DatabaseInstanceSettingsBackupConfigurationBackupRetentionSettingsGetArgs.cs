// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Sql.Inputs
{

    public sealed class DatabaseInstanceSettingsBackupConfigurationBackupRetentionSettingsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Depending on the value of retention_unit, this is used to determine if a backup needs to be deleted. If retention_unit
        /// is 'COUNT', we will retain this many backups.
        /// </summary>
        [Input("retainedBackups", required: true)]
        public Input<int> RetainedBackups { get; set; } = null!;

        /// <summary>
        /// The unit that 'retained_backups' represents. Defaults to `COUNT`.
        /// </summary>
        [Input("retentionUnit")]
        public Input<string>? RetentionUnit { get; set; }

        public DatabaseInstanceSettingsBackupConfigurationBackupRetentionSettingsGetArgs()
        {
        }
    }
}