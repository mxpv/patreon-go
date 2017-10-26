package patreon

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetchPledges(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/oauth2/api/campaigns/123/pledges", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, fetchPledgesResp)
	})

	resp, err := client.FetchPledges("123")
	require.NoError(t, err)

	require.Equal(t, 169, resp.Meta.Count)

	require.NotEmpty(t, resp.Links.First)
	require.NotEmpty(t, resp.Links.Next)

	// Attributes

	require.Equal(t, 2, len(resp.Data))
	require.Equal(t, "pledge", resp.Data[0].Type)
	require.Equal(t, "61272355", resp.Data[0].ID)
	require.Equal(t, 100, resp.Data[0].Attributes.AmountCents)
	require.Equal(t, 100, resp.Data[0].Attributes.PledgeCapCents)
	require.True(t, resp.Data[0].Attributes.PatronPaysFees)

	// Relationships

	creator := resp.Data[1].Relationships.Creator
	require.NotNil(t, creator)
	require.Equal(t, "12321312", creator.Data.ID)
	require.Equal(t, "user", creator.Data.Type)
	require.Equal(t, "https://www.patreon.com/api/user/12321312", creator.Links.Related)

	patron := resp.Data[1].Relationships.Patron
	require.NotNil(t, patron)
	require.Equal(t, "12321", patron.Data.ID)
	require.Equal(t, "user", patron.Data.Type)
	require.Equal(t, "https://www.patreon.com/api/user/12321", patron.Links.Related)

	reward := resp.Data[1].Relationships.Reward
	require.NotNil(t, reward)
	require.Equal(t, "21321321321", reward.Data.ID)
	require.Equal(t, "reward", reward.Data.Type)
	require.Equal(t, "https://www.patreon.com/api/rewards/21321321321", reward.Links.Related)
}

const fetchPledgesResp = `
{
    "data": [
        {
            "attributes": {
                "amount_cents": 100,
                "created_at": "2016-06-06T13:58:56+00:00",
                "declined_since": null,
                "patron_pays_fees": true,
                "pledge_cap_cents": 100
            },
            "id": "61272355",
            "type": "pledge"
        },
        {
            "attributes": {
                "total_historical_amount_cents": 100,
                "unread_count": 8
            },
            "id": "6628990",
			"relationships": {
				"creator": {
					"data": {
						"id": "12321312",
						"type": "user"
					},
					"links": {
						"related": "https://www.patreon.com/api/user/12321312"
					}
				},
				"patron": {
					"data": {
						"id": "12321",
						"type": "user"
					},
					"links": {
						"related": "https://www.patreon.com/api/user/12321"
					}
				},
				"reward": {
					"data": {
						"id": "21321321321",
						"type": "reward"
					},
					"links": {
						"related": "https://www.patreon.com/api/rewards/21321321321"
					}
				}
			},
            "type": "pledge"
        }
    ],
    "links": {
        "first": "https://www.patreon.com/api/oauth2/api/campaigns/21980312/pledges?page%5Bcount%5D=2&sort=created",
        "next": "https://www.patreon.com/api/oauth2/api/campaigns/21980312/pledges?page%5Bcount%5D=2&sort=created&page%5Bcursor%5D=2017-07-03T23%3A25%3A08.519452%2B00%3A00"
    },
    "meta": {
        "count": 169
    }
}
`
