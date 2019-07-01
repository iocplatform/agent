package feed

import "github.com/iocplatform/agent/pkg/feed/api"

// Builder is a collectory factory
type Builder func(map[string]interface{}) (api.Feed, error)

// Feeds is the map of supported collectors
var Feeds = map[string]Builder{}

// Add a collector builder
func Add(name string, builder Builder) {
	Feeds[name] = builder
}
