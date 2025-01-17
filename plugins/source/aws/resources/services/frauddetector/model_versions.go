// Code generated by codegen; DO NOT EDIT.

package frauddetector

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ModelVersions() *schema.Table {
	return &schema.Table{
		Name:        "aws_frauddetector_model_versions",
		Description: `https://docs.aws.amazon.com/frauddetector/latest/api/API_ModelVersionDetail.html`,
		Resolver:    fetchFrauddetectorModelVersions,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "created_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreatedTime"),
			},
			{
				Name:     "external_events_detail",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ExternalEventsDetail"),
			},
			{
				Name:     "ingested_events_detail",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IngestedEventsDetail"),
			},
			{
				Name:     "last_updated_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastUpdatedTime"),
			},
			{
				Name:     "model_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ModelId"),
			},
			{
				Name:     "model_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ModelType"),
			},
			{
				Name:     "model_version_number",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ModelVersionNumber"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "training_data_schema",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TrainingDataSchema"),
			},
			{
				Name:     "training_data_source",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TrainingDataSource"),
			},
			{
				Name:     "training_result",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TrainingResult"),
			},
			{
				Name:     "training_result_v2",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TrainingResultV2"),
			},
		},
	}
}
