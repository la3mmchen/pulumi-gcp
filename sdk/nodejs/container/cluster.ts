// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Creates a Google Kubernetes Engine (GKE) cluster. For more information see
 * [the official documentation](https://cloud.google.com/container-engine/docs/clusters)
 * and
 * [API](https://cloud.google.com/container-engine/reference/rest/v1/projects.zones.clusters).
 * 
 * ~> **Note:** All arguments including the username and password will be stored in the raw state as plain-text.
 * [Read more about sensitive data in state](/docs/state/sensitive-data.html).
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState): Cluster {
        return new Cluster(name, <any>state, { id });
    }

    /**
     * The list of additional Google Compute Engine
     * locations in which the cluster's nodes should be located. If additional zones are
     * configured, the number of nodes specified in `initial_node_count` is created in
     * all specified zones.
     */
    public readonly additionalZones: pulumi.Output<string[]>;
    /**
     * The configuration for addons supported by GKE.
     * Structure is documented below.
     */
    public readonly addonsConfig: pulumi.Output<{ horizontalPodAutoscaling: { disabled?: boolean }, httpLoadBalancing: { disabled?: boolean }, kubernetesDashboard: { disabled?: boolean }, networkPolicyConfig: { disabled?: boolean } }>;
    /**
     * The IP address range of the kubernetes pods in
     * this cluster. Default is an automatically assigned CIDR.
     */
    public readonly clusterIpv4Cidr: pulumi.Output<string>;
    /**
     * Description of the cluster.
     */
    public readonly description: pulumi.Output<string | undefined>;
    /**
     * Whether to enable Kubernetes Alpha features for
     * this cluster. Note that when this option is enabled, the cluster cannot be upgraded
     * and will be automatically deleted after 30 days.
     */
    public readonly enableKubernetesAlpha: pulumi.Output<boolean | undefined>;
    /**
     * Whether the ABAC authorizer is enabled for this cluster.
     * When enabled, identities in the system, including service accounts, nodes, and controllers,
     * will have statically granted permissions beyond those provided by the RBAC configuration or IAM.
     * Defaults to `false`
     */
    public readonly enableLegacyAbac: pulumi.Output<boolean | undefined>;
    /**
     * The IP address of this cluster's Kubernetes master.
     */
    public /*out*/ readonly endpoint: pulumi.Output<string>;
    /**
     * The number of nodes to create in this
     * cluster (not including the Kubernetes master). Must be set if `node_pool` is not set.
     */
    public readonly initialNodeCount: pulumi.Output<number | undefined>;
    /**
     * List of instance group URLs which have been assigned
     * to the cluster.
     */
    public /*out*/ readonly instanceGroupUrls: pulumi.Output<string[]>;
    /**
     * Configuration for cluster IP allocation. As of now, only pre-allocated subnetworks (custom type with secondary ranges) are supported.
     * This will activate IP aliases. See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/ip-aliases)
     * Structure is documented below.
     */
    public readonly ipAllocationPolicy: pulumi.Output<{ clusterSecondaryRangeName?: string, servicesSecondaryRangeName?: string } | undefined>;
    /**
     * The logging service that the cluster should
     * write logs to. Available options include `logging.googleapis.com` and
     * `none`. Defaults to `logging.googleapis.com`
     */
    public readonly loggingService: pulumi.Output<string>;
    /**
     * The maintenance policy to use for the cluster. Structure is
     * documented below.
     */
    public readonly maintenancePolicy: pulumi.Output<{ dailyMaintenanceWindow: { duration: string, startTime: string } } | undefined>;
    /**
     * The authentication information for accessing the
     * Kubernetes master. Structure is documented below.
     */
    public readonly masterAuth: pulumi.Output<{ clientCertificate: string, clientCertificateConfig?: { issueClientCertificate: boolean }, clientKey: string, clusterCaCertificate: string, password: string, username: string }>;
    /**
     * The desired configuration options
     * for master authorized networks. Omit the nested `cidr_blocks` attribute to disallow
     * external access (except the cluster node IPs, which GKE automatically whitelists).
     */
    public readonly masterAuthorizedNetworksConfig: pulumi.Output<{ cidrBlocks: { cidrBlock: string, displayName?: string }[] } | undefined>;
    /**
     * ) Specifies a private
     * [RFC1918](https://tools.ietf.org/html/rfc1918) block for the master's VPC. The master range must not overlap with any subnet in your cluster's VPC.
     * The master and your cluster use VPC peering. Must be specified in CIDR notation and must be `/28` subnet.
     */
    public readonly masterIpv4CidrBlock: pulumi.Output<string | undefined>;
    /**
     * The current version of the master in the cluster. This may
     * be different than the `min_master_version` set in the config if the master
     * has been updated by GKE.
     */
    public /*out*/ readonly masterVersion: pulumi.Output<string>;
    /**
     * The minimum version of the master. GKE
     * will auto-update the master to new versions, so this does not guarantee the
     * current master version--use the read-only `master_version` field to obtain that.
     * If unset, the cluster's version will be set by GKE to the version of the most recent
     * official release (which is not necessarily the latest version).
     */
    public readonly minMasterVersion: pulumi.Output<string | undefined>;
    /**
     * The monitoring service that the cluster
     * should write metrics to.
     * Automatically send metrics from pods in the cluster to the Google Cloud Monitoring API.
     * VM metrics will be collected by Google Compute Engine regardless of this setting
     * Available options include
     * `monitoring.googleapis.com` and `none`. Defaults to
     * `monitoring.googleapis.com`
     */
    public readonly monitoringService: pulumi.Output<string>;
    /**
     * The name of the cluster, unique within the project and
     * zone.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The name or self_link of the Google Compute Engine
     * network to which the cluster is connected. For Shared VPC, set this to the self link of the
     * shared network.
     */
    public readonly network: pulumi.Output<string | undefined>;
    /**
     * Configuration options for the
     * [NetworkPolicy](https://kubernetes.io/docs/concepts/services-networking/networkpolicies/)
     * feature. Structure is documented below.
     */
    public readonly networkPolicy: pulumi.Output<{ enabled?: boolean, provider?: string }>;
    /**
     * Parameters used in creating the cluster's nodes.
     * Structure is documented below.
     */
    public readonly nodeConfig: pulumi.Output<{ diskSizeGb: number, guestAccelerators: { count: number, type: string }[], imageType: string, labels?: {[key: string]: string}, localSsdCount: number, machineType: string, metadata?: {[key: string]: string}, minCpuPlatform?: string, oauthScopes: string[], preemptible?: boolean, serviceAccount: string, tags?: string[], taints?: { effect: string, key: string, value: string }[], workloadMetadataConfig?: { nodeMetadata: string } }>;
    /**
     * List of node pools associated with this cluster.
     * See [google_container_node_pool](container_node_pool.html) for schema.
     */
    public readonly nodePools: pulumi.Output<{ autoscaling?: { maxNodeCount: number, minNodeCount: number }, initialNodeCount: number, instanceGroupUrls: string[], management: { autoRepair?: boolean, autoUpgrade?: boolean }, name: string, namePrefix: string, nodeConfig: { diskSizeGb: number, guestAccelerators: { count: number, type: string }[], imageType: string, labels?: {[key: string]: string}, localSsdCount: number, machineType: string, metadata?: {[key: string]: string}, minCpuPlatform?: string, oauthScopes: string[], preemptible?: boolean, serviceAccount: string, tags?: string[], taints?: { effect: string, key: string, value: string }[], workloadMetadataConfig?: { nodeMetadata: string } }, nodeCount: number, version: string }[]>;
    /**
     * The Kubernetes version on the nodes. Must either be unset
     * or set to the same value as `min_master_version` on create. Defaults to the default
     * version set by GKE which is not necessarily the latest version.
     */
    public readonly nodeVersion: pulumi.Output<string>;
    /**
     * ) Configuration for the
     * [PodSecurityPolicy](https://cloud.google.com/kubernetes-engine/docs/how-to/pod-security-policies) feature.
     * Structure is documented below.
     */
    public readonly podSecurityPolicyConfig: pulumi.Output<{ enabled: boolean } | undefined>;
    /**
     * ) If true, a
     * [private cluster](https://cloud.google.com/kubernetes-engine/docs/how-to/private-clusters) will be created, which makes
     * the master inaccessible from the public internet and nodes do not get public IP addresses either. It is mandatory to specify
     * `master_ipv4_cidr_block` and `ip_allocation_policy` with this option.
     */
    public readonly privateCluster: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    public readonly region: pulumi.Output<string>;
    /**
     * If true, deletes the default node pool upon cluster creation.
     */
    public readonly removeDefaultNodePool: pulumi.Output<boolean | undefined>;
    /**
     * The name or self_link of the Google Compute Engine subnetwork in
     * which the cluster's instances are launched.
     */
    public readonly subnetwork: pulumi.Output<string>;
    /**
     * The zone that the master and the number of nodes specified
     * in `initial_node_count` should be created in. Only one of `zone` and `region`
     * may be set. If neither zone nor region are set, the provider zone is used.
     */
    public readonly zone: pulumi.Output<string>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ClusterArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ClusterState = argsOrState as ClusterState | undefined;
            inputs["additionalZones"] = state ? state.additionalZones : undefined;
            inputs["addonsConfig"] = state ? state.addonsConfig : undefined;
            inputs["clusterIpv4Cidr"] = state ? state.clusterIpv4Cidr : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["enableKubernetesAlpha"] = state ? state.enableKubernetesAlpha : undefined;
            inputs["enableLegacyAbac"] = state ? state.enableLegacyAbac : undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["initialNodeCount"] = state ? state.initialNodeCount : undefined;
            inputs["instanceGroupUrls"] = state ? state.instanceGroupUrls : undefined;
            inputs["ipAllocationPolicy"] = state ? state.ipAllocationPolicy : undefined;
            inputs["loggingService"] = state ? state.loggingService : undefined;
            inputs["maintenancePolicy"] = state ? state.maintenancePolicy : undefined;
            inputs["masterAuth"] = state ? state.masterAuth : undefined;
            inputs["masterAuthorizedNetworksConfig"] = state ? state.masterAuthorizedNetworksConfig : undefined;
            inputs["masterIpv4CidrBlock"] = state ? state.masterIpv4CidrBlock : undefined;
            inputs["masterVersion"] = state ? state.masterVersion : undefined;
            inputs["minMasterVersion"] = state ? state.minMasterVersion : undefined;
            inputs["monitoringService"] = state ? state.monitoringService : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["networkPolicy"] = state ? state.networkPolicy : undefined;
            inputs["nodeConfig"] = state ? state.nodeConfig : undefined;
            inputs["nodePools"] = state ? state.nodePools : undefined;
            inputs["nodeVersion"] = state ? state.nodeVersion : undefined;
            inputs["podSecurityPolicyConfig"] = state ? state.podSecurityPolicyConfig : undefined;
            inputs["privateCluster"] = state ? state.privateCluster : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["removeDefaultNodePool"] = state ? state.removeDefaultNodePool : undefined;
            inputs["subnetwork"] = state ? state.subnetwork : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            inputs["additionalZones"] = args ? args.additionalZones : undefined;
            inputs["addonsConfig"] = args ? args.addonsConfig : undefined;
            inputs["clusterIpv4Cidr"] = args ? args.clusterIpv4Cidr : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["enableKubernetesAlpha"] = args ? args.enableKubernetesAlpha : undefined;
            inputs["enableLegacyAbac"] = args ? args.enableLegacyAbac : undefined;
            inputs["initialNodeCount"] = args ? args.initialNodeCount : undefined;
            inputs["ipAllocationPolicy"] = args ? args.ipAllocationPolicy : undefined;
            inputs["loggingService"] = args ? args.loggingService : undefined;
            inputs["maintenancePolicy"] = args ? args.maintenancePolicy : undefined;
            inputs["masterAuth"] = args ? args.masterAuth : undefined;
            inputs["masterAuthorizedNetworksConfig"] = args ? args.masterAuthorizedNetworksConfig : undefined;
            inputs["masterIpv4CidrBlock"] = args ? args.masterIpv4CidrBlock : undefined;
            inputs["minMasterVersion"] = args ? args.minMasterVersion : undefined;
            inputs["monitoringService"] = args ? args.monitoringService : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["networkPolicy"] = args ? args.networkPolicy : undefined;
            inputs["nodeConfig"] = args ? args.nodeConfig : undefined;
            inputs["nodePools"] = args ? args.nodePools : undefined;
            inputs["nodeVersion"] = args ? args.nodeVersion : undefined;
            inputs["podSecurityPolicyConfig"] = args ? args.podSecurityPolicyConfig : undefined;
            inputs["privateCluster"] = args ? args.privateCluster : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["removeDefaultNodePool"] = args ? args.removeDefaultNodePool : undefined;
            inputs["subnetwork"] = args ? args.subnetwork : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["endpoint"] = undefined /*out*/;
            inputs["instanceGroupUrls"] = undefined /*out*/;
            inputs["masterVersion"] = undefined /*out*/;
        }
        super("gcp:container/cluster:Cluster", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * The list of additional Google Compute Engine
     * locations in which the cluster's nodes should be located. If additional zones are
     * configured, the number of nodes specified in `initial_node_count` is created in
     * all specified zones.
     */
    readonly additionalZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The configuration for addons supported by GKE.
     * Structure is documented below.
     */
    readonly addonsConfig?: pulumi.Input<{ horizontalPodAutoscaling?: pulumi.Input<{ disabled?: pulumi.Input<boolean> }>, httpLoadBalancing?: pulumi.Input<{ disabled?: pulumi.Input<boolean> }>, kubernetesDashboard?: pulumi.Input<{ disabled?: pulumi.Input<boolean> }>, networkPolicyConfig?: pulumi.Input<{ disabled?: pulumi.Input<boolean> }> }>;
    /**
     * The IP address range of the kubernetes pods in
     * this cluster. Default is an automatically assigned CIDR.
     */
    readonly clusterIpv4Cidr?: pulumi.Input<string>;
    /**
     * Description of the cluster.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Whether to enable Kubernetes Alpha features for
     * this cluster. Note that when this option is enabled, the cluster cannot be upgraded
     * and will be automatically deleted after 30 days.
     */
    readonly enableKubernetesAlpha?: pulumi.Input<boolean>;
    /**
     * Whether the ABAC authorizer is enabled for this cluster.
     * When enabled, identities in the system, including service accounts, nodes, and controllers,
     * will have statically granted permissions beyond those provided by the RBAC configuration or IAM.
     * Defaults to `false`
     */
    readonly enableLegacyAbac?: pulumi.Input<boolean>;
    /**
     * The IP address of this cluster's Kubernetes master.
     */
    readonly endpoint?: pulumi.Input<string>;
    /**
     * The number of nodes to create in this
     * cluster (not including the Kubernetes master). Must be set if `node_pool` is not set.
     */
    readonly initialNodeCount?: pulumi.Input<number>;
    /**
     * List of instance group URLs which have been assigned
     * to the cluster.
     */
    readonly instanceGroupUrls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Configuration for cluster IP allocation. As of now, only pre-allocated subnetworks (custom type with secondary ranges) are supported.
     * This will activate IP aliases. See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/ip-aliases)
     * Structure is documented below.
     */
    readonly ipAllocationPolicy?: pulumi.Input<{ clusterSecondaryRangeName?: pulumi.Input<string>, servicesSecondaryRangeName?: pulumi.Input<string> }>;
    /**
     * The logging service that the cluster should
     * write logs to. Available options include `logging.googleapis.com` and
     * `none`. Defaults to `logging.googleapis.com`
     */
    readonly loggingService?: pulumi.Input<string>;
    /**
     * The maintenance policy to use for the cluster. Structure is
     * documented below.
     */
    readonly maintenancePolicy?: pulumi.Input<{ dailyMaintenanceWindow: pulumi.Input<{ duration?: pulumi.Input<string>, startTime: pulumi.Input<string> }> }>;
    /**
     * The authentication information for accessing the
     * Kubernetes master. Structure is documented below.
     */
    readonly masterAuth?: pulumi.Input<{ clientCertificate?: pulumi.Input<string>, clientCertificateConfig?: pulumi.Input<{ issueClientCertificate: pulumi.Input<boolean> }>, clientKey?: pulumi.Input<string>, clusterCaCertificate?: pulumi.Input<string>, password: pulumi.Input<string>, username: pulumi.Input<string> }>;
    /**
     * The desired configuration options
     * for master authorized networks. Omit the nested `cidr_blocks` attribute to disallow
     * external access (except the cluster node IPs, which GKE automatically whitelists).
     */
    readonly masterAuthorizedNetworksConfig?: pulumi.Input<{ cidrBlocks?: pulumi.Input<{ cidrBlock: pulumi.Input<string>, displayName?: pulumi.Input<string> }[]> }>;
    /**
     * ) Specifies a private
     * [RFC1918](https://tools.ietf.org/html/rfc1918) block for the master's VPC. The master range must not overlap with any subnet in your cluster's VPC.
     * The master and your cluster use VPC peering. Must be specified in CIDR notation and must be `/28` subnet.
     */
    readonly masterIpv4CidrBlock?: pulumi.Input<string>;
    /**
     * The current version of the master in the cluster. This may
     * be different than the `min_master_version` set in the config if the master
     * has been updated by GKE.
     */
    readonly masterVersion?: pulumi.Input<string>;
    /**
     * The minimum version of the master. GKE
     * will auto-update the master to new versions, so this does not guarantee the
     * current master version--use the read-only `master_version` field to obtain that.
     * If unset, the cluster's version will be set by GKE to the version of the most recent
     * official release (which is not necessarily the latest version).
     */
    readonly minMasterVersion?: pulumi.Input<string>;
    /**
     * The monitoring service that the cluster
     * should write metrics to.
     * Automatically send metrics from pods in the cluster to the Google Cloud Monitoring API.
     * VM metrics will be collected by Google Compute Engine regardless of this setting
     * Available options include
     * `monitoring.googleapis.com` and `none`. Defaults to
     * `monitoring.googleapis.com`
     */
    readonly monitoringService?: pulumi.Input<string>;
    /**
     * The name of the cluster, unique within the project and
     * zone.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name or self_link of the Google Compute Engine
     * network to which the cluster is connected. For Shared VPC, set this to the self link of the
     * shared network.
     */
    readonly network?: pulumi.Input<string>;
    /**
     * Configuration options for the
     * [NetworkPolicy](https://kubernetes.io/docs/concepts/services-networking/networkpolicies/)
     * feature. Structure is documented below.
     */
    readonly networkPolicy?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, provider?: pulumi.Input<string> }>;
    /**
     * Parameters used in creating the cluster's nodes.
     * Structure is documented below.
     */
    readonly nodeConfig?: pulumi.Input<{ diskSizeGb?: pulumi.Input<number>, guestAccelerators?: pulumi.Input<{ count: pulumi.Input<number>, type: pulumi.Input<string> }[]>, imageType?: pulumi.Input<string>, labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, localSsdCount?: pulumi.Input<number>, machineType?: pulumi.Input<string>, metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, minCpuPlatform?: pulumi.Input<string>, oauthScopes?: pulumi.Input<pulumi.Input<string>[]>, preemptible?: pulumi.Input<boolean>, serviceAccount?: pulumi.Input<string>, tags?: pulumi.Input<pulumi.Input<string>[]>, taints?: pulumi.Input<{ effect: pulumi.Input<string>, key: pulumi.Input<string>, value: pulumi.Input<string> }[]>, workloadMetadataConfig?: pulumi.Input<{ nodeMetadata: pulumi.Input<string> }> }>;
    /**
     * List of node pools associated with this cluster.
     * See [google_container_node_pool](container_node_pool.html) for schema.
     */
    readonly nodePools?: pulumi.Input<{ autoscaling?: pulumi.Input<{ maxNodeCount: pulumi.Input<number>, minNodeCount: pulumi.Input<number> }>, initialNodeCount?: pulumi.Input<number>, instanceGroupUrls?: pulumi.Input<pulumi.Input<string>[]>, management?: pulumi.Input<{ autoRepair?: pulumi.Input<boolean>, autoUpgrade?: pulumi.Input<boolean> }>, name?: pulumi.Input<string>, namePrefix?: pulumi.Input<string>, nodeConfig?: pulumi.Input<{ diskSizeGb?: pulumi.Input<number>, guestAccelerators?: pulumi.Input<{ count: pulumi.Input<number>, type: pulumi.Input<string> }[]>, imageType?: pulumi.Input<string>, labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, localSsdCount?: pulumi.Input<number>, machineType?: pulumi.Input<string>, metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, minCpuPlatform?: pulumi.Input<string>, oauthScopes?: pulumi.Input<pulumi.Input<string>[]>, preemptible?: pulumi.Input<boolean>, serviceAccount?: pulumi.Input<string>, tags?: pulumi.Input<pulumi.Input<string>[]>, taints?: pulumi.Input<{ effect: pulumi.Input<string>, key: pulumi.Input<string>, value: pulumi.Input<string> }[]>, workloadMetadataConfig?: pulumi.Input<{ nodeMetadata: pulumi.Input<string> }> }>, nodeCount?: pulumi.Input<number>, version?: pulumi.Input<string> }[]>;
    /**
     * The Kubernetes version on the nodes. Must either be unset
     * or set to the same value as `min_master_version` on create. Defaults to the default
     * version set by GKE which is not necessarily the latest version.
     */
    readonly nodeVersion?: pulumi.Input<string>;
    /**
     * ) Configuration for the
     * [PodSecurityPolicy](https://cloud.google.com/kubernetes-engine/docs/how-to/pod-security-policies) feature.
     * Structure is documented below.
     */
    readonly podSecurityPolicyConfig?: pulumi.Input<{ enabled: pulumi.Input<boolean> }>;
    /**
     * ) If true, a
     * [private cluster](https://cloud.google.com/kubernetes-engine/docs/how-to/private-clusters) will be created, which makes
     * the master inaccessible from the public internet and nodes do not get public IP addresses either. It is mandatory to specify
     * `master_ipv4_cidr_block` and `ip_allocation_policy` with this option.
     */
    readonly privateCluster?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    /**
     * If true, deletes the default node pool upon cluster creation.
     */
    readonly removeDefaultNodePool?: pulumi.Input<boolean>;
    /**
     * The name or self_link of the Google Compute Engine subnetwork in
     * which the cluster's instances are launched.
     */
    readonly subnetwork?: pulumi.Input<string>;
    /**
     * The zone that the master and the number of nodes specified
     * in `initial_node_count` should be created in. Only one of `zone` and `region`
     * may be set. If neither zone nor region are set, the provider zone is used.
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * The list of additional Google Compute Engine
     * locations in which the cluster's nodes should be located. If additional zones are
     * configured, the number of nodes specified in `initial_node_count` is created in
     * all specified zones.
     */
    readonly additionalZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The configuration for addons supported by GKE.
     * Structure is documented below.
     */
    readonly addonsConfig?: pulumi.Input<{ horizontalPodAutoscaling?: pulumi.Input<{ disabled?: pulumi.Input<boolean> }>, httpLoadBalancing?: pulumi.Input<{ disabled?: pulumi.Input<boolean> }>, kubernetesDashboard?: pulumi.Input<{ disabled?: pulumi.Input<boolean> }>, networkPolicyConfig?: pulumi.Input<{ disabled?: pulumi.Input<boolean> }> }>;
    /**
     * The IP address range of the kubernetes pods in
     * this cluster. Default is an automatically assigned CIDR.
     */
    readonly clusterIpv4Cidr?: pulumi.Input<string>;
    /**
     * Description of the cluster.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Whether to enable Kubernetes Alpha features for
     * this cluster. Note that when this option is enabled, the cluster cannot be upgraded
     * and will be automatically deleted after 30 days.
     */
    readonly enableKubernetesAlpha?: pulumi.Input<boolean>;
    /**
     * Whether the ABAC authorizer is enabled for this cluster.
     * When enabled, identities in the system, including service accounts, nodes, and controllers,
     * will have statically granted permissions beyond those provided by the RBAC configuration or IAM.
     * Defaults to `false`
     */
    readonly enableLegacyAbac?: pulumi.Input<boolean>;
    /**
     * The number of nodes to create in this
     * cluster (not including the Kubernetes master). Must be set if `node_pool` is not set.
     */
    readonly initialNodeCount?: pulumi.Input<number>;
    /**
     * Configuration for cluster IP allocation. As of now, only pre-allocated subnetworks (custom type with secondary ranges) are supported.
     * This will activate IP aliases. See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/ip-aliases)
     * Structure is documented below.
     */
    readonly ipAllocationPolicy?: pulumi.Input<{ clusterSecondaryRangeName?: pulumi.Input<string>, servicesSecondaryRangeName?: pulumi.Input<string> }>;
    /**
     * The logging service that the cluster should
     * write logs to. Available options include `logging.googleapis.com` and
     * `none`. Defaults to `logging.googleapis.com`
     */
    readonly loggingService?: pulumi.Input<string>;
    /**
     * The maintenance policy to use for the cluster. Structure is
     * documented below.
     */
    readonly maintenancePolicy?: pulumi.Input<{ dailyMaintenanceWindow: pulumi.Input<{ duration?: pulumi.Input<string>, startTime: pulumi.Input<string> }> }>;
    /**
     * The authentication information for accessing the
     * Kubernetes master. Structure is documented below.
     */
    readonly masterAuth?: pulumi.Input<{ clientCertificate?: pulumi.Input<string>, clientCertificateConfig?: pulumi.Input<{ issueClientCertificate: pulumi.Input<boolean> }>, clientKey?: pulumi.Input<string>, clusterCaCertificate?: pulumi.Input<string>, password: pulumi.Input<string>, username: pulumi.Input<string> }>;
    /**
     * The desired configuration options
     * for master authorized networks. Omit the nested `cidr_blocks` attribute to disallow
     * external access (except the cluster node IPs, which GKE automatically whitelists).
     */
    readonly masterAuthorizedNetworksConfig?: pulumi.Input<{ cidrBlocks?: pulumi.Input<{ cidrBlock: pulumi.Input<string>, displayName?: pulumi.Input<string> }[]> }>;
    /**
     * ) Specifies a private
     * [RFC1918](https://tools.ietf.org/html/rfc1918) block for the master's VPC. The master range must not overlap with any subnet in your cluster's VPC.
     * The master and your cluster use VPC peering. Must be specified in CIDR notation and must be `/28` subnet.
     */
    readonly masterIpv4CidrBlock?: pulumi.Input<string>;
    /**
     * The minimum version of the master. GKE
     * will auto-update the master to new versions, so this does not guarantee the
     * current master version--use the read-only `master_version` field to obtain that.
     * If unset, the cluster's version will be set by GKE to the version of the most recent
     * official release (which is not necessarily the latest version).
     */
    readonly minMasterVersion?: pulumi.Input<string>;
    /**
     * The monitoring service that the cluster
     * should write metrics to.
     * Automatically send metrics from pods in the cluster to the Google Cloud Monitoring API.
     * VM metrics will be collected by Google Compute Engine regardless of this setting
     * Available options include
     * `monitoring.googleapis.com` and `none`. Defaults to
     * `monitoring.googleapis.com`
     */
    readonly monitoringService?: pulumi.Input<string>;
    /**
     * The name of the cluster, unique within the project and
     * zone.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name or self_link of the Google Compute Engine
     * network to which the cluster is connected. For Shared VPC, set this to the self link of the
     * shared network.
     */
    readonly network?: pulumi.Input<string>;
    /**
     * Configuration options for the
     * [NetworkPolicy](https://kubernetes.io/docs/concepts/services-networking/networkpolicies/)
     * feature. Structure is documented below.
     */
    readonly networkPolicy?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, provider?: pulumi.Input<string> }>;
    /**
     * Parameters used in creating the cluster's nodes.
     * Structure is documented below.
     */
    readonly nodeConfig?: pulumi.Input<{ diskSizeGb?: pulumi.Input<number>, guestAccelerators?: pulumi.Input<{ count: pulumi.Input<number>, type: pulumi.Input<string> }[]>, imageType?: pulumi.Input<string>, labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, localSsdCount?: pulumi.Input<number>, machineType?: pulumi.Input<string>, metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, minCpuPlatform?: pulumi.Input<string>, oauthScopes?: pulumi.Input<pulumi.Input<string>[]>, preemptible?: pulumi.Input<boolean>, serviceAccount?: pulumi.Input<string>, tags?: pulumi.Input<pulumi.Input<string>[]>, taints?: pulumi.Input<{ effect: pulumi.Input<string>, key: pulumi.Input<string>, value: pulumi.Input<string> }[]>, workloadMetadataConfig?: pulumi.Input<{ nodeMetadata: pulumi.Input<string> }> }>;
    /**
     * List of node pools associated with this cluster.
     * See [google_container_node_pool](container_node_pool.html) for schema.
     */
    readonly nodePools?: pulumi.Input<{ autoscaling?: pulumi.Input<{ maxNodeCount: pulumi.Input<number>, minNodeCount: pulumi.Input<number> }>, initialNodeCount?: pulumi.Input<number>, instanceGroupUrls?: pulumi.Input<pulumi.Input<string>[]>, management?: pulumi.Input<{ autoRepair?: pulumi.Input<boolean>, autoUpgrade?: pulumi.Input<boolean> }>, name?: pulumi.Input<string>, namePrefix?: pulumi.Input<string>, nodeConfig?: pulumi.Input<{ diskSizeGb?: pulumi.Input<number>, guestAccelerators?: pulumi.Input<{ count: pulumi.Input<number>, type: pulumi.Input<string> }[]>, imageType?: pulumi.Input<string>, labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, localSsdCount?: pulumi.Input<number>, machineType?: pulumi.Input<string>, metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, minCpuPlatform?: pulumi.Input<string>, oauthScopes?: pulumi.Input<pulumi.Input<string>[]>, preemptible?: pulumi.Input<boolean>, serviceAccount?: pulumi.Input<string>, tags?: pulumi.Input<pulumi.Input<string>[]>, taints?: pulumi.Input<{ effect: pulumi.Input<string>, key: pulumi.Input<string>, value: pulumi.Input<string> }[]>, workloadMetadataConfig?: pulumi.Input<{ nodeMetadata: pulumi.Input<string> }> }>, nodeCount?: pulumi.Input<number>, version?: pulumi.Input<string> }[]>;
    /**
     * The Kubernetes version on the nodes. Must either be unset
     * or set to the same value as `min_master_version` on create. Defaults to the default
     * version set by GKE which is not necessarily the latest version.
     */
    readonly nodeVersion?: pulumi.Input<string>;
    /**
     * ) Configuration for the
     * [PodSecurityPolicy](https://cloud.google.com/kubernetes-engine/docs/how-to/pod-security-policies) feature.
     * Structure is documented below.
     */
    readonly podSecurityPolicyConfig?: pulumi.Input<{ enabled: pulumi.Input<boolean> }>;
    /**
     * ) If true, a
     * [private cluster](https://cloud.google.com/kubernetes-engine/docs/how-to/private-clusters) will be created, which makes
     * the master inaccessible from the public internet and nodes do not get public IP addresses either. It is mandatory to specify
     * `master_ipv4_cidr_block` and `ip_allocation_policy` with this option.
     */
    readonly privateCluster?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    /**
     * If true, deletes the default node pool upon cluster creation.
     */
    readonly removeDefaultNodePool?: pulumi.Input<boolean>;
    /**
     * The name or self_link of the Google Compute Engine subnetwork in
     * which the cluster's instances are launched.
     */
    readonly subnetwork?: pulumi.Input<string>;
    /**
     * The zone that the master and the number of nodes specified
     * in `initial_node_count` should be created in. Only one of `zone` and `region`
     * may be set. If neither zone nor region are set, the provider zone is used.
     */
    readonly zone?: pulumi.Input<string>;
}