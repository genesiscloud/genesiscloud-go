package genesiscloud

import (
	"fmt"

	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
)

const DefaultEndpoint = "https://api.genesiscloud.com/compute/v1"

type ClientConfig struct {
	Endpoint string
	Token    string
}

func NewGenesisCloudClient(config ClientConfig, opts ...ClientOption) (*ClientWithResponses, error) {
	if config.Endpoint == "" {
		config.Endpoint = DefaultEndpoint
	}

	if config.Token == "" {
		return nil, fmt.Errorf("ClientOptions.Token is required")
	}

	apiKeyProvider, err := securityprovider.NewSecurityProviderApiKey("header", "X-Auth-Token", config.Token)
	if err != nil {
		return nil, err
	}

	client, err := NewClientWithResponses(config.Endpoint, append(opts, WithRequestEditorFn(apiKeyProvider.Intercept))...)
	if err != nil {
		return nil, err
	}

	return client, nil
}
