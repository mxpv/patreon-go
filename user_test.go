package patreon_go

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
	require.Equal(t, "3232132131", resp.Data.Id)
	require.Equal(t, "max@gmail.com", resp.Data.Attributes.Email)
	require.Equal(t, "max", resp.Data.Attributes.Facebook)
	require.Equal(t, "1312321312", resp.Data.Attributes.FacebookId)
	require.Equal(t, "Max", resp.Data.Attributes.FirstName)
	require.Equal(t, "Max", resp.Data.Attributes.LastName)
	require.Equal(t, "Max", resp.Data.Attributes.FullName)
	require.Equal(t, 1, resp.Data.Attributes.Gender)
	require.True(t, resp.Data.Attributes.HasPassword)
	require.Equal(t, "https://c8.patreon.com/2/400/3232132131", resp.Data.Attributes.ImageURL)
	require.Equal(t, "https://c8.patreon.com/2/100/3232132131", resp.Data.Attributes.ThumbURL)
	require.True(t, resp.Data.Attributes.IsDeleted)
	require.True(t, resp.Data.Attributes.IsEmailVerified)
	require.True(t, resp.Data.Attributes.IsNuked)
	require.True(t, resp.Data.Attributes.IsSuspended)
	require.Equal(t, "pod_sync", resp.Data.Attributes.Twitter)
	require.Equal(t, "https://www.patreon.com/podsync", resp.Data.Attributes.URL)
	require.Equal(t, "podsync", resp.Data.Attributes.Vanity)
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
        "type": "user"
    },
    "links": {
        "self": "https://www.patreon.com/api/user/3232132131"
    }
}
`
