package parser

import "github.com/iocplatform/agent/pkg/parser/api"

// Builder is a parser factory
type Builder func(map[string]interface{}) (api.Parser, error)

// Parsers is the map of supported parsers
var Parsers = map[string]Builder{}

// Add a collector builder
func Add(name string, builder Builder) {
	Parsers[name] = builder
}
