package patreon

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParseIncludes(t *testing.T) {
	includes := Includes{}
	err := json.Unmarshal([]byte(includesJson), &includes)
	require.NoError(t, err)
	require.Len(t, includes.Items, 6)

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

	pledge, ok := includes.Items[4].(*Pledge)
	require.True(t, ok)
	require.Equal(t, 100, pledge.Attributes.AmountCents)
	require.True(t, pledge.Attributes.CreatedAt.Valid)
	require.Equal(t, time.Date(2017, 6, 20, 23, 21, 34, 514822000, time.UTC).Unix(), pledge.Attributes.CreatedAt.Unix())
	require.False(t, pledge.Attributes.DeclinedSince.Valid)
	require.True(t, pledge.Attributes.PatronPaysFees)
	require.Equal(t, 100, pledge.Attributes.PledgeCapCents)

	card, ok := includes.Items[5].(*Card)
	require.True(t, ok)
	require.Equal(t, "bt_12312312", card.Id)
	require.Equal(t, "card", card.Type)
	require.Equal(t, "PayPal", card.Attributes.CardType)
	require.True(t, card.Attributes.HasFailedPayment)
	require.True(t, card.Attributes.IsVerified)
	require.Equal(t, "12312312", card.Attributes.Number)
	require.Equal(t, "bt_12312312", card.Attributes.PaymentToken)
	require.Equal(t, 12312312, card.Attributes.PaymentTokenID)
	require.NotNil(t, card.Relationships.User)
	require.Equal(t, "https://www.patreon.com/api/user/4221587", card.Relationships.User.Links.Related)
	require.Equal(t, "12312312", card.Relationships.User.Data.Id)
	require.Equal(t, "user", card.Relationships.User.Data.Type)
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
	},
	{
		"attributes": {
			"amount_cents": 100,
			"created_at": "2017-06-20T23:21:34.514822+00:00",
			"declined_since": null,
			"patron_pays_fees": true,
			"pledge_cap_cents": 100
		},
		"id": "2321312",
		"relationships": {
			"card": {
				"data": {
					"id": "bt_1231232132",
					"type": "card"
				},
				"links": {
					"related": "https://www.patreon.com/api/cards/bt_123123213"
				}
			},
			"creator": {
				"data": {
					"id": "12312321321312",
					"type": "user"
				},
				"links": {
					"related": "https://www.patreon.com/api/user/12312321321312"
				}
			},
			"patron": {
				"data": {
					"id": "213213213",
					"type": "user"
				},
				"links": {
					"related": "https://www.patreon.com/api/user/213213213"
				}
			},
			"reward": {
				"data": {
					"id": "12312321321",
					"type": "reward"
				},
				"links": {
					"related": "https://www.patreon.com/api/rewards/12312321321"
				}
			}
		},
		"type": "pledge"
	},
	{
		"attributes": {
			"card_type": "PayPal",
			"created_at": "2016-12-24T18:18:22+00:00",
			"expiration_date": null,
			"has_a_failed_payment": true,
			"is_verified": true,
			"number": "12312312",
			"payment_token": "bt_12312312",
			"payment_token_id": 12312312
		},
		"id": "bt_12312312",
		"relationships": {
			"user": {
				"data": {
					"id": "12312312",
					"type": "user"
				},
				"links": {
					"related": "https://www.patreon.com/api/user/4221587"
				}
			}
		},
		"type": "card"
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
