package dispatcher

import "context"

// Dispatcher is the dispatcher contract holder
type Dispatcher interface {
	Dispatch(ctx context.Context, record map[string]interface{}) error
}
