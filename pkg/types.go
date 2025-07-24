package pkg

import "context"

type Task interface {
	Name() string
	Execute(ctx context.Context) error
}
