package account

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/zoom/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/himalayan-institute/zoom-lib-golang"
)

func Settings() *schema.Table {
	return &schema.Table{
		Name:        "zoom_account_settings",
		Description: "Zoom account settings.",
		Resolver:    fetchAccountSettings,
		Columns: []schema.Column{
			{
				Name:        "schedule_meeting",
				Description: "Schedule meeting settings",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Settings.ScheduleMeeting"),
			},
			{
				Name:        "in_meeting",
				Description: "In meeting settings.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Settings.InMeeting"),
			},
			{
				Name:        "email_notification",
				Description: "Email notification settings.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Settings.EmailNotification"),
			},
			{
				Name:        "zoom_rooms",
				Description: "Zoom rooms settings.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Settings.ZoomRooms"),
			},
			{
				Name:        "security",
				Description: "Security settings.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Settings.Security"),
			},
			{
				Name:        "recording",
				Description: "Recording settings.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Settings.Recording"),
			},
			{
				Name:        "telephony",
				Description: "Telephony settings.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Settings.Telephony"),
			},
			{
				Name:        "tsp",
				Description: "Telephony Service Provider settings.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Settings.TSP"),
			},
			{
				Name:        "integration",
				Description: "Integration settings.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Settings.Integration"),
			},
			{
				Name:        "feature",
				Description: "Feature settings",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Settings.Feature"),
			},
			{
				Name:        "meeting_security",
				Description: "Meeting security settings.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Settings.MeetingSecurity"),
			},
			{
				Name:        "meeting_authentication",
				Description: "Meeting authentication settings.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "recording_authentication",
				Description: "Recording authentication settings.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "trusted_domains",
				Description: "A list of the account's trusted domains.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("TrustedDomains.TrustedDomains"),
			},
			{
				Name:        "managed_domains",
				Description: "Information about the managed domains.",
				Type:        schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchAccountSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)

	settings, err := svc.GetAccountSettings(zoom.GetAccountSettingsOpts{
		AccountID: "me",
	})
	if err != nil {
		return diag.WrapError(err)
	}

	meetingAuth, err := svc.GetAccountSettings(zoom.GetAccountSettingsOpts{
		AccountID: "me",
		Option:    "meeting_authentication",
	})
	if err != nil {
		return diag.WrapError(err)
	}

	recordingAuth, err := svc.GetAccountSettings(zoom.GetAccountSettingsOpts{
		AccountID: "me",
		Option:    "recording_authentication",
	})
	if err != nil {
		return diag.WrapError(err)
	}

	meetingSecurity, err := svc.GetAccountSettings(zoom.GetAccountSettingsOpts{
		AccountID: "me",
		Option:    "meeting_security",
	})
	if err != nil {
		return diag.WrapError(err)
	}
	settings.MeetingSecurity = meetingSecurity

	security, err := svc.GetAccountSettings(zoom.GetAccountSettingsOpts{
		AccountID: "me",
		Option:    "security",
	})
	if err != nil {
		return diag.WrapError(err)
	}
	settings.Security = security

	managed, err := svc.GetAccountManagedDomains(zoom.GetAccountManagedDomainsOpts{AccountID: "me"})
	if err != nil {
		return diag.WrapError(err)
	}

	trusted, err := svc.GetAccountTrustedDomains(zoom.GetAccountTrustedDomainsOpts{AccountID: "me"})
	if err != nil {
		return diag.WrapError(err)
	}

	res <- UnifiedSettings{
		Settings: settings,
		MeetingAuthentication: Authentication{
			Enabled:               meetingAuth.MeetingAuthentication,
			AuthenticationOptions: meetingAuth.AuthenticationOptions,
		},
		RecordingAuthentication: Authentication{
			Enabled:               recordingAuth.MeetingAuthentication,
			AuthenticationOptions: recordingAuth.AuthenticationOptions,
		},
		ManagedDomains: managed,
		TrustedDomains: trusted,
	}
	return nil
}
