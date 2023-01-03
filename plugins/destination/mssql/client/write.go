package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) WriteTableBatch(context.Context, *schema.Table, [][]any) error {
	panic("WriteTableBatch not implemented")
}

func (c *Client) Write(context.Context, schema.Tables, <-chan *destination.ClientResource) error {
	panic("Write not implemented")
}

func (c *Client) Metrics() destination.Metrics {
	return c.metrics
}
