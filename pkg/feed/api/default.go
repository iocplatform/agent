package api

import (
	parser_api "github.com/iocplatform/agent/pkg/parser/api"
	puller_api "github.com/iocplatform/agent/pkg/puller/api"
)

// DefaultFeed is the default feed implementation contains all common attibutes and methods.
type DefaultFeed struct {
	Type string
	Name string
	Tags []string

	puller puller_api.Puller
	parser parser_api.Parser
}

// GetType returns the feed type
func (f *DefaultFeed) GetType() string {
	return f.Type
}

// GetName returns the feed name
func (f *DefaultFeed) GetName() string {
	return f.Name
}

// GetTags returns the feed tags
func (f *DefaultFeed) GetTags() []string {
	return f.Tags
}

// SetPuller is ued to update puller reference
func (f *DefaultFeed) SetPuller(puller puller_api.Puller) {
	f.puller = puller
}

// Puller returns the puller instance
func (f *DefaultFeed) Puller() puller_api.Puller {
	return f.puller
}

// SetParser is used to update parser reference
func (f *DefaultFeed) SetParser(parser parser_api.Parser) {
	f.parser = parser
}

// Parser returns the parser instance
func (f *DefaultFeed) Parser() parser_api.Parser {
	return f.parser
}
