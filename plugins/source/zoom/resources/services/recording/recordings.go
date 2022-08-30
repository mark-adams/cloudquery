package recording

import (
	"context"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/zoom/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/himalayan-institute/zoom-lib-golang"
)

func Recordings() *schema.Table {
	return &schema.Table{
		Name:        "zoom_cloud_recordings",
		Description: "Zoom meting recordings.",
		Resolver:    fetchCloudRecordings,
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
				Description: "Meeting ID - also known as the meeting number.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "account_id",
				Description: "Unique Identifier of the user account.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountID"),
			},
			{
				Name:        "host_id",
				Description: "ID of the user set as host of meeting.",
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
				Description: "The recording's associated type of meeting or webinar.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "start_time",
				Description: "The time at which the meeting started.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("StartTime.Time"),
			},
			{
				Name:        "duration",
				Description: "Meeting duration.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "total_size",
				Description: "The total file size of the recording.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "recording_count",
				Description: "Number of recording files.",
				Type:        schema.TypeBigInt,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchCloudRecordings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)

	listUsers := zoom.ListUsersOptions{}
	for {
		users, err := svc.ListUsers(listUsers)
		if err != nil {
			return diag.WrapError(err)
		}

		for _, user := range users.Users {
			param300 := 300
			listInput := zoom.ListAllRecordingsOptions{
				UserID:   user.ID,
				PageSize: &param300,
				From:     time.Now().UTC().AddDate(0, -30, 0).Format("2006-01-02"),
				To:       time.Now().UTC().Format("2006-01-02"),
			}
			for {
				recordings, err := svc.ListAllRecordings(listInput)
				if err != nil {
					return diag.WrapError(err)
				}
				res <- recordings.Meetings
				if recordings.NextPageToken == "" {
					break
				}
				listInput.NextPageToken = recordings.NextPageToken
			}
		}
		if users.NextPageToken == "" {
			return nil
		}
		listUsers.NextPageToken = &users.NextPageToken
	}
}
