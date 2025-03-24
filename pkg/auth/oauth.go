package auth

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

// AuthCodeURL returns a URL to the OAuth 2.0 provider's consent page
func AuthCodeURL(state string, config *oauth2.Config) string {
	return config.AuthCodeURL(state)
}

// Exchange converts an authorization code into a token
func Exchange(ctx context.Context, code string, config *oauth2.Config) (*oauth2.Token, error) {
	return config.Exchange(ctx, code)
}

// Client returns an HTTP client using the specified token
func Client(ctx context.Context, token *oauth2.Token, config *oauth2.Config) *http.Client {
	return config.Client(ctx, token)
}
