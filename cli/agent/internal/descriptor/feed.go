package descriptor

// Feed is the a feed information holder for agent
type Feed struct {
	// Feed type
	Type string `yaml:"type"`

	// Feed name
	Name string `yaml:"name"`

	// Overrides agent settings
	Overrides map[string]interface{} `yaml:"overrides"`

	// Affected tags
	Tags []string `yaml:"tags"`
}
