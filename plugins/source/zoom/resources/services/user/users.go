package user

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/zoom/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/himalayan-institute/zoom-lib-golang"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:        "zoom_users",
		Description: "User represents an account user.s",
		Resolver:    fetchMeetingUsers,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "id",
				Description: "The user's ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "first_name",
				Description: "The user's first name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_name",
				Description: "The user's last name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "email",
				Description: "The user's email address.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The user's assigned plan type.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "status",
				Description: "The user's status.",
				Type:        schema.TypeString,
			},
			{
				Name:        "personal_meeting_id",
				Description: "The user's Personal Meeting ID.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("PMI"),
			},
			{
				Name:        "timezone",
				Description: "The user's timezone.",
				Type:        schema.TypeString,
			},
			{
				Name:        "dept",
				Description: "The user's department.",
				Type:        schema.TypeString,
			},
			{
				Name:        "role_id",
				Description: "The unique ID of the user's assigned role.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RoleID"),
			},
			{
				Name:        "created_at",
				Description: "The time at which the user's account was created.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("CreatedAt.Time"),
			},
			{
				Name:        "last_login_time",
				Description: "The user's last login time.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("LastLoginTime.Time"),
			},
			{
				Name:        "last_client_version",
				Description: "The last client version that user used to log in.",
				Type:        schema.TypeString,
			},
			{
				Name:        "group_ids",
				Description: "The IDs of groups where the user is a member.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("GroupIDs"),
			},
			{
				Name:        "im_group_ids",
				Description: "The IDs of IM directory groups where the user is a member.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("IMGroupIDs"),
			},
			{
				Name:        "verified",
				Description: "Whether the user's email address for the Zoom account is verified.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "custom_attributes",
				Description: "Information about the user's custom attributes.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "plan_united_type",
				Description: "The user's Zoom United plan.",
				Type:        schema.TypeString,
			},
			{
				Name:     "use_pmi",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("UsePMI"),
			},
			{
				Name:        "language",
				Description: "Default language for the Zoom Web Portal.",
				Type:        schema.TypeString,
			},
			{
				Name:        "phone_numbers",
				Description: "The user's phone numbers.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "vanity_url",
				Description: "Personal meeting room URL, if the user has one.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VanityURL"),
			},
			{
				Name:        "personal_meeting_url",
				Description: "User's personal meeting url.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PersonalMeetingURL"),
			},
			{
				Name:        "pic_url",
				Description: "The URL for user's profile picture.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PicURL"),
			},
			{
				Name:        "cms_user_id",
				Description: "CMS ID of the user.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CMSUserID"),
			},
			{
				Name:        "account_id",
				Description: "User's account ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountID"),
			},
			{
				Name:        "jid",
				Description: "The user's Jabber ID (JID).",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("JID"),
			},
			{
				Name:        "job_title",
				Description: "The user's job title.",
				Type:        schema.TypeString,
			},
			{
				Name:        "company",
				Description: "The user's company.",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "User's location.",
				Type:        schema.TypeString,
			},
			{
				Name:        "login_type",
				Description: "The user's login method.",
				Type:        schema.TypeBigInt,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchMeetingUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)

	includeFields := []string{"custom_attributes"}
	input := zoom.ListUsersOptions{
		PageSize:      300,
		IncludeFields: &includeFields,
	}
	for {
		output, err := svc.ListUsers(input)
		if err != nil {
			return diag.WrapError(err)
		}
		for _, userBasicInfo := range output.Users {
			// need to make separate GET call as some info isn't returned on ListUsers API (e.g., `use_pmi`)
			user, err := svc.GetUser(zoom.GetUserOpts{EmailOrID: userBasicInfo.ID})
			if err != nil {
				return diag.WrapError(err)
			}
			res <- user
		}
		if output.NextPageToken == "" {
			return nil
		}
		input.NextPageToken = &output.NextPageToken
	}
}
