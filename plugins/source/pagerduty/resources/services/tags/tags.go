// Code generated by codegen; DO NOT EDIT.

package tags

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func Tags() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_tags",
		Description: `https://developer.pagerduty.com/api-reference/e44b160c69bf3-list-tags`,
		Resolver:    fetchTags,
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
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "summary",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Summary"),
			},
			{
				Name:     "self",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Self"),
			},
			{
				Name:     "html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HTMLURL"),
			},
			{
				Name:     "label",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Label"),
			},
		},
	}
}
