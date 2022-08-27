package client

import (
	"errors"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	"github.com/zoom-lib-golang/zoom-lib-golang"
)

type Client struct {
	// This is a client that you need to create and initialize in Configure
	// It will be passed for each resource fetcher.
	logger hclog.Logger

	Zoom *zoom.Client
}

func (c *Client) Logger() hclog.Logger {
	return c.logger
}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, diag.Diagnostics) {
	providerConfig := config.(*Config)
	// validate provider config

	// TODO: add OAuth option as well
	if providerConfig.JWT == nil {
		return nil, diag.FromError(errors.New("missing authentication section `jwt` in configuration"), diag.ACCESS)
	}

	client, err := providerConfig.JWT.Client()
	if err != nil {
		return nil, diag.FromError(err, diag.ACCESS)
	}

	// Init your client and 3rd party clients using the user's configuration
	// passed by the SDK providerConfig
	return &Client{
		logger: logger,
		Zoom:   client,
	}, nil
}
