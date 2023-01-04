package client

import (
	"context"
	"database/sql"
	"strings"

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

type (
	msSQLTableInfo struct {
		name    string
		columns []msSQLColumn
	}
	msSQLColumn struct {
		name string
		typ  string
	}
)

func (c *Client) getTableColumns(ctx context.Context, tableName string) (*msSQLTableInfo, error) {
	tc := msSQLTableInfo{
		name: tableName,
	}

	const sqlSelectColumnTypes = `select column_name, data_type from information_schema.columns
where table_name = $tableName
order by ordinal_position asc`
	rows, err := c.db.QueryContext(ctx, sqlSelectColumnTypes, sql.Named("tableName", tableName))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.NextResultSet() {
		for rows.Next() {
			var name string
			var typ string
			if err := rows.Scan(&name, &typ); err != nil {
				return nil, err
			}
			tc.columns = append(tc.columns, msSQLColumn{
				name: strings.ToLower(name),
				typ:  strings.ToLower(typ),
			})
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &tc, nil
}
