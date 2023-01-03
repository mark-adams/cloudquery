package client

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []any) (err error) {
	tx, err := c.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  true,
	})
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	rows, err := tx.QueryContext(ctx,
		fmt.Sprintf(`select * from %s where %s = $sourceName order by %s asc`,
			sanitizeIdentifier(table.Name),
			sanitizeIdentifier(schema.CqSourceNameColumn.Name),
			sanitizeIdentifier(schema.CqSyncTimeColumn.Name),
		),
		sql.Named("sourceName", sourceName),
	)
	if err != nil {
		return err
	}

	cols, err := rows.Columns()
	if err != nil {
		return err
	}
	for rows.NextResultSet() {
		for rows.Next() {
			row := make([]any, len(cols))
			//row := scanArray(len(cols))
			if err := rows.Scan(row...); err != nil {
				return err
			}
			res <- extractScanned(row)
		}
	}
	if err := rows.Close(); err != nil {
		return err
	}
	return tx.Commit()
}

func scanArray(l int) []any {
	res := make([]any, 0, l)
	for i := 0; i < l; i++ {
		res = append(res, new(any))
	}
	return res
}

func extractScanned(scanned []any) []any {
	res := make([]any, len(scanned))
	for i, s := range scanned {
		res[i] = *s.(*any)
	}
	return res
}
