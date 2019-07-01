package generic

import (
	parser_api "github.com/iocplatform/agent/pkg/parser/api"
	puller_api "github.com/iocplatform/agent/pkg/puller/api"
)

// Options configures a generic feed.
type Options struct {
	name           string
	tags           []string
	observableType string
	puller         puller_api.Puller
	parser         parser_api.Parser
	defaults       map[string]interface{}
}

// Option configures how set up the connection
type Option func(*Options)

// WithName returns an Option function which sets the feed name
func WithName(name string) Option {
	return func(o *Options) {
		o.name = name
	}
}

// WithPuller returns an Option function which sets the puller instance
func WithPuller(puller puller_api.Puller) Option {
	return func(o *Options) {
		o.puller = puller
	}
}

// WithParser returns an Option function which sets the parser instance
func WithParser(parser parser_api.Parser) Option {
	return func(o *Options) {
		o.parser = parser
	}
}

// WithObservable returns an Option function which sets the observable type
func WithObservable(_type string) Option {
	return func(o *Options) {
		o.observableType = _type
	}
}

// WithDefaults returns an Option function which sets the defaults values
func WithDefaults(defaults map[string]interface{}) Option {
	return func(o *Options) {
		o.defaults = defaults
	}
}

// WithTags returns an Option function which sets the default tags
func WithTags(tags []string) Option {
	return func(o *Options) {
		o.tags = tags
	}
}
