package account

import "github.com/himalayan-institute/zoom-lib-golang"

type (
	UnifiedSettings struct {
		Settings                zoom.AccountSettings
		MeetingAuthentication   Authentication             `json:"meeting_authentication"`
		RecordingAuthentication Authentication             `json:"recording_authentication"`
		ManagedDomains          zoom.AccountManagedDomains `json:"managed_domains"`
		TrustedDomains          zoom.AccountTrustedDomains `json:"trusted_domains"`
	}
	Authentication struct {
		Enabled               bool        `json:"enabled"`
		AuthenticationOptions interface{} `json:"authentication_options"`
	}
)
