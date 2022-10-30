// Code generated by codegen; DO NOT EDIT.

package qldb

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Ledgers() *schema.Table {
	return &schema.Table{
		Name:                "aws_qldb_ledgers",
		Resolver:            fetchQldbLedgers,
		PreResourceResolver: getLedger,
		Multiplex:           client.ServiceAccountRegionMultiplexer("qldb"),
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
				Name:        "tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveQldbLedgerTags,
				Description: `The tags associated with the pipeline.`,
			},
			{
				Name: "arn",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "creation_date_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationDateTime"),
			},
			{
				Name:     "deletion_protection",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DeletionProtection"),
			},
			{
				Name:     "encryption_description",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EncryptionDescription"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "permissions_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PermissionsMode"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "result_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResultMetadata"),
			},
		},

		Relations: []*schema.Table{
			LedgerJournalKinesisStreams(),
			LedgerJournalS3Exports(),
		},
	}
}
