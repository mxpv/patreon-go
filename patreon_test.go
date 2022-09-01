package patreon

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
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
		WithCursor("123"),
	)

	require.NoError(t, err)
	require.Equal(t, "https://api.patreon.com/path?fields%5Bpledge%5D=total_historical_amount_cents%2Cunread_count&include=patron%2Creward%2Ccreator&page%5Bcount%5D=10&page%5Bcursor%5D=123", url)
}

func TestBuildURLWithInvalidPath(t *testing.T) {
	client := &Client{}

	url, err := client.buildURL("")
	require.Error(t, err)
	require.Empty(t, url)
}

func TestClient(t *testing.T) {
	tc := oauth2.NewClient(context.Background(), nil)
	client := NewClient(tc)
	require.Equal(t, tc, client.Client())
}
