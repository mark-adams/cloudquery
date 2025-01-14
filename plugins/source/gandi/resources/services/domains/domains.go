// Code generated by codegen; DO NOT EDIT.

package domains

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func Domains() *schema.Table {
	return &schema.Table{
		Name:                "gandi_domains",
		Resolver:            fetchDomains,
		PreResourceResolver: getDomain,
		Columns: []schema.Column{
			{
				Name:     "autorenew",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AutoRenew"),
			},
			{
				Name:     "can_tld_lock",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("CanTLDLock"),
			},
			{
				Name:     "contacts",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Contacts"),
			},
			{
				Name:     "dates",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Dates"),
			},
			{
				Name:     "fqdn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FQDN"),
			},
			{
				Name:     "fqdn_unicode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FQDNUnicode"),
			},
			{
				Name:     "href",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Href"),
			},
			{
				Name:     "nameservers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Nameservers"),
			},
			{
				Name:     "services",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Services"),
			},
			{
				Name:     "sharing_space",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SharingSpace"),
			},
			{
				Name:     "status",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "tld",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TLD"),
			},
			{
				Name:     "authinfo",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AuthInfo"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "sharing_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SharingID"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "trustee_roles",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("TrusteeRoles"),
			},
		},

		Relations: []*schema.Table{
			DomainLiveDNS(),
			DomainWebRedirections(),
			DomainGlueRecords(),
			DomainDNSSecKeys(),
		},
	}
}
