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

	mux.HandleFunc("/oauth2/api/current_user/campaigns", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, fetchCampaignResp)
	})

	resp, err := client.FetchCampaign()
	require.NoError(t, err)

	require.Equal(t, 1, len(resp.Data))
	require.Equal(t, "campaign", resp.Data[0].Type)
	require.Equal(t, "278915", resp.Data[0].ID)
	require.Equal(t, "new podcasting experience - Podsync", resp.Data[0].Attributes.CreationName)

	// Attributes

	attrs := resp.Data[0].Attributes
	require.Equal(t, 8, attrs.CreationCount)
	require.True(t, attrs.DisplayPatronGoals)
	require.NotEmpty(t, attrs.ImageSmallURL)
	require.NotEmpty(t, attrs.ImageURL)
	require.True(t, attrs.IsChargedImmediately)
	require.True(t, attrs.IsMonthly)
	require.True(t, attrs.IsNsfw)
	require.True(t, attrs.IsPlural)
	require.Equal(t, 123121, attrs.PatronCount)
	require.Equal(t, "month", attrs.PayPerName)
	require.Equal(t, 12321312, attrs.PledgeSum)
	require.NotEmpty(t, attrs.Summary)
	require.NotEmpty(t, attrs.PledgeURL)
	require.NotEmpty(t, attrs.ThanksMsg)

	// Relationships

	creator := resp.Data[0].Relationships.Creator
	require.NotNil(t, creator)
	require.Equal(t, "2343242423", creator.Data.ID)
	require.Equal(t, "user", creator.Data.Type)
	require.Equal(t, "https://www.patreon.com/api/user/2343242423", creator.Links.Related)

	categories := resp.Data[0].Relationships.Categories
	require.NotNil(t, categories)
	require.Equal(t, 1, len(categories.Data))
	require.Equal(t, "7", categories.Data[0].ID)
	require.Equal(t, "category", categories.Data[0].Type)

	// Includes

	user, ok := resp.Included.Items[0].(*User)
	require.True(t, ok)
	require.Equal(t, "2822191", user.ID)
	require.Equal(t, "user", user.Type)
	require.Equal(t, "podsync", user.Attributes.Vanity)

	reward, ok := resp.Included.Items[1].(*Reward)
	require.True(t, ok)
	require.Equal(t, "12312312", reward.ID)
	require.Equal(t, "reward", reward.Type)
	require.Equal(t, 100, reward.Attributes.Amount)

	goal, ok := resp.Included.Items[2].(*Goal)
	require.True(t, ok)
	require.Equal(t, "2131231", goal.ID)
	require.Equal(t, "goal", goal.Type)
	require.Equal(t, 1000, goal.Attributes.Amount)
}

const fetchCampaignResp = `
{
    "data": [
        {
            "attributes": {
                "created_at": "2016-02-02T19:58:18+00:00",
                "creation_count": 8,
                "creation_name": "new podcasting experience - Podsync",
                "discord_server_id": null,
                "display_patron_goals": true,
                "earnings_visibility": null,
                "image_small_url": "https://c10.patreon.com/3/eyJoIjoxMjgwLCJ3IjoxMjgwfQ%3D%3D/patreon-user/AS2N2NrZauWDuhVcuua87P7QtSOdCtPWiazP99SpvJWWHn8d4GvZI56AHqTn94g2_large_2.png?token-time=2145916800&token-hash=VNcCmkq6bOjbjizpwNHePu3aNqujVMKJYdfaPxoz3_c%3D",
                "image_url": "https://c10.patreon.com/3/eyJ3IjoxOTIwfQ%3D%3D/patreon-user/AS2N2NrZauWDuhVcuua87P7QtSOdCtPWiazP99SpvJWWHn8d4GvZI56AHqTn94g2_large_2.png?token-time=2145916800&token-hash=4KxOxPVCtGwPskLYr8BZGyZW94VwAKbD7j9RDHcvf0E%3D",
                "is_charged_immediately": true,
                "is_monthly": true,
                "is_nsfw": true,
                "is_plural": true,
                "main_video_embed": "",
                "main_video_url": "",
                "one_liner": null,
                "outstanding_payment_amount_cents": 0,
                "patron_count": 123121,
                "pay_per_name": "month",
                "pledge_sum": 12321312,
                "pledge_url": "/bePatron?c=278915",
                "published_at": "2016-02-02T20:11:19+00:00",
                "summary": "<a href=\"http://podsync.net/\" rel=\"nofollow\">Podsync</a> - is a simple, free service that lets you listen to any YouTube / Vimeo channels, playlists or user videos in podcast format.<br><br><strong>Idea:</strong><br>Podcast applications have a rich functionality for content delivery - automatic download of new episodes, remembering last played position, sync between devices and offline listening. This functionality is not available on YouTube and Vimeo. So the aim of\u00a0<a href=\"http://podsync.net/\" rel=\"nofollow\">Podsync</a> is to make your life easier and enable you to view/listen to content on any device in podcast client.<br><br>It's my hobby project, so to continue to support and improve it, I need your help. Your money will go into paying my server bills and adding new features.<br><br>",
                "thanks_embed": "",
                "thanks_msg": "You are awesome!",
                "thanks_video_url": null
            },
            "id": "278915",
            "type": "campaign",
            "relationships": {
                "categories": {
                    "data": [
                        {
                            "id": "7",
                            "type": "category"
                        }
                    ]
                },
                "creator": {
                    "data": {
                        "id": "2343242423",
                        "type": "user"
                    },
                    "links": {
                        "related": "https://www.patreon.com/api/user/2343242423"
                    }
                }
            }
        }
    ],
    "included": [
        {
            "attributes": {
                "vanity": "podsync"
            },
            "id": "2822191",
            "relationships": {},
            "type": "user"
        },
        {
            "attributes": {
                "amount": 100
            },
            "id": "12312312",
            "relationships": {},
            "type": "reward"
        },
        {
            "attributes": {
                "amount": 1000
            },
            "id": "2131231",
            "type": "goal"
        }
    ]
}
`
