package patreon

var (
	// MemberDefaultIncludes specifies default includes for Member.
	MemberDefaultIncludes = []string{"address", "campaign", "currently_entitled_tiers", "user"}

	// MemberAttributes is all fields in the Member Attributes struct
	MemberAttributes = []string{
		"CampaignLifetimeSupportCents", "CurrentlyEntitledAmountCents",
		"Email", "FullName", "IsFollower", "LastChargeDate",
		"LastChargeStatus", "LifetimeSupportCents", "NextChargeDate", "Note",
		"PatronStatus", "PledgeCadence", "PledgeRelationshipStart", "WillPayAmountCents",
	}
)

// Member is the record of a user's membership to a campaign.
// Remains consistent across months of pledging.
type Member struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		CampaignLifetimeSupportCents int      `json:"campaign_lifetime_support_cents"`
		CurrentlyEntitledAmountCents int      `json:"currently_entitled_amount_cents"`
		Email                        string   `json:"email"`
		FullName                     string   `json:"full_name"`
		IsFollower                   bool     `json:"is_follower"`
		LastChargeDate               NullTime `json:"last_charge_date"`
		LastChargeStatus             string   `json:"last_charge_status"`
		LifetimeSupportCents         int      `json:"lifetime_support_cents"`
		NextChargeDate               NullTime `json:"next_charge_date"`
		Note                         string   `json:"note"`
		PatronStatus                 string   `json:"patron_status"`
		PledgeCadence                int      `json:"pledge_cadence"`
		PledgeRelationshipStart      NullTime `json:"pledge_relationship_start"`
		WillPayAmountCents           int      `json:"will_pay_amount_cents"`
	} `json:"attributes"`
	Relationships struct {
		Address                *AddressRelationship     `json:"address,omitempty"`
		Campaign               *CampaignRelationship    `json:"campaign,omitempty"`
		CurrentlyEntitledTiers *TiersRelationship       `json:"currently_entitled_tiers,omitempty"`
		PledgeHistory          *PledgeEventRelationship `json:"pledge_history,omitempty"`
		User                   *UserRelationship        `json:"user,omitempty"`
	} `json:"relationships"`
}

// MemberResponse wraps Patreon's fetch benefit API response
type MemberResponse struct {
	Data     Member   `json:"data"`
	Included Includes `json:"included"`
	Links    struct {
		Self string `json:"self"`
	} `json:"links"`
}

// MembersResponse wraps Patreon's fetch benefit API response
type MembersResponse struct {
	Data     []Member `json:"data"`
	Included Includes `json:"included"`
	Links    struct {
		Self string `json:"self"`
	} `json:"links"`
}
