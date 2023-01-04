package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) autoMigrateTable(ctx context.Context, table *schema.Table) error {
	c.logger.Info().Str("table", table.Name).Msg("Table exists, auto-migrating")

	panic("implement")
}
