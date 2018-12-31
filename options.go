package patreon

import (
	"net/url"
	"strings"
)

type options struct {
	fields  map[string]string
	include string
	size    int
	cursor  string
}

type RequestOpt func(*options)

// WithFields specifies the resource attributes you want to be returned by API.
func WithFields(resource string, fields ...string) RequestOpt {
	return func(o *options) {
		if o.fields == nil {
			o.fields = make(map[string]string)
		}
		o.fields[resource] = strings.Join(fields, ",")
	}
}

// WithIncludes specifies the related resources you want to be returned by API.
func WithIncludes(include ...string) RequestOpt {
	return func(o *options) {
		o.include = strings.Join(include, ",")
	}
}

// WithPageSize specifies the number of items to return.
func WithPageSize(size int) RequestOpt {
	return func(o *options) {
		o.size = size
	}
}

// WithCursor controls cursor-based pagination. Cursor will also be extracted from navigation links for convenience.
func WithCursor(cursor string) RequestOpt {
	return func(o *options) {
		u, err := url.ParseRequestURI(cursor)
		if err == nil {
			cursor = u.Query().Get("page[cursor]")
		}

		o.cursor = cursor
	}
}

func getOptions(opts ...RequestOpt) options {
	cfg := options{}
	for _, fn := range opts {
		fn(&cfg)
	}

	return cfg
}
