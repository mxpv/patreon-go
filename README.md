[![Build Status](https://travis-ci.org/mxpv/patreon-go.svg?branch=master)](https://travis-ci.org/mxpv/patreon-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/mxpv/patreon-go)](https://goreportcard.com/report/github.com/mxpv/patreon-go)
[![codecov](https://codecov.io/gh/mxpv/patreon-go/branch/master/graph/badge.svg)](https://codecov.io/gh/mxpv/patreon-go)
[![MIT license](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)

# patreon-go

`patreon-go` is a Go client library for accessing the [Patreon API](https://www.patreon.com/platform/documentation/api).

## Basic example ##

```go
import "github.com/mxpv/patreon-go"

func main() {
	client := patreon.NewClient(nil)
  
	user, err := client.FetchUser()
	if err != nil {
		// ...
	}

	print(user.Data.Id)
}
```

## Authentication ##

The `patreon-go` library does not directly handle authentication. Instead, when creating a new client, pass an `http.Client` that can handle authentication for you, most likely you will need [oauth2](https://github.com/golang/oauth2) package.

Here is an example with static token:

```go
import (
	"github.com/mxpv/patreon-go"
	"golang.org/x/oauth2"
)

func NewPatreonClient(ctx context.Context, token string) *patreon.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	
	client := patreon.NewClient(tc)
	return client
}
```

With username and password:

```go
func NewPatreonClient(ctx context.Context, username, password string) (*patreon.Client, error) {
	config := oauth2.Config{
		ClientID:     "<client_id>",
		ClientSecret: "<clinet_secret>",
		Endpoint: oauth2.Endpoint{
			AuthURL:  patreon.AuthorizationURL,
			TokenURL: patreon.AccessTokenURL,
		},
		RedirectURL: "<redirect_url>",
		Scopes: []string{
			"users",
			"pledges-to-me",
			"my-campaign",
		},
	}
	
	token, err := config.PasswordCredentialsToken(ctx, username, password)
	if err != nil {
		return nil, err
	}

	client := config.Client(ctx, token)
	return patreon.NewClient(client), nil
}
```
