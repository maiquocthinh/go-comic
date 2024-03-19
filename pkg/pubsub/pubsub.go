package pubsub

import (
	"context"
)

type Sub interface {
	UnSubscribe(ctx context.Context)
	ListenAndReceive(hdl func(msg interface{}))
}

type PubSub interface {
	Publish(ctx context.Context, channel, msg string) error
	Subscribe(ctx context.Context, chanel string) Sub
}
