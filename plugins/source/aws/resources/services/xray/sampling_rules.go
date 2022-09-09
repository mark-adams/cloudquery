package xray

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SamplingRules() *schema.Table {
	return &schema.Table{
		Name:        "aws_xray_sampling_rules",
		Description: "A SamplingRule (https://docsawsamazoncom/xray/latest/api/API_SamplingRulehtml) and its metadata",
		Resolver:    fetchXraySamplingRules,
		Multiplex:   client.ServiceAccountRegionMultiplexer("xray"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveXraySamplingRuleTags,
			},
			{
				Name:        "created_at",
				Description: "When the rule was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "modified_at",
				Description: "When the rule was last modified",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:     "sampling_rule",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SamplingRule"),
			},
			{
				Name:     "sampling_rule",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SamplingRule"),
			},
			{
				Name:     "sampling_rule",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SamplingRule"),
			},
			{
				Name:     "sampling_rule",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SamplingRule"),
			},
			{
				Name:     "sampling_rule",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SamplingRule"),
			},
			{
				Name:     "sampling_rule",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SamplingRule"),
			},
			{
				Name:     "sampling_rule",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SamplingRule"),
			},
			{
				Name:     "sampling_rule",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SamplingRule"),
			},
			{
				Name:     "sampling_rule",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SamplingRule"),
			},
			{
				Name:        "version",
				Description: "The version of the sampling rule format (1)",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("SamplingRule.Version"),
			},
			{
				Name:     "sampling_rule",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SamplingRule"),
			},
			{
				Name:     "sampling_rule",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SamplingRule"),
			},
			{
				Name:     "sampling_rule",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SamplingRule"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchXraySamplingRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	paginator := xray.NewGetSamplingRulesPaginator(meta.(*client.Client).Services().Xray, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- v.SamplingRuleRecords
	}
	return nil
}
func resolveXraySamplingRuleTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	sr := resource.Item.(types.SamplingRuleRecord)
	cl := meta.(*client.Client)
	svc := cl.Services().Xray
	params := xray.ListTagsForResourceInput{ResourceARN: sr.SamplingRule.RuleARN}

	output, err := svc.ListTagsForResource(ctx, &params)
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}

	tags := map[string]string{}
	client.TagsIntoMap(output.Tags, tags)

	return resource.Set(c.Name, tags)
}
