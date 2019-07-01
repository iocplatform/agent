package descriptor

// Agent is the information holder for agent definition
type Agent struct {
	Enabled     bool     `yaml:"enabled"`
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Tags        []string `yaml:"tags"`

	// Default attributes
	Defaults map[string]interface{} `yaml:"defaults"`

	// Puller definition
	Puller Puller `yaml:"puller"`

	// Parser definition
	Parser Parser `yaml:"parser"`

	// Feeds
	Feeds []*Feed `yaml:"feeds"`
}
