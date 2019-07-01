package http

import (
	"time"

	"github.com/jmespath/go-jmespath"
)

// Options configures a collector.
type Options struct {
	url           string
	method        string
	parameters    map[string]string
	headers       map[string]string
	timeout       time.Duration
	jmesPath      *jmespath.JMESPath
	tlsSkipVerify bool
	json          bool
	lines         bool
	br2nl         bool
}

// Option configures how set up the connection
type Option func(*Options)

// WithURL returns an Option function which sets the url
func WithURL(url string) Option {
	return func(o *Options) {
		o.url = url
	}
}

// WithMethod returns an Option function which sets the method used to query the url
func WithMethod(method string) Option {
	return func(o *Options) {
		o.method = method
	}
}

// WithParameters returns an Option function which sets the parameters hash map
func WithParameters(params map[string]string) Option {
	return func(o *Options) {
		o.parameters = params
	}
}

// WithTLSSkipVerify returns an Option function which disable TLS verification
func WithTLSSkipVerify(status bool) Option {
	return func(o *Options) {
		o.tlsSkipVerify = status
	}
}

// WithTimeout returns an Option function which sets the query timeout
func WithTimeout(duration time.Duration) Option {
	return func(o *Options) {
		o.timeout = duration
	}
}

// WithReadLines returns an Option function which read response line by line
func WithReadLines() Option {
	return func(o *Options) {
		o.lines = true
		o.json = false
	}
}

// WithJSON returns an Option function which read response as json object
func WithJSON(expression string) Option {
	return func(o *Options) {
		o.json = true
		o.lines = false
		o.jmesPath = jmespath.MustCompile(expression)
	}
}

// WithHeaders returns an Option function which set the request headers
func WithHeaders(headers map[string]string) Option {
	return func(o *Options) {
		o.headers = headers
	}
}

// WithBR2NLPreprocessor returns an Option function which set the br2nl preprocessor status
func WithBR2NLPreprocessor(status bool) Option {
	return func(o *Options) {
		o.br2nl = status
	}
}
