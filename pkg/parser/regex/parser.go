package regex

import (
	"strings"

	"github.com/iocplatform/agent/pkg/parser/api"
)

type regexParser struct {
	api.DefaultParser

	dopts regexOptions
}

// New returns an instance of regex parser
func New(opts ...Option) (api.Parser, error) {
	parser := &regexParser{
		DefaultParser: api.DefaultParser{
			Type: "regex",
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

func (p *regexParser) Feed(input []byte) (map[string]interface{}, error) {
	p.DefaultParser.LineProcessed++

	if p.dopts.skipLines > 0 {
		if p.DefaultParser.GetProcessedCount() <= p.dopts.skipLines {
			p.DefaultParser.LineIgnored++
			return nil, api.ErrDataIgnored
		}
	}

	result := map[string]interface{}{}
	result["raw"] = input

	// Check ignores rules
	if p.dopts.ignoreRule != nil {
		// Check if record match
		if p.dopts.ignoreRule.Match(input) {
			p.DefaultParser.LineIgnored++
			return nil, api.ErrDataIgnored
		}
	}

	if p.dopts.rule != nil {
		// Check if record match
		if !p.dopts.rule.Match(input) {
			p.DefaultParser.LineIgnored++
			return nil, api.ErrDataNotMatchExtractionRule
		}
	} else {
		return nil, api.ErrNoRuleDefined
	}

	// Extract groups
	cols := p.dopts.rule.FindSubmatch(input)

	// Check field count
	if len(p.dopts.fields) != len(cols)-1 {
		return nil, api.ErrDataNoMatchDeclaredFields
	}

	// If user use named match groups
	if len(p.dopts.rule.SubexpNames()) > 0 {
		for i, key := range p.dopts.rule.SubexpNames() {
			if len(key) > 0 {
				result[key] = strings.TrimSpace(string(cols[i]))
			}
		}
	} else {
		return nil, api.ErrDataEmpty
	}

	p.DefaultParser.LineParsed++
	return result, nil
}
