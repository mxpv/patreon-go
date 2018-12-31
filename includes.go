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

// Includes wraps 'includes' JSON field to handle objects of different type within an array.
type IncludedItems struct {
	Items []*jsonItem
}

// UnmarshalJSON deserializes 'includes' field into the appropriate structs depending on the 'type' field.
// See http://gregtrowbridge.com/golang-json-serialization-with-interfaces/ for implementation details.
func (i *IncludedItems) UnmarshalJSON(b []byte) error {
	var items []*jsonItem
	if err := json.Unmarshal(b, &items); err != nil {
		return err
	}

	for _, item := range items {
		switch item.Type {
		case "address":
			item.Attr = &AddressAttributes{}
		case "campaign":
			item.Attr = &CampaignAttributes{}
		case "tier":
			item.Attr = &TierAttributes{}
		case "user":
			item.Attr = &UserAttributes{}
		default:
			return fmt.Errorf("unsupported type '%s'", item.Type)
		}

		if err := json.Unmarshal(item.Raw, item.Attr); err != nil {
			return err
		}
	}

	i.Items = items
	return nil
}
