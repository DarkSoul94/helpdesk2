package helpdesk

import "context"

// Usecase ...
type Usecase interface {
	HelloWorld(ctx context.Context)
}
