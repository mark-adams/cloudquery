package role

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/zoom/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/himalayan-institute/zoom-lib-golang"
)

func Roles() *schema.Table {
	return &schema.Table{
		Name:        "zoom_roles",
		Description: "Role represents an account role.",
		Resolver:    fetchMeetingRoles,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "id",
				Description: "Role ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Name of the role.",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "Description of the role.",
				Type:        schema.TypeString,
			},
			{
				Name:        "total_members",
				Description: "Total members assigned to that role.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "privileges",
				Description: "Privileges assigned to the role.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "sub_account_privileges",
				Description: "This field will only be displayed to accounts that are enrolled in a partner plan and follow the master accounts and sub accounts structure.",
				Type:        schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchMeetingRoles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)

	list, err := svc.ListRoles()
	if err != nil {
		return diag.WrapError(err)
	}

	for _, role := range list.Roles {
		fullRole, err := svc.GetRole(zoom.GetRoleOpts{ID: role.ID})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- fullRole
	}
	return nil
}
