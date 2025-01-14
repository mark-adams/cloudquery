// Code generated by codegen; DO NOT EDIT.

package downtimes

import (
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Downtimes() *schema.Table {
	return &schema.Table{
		Name:      "datadog_downtimes",
		Resolver:  fetchDowntimes,
		Multiplex: client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_name",
				Type:     schema.TypeString,
				Resolver: client.ResolveAccountName,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "active",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Active"),
			},
			{
				Name:     "active_child",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ActiveChild"),
			},
			{
				Name:     "canceled",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Canceled"),
			},
			{
				Name:     "creator_id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CreatorId"),
			},
			{
				Name:     "disabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Disabled"),
			},
			{
				Name:     "downtime_type",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DowntimeType"),
			},
			{
				Name:     "end",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("End"),
			},
			{
				Name:     "message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Message"),
			},
			{
				Name:     "monitor_id",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MonitorId"),
			},
			{
				Name:     "monitor_tags",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("MonitorTags"),
			},
			{
				Name:     "mute_first_recovery_notification",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("MuteFirstRecoveryNotification"),
			},
			{
				Name:     "parent_id",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ParentId"),
			},
			{
				Name:     "recurrence",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Recurrence"),
			},
			{
				Name:     "scope",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Scope"),
			},
			{
				Name:     "start",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Start"),
			},
			{
				Name:     "timezone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Timezone"),
			},
			{
				Name:     "updater_id",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("UpdaterId"),
			},
		},
	}
}
