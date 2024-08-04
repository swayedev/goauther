package provider

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/oauth2"
)

type OAuthContent struct {
	Token   *oauth2.Token
	Content []byte
}

// GetOAuthInfo retrieves the OAuth information from the specified API using the provided token.
// It makes a GET request to the API with the token's access token as the authorization header.
// The response body is read and returned as a byte slice.
// If any error occurs during the process, it is returned along with a nil byte slice.
func GetOAuthInfo(tokenType string, token *oauth2.Token, api string) ([]byte, error) {
	authToken := "token " + token.AccessToken
	if tokenType == "Bearer" {
		authToken = "Bearer " + token.AccessToken
	}

	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err.Error())
	}

	req.Header.Add("Authorization", authToken)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}

// GetOAuthInfoWithToken retrieves OAuth information with the provided token.
// It makes a request to the specified API using the given token and token type.
// It returns the OAuth content and any error encountered during the process.
func GetOAuthInfoWithToken(tokenType string, token *oauth2.Token, api string) (*OAuthContent, error) {
	contents, err := GetOAuthInfo(tokenType, token, api)
	if err != nil {
		return nil, err
	}

	return &OAuthContent{
		Token:   token,
		Content: contents,
	}, nil
}
