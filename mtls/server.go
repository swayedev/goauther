package mtls

import (
	"crypto/tls"
	"net/http"
)

// Server is a wrapper around http.Server with mTLS support.
type Server struct {
	http.Server
}

// NewServer creates a new mTLS-enabled server.
func NewServer(addr string, config *Config, handler http.Handler) (*Server, error) {
	tlsConfig, err := config.GetTLSConfig(tls.RequireAndVerifyClientCert)
	if err != nil {
		return nil, err
	}

	server := &http.Server{
		Addr:      addr,
		Handler:   handler,
		TLSConfig: tlsConfig,
	}

	return &Server{*server}, nil
}

// ListenAndServeTLS starts the mTLS server.
func (s *Server) ListenAndServeTLS() error {
	return s.Server.ListenAndServeTLS("", "")
}
