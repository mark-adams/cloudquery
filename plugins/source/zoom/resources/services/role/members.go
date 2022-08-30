package role

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/zoom/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/himalayan-institute/zoom-lib-golang"
)

func Members() *schema.Table {
	return &schema.Table{
		Name:        "zoom_role_members",
		Description: "The list of all members that are assigned a specific role.",
		Resolver:    fetchRoleMembers,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"role_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "role_id",
				Description: "Role ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RoleID"),
			},
			{
				Name:        "id",
				Description: "Member ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Member.ID"),
			},
			{
				Name:        "email",
				Description: "Member email.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Member.Email"),
			},
			{
				Name:        "first_name",
				Description: "Member first name.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Member.FirstName"),
			},
			{
				Name:        "last_name",
				Description: "Member last name.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Member.LastName"),
			},
			{
				Name:        "type",
				Description: "Member type.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Member.Type"),
			},
			{
				Name:        "department",
				Description: "Member department.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Member.Department"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchRoleMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)

	list, err := svc.ListRoles()
	if err != nil {
		return diag.WrapError(err)
	}

	for _, role := range list.Roles {
		listMembers := zoom.ListRoleMembersOptions{RoleID: role.ID}
		for {
			roleMembers, err := svc.ListRoleMembers(listMembers)
			if err != nil {
				return diag.WrapError(err)
			}
			for _, member := range roleMembers.Members {
				res <- Member{
					RoleID: role.ID,
					Member: member,
				}
			}
			if roleMembers.NextPageToken == "" {
				break
			}
			listMembers.NextPageToken = roleMembers.NextPageToken
		}
	}
	return nil
}
