package api

// DefaultPuller implements default puller behavior
type DefaultPuller struct {
	Name string
}

// GetName returns the puller name
func (c *DefaultPuller) GetName() string {
	return c.Name
}
