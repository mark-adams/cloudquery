package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/csv/client"
	"github.com/cloudquery/plugin-sdk/serve"
)

func main() {
	serve.Destination(client.New())
}
