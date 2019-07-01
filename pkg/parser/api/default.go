package api

// DefaultParser implementation
type DefaultParser struct {
	Type string `yaml:"type"`

	LineProcessed int
	LineParsed    int
	LineIgnored   int
}

// GetName returns the parser name
func (p *DefaultParser) GetName() string {
	return p.Type
}

// GetProcessedCount returns the line processed count
func (p *DefaultParser) GetProcessedCount() int {
	return p.LineProcessed
}

// GetParsedCount returns the line parsed count
func (p *DefaultParser) GetParsedCount() int {
	return p.LineParsed
}

// GetIgnoredCount returns the line ignored count
func (p *DefaultParser) GetIgnoredCount() int {
	return p.LineIgnored
}

// Reset all counter
func (p *DefaultParser) Reset() {
	p.LineProcessed = 0
	p.LineParsed = 0
	p.LineIgnored = 0
}
