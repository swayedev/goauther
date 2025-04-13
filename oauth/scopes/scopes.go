package scopes

type OAuthScopes struct {
	Profile string `json:"profile"`
	Email   string `json:"email"`
	Openid  string `json:"openid"`
	Read    string `json:"read"`
	Write   string `json:"write"`
	Admin   string `json:"admin"`
}

func (o *OAuthScopes) SetProfile(profile string) {
	o.Profile = profile
}

func (o *OAuthScopes) SetEmail(email string) {
	o.Email = email
}

func (o *OAuthScopes) SetOpenid(openid string) {
	o.Openid = openid
}

func (o *OAuthScopes) SetRead(read string) {
	o.Read = read
}

func (o *OAuthScopes) SetWrite(write string) {
	o.Write = write
}

func (o *OAuthScopes) SetAdmin(admin string) {
	o.Admin = admin
}
