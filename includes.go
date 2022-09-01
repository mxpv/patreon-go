package patreon

import (
	"encoding/json"
	"fmt"
)

// Includes wraps 'includes' JSON field to handle objects of different type within an array.
type Includes struct {
	Items []interface{}
}

// UnmarshalJSON deserializes 'includes' field into the appropriate structs depending on the 'type' field.
// See http://gregtrowbridge.com/golang-json-serialization-with-interfaces/ for implementation details.
func (i *Includes) UnmarshalJSON(b []byte) error {
	var items []*json.RawMessage
	if err := json.Unmarshal(b, &items); err != nil {
		return err
	}

	count := len(items)
	i.Items = make([]interface{}, count)

	s := struct {
		Type string `json:"type"`
	}{}

	for idx, raw := range items {
		if err := json.Unmarshal(*raw, &s); err != nil {
			return err
		}

		var obj interface{}

		// Depending on the type, we can run json.Unmarshal again on the same byte slice
		// But this time, we'll pass in the appropriate struct instead of a map
		if s.Type == "user" {
			obj = &User{}
		} else if s.Type == "tier" {
			obj = &Tier{}
		} else if s.Type == "goal" {
			obj = &Goal{}
		} else if s.Type == "campaign" {
			obj = &Campaign{}
		} else if s.Type == "benefit" {
			obj = &Benefit{}
		} else if s.Type == "membership" {
			obj = &Member{}
		} else if s.Type == "address" {
			obj = &Address{}
		} else if s.Type == "patron" {
			obj = &User{}
		} else if s.Type == "webhook" {
			obj = &Webhook{}
		} else if s.Type == "deliverable" {
			obj = &Deliverable{}
		} else {
			return fmt.Errorf("unsupported type '%s'", s.Type)
		}

		if err := json.Unmarshal(*raw, obj); err != nil {
			return err
		}

		i.Items[idx] = obj
	}

	return nil
}
