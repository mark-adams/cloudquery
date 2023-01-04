package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) createTableIfNotExist(ctx context.Context, table *schema.Table) error {
	var sb strings.Builder
	sb.WriteString("CREATE TABLE IF NOT EXISTS ")
	sb.WriteString(sanitizeIdentifier(table.Name))

	sb.WriteString(" (")
	sb.WriteString(strings.Join(c.getColumnDefinitions(table.Columns), ", "))

	sb.WriteString(", CONSTRAINT ")
	sb.WriteString(sanitizeIdentifier(table.Name + "_cqpk"))
	sb.WriteString(" PRIMARY KEY (")
	primaryKeys := c.getPrimaryKeys(table.Columns)
	if len(primaryKeys) == 0 {
		// if no primary keys are defined, add a PK constraint for _cq_id
		primaryKeys = []string{"_cq_id"}
	}
	sb.WriteString(strings.Join(primaryKeys, ","))
	sb.WriteString("))")
	_, err := c.db.ExecContext(ctx, sb.String())
	if err != nil {
		return fmt.Errorf("failed to create table %s: %w", table.Name, err)
	}
	return nil
}
