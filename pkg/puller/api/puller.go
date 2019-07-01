package api

import (
	"context"

	"github.com/iocplatform/agent/pkg/dispatcher"
)

// Puller defines pull strategy contract
type Puller interface {
	GetName() string
	SetParameters(map[string]interface{})
	Pull(ctx context.Context, d dispatcher.Dispatcher) error
}
