package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
)

func (c *Client) enabledPks() bool {
	return c.spec.WriteMode == specs.WriteModeOverwrite || c.spec.WriteMode == specs.WriteModeOverwriteDeleteStale
}

func (c *Client) getPrimaryKeys(columns schema.ColumnList) []string {
	if !c.enabledPks() {
		return nil
	}

	var result []string
	for _, column := range columns {
		if column.CreationOptions.PrimaryKey {
			result = append(result, column.Name)
		}
	}

	return result
}

func (c *Client) getColumnDefinitions(columns schema.ColumnList) []string {
	definitions := make([]string, len(columns))
	for i, column := range columns {
		sqlType := msSQLType(column.Type)
		columnName := sanitizeIdentifier(column.Name)
		fieldDef := columnName + " " + sqlType
		if column.Name == "_cq_id" {
			// _cq_id column should always have a "unique not null" constraint
			fieldDef += " UNIQUE NOT NULL"
		}
		definitions[i] = fieldDef
	}
	return definitions
}
