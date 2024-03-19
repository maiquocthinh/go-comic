package redis_pubsub

import (
	"context"
	"github.com/maiquocthinh/go-comic/pkg/pubsub"
	"github.com/redis/go-redis/v9"
)

type redisPubSub struct {
	redisClient *redis.Client
}

func NewRedisPubSub(redisClient *redis.Client) *redisPubSub {
	return &redisPubSub{redisClient: redisClient}
}

func (ps *redisPubSub) Publish(ctx context.Context, channel, msg string) error {
	return ps.redisClient.Publish(ctx, channel, msg).Err()
}

func (ps *redisPubSub) Subscribe(ctx context.Context, chanel string) pubsub.Sub {
	return &redisSub{
		chanel:  chanel,
		redisPS: ps.redisClient.Subscribe(ctx, chanel),
	}
}

type redisSub struct {
	chanel  string
	redisPS *redis.PubSub
}

func (sl *redisSub) UnSubscribe(ctx context.Context) {
	sl.redisPS.Unsubscribe(ctx, sl.chanel)
}

func (sl *redisSub) ListenAndReceive(hdl func(msg interface{})) {
	ch := sl.redisPS.Channel()
	for msg := range ch {
		hdl(msg.Payload)
	}
}
