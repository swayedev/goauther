package oauth

import "github.com/golang-jwt/jwt/v5"

// create client

// create token
func CreateToken() (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "admin",
		"password": "admin",
	})

	token, err := claims.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return token, nil
}

// parse token
func ParseToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
}
