package patreon

import "time"

// Default includes for Pledge.
const PledgeDefaultRelations = "patron,reward,creator,address,pledge_vat_location"

// Pledge represents Patreon's pledge.
// Valid relationships: patron, reward, creator, address (?), card (?), pledge_vat_location (?).
type Pledge struct {
	Type       string `json:"type"`
	Id         string `json:"id"`
	Attributes struct {
		AmountCents    int       `json:"amount_cents"`
		CreatedAt      time.Time `json:"created_at"`
		PledgeCapCents int       `json:"pledge_cap_cents"`
		PatronPaysFees bool      `json:"patron_pays_fees"`
	} `json:"attributes"`
	Relationships struct {
		Patron  *PatronRelationship  `json:"patron"`
		Reward  *RewardRelationship  `json:"reward"`
		Creator *CreatorRelationship `json:"creator"`
	} `json:"relationships"`
}

type PledgeResponse struct {
	Data     []Pledge `json:"data"`
	Included Includes `json:"included"`
	Links    struct {
		First string `json:"first"`
		Next  string `json:"next"`
	} `json:"links"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
}
