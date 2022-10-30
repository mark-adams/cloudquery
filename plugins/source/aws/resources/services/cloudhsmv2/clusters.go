// Code generated by codegen; DO NOT EDIT.

package cloudhsmv2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:        "aws_cloudhsmv2_clusters",
		Description: "https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_Cluster.html",
		Resolver:    fetchCloudhsmv2Clusters,
		Multiplex:   client.ServiceAccountRegionMultiplexer("cloudhsmv2"),
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
				Resolver: resolveClusterArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTagField("TagList"),
			},
			{
				Name:     "backup_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BackupPolicy"),
			},
			{
				Name:     "backup_retention_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BackupRetentionPolicy"),
			},
			{
				Name:     "certificates",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Certificates"),
			},
			{
				Name:     "cluster_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterId"),
			},
			{
				Name:     "create_timestamp",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreateTimestamp"),
			},
			{
				Name:     "hsm_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HsmType"),
			},
			{
				Name:     "hsms",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Hsms"),
			},
			{
				Name:     "pre_co_password",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PreCoPassword"),
			},
			{
				Name:     "security_group",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecurityGroup"),
			},
			{
				Name:     "source_backup_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceBackupId"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "state_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StateMessage"),
			},
			{
				Name:     "subnet_mapping",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SubnetMapping"),
			},
			{
				Name:     "vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcId"),
			},
		},
	}
}
