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

	require.Equal(t, 2, len(resp.Data))
	require.Equal(t, "pledge", resp.Data[0].Type)
	require.Equal(t, "61272355", resp.Data[0].Id)
	require.Equal(t, 100, resp.Data[0].Attributes.AmountCents)
	require.Equal(t, 100, resp.Data[0].Attributes.PledgeCapCents)
	require.True(t, resp.Data[0].Attributes.PatronPaysFees)
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
