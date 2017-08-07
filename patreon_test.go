package patreon_go

import (
	"net/http"
	"net/http/httptest"
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
