package patreon

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseIncludes(t *testing.T) {
	includes := IncludedItems{}
	err := json.Unmarshal([]byte(includesJson), &includes)
	require.NoError(t, err)
	require.Len(t, includes.Items, 3)

	address, ok := includes.Items[0].Attr.(*AddressAttributes)
	require.True(t, ok)

	assert.Equal(t, "Platform Team", address.Addressee)
	assert.Equal(t, "San Francisco", address.City)
	assert.Equal(t, "US", address.Country)

	user, ok := includes.Items[1].Attr.(*UserAttributes)
	require.True(t, ok)

	assert.Equal(t, "Platform Team", user.FullName)
	assert.True(t, user.HidePledges)

	tier, ok := includes.Items[2].Attr.(*TierAttributes)
	require.True(t, ok)

	assert.Equal(t, "Tshirt Tier", tier.Title)
}

func TestParseUnsupportedInclude(t *testing.T) {
	includes := IncludedItems{}
	err := json.Unmarshal([]byte(unknownIncludeJson), &includes)
	require.Error(t, err)
	require.EqualError(t, err, "unsupported type 'unknown'")
}

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
		"attributes": {},
		"id": "12312312",
		"relationships": {},
		"type": "unknown"
	}
]
`
