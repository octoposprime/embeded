package notify

import (
	"context"
	"sync"

	observer "github.com/octoposprime/op-go-os/pkg/observer"
	"github.com/octoposprime/op-go-os/tool/logger"
)

var lock = &sync.Mutex{}

var notifierInstance *Notifier

func init() {
	GetNotifierInstance()
}

type Notifier struct {
	Subscribers map[string]observer.Observable
}

func GetNotifierInstance() *Notifier {
	if notifierInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if notifierInstance == nil {
			notifierInstance = &Notifier{}
			notifierInstance.Subscribers = make(map[string]observer.Observable)
		}
	}
	return notifierInstance
}

func (n *Notifier) Subscribe(ctx context.Context, observer observer.Observable, keyId string) error {
	n.Subscribers[keyId] = observer
	logger.GetLoggerInstance().Debugf("Subscribed to Notifier with keyId: %s\n", keyId)
	return nil
}

func (n *Notifier) UnSubscribe(ctx context.Context, keyId string) error {
	delete(n.Subscribers, keyId)
	return nil
}

func (n *Notifier) Notify(ctx context.Context) error {
	return nil
}
