package descriptor

import (
	"context"
	"fmt"
	"strings"
)

// Run a agent definition
func Run(ctx context.Context, agent *Agent) error {
	if !agent.Enabled {
		return fmt.Errorf("agent is disabled by configuration")
	}

	// Get defaults attributes
	defaults := agent.Defaults

	// Run puller
	puller := agent.Puller.Concrete()
	if puller == nil {
		return fmt.Errorf("puller not defined")
	}

	// Run parser
	parser := agent.Parser.Concrete()
	if parser == nil {
		return fmt.Errorf("parser not defined")
	}

	// Run all feeds
	for _, f := range agent.Feeds {
		parser.Reset()

		// Overrides default parameters
		puller.SetParameters(f.Overrides)

		// Extract tags
		var tags []string
		if value, ok := f.Overrides["tags"]; ok {
			for _, t := range value.([]interface{}) {
				tags = append(tags, strings.ToLower(strings.TrimSpace(t.(string))))
			}
		}

		// Create a feed object
		feed, err := f.Build(
			fmt.Sprintf("feed:%s:%s", strings.ToLower(strings.TrimSpace(agent.Name)), strings.ToLower(strings.TrimSpace(f.Name))),
			tags,
			"",
			puller, parser, defaults)
		if err != nil {
			return err
		}

		// Process the feed
		_, err = feed.Run(ctx)
		if err != nil {
			return err
		}
	}

	// No error
	return nil
}
