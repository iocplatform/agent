package api

import "errors"

var (
	// ErrNoPullerDefined is raised when no puller is affected to the feed
	ErrNoPullerDefined = errors.New("no puller defined")
	// ErrNoParserDefined is raised when no parser is affected the to feed
	ErrNoParserDefined = errors.New("no parser defined")
)
