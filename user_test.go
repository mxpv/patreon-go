package patreon

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetchUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/oauth2/api/current_user", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, currentUserResp)
	})

	resp, err := client.FetchUser()
	require.NoError(t, err)
	require.Equal(t, "https://www.patreon.com/api/user/3232132131", resp.Links.Self)
	require.Equal(t, "user", resp.Data.Type)
	require.Equal(t, "3232132131", resp.Data.ID)

	// Attributes

	attrs := resp.Data.Attributes
	require.Equal(t, "max@gmail.com", attrs.Email)
	require.Equal(t, "max", attrs.Facebook)
	require.Equal(t, "1312321312", attrs.FacebookId)
	require.Equal(t, "Max", attrs.FirstName)
	require.Equal(t, "Max", attrs.LastName)
	require.Equal(t, "Max", attrs.FullName)
	require.Equal(t, 1, attrs.Gender)
	require.True(t, attrs.HasPassword)
	require.Equal(t, "https://c8.patreon.com/2/400/3232132131", attrs.ImageURL)
	require.Equal(t, "https://c8.patreon.com/2/100/3232132131", attrs.ThumbURL)
	require.True(t, attrs.IsDeleted)
	require.True(t, attrs.IsEmailVerified)
	require.True(t, attrs.IsNuked)
	require.True(t, attrs.IsSuspended)
	require.Equal(t, "pod_sync", attrs.Twitter)
	require.Equal(t, "https://www.patreon.com/podsync", attrs.URL)
	require.Equal(t, "podsync", attrs.Vanity)

	// Relationships

	pledges := resp.Data.Relationships.Pledges
	require.NotNil(t, pledges)
	require.Len(t, pledges.Data, 1)
	require.Equal(t, "2444714", pledges.Data[0].ID)
	require.Equal(t, "pledge", pledges.Data[0].Type)
}

const currentUserResp = `
{
    "data": {
        "attributes": {
            "about": "",
            "created": "2016-02-02T19:56:14+00:00",
            "discord_id": null,
            "email": "max@gmail.com",
            "facebook": "max",
            "facebook_id": "1312321312",
            "first_name": "Max",
            "full_name": "Max",
            "gender": 1,
            "has_password": true,
            "image_url": "https://c8.patreon.com/2/400/3232132131",
            "is_deleted": true,
            "is_email_verified": true,
            "is_nuked": true,
            "is_suspended": true,
            "last_name": "Max",
            "social_connections": {
                "deviantart": null,
                "discord": null,
                "facebook": null,
                "spotify": null,
                "twitch": null,
                "twitter": null,
                "youtube": null
            },
            "thumb_url": "https://c8.patreon.com/2/100/3232132131",
            "twitch": null,
            "twitter": "pod_sync",
            "url": "https://www.patreon.com/podsync",
            "vanity": "podsync",
            "youtube": null
        },
        "id": "3232132131",
        "relationships": {
            "pledges": {
                "data": [
                	{
                    	"id": "2444714",
                        "type": "pledge"
                    }
                ]
            }
        },
        "type": "user"
    },
    "links": {
        "self": "https://www.patreon.com/api/user/3232132131"
    }
}
`
