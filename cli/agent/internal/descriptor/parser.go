package descriptor

import (
	"github.com/iocplatform/agent/pkg/parser/api"
	"github.com/iocplatform/agent/pkg/parser/csv"
	"github.com/iocplatform/agent/pkg/parser/regex"
)

// Parser is the information holder for agent's parser definition
type Parser struct {
	Type       string                 `yaml:"type"`
	Parameters map[string]interface{} `yaml:"parameters"`

	// Concrete parser implementation
	concrete api.Parser `yaml:"-"`
}

// Concrete implementation
func (p *Parser) Concrete() api.Parser {
	return p.concrete
}

// Build a conrete implementation from definition
func (p *Parser) Build() (api.Parser, error) {
	var err error

	switch p.Type {
	case "csv":
		p.concrete, err = csv.Build(p.Parameters)
	case "regex":
		p.concrete, err = regex.Build(p.Parameters)
	default:
	}

	return p.concrete, err
}
