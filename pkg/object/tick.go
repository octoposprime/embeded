package object

import "context"

type Tick struct {
}

func (t *Tick) Start(ctx context.Context) error {
	return nil
}
