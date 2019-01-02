package patreon

import (
	"encoding/json"
	"fmt"
)

type jsonItem struct {
	ID   string          `json:"id"`
	Type string          `json:"type"`
	Raw  json.RawMessage `json:"attributes"`
	Attr interface{}     `json:"-"`
}

func (i *jsonItem) toCampaign() (*Campaign, error) {
	if i.Type != "campaign" {
		return nil, fmt.Errorf("can't convert %q to campaign", i.Type)
	}

	campaign := &Campaign{ID: i.ID}

	if i.Attr != nil {
		attrs, ok := i.Attr.(*CampaignAttributes)
		if !ok {
			return nil, fmt.Errorf("unable to cast %q (id %q)", i.Type, i.ID)
		}

		campaign.CampaignAttributes = attrs
	}

	return campaign, nil
}

func (i *jsonItem) toMember() (*Member, error) {
	if i.Type != "memberships" {
		return nil, fmt.Errorf("can't convert %q to memberships", i.Type)
	}

	member := &Member{ID: i.ID}

	if i.Attr != nil {
		attrs, ok := i.Attr.(*MemberAttributes)
		if !ok {
			return nil, fmt.Errorf("unable to cast %q (id %q)", i.Type, i.ID)
		}

		member.MemberAttributes = attrs
	}

	return member, nil
}

func (i *jsonItem) toTier() (*Tier, error) {
	if i.Type != "tier" {
		return nil, fmt.Errorf("can't convert %q to tier", i.Type)
	}

	tier := &Tier{ID: i.ID}

	if i.Attr != nil {
		attrs, ok := i.Attr.(*TierAttributes)
		if !ok {
			return nil, fmt.Errorf("unable to cast %q (id %q)", i.Type, i.ID)
		}

		tier.TierAttributes = attrs
	}

	return tier, nil
}

func (i *jsonItem) toGoal() (*Goal, error) {
	if i.Type != "goal" {
		return nil, fmt.Errorf("can't convert %q to goal", i.Type)
	}

	goal := &Goal{ID: i.ID}

	if i.Attr != nil {
		attrs, ok := i.Attr.(*GoalAttributes)
		if !ok {
			return nil, fmt.Errorf("unable to cast %q (id %q)", i.Type, i.ID)
		}

		goal.GoalAttributes = attrs
	}

	return goal, nil
}

func (i *jsonItem) toUser() (*User, error) {
	if i.Type != "user" {
		return nil, fmt.Errorf("can't convert %q to user", i.Type)
	}

	user := &User{ID: i.ID}

	if i.Attr != nil {
		attrs, ok := i.Attr.(*UserAttributes)
		if !ok {
			return nil, fmt.Errorf("unable to cast %q (id %q)", i.Type, i.ID)
		}

		user.UserAttributes = attrs
	}

	return user, nil
}

func (i *jsonItem) toBenefit() (*Benefit, error) {
	if i.Type != "benefit" {
		return nil, fmt.Errorf("can't convert %q to benefit", i.Type)
	}

	benefit := &Benefit{ID: i.ID}

	if i.Attr != nil {
		attrs, ok := i.Attr.(*BenefitAttributes)
		if !ok {
			return nil, fmt.Errorf("unable to cast %q (id %q)", i.Type, i.ID)
		}

		benefit.BenefitAttributes = attrs
	}

	return benefit, nil
}

type relationItem struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type data struct {
	Data *relationItem `json:"data"`
}

type dataArray struct {
	Data []*relationItem `json:"data"`
}

// Includes wraps 'includes' JSON field to handle objects of different type within an array.
type includedItems struct {
	Items []*jsonItem
}

// UnmarshalJSON deserializes 'includes' field into the appropriate structs depending on the 'type' field.
// See http://gregtrowbridge.com/golang-json-serialization-with-interfaces/ for implementation details.
func (i *includedItems) UnmarshalJSON(b []byte) error {
	var items []*jsonItem
	if err := json.Unmarshal(b, &items); err != nil {
		return err
	}

	for _, item := range items {
		// Check if empty JSON block '{}'
		if len(item.Raw) == 2 && item.Raw[0] == 123 && item.Raw[1] == 125 {
			continue
		}

		switch item.Type {
		case "address":
			item.Attr = &AddressAttributes{}
		case "campaign":
			item.Attr = &CampaignAttributes{}
		case "goal":
			item.Attr = &GoalAttributes{}
		case "tier":
			item.Attr = &TierAttributes{}
		case "user":
			item.Attr = &UserAttributes{}
		default:
			return fmt.Errorf("unsupported type '%s'", item.Type)
		}

		// Deserialize attributes
		if err := json.Unmarshal(item.Raw, item.Attr); err != nil {
			return err
		}
	}

	i.Items = items
	return nil
}

func (i *includedItems) findBy(relation *relationItem) (*jsonItem, error) {
	// nil relation means no relation and no error
	if relation == nil {
		return nil, nil
	}

	for _, data := range i.Items {
		if data.Type == relation.Type && data.ID == relation.ID {
			return data, nil
		}
	}

	return nil, fmt.Errorf("can't find relation with id %q (type %q)", relation.ID, relation.Type)
}
