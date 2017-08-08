package patreon

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/oauth2"
)

var testAccessToken = os.Getenv("PATREON_ACCESS_TOKEN")

// Fetches the list of pledges with corresponding users.
// This example is a port of PHP version https://github.com/Patreon/patreon-php/blob/master/examples/patron-list.php
func Example_fetchPatronsAndPledges() {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: testAccessToken})
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	// Create client with static access token
	client := NewClient(tc)

	// Get your campaign data
	campaignResponse, err := client.FetchCampaign()
	if err != nil {
		return
	}

	campaignId := campaignResponse.Data[0].Id

	cursor := ""
	page := 1

	for {
		pledgesResponse, err := client.FetchPledges(campaignId,
			WithPageSize(25),
			WithCursor(cursor))

		if err != nil {
			panic(err)
			return
		}

		// Get all the users in an easy-to-lookup way
		users := make(map[string]*User)
		for _, item := range pledgesResponse.Included.Items {
			u, ok := item.(*User)
			if !ok {
				continue
			}

			users[u.Id] = u
		}

		fmt.Printf("Page %d\r\n", page)

		// Loop over the pledges to get e.g. their amount and user name
		for _, pledge := range pledgesResponse.Data {
			amount := pledge.Attributes.AmountCents
			patronId := pledge.Relationships.Patron.Data.Id
			patronFullName := users[patronId].Attributes.FullName

			fmt.Printf("%s is pledging %d cents\r\n", patronFullName, amount)
		}

		// Get the link to the next page of pledges
		nextLink := pledgesResponse.Links.Next
		if nextLink == "" {
			break
		}

		cursor = nextLink
		page++
	}

	fmt.Print("Done!")
}

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
		RefreshToken: "<current_refresh_token>",
		// Must be non-nil, otherwise token will not be expired
		Expiry: time.Now().Add(-24 * time.Hour),
	}

	tc := config.Client(oauth2.NoContext, &token)

	client := NewClient(tc)
	_, err := client.FetchUser()
	if err != nil {
		panic(err)
	}

	print("OK")
}
