package cdn

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func endpoints() *schema.Table {
	return &schema.Table{
		Name:      "azure_cdn_endpoints",
		Resolver:  fetchEndpoints,
		Transform: transformers.TransformWithStruct(&armcdn.Endpoint{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
