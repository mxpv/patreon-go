package patreon

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetchCampaign(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/oauth2/v2/campaigns", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, fetchCampaignResp)
	})

	resp, err := client.FetchCampaigns()
	if err != nil {
		panic(err)
	}
	require.NoError(t, err)

	require.Equal(t, 1, len(resp.Data))
	require.Equal(t, "campaign", resp.Data[0].Type)
	require.Equal(t, "8636299", resp.Data[0].ID)
	require.Equal(t, "outstanding coding projects", resp.Data[0].Attributes.CreationName)

	// Attributes

	attrs := resp.Data[0].Attributes
	require.NotEmpty(t, attrs.ImageSmallURL)
	require.NotEmpty(t, attrs.ImageURL)
	// require.True(t, attrs.IsChargedImmediately)
	// require.True(t, attrs.IsMonthly)
	require.False(t, attrs.IsNsfw)
	require.Equal(t, 123121, attrs.PatronCount)
	require.Equal(t, "month", attrs.PayPerName)
	require.NotEmpty(t, attrs.Summary)
	require.NotEmpty(t, attrs.PledgeURL)
	require.NotEmpty(t, attrs.ThanksMsg)

	// Relationships

	creator := resp.Data[0].Relationships.Creator
	require.NotNil(t, creator)
	require.Equal(t, "2343242423", creator.Data.ID)
	require.Equal(t, "user", creator.Data.Type)
	require.Equal(t, "https://www.patreon.com/api/oauth2/v2/user/2343242423", creator.Links.Related)

	// Includes

	user, ok := resp.Included.Items[0].(*User)
	require.True(t, ok)
	require.Equal(t, "2343242423", user.ID)
	require.Equal(t, "user", user.Type)

	tier, ok := resp.Included.Items[1].(*Tier)
	require.True(t, ok)
	require.Equal(t, "8606545", tier.ID)
	require.Equal(t, "tier", tier.Type)

}

const fetchCampaignResp = `
{
    "data": [
      {
        "attributes": {
          "created_at": "2022-05-05T18:33:45.000+00:00",
          "creation_name": "outstanding coding projects",
          "discord_server_id": null,
          "google_analytics_id": null,
          "has_rss": false,
          "has_sent_rss_notify": false,
          "image_small_url": "https://c10.patreonusercontent.com/4/patreon-media/p/campaign/8636299/1035867e95234da7b561610c47ca7ed4/eyJ3IjoxOTIwLCJ3ZSI6MX0%3D/1.jpg?token-time=1664841600&token-hash=OZAH1tyGDklRD4891PGP0ULGoSs3KeecPeqtMJC96qs%3D",
          "image_url": "https://c10.patreonusercontent.com/4/patreon-media/p/campaign/8636299/1035867e95234da7b561610c47ca7ed4/eyJ3IjoxOTIwLCJ3ZSI6MX0%3D/1.jpg?token-time=1664841600&token-hash=OZAH1tyGDklRD4891PGP0ULGoSs3KeecPeqtMJC96qs%3D",
          "is_charged_immediately": false,
          "is_monthly": true,
          "is_nsfw": false,
          "main_video_embed": null,
          "main_video_url": null,
          "one_liner": null,
          "patron_count": 123121,
          "pay_per_name": "month",
          "pledge_url": "/join/austinhub",
          "published_at": "2022-08-31T22:43:52.000+00:00",
          "rss_artwork_url": null,
          "rss_feed_title": null,
          "summary": "Austin Hub is a great source for challenging and rewarding software development tutorials. Join Austin Hub on the journey to creating products.",
          "thanks_embed": null,
          "thanks_msg": "Thank you!",
          "thanks_video_url": null
        },
        "id": "8636299",
        "relationships": {
          "benefits": {
            "data": [
              { "id": "10456246", "type": "benefit" },
              { "id": "10456346", "type": "benefit" },
              { "id": "10456319", "type": "benefit" },
              { "id": "10456374", "type": "benefit" },
              { "id": "10456244", "type": "benefit" },
              { "id": "10456402", "type": "benefit" }
            ]
          },
          "creator": {
            "data": { "id": "2343242423", "type": "user" },
            "links": {
              "related": "https://www.patreon.com/api/oauth2/v2/user/2343242423"
            }
          },
          "goals": { "data": [] },
          "tiers": {
            "data": [
              { "id": "8606545", "type": "tier" },
              { "id": "8606546", "type": "tier" },
              { "id": "8606547", "type": "tier" }
            ]
          }
        },
        "type": "campaign"
      }
    ],
    "included": [
      { "attributes": {}, "id": "2343242423", "type": "user" },
      { "attributes": {}, "id": "8606545", "type": "tier" },
      { "attributes": {}, "id": "8606546", "type": "tier" },
      { "attributes": {}, "id": "8606547", "type": "tier" },
      { "attributes": {}, "id": "10456246", "type": "benefit" },
      { "attributes": {}, "id": "10456346", "type": "benefit" },
      { "attributes": {}, "id": "10456319", "type": "benefit" },
      { "attributes": {}, "id": "10456374", "type": "benefit" },
      { "attributes": {}, "id": "10456244", "type": "benefit" },
      { "attributes": {}, "id": "10456402", "type": "benefit" }
    ],
    "meta": { "pagination": { "total": 1 } }
  }
  
`
