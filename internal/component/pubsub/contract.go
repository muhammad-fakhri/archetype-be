package pubsub

import (
	"context"
)

type Publisher interface {
	Publish(ctx context.Context, payload interface{}, attributes map[string]string) (err error)
}

type Subscriber interface {
	Start() (stopFunc func())
	Stop()
}
