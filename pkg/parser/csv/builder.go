package csv

import (
	"strings"

	"github.com/iocplatform/agent/pkg/parser"
	"github.com/iocplatform/agent/pkg/parser/api"
)

// Build a parser from a parameter map
func Build(parameters map[string]interface{}) (api.Parser, error) {
	var options []Option

	for key, value := range parameters {
		switch key {
		case "skip_lines":
			options = append(options, WithSkipLines(value.(int)))
		case "separator":
			options = append(options, WithSeparator(value.(string)))
		case "fields":
			var fields []string
			for _, f := range value.([]interface{}) {
				fields = append(fields, f.(string))
			}
			options = append(options, WithFields(fields...))
		case "tag_field":
			options = append(options, WithTagField(strings.TrimSpace(value.(string))))
		case "tag_separator":
			options = append(options, WithTagSeparator(value.(string)))
		case "tag_map":
			tagMap := map[string][]string{}
			for k, v := range value.(map[interface{}]interface{}) {
				for _, t := range v.([]interface{}) {
					tagMap[k.(string)] = append(tagMap[k.(string)], t.(string))
				}
			}
			options = append(options, WithTagMap(tagMap))
		case "date_field":
			options = append(options, WithDateField(value.(string)))
		case "date_format":
			options = append(options, WithDateFormat(value.(string)))
		case "type_field":
			options = append(options, WithTypeField(value.(string)))
		default:
		}
	}

	return New(options...)
}

func init() {
	parser.Add("csv", func(parameters map[string]interface{}) (api.Parser, error) {
		return Build(parameters)
	})
}
