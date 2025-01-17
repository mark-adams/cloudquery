// Code generated by codegen; DO NOT EDIT.

package users

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func UserContactMethods() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_user_contact_methods",
		Description: `https://developer.pagerduty.com/api-reference/50d46c0eb020d-list-a-user-s-contact-methods`,
		Resolver:    fetchUserContactMethods,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "summary",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Summary"),
			},
			{
				Name:     "self",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Self"),
			},
			{
				Name:     "html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HTMLURL"),
			},
			{
				Name:     "label",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Label"),
			},
			{
				Name:     "address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Address"),
			},
			{
				Name:     "send_short_email",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SendShortEmail"),
			},
			{
				Name:     "send_html_email",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SendHTMLEmail"),
			},
			{
				Name:     "blacklisted",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Blacklisted"),
			},
			{
				Name:     "country_code",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CountryCode"),
			},
			{
				Name:     "enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Enabled"),
			},
		},
	}
}
