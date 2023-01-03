package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"

	mssql "github.com/microsoft/go-mssqldb"
)

type Client struct {
	conn mssql.Conn

	metrics destination.Metrics

	destination.UnimplementedUnmanagedWriter
	destination.DefaultReverseTransformer
}

func (c *Client) Close(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

var _ destination.Client = (*Client)(nil)

func New(ctx context.Context, logger zerolog.Logger, destSpec specs.Destination) (destination.Client, error) {

	return nil, nil
}
