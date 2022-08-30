package account

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/zoom/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/himalayan-institute/zoom-lib-golang"
)

func LockedSettings() *schema.Table {
	return &schema.Table{
		Name:        "zoom_account_locked_settings",
		Description: "Locked Zoom account settings.",
		Resolver:    fetchAccountLockedSettings,
		Columns: []schema.Column{
			{
				Name:        "schedule_meeting",
				Description: "Schedule meeting settings",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "in_meeting",
				Description: "In meeting settings.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "email_notification",
				Description: "Email notification settings.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "zoom_rooms",
				Description: "Zoom rooms settings.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "recording",
				Description: "Recording settings.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "telephony",
				Description: "Telephony settings.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "tsp",
				Description: "Telephony Service Provider settings.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("TSP"),
			},
			{
				Name:        "meeting_security",
				Description: "Meeting security settings.",
				Type:        schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchAccountLockedSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)

	settings, err := svc.GetAccountLockSettings(zoom.GetAccountLockSettingsOpts{
		AccountID: "me",
	})
	if err != nil {
		return diag.WrapError(err)
	}

	meeting, err := svc.GetAccountLockSettings(zoom.GetAccountLockSettingsOpts{
		AccountID: "me",
		Option:    "meeting_security",
	})
	if err != nil {
		return diag.WrapError(err)
	}

	settings.MeetingSecurity = meeting

	res <- settings
	return nil
}
