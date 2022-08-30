package group

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/zoom/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Groups() *schema.Table {
	return &schema.Table{
		Name:        "zoom_groups",
		Description: "Group represents an account group.",
		Resolver:    fetchGroups,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "id",
				Description: "Group ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Group name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "total_members",
				Description: "Total number of members in this group.",
				Type:        schema.TypeBigInt,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)

	groups, err := svc.ListGroups()
	if err != nil {
		return diag.WrapError(err)
	}
	res <- groups.Groups
	return nil
}
