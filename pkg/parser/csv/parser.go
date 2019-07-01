package csv

import (
	"bytes"
	"encoding/csv"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/tv42/slug"

	"github.com/iocplatform/agent/pkg/parser/api"
)

type csvParser struct {
	api.DefaultParser

	dopts csvOptions
}

// New build a CSV parser object
func New(opts ...Option) (api.Parser, error) {
	parser := &csvParser{
		DefaultParser: api.DefaultParser{
			Type: "csv",
		},
	}

	// Reset counter
	parser.Reset()

	// Assign all options
	for _, opt := range opts {
		opt(&parser.dopts)
	}

	return parser, nil
}

// -----------------------------------------------------------------------------

func (p *csvParser) Feed(input []byte) (map[string]interface{}, error) {
	p.DefaultParser.LineProcessed++

	if p.dopts.skipLines > 0 {
		if p.DefaultParser.GetProcessedCount() < p.dopts.skipLines {
			p.DefaultParser.LineIgnored++
			return nil, api.ErrDataIgnored
		}
	}

	result := map[string]interface{}{}

	// CSV Reader
	reader := csv.NewReader(bytes.NewReader(input))
	reader.Comma, _ = utf8.DecodeRuneInString(p.dopts.separator)

	// Get all columns
	cols, err := reader.Read()
	if err != nil {
		return nil, err
	}

	// Check field count
	if len(p.dopts.fields) != len(cols) {
		p.DefaultParser.LineIgnored++
		return nil, api.ErrDataNoMatchDeclaredFields
	}

	// Map all columns to declared fields
	for i, key := range p.dopts.fields {
		result[key] = cols[i]
	}

	// Parse tags
	if len(p.dopts.tagField) > 0 {

		var rawTags []string

		// If a separator is defined
		if len(p.dopts.tagSeparator) > 0 {
			rawTags = strings.Split(result[p.dopts.tagField].(string), p.dopts.tagSeparator)
		} else {
			rawTags = append(rawTags, result[p.dopts.tagField].(string))
		}

		// Slugify all tags
		for i, tag := range rawTags {
			rawTags[i] = slug.Slug(tag)
		}

		// Translate tags if there is a translation map
		newTags := []string{}
		for tagReplacement, mappedTags := range p.dopts.tagMap {
			for _, origTag := range rawTags {
				if tagReplacement == origTag {
					for _, tag := range mappedTags {
						newTags = append(newTags, tag)
					}
				}
			}
		}

		if len(newTags) > 0 {
			result["tags"] = newTags
		} else {
			result["tags"] = rawTags
		}
	}

	// Parse date
	if len(p.dopts.dateField) > 0 {
		dateRaw, ok := result[p.dopts.dateField]
		if !ok {
			return nil, api.ErrDateFieldNotFound
		}

		fieldDate := time.Now().UTC()
		if len(p.dopts.dateFormat) > 0 {
			if p.dopts.dateFormat == "unix" {
				ts, _ := strconv.ParseInt(dateRaw.(string), 10, 64)
				fieldDate = time.Unix(ts, 0)
			} else {
				// Try to parse
				fieldDate, err = time.Parse(p.dopts.dateFormat, dateRaw.(string))
				if err != nil {
					return nil, err
				}
			}
		}

		// Overrides the timestamp
		result["_timestamp"] = fieldDate.UTC().Unix()
	} else {
		// Overrides the timestamp
		result["_timestamp"] = time.Now().UTC().Unix()
	}

	// Assign the raw input
	result["raw"] = input

	p.DefaultParser.LineParsed++
	return result, nil
}
