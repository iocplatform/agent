package generic

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/iocplatform/agent/pkg/feed/api"
)

type genericFeed struct {
	api.DefaultFeed

	dopts Options
}

// New build a generic feed object
func New(opts ...Option) (api.Feed, error) {
	feed := &genericFeed{
		DefaultFeed: api.DefaultFeed{
			Type: "generic",
		},
	}

	// Assign all options
	for _, opt := range opts {
		opt(&feed.dopts)
	}

	// Internal values

	if feed.dopts.puller != nil {
		feed.SetPuller(feed.dopts.puller)
	}
	if feed.dopts.parser != nil {
		feed.SetParser(feed.dopts.parser)
	}
	if len(strings.TrimSpace(feed.dopts.name)) > 0 {
		feed.Name = feed.dopts.name
	}
	if len(feed.dopts.tags) > 0 {
		feed.Tags = feed.dopts.tags
	}

	return feed, nil
}

//- Implementations ------------------------------------------------------------

func (f *genericFeed) Run(ctx context.Context) (*api.FeedStatistics, error) {
	// Check execution condition
	puller := f.Puller()
	parser := f.Parser()

	if puller == nil {
		return nil, api.ErrNoPullerDefined
	}
	if parser == nil {
		return nil, api.ErrNoParserDefined
	}

	err := puller.Pull(ctx, f)
	if err != nil {
		return nil, fmt.Errorf("Error during collect process : %v", err)
	}

	// Build process metrics
	metrics := &api.FeedStatistics{
		ProcessedLineCount: parser.GetProcessedCount(),
		IgnoredLineCount:   parser.GetIgnoredCount(),
		ParsedLineCount:    parser.GetParsedCount(),
	}

	return metrics, nil
}

func (f *genericFeed) Dispatch(ctx context.Context, fields map[string]interface{}) error {
	// Add feed name
	fields["id"] = f.GetName()

	// Enhance fields with defaults values
	for k, v := range f.dopts.defaults {
		switch k {
		case "tlp", "threat", "severity", "boost", "confidence", "details_url", "observable":
			if _, ok := fields[k]; !ok {
				fields[k] = v
			}
		default:
			log.Printf("Ignored feed attributes '%s' => '%v'", k, v)
		}
	}

	// Switch strategy from fields content
	switch fields["content_type"] {
	case "lines":
		lines := fields["body"].([]string)

		// Do some cleanup to prevent fields propagation
		delete(fields, "body")
		delete(fields, "content_type")

		for i, l := range lines {
			item, err := f.dopts.parser.Feed([]byte(l))
			if err != nil {
				log.Printf("Unable to parse line %d - '%s' : %v", i, l, err)
				continue
			}

			// Add feed tags
			if _, ok := item["tags"]; !ok {
				item["tags"] = f.Tags
			}

			// Assign observable type
			fields["observable"] = f.dopts.observableType

			spew.Dump(item)
		}

	default:
		log.Printf("Unhandled message body '%s'", fields["content_type"])
	}

	return nil
}
