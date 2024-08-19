package mtls

import (
	"crypto/tls"
	"crypto/x509"
	"os"
)

// Config holds the configuration for mTLS.
type Config struct {
	CertFile string // Path to the certificate file
	KeyFile  string // Path to the private key file
	CAFile   string // Path to the CA certificate file
}

// GetTLSConfig generates a tls.Config based on the provided configuration.
func (c *Config) GetTLSConfig(clientAuth tls.ClientAuthType) (*tls.Config, error) {
	// Load server's certificate and private key
	cert, err := tls.LoadX509KeyPair(c.CertFile, c.KeyFile)
	if err != nil {
		return nil, err
	}

	// Load CA certificate
	caCert, err := os.ReadFile(c.CAFile)
	if err != nil {
		return nil, err
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
