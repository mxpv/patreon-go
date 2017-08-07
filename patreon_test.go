package patreon

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *Client
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = NewClient(nil)
	client.baseURL = server.URL
}

func teardown() {
	server.Close()
}

func TestBuildURL(t *testing.T) {
	client := NewClient(nil)

	url, err := client.buildURL("/path",
		WithIncludes("patron", "reward", "creator"),
		WithFields("pledge", "total_historical_amount_cents", "unread_count"),
		WithPageSize(10),
	)

	require.NoError(t, err)
	require.Equal(t, "https://api.patreon.com/path?fields%5Bpledge%5D=total_historical_amount_cents%2Cunread_count&include=patron%2Creward%2Ccreator&page%5Bcount%5D=10", url)
}
