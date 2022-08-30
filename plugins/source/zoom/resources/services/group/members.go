package group

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/zoom/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/himalayan-institute/zoom-lib-golang"
)

func Members() *schema.Table {
	return &schema.Table{
		Name:        "zoom_group_members",
		Description: "The list of all members that are a part of a specific group.",
		Resolver:    fetchGroupMembers,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"group_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "group_id",
				Description: "Group ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("GroupID"),
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

func fetchGroupMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)

	list, err := svc.ListGroups()
	if err != nil {
		return diag.WrapError(err)
	}

	for _, group := range list.Groups {
		listMembers := zoom.ListGroupMembersOptions{GroupID: group.ID}
		for {
			groupMembers, err := svc.ListGroupMembers(listMembers)
			if err != nil {
				return diag.WrapError(err)
			}
			for _, member := range groupMembers.Members {
				res <- Member{
					GroupID: group.ID,
					Member:  member,
				}
			}
			if groupMembers.NextPageToken == "" {
				break
			}
			listMembers.NextPageToken = &groupMembers.NextPageToken
		}
	}
	return nil
}
