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
	require.Equal(t, "2822191", user.Id)
	require.Equal(t, "user", user.Type)
	require.Equal(t, "podsync", user.Attributes.Vanity)

	reward, ok := includes.Items[1].(*Reward)
	require.True(t, ok)
	require.Equal(t, "12312312", reward.Id)
	require.Equal(t, "reward", reward.Type)
	require.Equal(t, 100, reward.Attributes.Amount)

	goal, ok := includes.Items[2].(*Goal)
	require.True(t, ok)
	require.Equal(t, "2131231", goal.Id)
	require.Equal(t, "goal", goal.Type)
	require.Equal(t, 1000, goal.Attributes.Amount)

	campaign, ok := includes.Items[3].(*Campaign)
	require.True(t, ok)
	require.Equal(t, "12312321", campaign.Id)
	require.Equal(t, "campaign", campaign.Type)
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
	},
	{
		"attributes": {},
		"id": "12312321",
		"type": "campaign"
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
