<!-- [![GoDoc](https://godoc.org/github.com/mxpv/patreon-go?status.svg)](https://godoc.org/github.com/mxpv/patreon-go/) -->

[![MIT license](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)
[![Patreon](https://img.shields.io/badge/support-patreon-E6461A.svg)](https://www.patreon.com/austinhub)

# patreon-go-wrapper

`patreon-go-wrapper` is a Go client library for accessing the [Patreon API V2](https://docs.patreon.com/#api).

> Forked from [patreon-go](https://github.com/mxpv/patreon-go) which has a great implementation for Patreon API V1

## How to import

The `patreon-go` package may be installed by running:

```
go get github.com/battle-of-the-states/patreon-go-wrapper
```

or

```
import "github.com/battle-of-the-states/patreon-go-wrapper"
```

## Basic example

```go
import "github.com/battle-of-the-states/patreon-go-wrapper"

func main() {
	client := patreon.NewClient(nil)

	user, err := client.FetchIdentity()
	if err != nil {
		// ...
	}

	fmt.Println(user.Data.Id)
}
```

## Authentication

The `patreon-go-wrapper` library does not directly handle authentication. Instead, when creating a new client, pass an `http.Client` that can handle authentication for you, most likely you will need [oauth2](https://github.com/golang/oauth2) package.

Here is an example with static token:

```go
import (
	"gopkg.in/mxpv/patreon-go.v1"
	"golang.org/x/oauth2"
)

func NewPatreonClient(ctx context.Context, token string) *patreon.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)

	client := patreon.NewClient(tc)
	return client
}
```

Automatically refresh token:

> Check the available scopes in the [Patreon Docs](https://docs.patreon.com/#note-to-those-with-v1-tokens)

```go
func NewPatreonClient() (*patreon.Client, error) {
	config := oauth2.Config{
		ClientID:     "<client_id>",
		ClientSecret: "<client_secret>",
		Endpoint: oauth2.Endpoint{
			AuthURL:  AuthorizationURL,
			TokenURL: AccessTokenURL,
		},
		Scopes: []string{"campaigns.members", "campaigns.members[email]", "campaigns.members.address", "campaigns", "identity.memberships", "identity", "identity[email]"},
	}

	token := oauth2.Token{
		AccessToken:  "<current_access_token>",
		RefreshToken: "<current_refresh_token>",
		// Must be non-nil, otherwise token will not be expired
		Expiry: time.Now().Add(-24 * time.Hour),
	}

	tc := config.Client(context.Background(), &token)

	client := NewClient(tc)
	_, err := client.FetchUser()

	return client, err
}
```

## Look & Feel

```go
func main() {
	patreonConfig := oauth2.Config{
		ClientID:     config("PATREON_CLIENT_ID"),
		ClientSecret: config("PATREON_CLIENT_SECRET"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  patreon.AuthorizationURL,
			TokenURL: patreon.AccessTokenURL,
		},
		Scopes: []string{"campaigns.members", "campaigns.members[email]", "campaigns.members.address", "campaigns", "identity.memberships", "identity", "identity[email]"},
	}

	token := oauth2.Token{
		AccessToken:  config("PATREON_ACCESS_TOKEN"),
		RefreshToken: config("PATREON_REFRESH_TOKEN"),
		// Must be non-nil, otherwise token will not be expired
		Expiry: time.Now().Add(2 * time.Hour),
	}

	tc := patreonConfig.Client(context.Background(), &token)

	client := patreon.NewClient(tc)

	fieldOpts := patreon.WithFields()
	includeOpts := patreon.WithIncludes("campaign", "memberships")

	user, err := client.FetchIdentity(fieldOpts, includeOpts)
	if err != nil {
		panic(err)
	}

	fmt.Println(user.Data.ID)
}
```
