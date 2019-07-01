package api

import "errors"

var (
	// ErrDataNoMatchDeclaredFields is raised when field count is different from the data field count
	ErrDataNoMatchDeclaredFields = errors.New("Data does not match declared fields")

	// ErrDataNotMatchExtractionRule is raised when no parser can process the data line
	ErrDataNotMatchExtractionRule = errors.New("Data does not match any parser rule")

	// ErrDataIgnored is raised when parser ignores the line
	ErrDataIgnored = errors.New("Data is ignored by the parser")

	// ErrNoRuleDefined is raised when parser no rule is defined
	ErrNoRuleDefined = errors.New("No rule defined for the parser")

	// ErrDataEmpty is raised when extraction returned no data
	ErrDataEmpty = errors.New("No data after extraction")

	// ErrDateFieldNotFound is raised when extraction does not contain the specified date field
	ErrDateFieldNotFound = errors.New("The given date field could not be found")

	// ErrTypeFieldNotFound is raised when extraction does not contains the specified type field
	ErrTypeFieldNotFound = errors.New("The given type field could not be found")
)
