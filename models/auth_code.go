package oauth

import "time"

type AuthCode struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	ClientId  string    `json:"client_id"`
	Scopes    []string  `json:"scopes"`
	Revoked   bool      `json:"revoked"`
	ExpiresAt time.Time `json:"expires_at"`
}
