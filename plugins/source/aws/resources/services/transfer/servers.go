package transfer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/transfer"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Servers() *schema.Table {
	return &schema.Table{
		Name:        "aws_transfer_servers",
		Description: "Describes the properties of a file transfer protocol-enabled server that was specified",
		Resolver:    fetchTransferServers,
		Multiplex:   client.ServiceAccountRegionMultiplexer("glue"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:            "arn",
				Description:     "Specifies the unique Amazon Resource Name (ARN) of the server",
				Type:            schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "certificate",
				Description: "Specifies the ARN of the Amazon Web ServicesCertificate Manager (ACM) certificate",
				Type:        schema.TypeString,
			},
			{
				Name:        "domain",
				Description: "Specifies the domain of the storage system that is used for file transfers",
				Type:        schema.TypeString,
			},
			{
				Name:        "endpoint_details_address_allocation_ids",
				Description: "A list of address allocation IDs that are required to attach an Elastic IP address to your server's endpoint",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("EndpointDetails.AddressAllocationIds"),
			},
			{
				Name:        "endpoint_details_security_group_ids",
				Description: "A list of security groups IDs that are available to attach to your server's endpoint",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("EndpointDetails.SecurityGroupIds"),
			},
			{
				Name:     "endpoint_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EndpointDetails"),
			},
			{
				Name:     "endpoint_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EndpointDetails"),
			},
			{
				Name:        "endpoint_details_vpc_id",
				Description: "The VPC ID of the VPC in which a server's endpoint will be hosted",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EndpointDetails.VpcId"),
			},
			{
				Name:        "endpoint_type",
				Description: "Defines the type of endpoint that your server is connected to",
				Type:        schema.TypeString,
			},
			{
				Name:        "host_key_fingerprint",
				Description: "Specifies the Base64-encoded SHA256 fingerprint of the server's host key",
				Type:        schema.TypeString,
			},
			{
				Name:     "identity_provider_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IdentityProviderDetails"),
			},
			{
				Name:     "identity_provider_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IdentityProviderDetails"),
			},
			{
				Name:     "identity_provider_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IdentityProviderDetails"),
			},
			{
				Name:     "identity_provider_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IdentityProviderDetails"),
			},
			{
				Name:        "identity_provider_type",
				Description: "The mode of authentication for a server",
				Type:        schema.TypeString,
			},
			{
				Name:        "logging_role",
				Description: "The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role that allows a server to turn on Amazon CloudWatch logging for Amazon S3 or Amazon EFSevents",
				Type:        schema.TypeString,
			},
			{
				Name:        "post_authentication_login_banner",
				Description: "Specifies a string to display when users connect to a server",
				Type:        schema.TypeString,
			},
			{
				Name:        "pre_authentication_login_banner",
				Description: "Specifies a string to display when users connect to a server",
				Type:        schema.TypeString,
			},
			{
				Name:     "protocol_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ProtocolDetails"),
			},
			{
				Name:     "protocol_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ProtocolDetails"),
			},
			{
				Name:     "protocol_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ProtocolDetails"),
			},
			{
				Name:     "protocol_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ProtocolDetails"),
			},
			{
				Name:        "protocols",
				Description: "Specifies the file transfer protocol or protocols over which your file transfer protocol client can connect to your server's endpoint",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "security_policy_name",
				Description: "Specifies the name of the security policy that is attached to the server",
				Type:        schema.TypeString,
			},
			{
				Name:        "server_id",
				Description: "Specifies the unique system-assigned identifier for a server that you instantiate",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "The condition of the server that was described",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Specifies the key-value pairs that you can use to search for and group servers that were assigned to the server that was described",
				Type:        schema.TypeJSON,
				Resolver:    resolveServersTags,
			},
			{
				Name:        "user_count",
				Description: "Specifies the number of users that are assigned to a server you specified with the ServerId",
				Type:        schema.TypeInt,
			},
			{
				Name:        "workflow_details",
				Description: "Specifies the workflow ID for the workflow to assign and the execution role that's used for executing the workflow",
				Type:        schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchTransferServers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Transfer
	input := transfer.ListServersInput{MaxResults: aws.Int32(1000)}
	for {
		result, err := svc.ListServers(ctx, &input)
		if err != nil {
			return err
		}
		for _, server := range result.Servers {
			desc, err := svc.DescribeServer(ctx, &transfer.DescribeServerInput{ServerId: server.ServerId})
			if err != nil {
				if cl.IsNotFoundError(err) {
					continue
				}
				return err
			}
			if desc.Server != nil {
				res <- desc.Server
			}
		}
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}
func resolveServersTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Transfer
	server := resource.Item.(*types.DescribedServer)
	input := transfer.ListTagsForResourceInput{Arn: server.Arn}
	tags := make(map[string]string)
	for {
		result, err := svc.ListTagsForResource(ctx, &input)
		if err != nil {
			if cl.IsNotFoundError(err) {
				continue
			}
			return err
		}
		client.TagsIntoMap(result.Tags, tags)
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return resource.Set(c.Name, tags)
}
