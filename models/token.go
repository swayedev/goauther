package models

import "time"

type AccessToken struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	ClientId  string    `json:"client_id"`
	Name      string    `json:"name"`
	Scopes    []string  `json:"scope"`
	Revoked   bool      `json:"revoked"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

type RefreshToken struct {
	Id            string    `json:"id"`
	AccessTokenId string    `json:"access_token_id"`
	Revoked       bool      `json:"revoked"`
	ExpiresAt     time.Time `json:"expires_at"`
}

GenerateAccessToken(data *server.AccessData, generaterefresh bool) (accesstoken string, refreshtoken string, err error) {
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