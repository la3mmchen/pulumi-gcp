// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * A Google Cloud Filestore instance.
 *
 * To get more information about Instance, see:
 *
 * * [API documentation](https://cloud.google.com/filestore/docs/reference/rest/v1beta1/projects.locations.instances/create)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/filestore/docs/creating-instances)
 *     * [Use with Kubernetes](https://cloud.google.com/filestore/docs/accessing-fileshares)
 *     * [Copying Data In/Out](https://cloud.google.com/filestore/docs/copying-data)
 *
 * ## Example Usage
 * ### Filestore Instance Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const instance = new gcp.filestore.Instance("instance", {
 *     fileShares: {
 *         capacityGb: 2660,
 *         name: "share1",
 *     },
 *     networks: [{
 *         modes: ["MODE_IPV4"],
 *         network: "default",
 *     }],
 *     tier: "PREMIUM",
 *     zone: "us-central1-b",
 * });
 * ```
 * ### Filestore Instance Full
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const instance = new gcp.filestore.Instance("instance", {
 *     zone: "us-central1-b",
 *     tier: "BASIC_SSD",
 *     fileShares: {
 *         capacityGb: 2660,
 *         name: "share1",
 *         nfsExportOptions: [
 *             {
 *                 ipRanges: ["10.0.0.0/24"],
 *                 accessMode: "READ_WRITE",
 *                 squashMode: "NO_ROOT_SQUASH",
 *             },
 *             {
 *                 ipRanges: ["10.10.0.0/24"],
 *                 accessMode: "READ_ONLY",
 *                 squashMode: "ROOT_SQUASH",
 *                 anonUid: 123,
 *                 anonGid: 456,
 *             },
 *         ],
 *     },
 *     networks: [{
 *         network: "default",
 *         modes: ["MODE_IPV4"],
 *         connectMode: "DIRECT_PEERING",
 *     }],
 * }, {
 *     provider: google_beta,
 * });
 * ```
 *
 * ## Import
 *
 * Instance can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:filestore/instance:Instance default projects/{{project}}/locations/{{zone}}/instances/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:filestore/instance:Instance default {{project}}/{{zone}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:filestore/instance:Instance default {{zone}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:filestore/instance:Instance default {{name}}
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:filestore/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * A description of the instance.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * File system shares on the instance. For this version, only a
     * single file share is supported.
     * Structure is documented below.
     */
    public readonly fileShares!: pulumi.Output<outputs.filestore.InstanceFileShares>;
    /**
     * Resource labels to represent user-provided metadata.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of the fileshare (16 characters or less)
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * VPC networks to which the instance is connected. For this version,
     * only a single network is supported.
     * Structure is documented below.
     */
    public readonly networks!: pulumi.Output<outputs.filestore.InstanceNetwork[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The service tier of the instance.
     * Possible values are `TIER_UNSPECIFIED`, `STANDARD`, `PREMIUM`, `BASIC_HDD`, `BASIC_SSD`, and `HIGH_SCALE_SSD`.
     */
    public readonly tier!: pulumi.Output<string>;
    /**
     * The name of the Filestore zone of the instance.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["fileShares"] = state ? state.fileShares : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networks"] = state ? state.networks : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["tier"] = state ? state.tier : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.fileShares === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fileShares'");
            }
            if ((!args || args.networks === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networks'");
            }
            if ((!args || args.tier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tier'");
            }
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["fileShares"] = args ? args.fileShares : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networks"] = args ? args.networks : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["tier"] = args ? args.tier : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Instance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    createTime?: pulumi.Input<string>;
    /**
     * A description of the instance.
     */
    description?: pulumi.Input<string>;
    /**
     * Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
     */
    etag?: pulumi.Input<string>;
    /**
     * File system shares on the instance. For this version, only a
     * single file share is supported.
     * Structure is documented below.
     */
    fileShares?: pulumi.Input<inputs.filestore.InstanceFileShares>;
    /**
     * Resource labels to represent user-provided metadata.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the fileshare (16 characters or less)
     */
    name?: pulumi.Input<string>;
    /**
     * VPC networks to which the instance is connected. For this version,
     * only a single network is supported.
     * Structure is documented below.
     */
    networks?: pulumi.Input<pulumi.Input<inputs.filestore.InstanceNetwork>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The service tier of the instance.
     * Possible values are `TIER_UNSPECIFIED`, `STANDARD`, `PREMIUM`, `BASIC_HDD`, `BASIC_SSD`, and `HIGH_SCALE_SSD`.
     */
    tier?: pulumi.Input<string>;
    /**
     * The name of the Filestore zone of the instance.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * A description of the instance.
     */
    description?: pulumi.Input<string>;
    /**
     * File system shares on the instance. For this version, only a
     * single file share is supported.
     * Structure is documented below.
     */
    fileShares: pulumi.Input<inputs.filestore.InstanceFileShares>;
    /**
     * Resource labels to represent user-provided metadata.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the fileshare (16 characters or less)
     */
    name?: pulumi.Input<string>;
    /**
     * VPC networks to which the instance is connected. For this version,
     * only a single network is supported.
     * Structure is documented below.
     */
    networks: pulumi.Input<pulumi.Input<inputs.filestore.InstanceNetwork>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The service tier of the instance.
     * Possible values are `TIER_UNSPECIFIED`, `STANDARD`, `PREMIUM`, `BASIC_HDD`, `BASIC_SSD`, and `HIGH_SCALE_SSD`.
     */
    tier: pulumi.Input<string>;
    /**
     * The name of the Filestore zone of the instance.
     */
    zone: pulumi.Input<string>;
}
