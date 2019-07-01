package descriptor

import (
	"fmt"

	"github.com/iocplatform/agent/pkg/puller/api"
	"github.com/iocplatform/agent/pkg/puller/http"
)

// Puller is the information holder for grabber's collector definition
type Puller struct {
	Type       string                 `yaml:"type"`
	Parameters map[string]interface{} `yaml:"parameters"`

	concrete api.Puller `yaml:"-"`
}

// Concrete implementation
func (p *Puller) Concrete() api.Puller {
	return p.concrete
}

// Build a conrete implementation from definition
func (p *Puller) Build() (api.Puller, error) {
	var err error

	switch p.Type {
	case "http":
		p.concrete, err = http.Build(p.Parameters)
	default:
	}

	if p.concrete == nil {
		return nil, fmt.Errorf("unsupported puller type '%s'", p.Type)
	}

	return p.concrete, err
}
