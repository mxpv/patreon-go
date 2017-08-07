package patreon

import "strings"

type options struct {
	fields  map[string]string
	include string
	size    int
	cursor  string
}

type requestOption func(*options)

func WithFields(resource string, fields ...string) requestOption {
	return func(o *options) {
		if o.fields == nil {
			o.fields = make(map[string]string)
		}
		o.fields[resource] = strings.Join(fields, ",")
	}
}

func WithIncludes(include ...string) requestOption {
	return func(o *options) {
		o.include = strings.Join(include, ",")
	}
}

func WithPageSize(size int) requestOption {
	return func(o *options) {
		o.size = size
	}
}

func WithCursor(cursor string) requestOption {
	return func(o *options) {
		o.cursor = cursor
	}
}

func getOptions(opts ...requestOption) options {
	cfg := options{}
	for _, fn := range opts {
		fn(&cfg)
	}

	return cfg
}
