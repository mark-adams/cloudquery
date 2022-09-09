package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Directories() *schema.Table {
	return &schema.Table{
		Name:        "aws_workspaces_directories",
		Description: "Describes a directory that is used with Amazon WorkSpaces.",
		Resolver:    fetchWorkspacesDirectories,
		Multiplex:   client.ServiceAccountRegionMultiplexer("workspaces"),
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
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the workspaces directory",
				Type:        schema.TypeString,
				Resolver: client.ResolveARN(client.WorkspacesService, func(resource *schema.Resource) ([]string, error) {
					return []string{"directory", *resource.Item.(types.WorkspaceDirectory).DirectoryId}, nil
				}),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "alias",
				Description: "The directory alias.",
				Type:        schema.TypeString,
			},
			{
				Name:        "customer_user_name",
				Description: "The user name for the service account.",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The directory identifier.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DirectoryId"),
			},
			{
				Name:        "name",
				Description: "The name of the directory.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DirectoryName"),
			},
			{
				Name:        "type",
				Description: "The directory type.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DirectoryType"),
			},
			{
				Name:        "dns_ip_addresses",
				Description: "The IP addresses of the DNS servers for the directory.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "iam_role_id",
				Description: "The identifier of the IAM role",
				Type:        schema.TypeString,
			},
			{
				Name:          "ip_group_ids",
				Description:   "The identifiers of the IP access control groups associated with the directory.",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "registration_code",
				Description: "The registration code for the directory",
				Type:        schema.TypeString,
			},
			{
				Name:        "change_compute_type",
				Description: "Specifies whether users can change the compute type (bundle) for their WorkSpace.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SelfservicePermissions.ChangeComputeType"),
			},
			{
				Name:     "selfservice_permissions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SelfservicePermissions"),
			},
			{
				Name:     "selfservice_permissions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SelfservicePermissions"),
			},
			{
				Name:     "selfservice_permissions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SelfservicePermissions"),
			},
			{
				Name:     "selfservice_permissions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SelfservicePermissions"),
			},
			{
				Name:        "state",
				Description: "The state of the directory's registration with Amazon WorkSpaces",
				Type:        schema.TypeString,
			},
			{
				Name:        "subnet_ids",
				Description: "The identifiers of the subnets used with the directory.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "tenancy",
				Description: "Specifies whether the directory is dedicated or shared",
				Type:        schema.TypeString,
			},
			{
				Name:        "device_type_android",
				Description: "Indicates whether users can use Android and Android-compatible Chrome OS devices to access their WorkSpaces.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("WorkspaceAccessProperties.DeviceTypeAndroid"),
			},
			{
				Name:     "workspace_access_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("WorkspaceAccessProperties"),
			},
			{
				Name:     "workspace_access_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("WorkspaceAccessProperties"),
			},
			{
				Name:     "workspace_access_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("WorkspaceAccessProperties"),
			},
			{
				Name:     "workspace_access_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("WorkspaceAccessProperties"),
			},
			{
				Name:     "workspace_access_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("WorkspaceAccessProperties"),
			},
			{
				Name:     "workspace_access_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("WorkspaceAccessProperties"),
			},
			{
				Name:     "workspace_access_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("WorkspaceAccessProperties"),
			},
			{
				Name:          "workspace_creation_properties",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("WorkspaceCreationProperties"),
				IgnoreInTests: true,
			},
			{
				Name:          "default_ou",
				Description:   "The organizational unit (OU) in the directory for the WorkSpace machine accounts.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("WorkspaceCreationProperties.DefaultOu"),
				IgnoreInTests: true,
			},
			{
				Name:     "workspace_creation_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("WorkspaceCreationProperties"),
			},
			{
				Name:     "workspace_creation_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("WorkspaceCreationProperties"),
			},
			{
				Name:     "workspace_creation_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("WorkspaceCreationProperties"),
			},
			{
				Name:     "workspace_creation_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("WorkspaceCreationProperties"),
			},
			{
				Name:        "workspace_security_group_id",
				Description: "The identifier of the security group that is assigned to new WorkSpaces.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchWorkspacesDirectories(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Workspaces
	input := workspaces.DescribeWorkspaceDirectoriesInput{}
	for {
		output, err := svc.DescribeWorkspaceDirectories(ctx, &input)
		if err != nil {
			return err
		}
		res <- output.Directories
		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	return nil
}
