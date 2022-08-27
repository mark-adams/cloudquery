package provider

import (
	"github.com/cloudquery/cloudquery/plugins/source/zoom/client"
	sdkprovider "github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	Version = "Development"
)

func Provider() *sdkprovider.Provider {
	return &sdkprovider.Provider{
		Version:     Version,
		Name:        "zoom",
		Configure:   client.Configure,
		ResourceMap: map[string]*schema.Table{},
		Config: func() sdkprovider.Config {
			return &client.Config{}
		},
	}
}
