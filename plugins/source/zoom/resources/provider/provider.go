package provider

import (
	"github.com/cloudquery/cloudquery/plugins/source/zoom/client"
	"github.com/cloudquery/cloudquery/plugins/source/zoom/resources/services/account"
	"github.com/cloudquery/cloudquery/plugins/source/zoom/resources/services/group"
	"github.com/cloudquery/cloudquery/plugins/source/zoom/resources/services/meeting"
	"github.com/cloudquery/cloudquery/plugins/source/zoom/resources/services/recording"
	"github.com/cloudquery/cloudquery/plugins/source/zoom/resources/services/role"
	"github.com/cloudquery/cloudquery/plugins/source/zoom/resources/services/user"
	sdkprovider "github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	Version = "Development"
)

func Provider() *sdkprovider.Provider {
	return &sdkprovider.Provider{
		Version:   Version,
		Name:      "zoom",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"zoom.account_settings":        account.Settings(),
			"zoom.account_locked_settings": account.LockedSettings(),
			"zoom.cloud_recordings":        recording.Recordings(),
			"zoom.group_members":           group.Members(),
			"zoom.groups":                  group.Groups(),
			"zoom.meetings":                meeting.Meetings(),
			"zoom.role_members":            role.Members(),
			"zoom.roles":                   role.Roles(),
			"zoom.users":                   user.Users(),
		},
		Config: func() sdkprovider.Config {
			return &client.Config{}
		},
	}
}
