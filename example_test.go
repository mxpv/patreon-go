package patreon

import (
	"context"
	"time"

	"golang.org/x/oauth2"
)

// Automatically refresh token
func Example_refreshToken() {
	config := oauth2.Config{
		ClientID:     "<client_id>",
		ClientSecret: "<client_secret>",
		Endpoint: oauth2.Endpoint{
			AuthURL:  AuthorizationURL,
			TokenURL: AccessTokenURL,
		},
		Scopes: []string{"users", "pledges-to-me", "my-campaign"},
	}

	token := oauth2.Token{
		AccessToken:  "<current_access_token>",
		RefreshToken: "<current_refresh_token>",
		// Must be non-nil, otherwise token will not be expired
		Expiry: time.Now().Add(-24 * time.Hour),
	}

	tc := config.Client(context.Background(), &token)

	client := NewClient(tc)
	_, err := client.FetchIdentity()
	if err != nil {
		panic(err)
	}

	print("OK")
}
