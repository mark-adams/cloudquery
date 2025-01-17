// Code generated by codegen; DO NOT EDIT.

package deployment

import (
	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Deployments() *schema.Table {
	return &schema.Table{
		Name:      "vercel_deployments",
		Resolver:  fetchDeployments,
		Multiplex: client.TeamMultiplex,
		Columns: []schema.Column{
			{
				Name:     "uid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("URL"),
			},
			{
				Name:     "source",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Source"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "inspector_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InspectorURL"),
			},
			{
				Name:     "is_rollback_candidate",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsRollbackCandidate"),
			},
			{
				Name:     "ready",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Ready"),
			},
			{
				Name:     "checks_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ChecksState"),
			},
			{
				Name:     "checks_conclusion",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ChecksConclusion"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "building_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("BuildingAt"),
			},
		},

		Relations: []*schema.Table{
			DeploymentChecks(),
		},
	}
}
