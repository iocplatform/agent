package api

import "context"

// Puller defines pull strategy contract
type Puller interface {
	Pull(ctx context.Context) error
}
