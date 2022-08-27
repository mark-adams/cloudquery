package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/zoom/resources/provider"
	"github.com/cloudquery/cq-provider-sdk/serve"
)

func main() {
	serve.Serve(&serve.Options{
		Name:     "zoom",
		Provider: provider.Provider(),
	})
}
