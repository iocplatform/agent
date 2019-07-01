package generic

import (
	"github.com/iocplatform/agent/pkg/feed"
	"github.com/iocplatform/agent/pkg/feed/api"
)

// Build a parser from a parameter map
func Build(parameters map[string]interface{}) (api.Feed, error) {
	var options []Option

	return New(options...)
}

func init() {
	feed.Add("generic", func(parameters map[string]interface{}) (api.Feed, error) {
		return Build(parameters)
	})
}
