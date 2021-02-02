// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "gcp:monitoring/alertPolicy:AlertPolicy":
		r, err = NewAlertPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:monitoring/customService:CustomService":
		r, err = NewCustomService(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:monitoring/dashboard:Dashboard":
		r, err = NewDashboard(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:monitoring/group:Group":
		r, err = NewGroup(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:monitoring/metricDescriptor:MetricDescriptor":
		r, err = NewMetricDescriptor(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:monitoring/notificationChannel:NotificationChannel":
		r, err = NewNotificationChannel(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:monitoring/slo:Slo":
		r, err = NewSlo(ctx, name, nil, pulumi.URN_(urn))
	case "gcp:monitoring/uptimeCheckConfig:UptimeCheckConfig":
		r, err = NewUptimeCheckConfig(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	return
}

func init() {
	version, err := gcp.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"gcp",
		"monitoring/alertPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"monitoring/customService",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"monitoring/dashboard",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"monitoring/group",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"monitoring/metricDescriptor",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"monitoring/notificationChannel",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"monitoring/slo",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"monitoring/uptimeCheckConfig",
		&module{version},
	)
}