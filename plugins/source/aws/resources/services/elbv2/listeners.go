// Code generated by codegen; DO NOT EDIT.

package elbv2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Listeners() *schema.Table {
	return &schema.Table{
		Name:        "aws_elbv2_listeners",
		Description: "https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_Listener.html",
		Resolver:    fetchElbv2Listeners,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticloadbalancing"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ListenerArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveElbv2listenerTags,
			},
			{
				Name:     "alpn_policy",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AlpnPolicy"),
			},
			{
				Name:     "certificates",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Certificates"),
			},
			{
				Name:     "default_actions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DefaultActions"),
			},
			{
				Name:     "load_balancer_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LoadBalancerArn"),
			},
			{
				Name:     "port",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Port"),
			},
			{
				Name:     "protocol",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Protocol"),
			},
			{
				Name:     "ssl_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SslPolicy"),
			},
		},

		Relations: []*schema.Table{
			ListenerCertificates(),
		},
	}
}
