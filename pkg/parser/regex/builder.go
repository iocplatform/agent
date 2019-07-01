package regex

import (
	"github.com/iocplatform/agent/pkg/parser"
	"github.com/iocplatform/agent/pkg/parser/api"
)

// Build a parser from a parameter map
func Build(parameters map[string]interface{}) (api.Parser, error) {
	var options []Option

	for key, value := range parameters {
		switch key {
		case "rule":
			options = append(options, WithRule(value.(string)))
		case "ignore_rule":
			options = append(options, WithIgnoreRule(value.(string)))
		case "fields":
			var fields []string
			for _, f := range value.([]interface{}) {
				fields = append(fields, f.(string))
			}
			options = append(options, WithFields(fields...))
		default:
		}
	}

	return New(options...)
}

func init() {
	parser.Add("regex", func(parameters map[string]interface{}) (api.Parser, error) {
		return Build(parameters)
	})
}
