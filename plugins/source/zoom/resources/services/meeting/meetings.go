package meeting

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/zoom/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/himalayan-institute/zoom-lib-golang"
)

func Meetings() *schema.Table {
	return &schema.Table{
		Name:        "zoom_meetings",
		Description: "Scheduled Zoom meetings.",
		Resolver:    fetchMeetingMeetings,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"uuid"}},
		Columns: []schema.Column{
			{
				Name:        "uuid",
				Description: "Unique Meeting ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("UUID"),
			},
			{
				Name:        "id",
				Description: "Meeting ID - also known as the meeting number in long (int64) format.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "host_id",
				Description: "ID of the user who is set as the host of the meeting.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HostID"),
			},
			{
				Name:        "topic",
				Description: "Meeting topic.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Meeting type.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "status",
				Description: "Meeting status.",
				Type:        schema.TypeString,
			},
			{
				Name:        "start_time",
				Description: "Meeting start time.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("StartTime.Time"),
			},
			{
				Name:        "duration",
				Description: "Meeting duration.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "timezone",
				Description: "Timezone to format the meeting start time.",
				Type:        schema.TypeString,
			},
			{
				Name:        "created_at",
				Description: "Time of creation.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("CreatedAt.Time"),
			},
			{
				Name:        "agenda",
				Description: "Meeting description.",
				Type:        schema.TypeString,
			},
			{
				Name:        "start_url",
				Description: "The URL using which a host or an alternative host can start the Meeting.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StartURL"),
			},
			{
				Name:        "join_url",
				Description: "URL for participants to join the meeting.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("JoinURL"),
			},
			{
				Name:        "password",
				Description: "Meeting passcode.",
				Type:        schema.TypeString,
			},
			{
				Name:        "h323_password",
				Description: "H.323/SIP room system passcode.",
				Type:        schema.TypeString,
			},
			{
				Name:        "encrypted_password",
				Description: "Encrypted passcode for third party endpoints (H323/SIP).",
				Type:        schema.TypeString,
			},
			{
				Name:        "personal_meeting_id",
				Description: "Personal Meeting ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("JoinURL"),
			},
			{
				Name:        "tracking_fields",
				Description: "",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "occurrences",
				Description: "Array of occurrence objects.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "settings",
				Description: "Meeting settings.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "recurrence",
				Description: "Recurrence object.",
				Type:        schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchMeetingMeetings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)

	listUsers := zoom.ListUsersOptions{}
	for {
		users, err := svc.ListUsers(listUsers)
		if err != nil {
			return diag.WrapError(err)
		}

		for _, user := range users.Users {
			param300 := 300
			listInput := zoom.ListMeetingsOptions{
				HostID:   user.ID,
				PageSize: &param300,
			}

			for {
				meetings, err := svc.ListMeetings(listInput)
				if err != nil {
					return diag.WrapError(err)
				}
				for _, meetingMeta := range meetings.Meetings {
					meetingInfo, err := svc.GetMeeting(zoom.GetMeetingOptions{MeetingID: meetingMeta.ID})
					if err != nil {
						return diag.WrapError(err)
					}
					res <- meetingInfo
				}
				if meetings.NextPageToken == "" {
					break
				}
				listInput.NextPageToken = &meetings.NextPageToken
			}
		}
		if users.NextPageToken == "" {
			return nil
		}
		listUsers.NextPageToken = &users.NextPageToken
	}
}
