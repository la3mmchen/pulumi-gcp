// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about a Google IAM Role.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const roleinfo = pulumi.output(gcp.iam.getRule({
 *     name: "roles/compute.viewer",
 * }));
 * 
 * export const theRolePermissions = roleinfo.apply(roleinfo => roleinfo.includedPermissions);
 * ```
 */
export function getRule(args: GetRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetRuleResult> {
    return pulumi.runtime.invoke("gcp:iam/getRule:getRule", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getRule.
 */
export interface GetRuleArgs {
    /**
     * The name of the Role to lookup in the form `roles/{ROLE_NAME}`, `organizations/{ORGANIZATION_ID}/roles/{ROLE_NAME}` or `projects/{PROJECT_ID}/roles/{ROLE_NAME}`
     */
    readonly name: string;
}

/**
 * A collection of values returned by getRule.
 */
export interface GetRuleResult {
    /**
     * specifies the list of one or more permissions to include in the custom role, such as - `iam.roles.get`
     */
    readonly includedPermissions: string[];
    /**
     * indicates the stage of a role in the launch lifecycle, such as `GA`, `BETA` or `ALPHA`.
     */
    readonly stage: string;
    /**
     * is a friendly title for the role, such as "Role Viewer"
     */
    readonly title: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}