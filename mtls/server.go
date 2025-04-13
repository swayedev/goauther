package mtls

import (
	"context"
	"crypto/tls"
	"net/http"
)

// Server is a wrapper around http.Server with mTLS support.
type Server struct {
	http.Server
	Config *Config
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

	return &Server{*server, config}, nil
}

// ListenAndServeTLS starts the mTLS server.
func (s *Server) ListenAndServeTLS() error {
	return s.Server.ListenAndServeTLS("", "")
}

// ReloadCerts reloads the server's certificates without restarting.
func (s *Server) ReloadCerts() error {
	newConfig, err := s.Config.ReloadCertificates()
	if err != nil {
		return err
	}
	s.TLSConfig = newConfig
	return nil
}

// GracefulShutdown gracefully shuts down the server.
func (s *Server) GracefulShutdown(ctx context.Context) error {
	return s.Shutdown(ctx)
}
