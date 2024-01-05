package oauth

import "time"

type Client struct {
	ID                   []byte
	UserId               string
	Name                 string
	Secret               string `json:"-"`
	Provider             string
	RedirectURIs         []string
	Revoked              bool
	PersonalAccessClient bool
	PasswordClient       bool
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

func CreateNewClient() {

}
