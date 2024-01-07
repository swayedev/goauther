package server

import (
	"errors"
	"net/http"
	"time"
)

// Server is an OAuth2 implementation
type Server struct {
	Config            *ServerConfig
	Storage           Storage
	AuthorizeTokenGen AuthorizeTokenGen
	AccessTokenGen    AccessTokenGen
	Now               func() time.Time
	Logger            Logger
}

// NewServer creates a new server instance
func NewServer(config *ServerConfig, storage Storage) *Server {
	return &Server{
		Config:            config,
		Storage:           storage,
		AuthorizeTokenGen: &AuthorizeTokenGenDefault{},
		AccessTokenGen:    &AccessTokenGenDefault{},
		Now:               time.Now,
		Logger:            &LoggerDefault{},
	}
}

// NewResponse creates a new response for the server
func (s *Server) NewResponse() *Response {
	r := NewResponse(s.Storage)
	r.ErrorStatusCode = s.Config.ErrorStatusCode
	return r
}

// getClientAuth checks client basic authentication in params if allowed,
// otherwise gets it from the header.
// Sets an error on the response if no auth is present or a server error occurs.
func (s Server) getClientAuth(w *Response, r *http.Request, allowQueryParams bool) *BasicAuth {

	if allowQueryParams {
		// Allow for auth without password
		if _, hasSecret := r.Form["client_secret"]; hasSecret {
			auth := &BasicAuth{
				Username: r.FormValue("client_id"),
				Password: r.FormValue("client_secret"),
			}
			if auth.Username != "" {
				return auth
			}
		}
	}

	auth, err := CheckBasicAuth(r)
	if err != nil {
		s.setErrorAndLog(w, E_INVALID_REQUEST, err, "get_client_auth=%s", "check auth error")
		return nil
	}
	if auth == nil {
		s.setErrorAndLog(w, E_INVALID_REQUEST, errors.New("Client authentication not sent"), "get_client_auth=%s", "client authentication not sent")
		return nil
	}
	return auth
}
