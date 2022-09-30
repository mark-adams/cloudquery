package client

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"path"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	Version = "Development"
)

type Spec struct {
	Directory string `json:"directory,omitempty"`
}

type Client struct {
	logger     zerolog.Logger
	spec       specs.Destination
	directory  string
	tables     map[string]*schema.Table
	files      map[string]io.Closer
	csvWriters map[string]*csv.Writer
	unflushed  map[string]int
}

func New() *Client {
	return &Client{
		logger:     log.With().Str("module", "test").Logger(),
		tables:     map[string]*schema.Table{},
		files:      map[string]io.Closer{},
		csvWriters: map[string]*csv.Writer{},
	}
}

func (c *Client) Name() string {
	return "csv"
}

func (c *Client) Version() string {
	return Version
}

func (c *Client) Initialize(ctx context.Context, spec specs.Destination) error {
	var csvSpec Spec
	c.spec = spec
	if err := spec.UnmarshalSpec(&csvSpec); err != nil {
		return fmt.Errorf("failed to unmarshal csv spec: %w", err)
	}
	if csvSpec.Directory == "" {
		return errors.New("directory is a required parameter")
	}
	c.directory = csvSpec.Directory
	return nil
}

func (c *Client) SetLogger(logger zerolog.Logger) {
	c.logger = logger
}

func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	for _, t := range tables {
		c.tables[t.Name] = t

		name := fmt.Sprintf("%s.csv", t.Name)
		f, err := os.Create(path.Join(c.directory, name))
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		c.files[t.Name] = f

		w := csv.NewWriter(f)
		headerRow := make([]string, len(t.Columns))
		for i, col := range t.Columns {
			headerRow[i] = col.Name
		}
		w.Write(headerRow)
		c.csvWriters[t.Name] = w
	}
	return nil
}

func (c *Client) Close(_ context.Context) error {
	var failedToFlush []string
	for t, w := range c.csvWriters {
		w.Flush()
		err := w.Error()
		if err != nil {
			failedToFlush = append(failedToFlush, t)
		}
	}

	var failedToClose []string
	for t, f := range c.files {
		err := f.Close()
		if err != nil {
			failedToClose = append(failedToClose, t)
		}
	}

	if len(failedToFlush) > 0 {
		return fmt.Errorf("failed to flush files for: %v", strings.Join(failedToClose, ","))
	}
	if len(failedToClose) > 0 {
		return fmt.Errorf("failed to close files for: %v", strings.Join(failedToClose, ","))
	}

	return nil
}

func (c *Client) Write(_ context.Context, tableName string, data map[string]interface{}) error {
	table, found := c.tables[tableName]
	if !found {
		return fmt.Errorf("table %v not found", tableName)
	}

	w := c.csvWriters[tableName]
	record := make([]string, len(table.Columns))
	for i, col := range table.Columns {
		v, err := valueToString(col.Type, data[col.Name])
		if err != nil {
			return fmt.Errorf("failed to convert value %v to string for col type %v: %w", data[col.Name], col.Type, err)
		}
		record[i] = v
	}
	if err := w.Write(record); err != nil {
		return fmt.Errorf("error writing record: %w", err)
	}

	return nil
}

func valueToString(t schema.ValueType, v interface{}) (string, error) {
	switch x := v.(type) {
	case bool:
		if x {
			return "true", nil
		}
		return "false", nil
	case int, int32, int64, uint, uint32, uint64:
		return fmt.Sprintf("%d", x), nil
	case float32, float64:
		if t == schema.TypeInt {
			return fmt.Sprintf("%d", x), nil
		}
		return fmt.Sprintf("%v", x), nil
	case string:
		return x, nil
	default:
		b, err := json.Marshal(v)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
}
