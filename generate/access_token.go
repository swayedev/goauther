package generate

import (
	"crypto/rsa"

	"github.com/golang-jwt/jwt/v5"
	"github.com/swayedev/oauth/server"
)

// JWT access token generator
type AccessTokenGenJWT struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func GenerateAccessToken(data *server.AccessData, generaterefresh bool) (accesstoken string, refreshtoken string, err error) {
	// generate JWT access token

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"cid": data.Client.GetId(),
		"exp": data.ExpireAt().Unix(),
	})

	accesstoken, err = token.SignedString(c.PrivateKey)
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

	refreshtoken, err = token.SignedString(c.PrivateKey)
	if err != nil {
		return "", "", err
	}
	return
}

func GenerateAccessToken(data *server.AccessData, generaterefresh bool) (accesstoken string, refreshtoken string, err error) {
	// generate JWT access token

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"cid": data.Client.GetId(),
		"exp": data.ExpireAt().Unix(),
	})

	accesstoken, err = token.SignedString(c.PrivateKey)
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

	refreshtoken, err = token.SignedString(c.PrivateKey)
	if err != nil {
		return "", "", err
	}
	return
}
