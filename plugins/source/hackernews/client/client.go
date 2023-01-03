package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/hackernews/client/services"
	"github.com/cloudquery/plugin-sdk/backend"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/hermanschaaf/hackernews"
	"github.com/rs/zerolog"
)

const (
	defaultMaxRetries = 5
	defaultBackoff    = 10 * time.Second
)

type Client struct {
	logger     zerolog.Logger
	sourceSpec specs.Source
	HackerNews services.HackernewsClient
	Spec       Spec
	Backend    backend.Backend
	maxRetries int
	backoff    time.Duration // backoff duration between retries (jitter will be added)
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.sourceSpec.Name
}

func Configure(ctx context.Context, logger zerolog.Logger, sourceSpec specs.Source, backend backend.Backend) (schema.ClientMeta, error) {
	var config Spec
	err := sourceSpec.UnmarshalSpec(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	config.SetDefaults()

	client := hackernews.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create hackernews client: %w", err)
	}

	return &Client{
		logger:     logger,
		sourceSpec: sourceSpec,
		Spec:       config,
		HackerNews: client,
		maxRetries: defaultMaxRetries,
		backoff:    defaultBackoff,
		Backend:    backend,
	}, nil
}