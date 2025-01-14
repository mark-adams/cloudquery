// Code generated by codegen; DO NOT EDIT.

package docdb

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func GlobalClusters() *schema.Table {
	return &schema.Table{
		Name:        "aws_docdb_global_clusters",
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_GlobalCluster.html`,
		Resolver:    fetchDocdbGlobalClusters,
		Multiplex:   client.ServiceAccountRegionMultiplexer("docdb"),
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
				Resolver: schema.PathResolver("GlobalClusterArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "database_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DatabaseName"),
			},
			{
				Name:     "deletion_protection",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DeletionProtection"),
			},
			{
				Name:     "engine",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Engine"),
			},
			{
				Name:     "engine_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EngineVersion"),
			},
			{
				Name:     "global_cluster_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GlobalClusterIdentifier"),
			},
			{
				Name:     "global_cluster_members",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("GlobalClusterMembers"),
			},
			{
				Name:     "global_cluster_resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GlobalClusterResourceId"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "storage_encrypted",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("StorageEncrypted"),
			},
		},
	}
}
