package patreon

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseIncludes(t *testing.T) {
	includes := Includes{}
	err := json.Unmarshal([]byte(includesJson), &includes)
	require.NoError(t, err)
	require.Len(t, includes.Items, 4)

	user, ok := includes.Items[0].(*User)
	require.True(t, ok)
	require.Equal(t, "2822191", user.ID)
	require.Equal(t, "user", user.Type)
	require.Equal(t, "podsync", user.Attributes.Vanity)

	goal, ok := includes.Items[1].(*Goal)
	require.True(t, ok)
	require.Equal(t, "2131231", goal.ID)
	require.Equal(t, "goal", goal.Type)

	campaign, ok := includes.Items[2].(*Campaign)
	require.True(t, ok)
	require.Equal(t, "12312321", campaign.ID)
	require.Equal(t, "campaign", campaign.Type)

	require.True(t, ok)
}

func TestParseUnsupportedInclude(t *testing.T) {
	includes := Includes{}
	err := json.Unmarshal([]byte(unknownIncludeJson), &includes)
	require.Error(t, err)
	require.Equal(t, "unsupported type 'unknown'", err.Error())
}

const includesJson = `
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
			"amount": 1000
		},
		"id": "2131231",
		"type": "goal"
	},
	{
		"attributes": {},
		"id": "12312321",
		"type": "campaign"
	},
	{
		"attributes": {
			"amount_cents": 100,
			"created_at": "2017-06-20T23:21:34.514822+00:00",
			"declined_since": null,
			"patron_pays_fees": true,
			"pledge_cap_cents": 100
		}
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
