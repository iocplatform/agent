package api

// Parser is the default parser contract
type Parser interface {
	GetName() string
	Feed(input []byte) (map[string]interface{}, error)
	GetParsedCount() int
	GetIgnoredCount() int
	GetProcessedCount() int
	Reset()
}
