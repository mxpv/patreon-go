package patreon

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWithCursorURL(t *testing.T) {
	fn := WithCursor("https://www.patreon.com/api/oauth2/api/campaigns/123456/pledges?page%5Bcount%5D=10&sort=created&page%5Bcursor%5D=2017-01-19T18%3A39%3A17%2B00%3A00")

	opt := options{}
	fn(&opt)

	require.Equal(t, "2017-01-19T18:39:17+00:00", opt.cursor)
}

func TestWithCursor(t *testing.T) {
	fn := WithCursor("2017-01-19T18:39:17+00:00")

	opt := options{}
	fn(&opt)

	require.Equal(t, "2017-01-19T18:39:17+00:00", opt.cursor)
}
