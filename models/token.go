package oauth

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
