package mtls

import (
	"crypto/tls"
	"net/http"
)

// Client is a wrapper around http.Client with mTLS support.
type Client struct {
	http.Client
}

// NewClient creates a new mTLS-enabled client.
func NewClient(config *Config) (*Client, error) {
	tlsConfig, err := config.GetTLSConfig(tls.NoClientCert)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	return &Client{*client}, nil
}
