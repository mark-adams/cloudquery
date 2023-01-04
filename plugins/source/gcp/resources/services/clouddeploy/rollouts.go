// Code generated by codegen; DO NOT EDIT.

package clouddeploy

import (
	"context"
	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/deploy/apiv1/deploypb"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"cloud.google.com/go/deploy/apiv1"
)

func Rollouts() *schema.Table {
	return &schema.Table{
		Name:        "gcp_clouddeploy_rollouts",
		Description: `https://cloud.google.com/deploy/docs/api/reference/rest/v1/projects.locations.deliveryPipelines.releases.rollouts#Rollout`,
		Resolver:    fetchRollouts,
		Multiplex:   client.ProjectMultiplexEnabledServices("clouddeploy.googleapis.com"),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "uid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Uid"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "annotations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Annotations"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "create_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("CreateTime"),
			},
			{
				Name:     "approve_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("ApproveTime"),
			},
			{
				Name:     "enqueue_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("EnqueueTime"),
			},
			{
				Name:     "deploy_start_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("DeployStartTime"),
			},
			{
				Name:     "deploy_end_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("DeployEndTime"),
			},
			{
				Name:     "target_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TargetId"),
			},
			{
				Name:     "approval_state",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("ApprovalState"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("State"),
			},
			{
				Name:     "failure_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FailureReason"),
			},
			{
				Name:     "deploying_build",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeployingBuild"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
			{
				Name:     "deploy_failure_cause",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("DeployFailureCause"),
			},
			{
				Name:     "phases",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Phases"),
			},
			{
				Name:     "metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Metadata"),
			},
		},

		Relations: []*schema.Table{
			JobRuns(),
		},
	}
}

func fetchRollouts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListRolloutsRequest{
		Parent: parent.Item.(*pb.Release).Name,
	}
	gcpClient, err := deploy.NewCloudDeployClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListRollouts(ctx, req, c.CallOptions...)
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