package mtls

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
)

// Config holds the configuration for mTLS.
type Config struct {
	CertFile string // Path to the certificate file
	KeyFile  string // Path to the private key file
	CAFile   string // Path to the CA certificate file
}

// LoadFromEnv allows loading config values from environment variables.
func (c *Config) LoadFromEnv() {
	if certPath := os.Getenv("MTLS_CERT_FILE"); certPath != "" {
		c.CertFile = certPath
	}
	if keyPath := os.Getenv("MTLS_KEY_FILE"); keyPath != "" {
		c.KeyFile = keyPath
	}
	if caPath := os.Getenv("MTLS_CA_FILE"); caPath != "" {
		c.CAFile = caPath
	}
}

// Validate checks that all necessary fields are set.
func (c *Config) Validate() error {
	if c.CertFile == "" || c.KeyFile == "" || c.CAFile == "" {
		return fmt.Errorf("certificate, key, and CA file paths must be provided")
	}
	return nil
}

// GetTLSConfig generates a tls.Config based on the provided configuration.
func (c *Config) GetTLSConfig(clientAuth tls.ClientAuthType) (*tls.Config, error) {
	// Validate the config
	if err := c.Validate(); err != nil {
		return nil, WrapConfigError("configuration validation failed", err)
	}

	// Load server's certificate and private key
	cert, err := tls.LoadX509KeyPair(c.CertFile, c.KeyFile)
	if err != nil {
		return nil, WrapConfigError("failed to load certificate", err)
	}

	// Load CA certificate
	caCert, err := os.ReadFile(c.CAFile)
	if err != nil {
		return nil, WrapConfigError("failed to load CA certificate", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup TLS configuration
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caCertPool,
		ClientAuth:   clientAuth,
	}

	return tlsConfig, nil
}

// ReloadCertificates reloads certificates from files.
func (c *Config) ReloadCertificates() (*tls.Config, error) {
	return c.GetTLSConfig(tls.RequireAndVerifyClientCert)
}

// MTLSConfigError is a custom error type for mTLS configuration errors.
type MTLSConfigError struct {
	Msg string
	Err error
}

func (e *MTLSConfigError) Error() string {
	return fmt.Sprintf("mTLS Config Error: %s - %v", e.Msg, e.Err)
}

func WrapConfigError(msg string, err error) error {
	return &MTLSConfigError{Msg: msg, Err: err}
}
