package patreon

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

const includesJson = `
[
        {
            "attributes": {
                "addressee": "Platform Team",
                "city": "San Francisco",
                "confirmed": true,
                "confirmed_at": null,
                "country": "US",
                "created_at": "2018-06-03T16:23:38+00:00",
                "line_1": "555 Main St",
                "line_2": "",
                "phone_number": null,
                "postal_code": "94103",
                "state": "CA"
            },
            "id": "123456",
            "type": "address"
        },
        {
            "attributes": {
                "full_name": "Platform Team",
                "hide_pledges": true
            },
            "id": "654321",
            "type": "user"
        },
        {
            "attributes": {
                "title": "Tshirt Tier"
            },
            "id": "99001122",
            "type": "tier"
        }
    ]
`

const unknownIncludeJson = `
[
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
			"vanity": "podsync"
		},
		"id": "12312312",
		"relationships": {},
		"type": "unknown"
	}
]
`

const emptyInclude = `
[{"attributes":{},"id":"278915","type":"campaign"}]
`

func TestParseIncludes(t *testing.T) {
	includes := includes{}
	err := json.Unmarshal([]byte(includesJson), &includes)
	require.NoError(t, err)

	require.Len(t, includes.addresses, 1)
	require.Len(t, includes.users, 1)
	require.Len(t, includes.tiers, 1)

	for _, address := range includes.addresses {
		assert.Equal(t, "Platform Team", address.Addressee)
		assert.Equal(t, "San Francisco", address.City)
		assert.Equal(t, "US", address.Country)
	}

	for _, user := range includes.users {
		assert.Equal(t, "Platform Team", user.FullName)
		assert.True(t, user.HidePledges)
	}

	for _, tier := range includes.tiers {
		assert.Equal(t, "Tshirt Tier", tier.Title)
	}
}

func TestParseUnsupportedInclude(t *testing.T) {
	includes := includes{}
	err := json.Unmarshal([]byte(unknownIncludeJson), &includes)
	require.Error(t, err)
	require.EqualError(t, err, "unsupported type 'unknown'")
}

func TestEmptyInclude(t *testing.T) {
	includes := includes{}
	err := json.Unmarshal([]byte(emptyInclude), &includes)
	require.NoError(t, err)
	require.Len(t, includes.campaigns, 1)

	item, ok := includes.campaigns["278915"]
	require.True(t, ok)

	require.Equal(t, "278915", item.ID)
	require.Nil(t, item.CampaignAttributes)
}
