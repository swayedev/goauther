package server

import (
	"crypto/ed25519"
	"crypto/rsa"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func (s Server) GenerateAccessToken(data *AccessData, generaterefresh bool) (accesstoken string, refreshtoken string, err error) {
	switch s.Certificate.GetType() {
	case "rsa":
		return s.GenerateAccessTokenWithRsa(data, generaterefresh)
	case "ed25519":
		return s.GenerateAccessTokenWithEd25519(data, generaterefresh)
	}
	return "", "", errors.New(E_INVALID_CERTIFICATE)
}

func (s Server) GenerateAccessTokenWithRsa(data *AccessData, generaterefresh bool) (accesstoken string, refreshtoken string, err error) {
	privKey, err := s.Certificate.ParsePrivateKey()
	if err != nil {
		return "", "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"cid": data.Client.GetId(),
		"exp": data.ExpireAt().Unix(),
	})

	accesstoken, err = token.SignedString(privKey.(rsa.PrivateKey))
	if err != nil {
		return "", "", err
	}

	if !generaterefresh {
		return
	}

	// generate JWT refresh token
	token = jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"cid": data.Client.GetId(),
	})

	refreshtoken, err = token.SignedString(privKey.(rsa.PrivateKey))
	if err != nil {
		return "", "", err
	}
	return
}

func (s Server) GenerateAccessTokenWithEd25519(data *AccessData, generaterefresh bool) (accesstoken string, refreshtoken string, err error) {
	privKey, err := s.Certificate.ParsePrivateKey()
	if err != nil {
		return "", "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, jwt.MapClaims{
		"cid": data.Client.GetId(),
		"exp": data.ExpireAt().Unix(),
	})

	accesstoken, err = token.SignedString(privKey.(ed25519.PrivateKey))
	if err != nil {
		return "", "", err
	}

	if !generaterefresh {
		return
	}

	// generate JWT refresh token
	token = jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"cid": data.Client.GetId(),
	})

	refreshtoken, err = token.SignedString(privKey.(ed25519.PrivateKey))
	if err != nil {
		return "", "", err
	}
	return
}
