package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
)

// Migrate relies on the CLI/client to lock before running migration.
func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	for _, table := range tables {
		c.logger.Info().Str("table", table.Name).Msg("Migrating table")
		if len(table.Columns) == 0 {
			c.logger.Info().Str("table", table.Name).Msg("Table with no columns, skipping")
			continue
		}
		tableExist, err := c.tableExists(ctx, table.Name)
		if err != nil {
			return fmt.Errorf("failed to check if table %s exists: %w", table.Name, err)
		}
		if tableExist {
			c.logger.Info().Str("table", table.Name).Msg("Table exists, auto-migrating")
			if err := c.autoMigrateTable(ctx, table); err != nil {
				return err
			}
		} else {
			c.logger.Debug().Str("table", table.Name).Msg("Table doesn't exist, creating")
			if err := c.createTableIfNotExist(ctx, table); err != nil {
				return err
			}
		}
		if err := c.Migrate(ctx, table.Relations); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) tableExists(ctx context.Context, table string) (bool, error) {
	const tableExistsQuery = `select count(*) from information_schema.tables where table_name = $tableName`

	var tableExist int
	if err := c.db.QueryRowContext(ctx, tableExistsQuery, table).Scan(&tableExist); err != nil {
		return false, fmt.Errorf("failed to check if table %s exists: %w", table, err)
	}
	return tableExist == 1, nil
}

func (c *Client) autoMigrateTable(ctx context.Context, table *schema.Table) error {
	panic("implement")
}
