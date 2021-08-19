// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Represents an environment for an agent. You can create multiple versions of your agent and publish them to separate environments.
 * When you edit an agent, you are editing the draft agent. At any point, you can save the draft agent as an agent version, which is an immutable snapshot of your agent.
 * When you save the draft agent, it is published to the default environment. When you create agent versions, you can publish them to custom environments. You can create a variety of custom environments for testing, development, production, etc.
 *
 * To get more information about Environment, see:
 *
 * * [API documentation](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.environments)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/dialogflow/cx/docs)
 *
 * ## Example Usage
 * ### Dialogflowcx Environment Full
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const agent = new gcp.diagflow.CxAgent("agent", {
 *     displayName: "dialogflowcx-agent",
 *     location: "global",
 *     defaultLanguageCode: "en",
 *     supportedLanguageCodes: [
 *         "fr",
 *         "de",
 *         "es",
 *     ],
 *     timeZone: "America/New_York",
 *     description: "Example description.",
 *     avatarUri: "https://cloud.google.com/_static/images/cloud/icons/favicons/onecloud/super_cloud.png",
 *     enableStackdriverLogging: true,
 *     enableSpellCorrection: true,
 *     speechToTextSettings: {
 *         enableSpeechAdaptation: true,
 *     },
 * });
 * const version1 = new gcp.diagflow.CxVersion("version1", {
 *     parent: agent.startFlow,
 *     displayName: "1.0.0",
 *     description: "version 1.0.0",
 * });
 * const development = new gcp.diagflow.CxEnvironment("development", {
 *     parent: agent.id,
 *     displayName: "Development",
 *     description: "Development Environment",
 *     versionConfigs: [{
 *         version: version1.id,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Environment can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:diagflow/cxEnvironment:CxEnvironment default {{parent}}/environments/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:diagflow/cxEnvironment:CxEnvironment default {{parent}}/{{name}}
 * ```
 */
export class CxEnvironment extends pulumi.CustomResource {
    /**
     * Get an existing CxEnvironment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CxEnvironmentState, opts?: pulumi.CustomResourceOptions): CxEnvironment {
        return new CxEnvironment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:diagflow/cxEnvironment:CxEnvironment';

    /**
     * Returns true if the given object is an instance of CxEnvironment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CxEnvironment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CxEnvironment.__pulumiType;
    }

    /**
     * The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The human-readable name of the environment (unique in an agent). Limit of 64 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The name of the environment.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The Agent to create an Environment for.
     * Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
     */
    public readonly parent!: pulumi.Output<string | undefined>;
    /**
     * Update time of this environment. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
     * fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * A list of configurations for flow versions. You should include version configs for all flows that are reachable from [Start Flow][Agent.start_flow] in the agent. Otherwise, an error will be returned.
     * Structure is documented below.
     */
    public readonly versionConfigs!: pulumi.Output<outputs.diagflow.CxEnvironmentVersionConfig[]>;

    /**
     * Create a CxEnvironment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CxEnvironmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CxEnvironmentArgs | CxEnvironmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CxEnvironmentState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parent"] = state ? state.parent : undefined;
            inputs["updateTime"] = state ? state.updateTime : undefined;
            inputs["versionConfigs"] = state ? state.versionConfigs : undefined;
        } else {
            const args = argsOrState as CxEnvironmentArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.versionConfigs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'versionConfigs'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["parent"] = args ? args.parent : undefined;
            inputs["versionConfigs"] = args ? args.versionConfigs : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(CxEnvironment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CxEnvironment resources.
 */
export interface CxEnvironmentState {
    /**
     * The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
     */
    description?: pulumi.Input<string>;
    /**
     * The human-readable name of the environment (unique in an agent). Limit of 64 characters.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The name of the environment.
     */
    name?: pulumi.Input<string>;
    /**
     * The Agent to create an Environment for.
     * Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
     */
    parent?: pulumi.Input<string>;
    /**
     * Update time of this environment. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
     * fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
     */
    updateTime?: pulumi.Input<string>;
    /**
     * A list of configurations for flow versions. You should include version configs for all flows that are reachable from [Start Flow][Agent.start_flow] in the agent. Otherwise, an error will be returned.
     * Structure is documented below.
     */
    versionConfigs?: pulumi.Input<pulumi.Input<inputs.diagflow.CxEnvironmentVersionConfig>[]>;
}

/**
 * The set of arguments for constructing a CxEnvironment resource.
 */
export interface CxEnvironmentArgs {
    /**
     * The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
     */
    description?: pulumi.Input<string>;
    /**
     * The human-readable name of the environment (unique in an agent). Limit of 64 characters.
     */
    displayName: pulumi.Input<string>;
    /**
     * The Agent to create an Environment for.
     * Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
     */
    parent?: pulumi.Input<string>;
    /**
     * A list of configurations for flow versions. You should include version configs for all flows that are reachable from [Start Flow][Agent.start_flow] in the agent. Otherwise, an error will be returned.
     * Structure is documented below.
     */
    versionConfigs: pulumi.Input<pulumi.Input<inputs.diagflow.CxEnvironmentVersionConfig>[]>;
}