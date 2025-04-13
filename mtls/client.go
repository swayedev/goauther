package mtls

import (
	"crypto/tls"
	"net/http"
)

// Client is a wrapper around http.Client with mTLS support.
type Client struct {
	http.Client
}

// NewClient creates a new mTLS-enabled client with optional configurations.
func NewClient(config *Config, options ...func(*http.Client)) (*Client, error) {
	tlsConfig, err := config.GetTLSConfig(tls.NoClientCert)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// Apply custom options to the client
	for _, option := range options {
		option(client)
	}

	return &Client{*client}, nil
}
