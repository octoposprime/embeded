package object

import (
	"context"

	notify "github.com/octoposprime/op-go-os/pkg/notify"
	_ "github.com/octoposprime/op-go-os/pkg/observer"
)

type Dlr struct {
	KeyId string
}

func NewDlr(keyId string) *Dlr {
	dlr := &Dlr{
		KeyId: keyId,
	}
	dlr.Subscribe(context.Background())
	return dlr
}

func (d *Dlr) Subscribe(ctx context.Context) error {
	notify.GetNotifierInstance().Subscribe(ctx, d, d.KeyId)
	return nil
}

func (d *Dlr) UnSubscribe(ctx context.Context) error {
	notify.GetNotifierInstance().UnSubscribe(ctx, d.KeyId)
	return nil
}
