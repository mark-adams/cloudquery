package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

type Route53HealthCheckWrapper struct {
	types.HealthCheck
	Tags map[string]string
}

func Route53HealthChecks() *schema.Table {
	return &schema.Table{
		Name:          "aws_route53_health_checks",
		Description:   "A complex type that contains information about one health check that is associated with the current AWS account.",
		Resolver:      fetchRoute53HealthChecks,
		Multiplex:     client.AccountMultiplex,
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "cloud_watch_alarm_configuration_dimensions",
				Description: "the metric that the CloudWatch alarm is associated with, a complex type that contains information about the dimensions for the metric.",
				Type:        schema.TypeJSON,
				Resolver:    resolveRoute53healthCheckCloudWatchAlarmConfigurationDimensions,
			},
			{
				Name:        "tags",
				Description: "The tags associated with the health check.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "caller_reference",
				Description: "A unique string that you specified when you created the health check.",
				Type:        schema.TypeString,
			},
			{
				Name:     "health_check_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HealthCheckConfig"),
			},
			{
				Name:     "health_check_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HealthCheckConfig"),
			},
			{
				Name:     "health_check_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HealthCheckConfig"),
			},
			{
				Name:        "child_health_checks",
				Description: "(CALCULATED Health Checks Only) A complex type that contains one ChildHealthCheck element for each health check that you want to associate with a CALCULATED health check.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("HealthCheckConfig.ChildHealthChecks"),
			},
			{
				Name:     "health_check_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HealthCheckConfig"),
			},
			{
				Name:     "health_check_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HealthCheckConfig"),
			},
			{
				Name:     "health_check_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HealthCheckConfig"),
			},
			{
				Name:     "health_check_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HealthCheckConfig"),
			},
			{
				Name:     "health_check_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HealthCheckConfig"),
			},
			{
				Name:     "health_check_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HealthCheckConfig"),
			},
			{
				Name:     "health_check_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HealthCheckConfig"),
			},
			{
				Name:     "health_check_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HealthCheckConfig"),
			},
			{
				Name:     "health_check_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HealthCheckConfig"),
			},
			{
				Name:     "health_check_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HealthCheckConfig"),
			},
			{
				Name:     "health_check_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HealthCheckConfig"),
			},
			{
				Name:     "health_check_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HealthCheckConfig"),
			},
			{
				Name:     "health_check_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HealthCheckConfig"),
			},
			{
				Name:     "health_check_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HealthCheckConfig"),
			},
			{
				Name:        "health_check_version",
				Description: "The version of the health check.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "id",
				Description: "The identifier that Amazon Route 53 assigned to the health check when you created it.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Id"),
			},
			{
				Name:     "cloud_watch_alarm_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CloudWatchAlarmConfiguration"),
			},
			{
				Name:     "cloud_watch_alarm_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CloudWatchAlarmConfiguration"),
			},
			{
				Name:     "cloud_watch_alarm_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CloudWatchAlarmConfiguration"),
			},
			{
				Name:     "cloud_watch_alarm_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CloudWatchAlarmConfiguration"),
			},
			{
				Name:     "cloud_watch_alarm_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CloudWatchAlarmConfiguration"),
			},
			{
				Name:     "cloud_watch_alarm_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CloudWatchAlarmConfiguration"),
			},
			{
				Name:     "cloud_watch_alarm_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CloudWatchAlarmConfiguration"),
			},
			{
				Name:     "linked_service",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LinkedService"),
			},
			{
				Name:     "linked_service",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LinkedService"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARNGlobal(client.Route53Service, func(resource *schema.Resource) ([]string, error) {
					return []string{"healthcheck", *resource.Item.(Route53HealthCheckWrapper).Id}, nil
				}),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchRoute53HealthChecks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config route53.ListHealthChecksInput
	c := meta.(*client.Client)
	svc := c.Services().Route53

	processHealthChecksBundle := func(healthChecks []types.HealthCheck) error {
		tagsCfg := &route53.ListTagsForResourcesInput{ResourceType: types.TagResourceTypeHealthcheck, ResourceIds: make([]string, 0, len(healthChecks))}
		for _, h := range healthChecks {
			tagsCfg.ResourceIds = append(tagsCfg.ResourceIds, *h.Id)
		}
		tagsResponse, err := svc.ListTagsForResources(ctx, tagsCfg)
		if err != nil {
			return err
		}
		for _, h := range healthChecks {
			wrapper := Route53HealthCheckWrapper{
				HealthCheck: h,
				Tags:        client.TagsToMap(getRoute53tagsByResourceID(*h.Id, tagsResponse.ResourceTagSets)),
			}
			res <- wrapper
		}
		return nil
	}

	for {
		response, err := svc.ListHealthChecks(ctx, &config)
		if err != nil {
			return err
		}

		for i := 0; i < len(response.HealthChecks); i += 10 {
			end := i + 10

			if end > len(response.HealthChecks) {
				end = len(response.HealthChecks)
			}
			zones := response.HealthChecks[i:end]
			err := processHealthChecksBundle(zones)
			if err != nil {
				return err
			}
		}

		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
func resolveRoute53healthCheckCloudWatchAlarmConfigurationDimensions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(Route53HealthCheckWrapper)

	if r.CloudWatchAlarmConfiguration == nil {
		return nil
	}
	tags := map[string]*string{}
	for _, t := range r.CloudWatchAlarmConfiguration.Dimensions {
		tags[*t.Name] = t.Value
	}
	return resource.Set(c.Name, tags)
}
