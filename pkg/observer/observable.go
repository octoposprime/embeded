package observer

import "context"

type Observable interface {
	Subscribe(ctx context.Context) error
	UnSubscribe(ctx context.Context) error
}
