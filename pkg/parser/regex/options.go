package regex

import "regexp"

// regexOptions configures a regex parser.
type regexOptions struct {
	fields     []string
	skipLines  int
	rule       *regexp.Regexp
	ignoreRule *regexp.Regexp
}

// Option configures how set up the connection
type Option func(*regexOptions)

// WithRule returns a regexOption which sets the extraction rule
func WithRule(rule string) Option {
	return func(o *regexOptions) {
		o.rule = regexp.MustCompile(rule)
	}
}

// WithIgnoreRule returns a regexOption which sets the ignore rule
func WithIgnoreRule(rule string) Option {
	return func(o *regexOptions) {
		o.ignoreRule = regexp.MustCompile(rule)
	}
}

// WithFields returns o regexOption which sets the fields attribute
func WithFields(fields ...string) Option {
	return func(o *regexOptions) {
		o.fields = fields
	}
}

// WithSkipLines returns o regexOption which sets the skip lines attribute.
func WithSkipLines(sl int) Option {
	return func(o *regexOptions) {
		o.skipLines = sl
	}
}
