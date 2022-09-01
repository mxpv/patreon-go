package patreon

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetchVanity(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/oauth2/v2/identity", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, currentyVanityResp)
	})

	resp, err := client.FetchIdentity()
	require.NoError(t, err)
	require.Equal(t, "https://www.patreon.com/api/oauth2/v2/user/3232132131", resp.Links.Self)
	require.Equal(t, "user", resp.Data.Type)
	require.Equal(t, "3232132131", resp.Data.ID)

	// Attributes

	attrs := resp.Data.Attributes
	require.Equal(t, "Super cool", attrs.About)
	require.True(t, attrs.CanSeeNSFW)
	// require.Equal(t, "2022-05-05T18:33:14.000+00:00", attrs.Created)
	require.Equal(t, "austin@gmail.com", attrs.Email)
	require.Equal(t, "Austin", attrs.FirstName)
	require.Equal(t, "Austin Austin", attrs.FullName)
	require.True(t, attrs.HidePledges)
	require.Equal(t, "https://c10.patreonusercontent.com/4/patreon-media/p/user/3232132131", attrs.ImageURL)
	require.True(t, attrs.IsEmailVerified)
	require.Equal(t, "Austin", attrs.LastName)
	require.Equal(t, 0, attrs.LikeCount)
	require.Equal(t, "https://c10.patreonusercontent.com/4/patreon-media/p/user/3232132131", attrs.ThumbURL)
	require.Equal(t, "https://www.patreon.com/user?u=3232132131", attrs.URL)
	require.Equal(t, "yurt", attrs.Vanity)

	// Relationships

	// pledges := resp.Data.Relationships.Pledges
	// require.NotNil(t, pledges)
	// require.Len(t, pledges.Data, 1)
	// require.Equal(t, "2444714", pledges.Data[0].ID)
	// require.Equal(t, "pledge", pledges.Data[0].Type)
}

const currentyVanityResp = `
{
    "data": {
      "attributes": {
        "about": "Super cool",
        "can_see_nsfw": true,
        "created": "2022-05-05T18:33:14.000+00:00",
        "email": "austin@gmail.com",
        "first_name": "Austin",
        "full_name": "Austin Austin",
        "hide_pledges": true,
        "image_url": "https://c10.patreonusercontent.com/4/patreon-media/p/user/3232132131",
        "is_email_verified": true,
        "last_name": "Austin",
        "like_count":0,
        "social_connections": {
          "deviantart": null,
          "discord": null,
          "facebook": null,
          "google": null,
          "instagram": null,
          "reddit": null,
          "spotify": null,
          "twitch": null,
          "twitter": null,
          "vimeo": null,
          "youtube": {
            "scopes": ["https://www.googleapis.com/auth/youtube.readonly"],
            "url": "https://youtube.com/channel/UC0qMcx_Yg_oiN5dVllNk71w",
            "user_id": "UC0qMcx_Yg_3232132131"
          }
        },
        "thumb_url": "https://c10.patreonusercontent.com/4/patreon-media/p/user/3232132131",
        "url": "https://www.patreon.com/user?u=3232132131",
        "vanity": "yurt"
      },
      "id": "3232132131",
      "type": "user"
    },
    "links": { "self": "https://www.patreon.com/api/oauth2/v2/user/3232132131" }
  }
`
