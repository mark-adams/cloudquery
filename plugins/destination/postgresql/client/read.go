package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/jackc/pgx/v5"
)

const (
	readSQL = "SELECT * FROM %s WHERE _cq_source_name = $1 order by _cq_sync_time asc"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []any) error {
	sql := fmt.Sprintf(readSQL, pgx.Identifier{table.Name}.Sanitize())
	rows, err := c.conn.Query(ctx, sql, sourceName)
	if err != nil {
		return err
	}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return err
		}
		res <- values
	}
	rows.Close()
	return nil
}
