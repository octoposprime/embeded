package object

import "context"

type Keyable interface {
	Start(ctx context.Context) error
}
