package patreon

import (
	"encoding/json"
	"fmt"
)

type baseItem struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type data struct {
	Data *baseItem `json:"data"`
}

type dataArray struct {
	Data []*baseItem `json:"data"`
}

type identityRelationships struct {
	Campaign    data      `json:"campaign"`
	Memberships dataArray `json:"memberships"`
}

type identityData struct {
	ID            string                `json:"id"`
	Attributes    *UserAttributes       `json:"attributes"`
	Relationships identityRelationships `json:"relationships"`
}

type identityResponse struct {
	Data     identityData `json:"data"`
	Included includes     `json:"included"`
}

type campaignRelationships struct {
	Benefits dataArray `json:"benefits"`
	Creator  data      `json:"creator"`
	Goals    dataArray `json:"goals"`
	Tiers    dataArray `json:"tiers"`
}

type campaignData struct {
	ID            string                `json:"id"`
	Attributes    *CampaignAttributes   `json:"attributes"`
	Relationships campaignRelationships `json:"relationships"`
}

type campaignListResponse struct {
	Data     []campaignData `json:"data"`
	Included includes       `json:"included"`
}

type campaignResponse struct {
	Data     campaignData `json:"data"`
	Included includes     `json:"included"`
}

type memberRelationships struct {
	Address                data      `json:"address"`
	Campaign               data      `json:"campaign"`
	CurrentlyEntitledTiers dataArray `json:"currently_entitled_tiers"`
	User                   data      `json:"user"`
}

type memberData struct {
	ID            string              `json:"id"`
	Attributes    *MemberAttributes   `json:"attributes"`
	Relationships memberRelationships `json:"relationships"`
}

type memberResponse struct {
	Data     memberData `json:"data"`
	Included includes   `json:"included"`
}

type memberListResponse struct {
	Data     []memberData `json:"data"`
	Included includes     `json:"included"`
}

// includes wraps 'includes' JSON field to handle objects of different type within an array.
type includes struct {
	addresses   map[string]*Address
	benefits    map[string]*Benefit
	campaigns   map[string]*Campaign
	goals       map[string]*Goal
	memberships map[string]*Member
	tiers       map[string]*Tier
	users       map[string]*User
}

type jsonItem struct {
	ID   string          `json:"id"`
	Type string          `json:"type"`
	Raw  json.RawMessage `json:"attributes"`
	Attr interface{}     `json:"-"`
}

// UnmarshalJSON deserializes 'includes' field into the appropriate structs depending on the 'type' field.
// See http://gregtrowbridge.com/golang-json-serialization-with-interfaces/ for implementation details.
func (i *includes) UnmarshalJSON(b []byte) error {
	var items []*jsonItem
	if err := json.Unmarshal(b, &items); err != nil {
		return err
	}

	for _, item := range items {
		// Check if empty JSON block '{}'
		isEmpty := len(item.Raw) == 2 && item.Raw[0] == 123 && item.Raw[1] == 125

		switch item.Type {
		case "address":
			address := &Address{
				ID: item.ID,
			}

			if !isEmpty {
				attr := &AddressAttributes{}

				item.Attr = attr
				if err := json.Unmarshal(item.Raw, item.Attr); err != nil {
					return err
				}

				address.AddressAttributes = attr
			}

			if i.addresses == nil {
				i.addresses = make(map[string]*Address)
			}

			i.addresses[address.ID] = address

		case "benefit":
			benefit := &Benefit{
				ID: item.ID,
			}

			if !isEmpty {
				attr := &BenefitAttributes{}

				item.Attr = attr
				if err := json.Unmarshal(item.Raw, item.Attr); err != nil {
					return err
				}

				benefit.BenefitAttributes = attr
			}

			if i.benefits == nil {
				i.benefits = make(map[string]*Benefit)
			}

			i.benefits[item.ID] = benefit

		case "campaign":
			campaign := &Campaign{
				ID: item.ID,
			}

			if !isEmpty {
				attr := &CampaignAttributes{}

				item.Attr = attr
				if err := json.Unmarshal(item.Raw, item.Attr); err != nil {
					return err
				}

				campaign.CampaignAttributes = attr
			}

			if i.campaigns == nil {
				i.campaigns = make(map[string]*Campaign)
			}

			i.campaigns[item.ID] = campaign

		case "goal":
			goal := &Goal{
				ID: item.ID,
			}

			if !isEmpty {
				attr := &GoalAttributes{}

				item.Attr = attr
				if err := json.Unmarshal(item.Raw, item.Attr); err != nil {
					return err
				}

				goal.GoalAttributes = attr
			}

			if i.goals == nil {
				i.goals = make(map[string]*Goal)
			}

			i.goals[item.ID] = goal

		case "memberships":
			member := &Member{
				ID: item.ID,
			}

			if !isEmpty {
				attr := &MemberAttributes{}

				item.Attr = attr
				if err := json.Unmarshal(item.Raw, item.Attr); err != nil {
					return err
				}

				member.MemberAttributes = attr
			}

			if i.memberships == nil {
				i.memberships = make(map[string]*Member)
			}

			i.memberships[item.ID] = member

		case "tier":
			tier := &Tier{
				ID: item.ID,
			}

			if !isEmpty {
				attr := &TierAttributes{}

				item.Attr = attr
				if err := json.Unmarshal(item.Raw, item.Attr); err != nil {
					return err
				}

				tier.TierAttributes = attr
			}

			if i.tiers == nil {
				i.tiers = make(map[string]*Tier)
			}

			i.tiers[item.ID] = tier

		case "user":
			user := &User{
				ID: item.ID,
			}

			if !isEmpty {
				attr := &UserAttributes{}

				item.Attr = attr
				if err := json.Unmarshal(item.Raw, item.Attr); err != nil {
					return err
				}

				user.UserAttributes = attr
			}

			if i.users == nil {
				i.users = make(map[string]*User)
			}

			i.users[item.ID] = user

		default:
			return fmt.Errorf("unsupported type '%s'", item.Type)
		}
	}

	return nil
}
