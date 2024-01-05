package server

import (
	"encoding/base64"

	"github.com/google/uuid"
)

// AuthorizeTokenGenDefault is the default authorization token generator
type AuthorizeTokenGenDefault struct {
}

// GenerateAuthorizeToken generates a base64-encoded UUID code
func (a *AuthorizeTokenGenDefault) GenerateAuthorizeToken(data *AuthorizeData) (ret string, err error) {
	token, _ := uuid.NewRandom()
	return base64.RawURLEncoding.EncodeToString(token[:]), nil
}

// AccessTokenGenDefault is the default authorization token generator
type AccessTokenGenDefault struct {
}

// GenerateAccessToken generates base64-encoded UUID access and refresh tokens
func (a *AccessTokenGenDefault) GenerateAccessToken(data *AccessData, generaterefresh bool) (accesstoken string, refreshtoken string, err error) {
	token, _ := uuid.NewRandom()
	accesstoken = base64.RawURLEncoding.EncodeToString(token[:])

	if generaterefresh {
		rtoken, _ := uuid.NewRandom()
		refreshtoken = base64.RawURLEncoding.EncodeToString(rtoken[:])
	}
	return
}
