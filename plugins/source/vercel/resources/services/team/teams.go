// Code generated by codegen; DO NOT EDIT.

package team

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func Teams() *schema.Table {
	return &schema.Table{
		Name:     "vercel_teams",
		Resolver: fetchTeams,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "slug",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Slug"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "avatar",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Avatar"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "creator_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreatorID"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "profiles",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Profiles"),
			},
			{
				Name:     "staging_prefix",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StagingPrefix"),
			},
		},

		Relations: []*schema.Table{
			TeamMembers(),
		},
	}
}
