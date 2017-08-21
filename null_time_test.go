package patreon

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNullTime_parseValidTime(t *testing.T) {
	s := &struct {
		Time NullTime `json:"time"`
	}{}

	err := json.Unmarshal([]byte(`{ "time": "2017-06-20T23:21:34.514822+00:00" }`), s)
	require.NoError(t, err)
	require.True(t, s.Time.Valid)
	require.Equal(t, time.Date(2017, 6, 20, 23, 21, 34, 514822000, time.UTC).Unix(), s.Time.Time.Unix())
}

func TestNullTime_parseInvalidTime(t *testing.T) {
	s := &struct {
		Time NullTime `json:"time"`
	}{}

	err := json.Unmarshal([]byte(`{ "time": null }`), s)
	require.NoError(t, err)
	require.False(t, s.Time.Valid)
}
