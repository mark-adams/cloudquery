package client

import (
	"errors"

	"github.com/himalayan-institute/zoom-lib-golang"
)

type Config struct {
	JWT *jwtConfig `yaml:"jwt"`
}

func (Config) Example() string {
	return `
		jwt:
			api_key:    <API Key HERE>
			api_secret: <API Secret HERE>
`
}

// The use of these credentials will expire on January 06, 2023.
type jwtConfig struct {
	APIKey    string `yaml:"api_key"`
	APISecret string `yaml:"api_secret"`
}

func (c jwtConfig) Client() (*zoom.Client, error) {
	if c.APIKey == "" {
		return nil, errors.New("missing API Key in configuration")
	}
	if c.APISecret == "" {
		return nil, errors.New("missing API Secret in configuration")
	}

	return zoom.NewClient(c.APIKey, c.APISecret), nil
}
