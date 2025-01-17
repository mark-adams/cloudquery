// Code generated by codegen; DO NOT EDIT.

package groups

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func Groups() *schema.Table {
	return &schema.Table{
		Name:     "okta_groups",
		Resolver: fetchGroups,
		Columns: []schema.Column{
			{
				Name:     "created",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Created"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "last_membership_updated",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastMembershipUpdated"),
			},
			{
				Name:     "last_updated",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastUpdated"),
			},
			{
				Name:     "object_class",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ObjectClass"),
			},
			{
				Name:     "profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Profile"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "_embedded",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Embedded"),
			},
			{
				Name:     "_links",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Links"),
			},
			{
				Name:     "additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AdditionalProperties"),
			},
		},

		Relations: []*schema.Table{
			GroupUsers(),
		},
	}
}
