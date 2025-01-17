// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"google.golang.org/api/iterator"

	pb "google.golang.org/genproto/googleapis/cloud/compute/v1"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"cloud.google.com/go/compute/apiv1"
)

func Addresses() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_addresses",
		Description: `https://cloud.google.com/compute/docs/reference/rest/v1/addresses#Address`,
		Resolver:    fetchAddresses,
		Multiplex:   client.ProjectMultiplexEnabledServices("compute.googleapis.com"),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Address"),
			},
			{
				Name:     "address_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AddressType"),
			},
			{
				Name:     "creation_timestamp",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreationTimestamp"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "ip_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IpVersion"),
			},
			{
				Name:     "ipv6_endpoint_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Ipv6EndpointType"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "network",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Network"),
			},
			{
				Name:     "network_tier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NetworkTier"),
			},
			{
				Name:     "prefix_length",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PrefixLength"),
			},
			{
				Name:     "purpose",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Purpose"),
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Region"),
			},
			{
				Name:     "self_link",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SelfLink"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "subnetwork",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Subnetwork"),
			},
			{
				Name:     "users",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Users"),
			},
		},
	}
}

func fetchAddresses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.AggregatedListAddressesRequest{
		Project: c.ProjectId,
	}
	gcpClient, err := compute.NewAddressesRESTClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.AggregatedList(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp.Value.Addresses

	}
	return nil
}
