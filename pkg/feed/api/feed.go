package api

import (
	"context"

	parser_api "github.com/iocplatform/agent/pkg/parser/api"
	puller_api "github.com/iocplatform/agent/pkg/puller/api"
)

// FeedStatistics is the fee process metrics holder
type FeedStatistics struct {
	ProcessedLineCount int
	IgnoredLineCount   int
	ParsedLineCount    int
}

// Feed is the default contract holder for feeds
type Feed interface {
	GetName() string
	GetTags() []string
	Parser() parser_api.Parser
	Puller() puller_api.Puller
	Run(ctx context.Context) (*FeedStatistics, error)
}
