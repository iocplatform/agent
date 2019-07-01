package descriptor

// Build a agent instance from agent definition
func Build(config Agent) (*Agent, error) {
	_, err := config.Puller.Build()
	if err != nil {
		return nil, err
	}

	_, err = config.Parser.Build()
	if err != nil {
		return nil, err
	}

	return &config, nil
}
