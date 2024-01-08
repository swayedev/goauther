package token

import (
	"crypto/ecdsa"

	"github.com/golang-jwt/jwt/v5"
)

// create client

// create token
func CreateToken(key *ecdsa.PrivateKey) (string, error) {
	//   key = /* Load key from somewhere, for example a file */
	t := jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			"iss":     "my-auth-server",
			"sub":     "john",
			"expires": 2,
		})
	return t.SignedString(key)
}

// parse token
func ParseToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
}
