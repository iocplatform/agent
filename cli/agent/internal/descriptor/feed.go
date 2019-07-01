package descriptor

import (
	parser_api "github.com/iocplatform/agent/pkg/parser/api"
	puller_api "github.com/iocplatform/agent/pkg/puller/api"

	"github.com/iocplatform/agent/pkg/feed/api"
	"github.com/iocplatform/agent/pkg/feed/generic"
)

// Feed is the a feed information holder for agent
type Feed struct {
	// Feed type
	Type string `yaml:"type"`

	// Feed name
	Name string `yaml:"name"`

	// Overrides agent settings
	Overrides map[string]interface{} `yaml:"overrides"`

	// Affected tags
	Tags []string `yaml:"tags"`

	// Concrete parser implementation
	concrete api.Feed `yaml:"-"`
}

// Concrete implementation
func (f *Feed) Concrete() api.Feed {
	return f.concrete
}

// Build the feed from specification
func (f *Feed) Build(name string, tags []string, observableType string, puller puller_api.Puller, parser parser_api.Parser, defaults map[string]interface{}) (api.Feed, error) {
	var err error

	switch f.Type {
	default:
		f.concrete, err = generic.New(
			generic.WithName(name),
			generic.WithTags(tags),
			generic.WithObservable(observableType),
			generic.WithPuller(puller),
			generic.WithParser(parser),
			generic.WithDefaults(defaults),
		)
	}

	return f.concrete, err
}
