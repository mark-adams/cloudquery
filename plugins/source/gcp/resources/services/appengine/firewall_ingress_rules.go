package appengine

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	appengine "cloud.google.com/go/appengine/apiv1"
)

func FirewallIngressRules() *schema.Table {
	return &schema.Table{
		Name:        "gcp_appengine_firewall_ingress_rules",
		Description: `https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.firewall.ingressRules#FirewallRule`,
		Resolver:    fetchFirewallIngressRules,
		Multiplex:   client.ProjectMultiplexEnabledServices("appengine.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.FirewallRule{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchFirewallIngressRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListIngressRulesRequest{
		Parent: "apps/" + c.ProjectId,
	}
	gcpClient, err := appengine.NewFirewallClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListIngressRules(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp
	}
	return nil
}
