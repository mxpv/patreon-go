package patreon

import (
	"encoding/json"
	"strings"
	"time"
)

// NullTime represents a time.Time that may be JSON "null".
// golang prior 1.8 doesn't support this scenario (fails with error: parsing time "null" as ""2006-01-02T15:04:05Z07:00"": cannot parse "null" as """)
type NullTime struct {
	time.Time
	Valid bool
}

// MarshalJSON implements json.Marshaler, it will encode null if this time is null.
func (t *NullTime) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return t.Time.MarshalJSON()
}

// UnmarshalJSON implements json.Unmarshaler with JSON "null" support
func (t *NullTime) UnmarshalJSON(data []byte) error {
	s := string(data)
	if strings.EqualFold(s, "null") {
		t.Valid = false
		return nil
	}

	err := json.Unmarshal(data, &t.Time)
	t.Valid = err == nil

	return err
}
