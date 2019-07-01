package puller

import "github.com/iocplatform/agent/pkg/puller/api"

// Builder is a puller factory
type Builder func(map[string]interface{}) (api.Puller, error)

// Pullers is the map of supported pullers
var Pullers = map[string]Builder{}

// Add a collector builder
func Add(name string, builder Builder) {
	Pullers[name] = builder
}
