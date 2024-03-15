package redis

import (
	"github.com/maiquocthinh/go-comic/config"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewRedisClient(cfg *config.RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:        cfg.Addr,
		Username:    cfg.Username,
		Password:    cfg.Password,
		DB:          cfg.DB,
		PoolSize:    cfg.PoolSize,
		PoolTimeout: time.Duration(cfg.PoolTimeout) * time.Second,
	})
}
