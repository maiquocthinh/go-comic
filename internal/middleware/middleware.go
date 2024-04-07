package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/maiquocthinh/go-comic/config"
	"github.com/maiquocthinh/go-comic/pkg/pubsub"
)

type authRedisRepository interface {
	IsTokenInBlackList(ctx context.Context, jti string) (bool, error)
}

type middlewareManager struct {
	cfg           *config.Config
	pubsub        pubsub.PubSub
	authRedisRepo authRedisRepository
}

func NewMiddlewareManager(cfg *config.Config, pubsub pubsub.PubSub, authRedisRepo authRedisRepository) MiddlewareManager {
	return &middlewareManager{
		cfg:           cfg,
		pubsub:        pubsub,
		authRedisRepo: authRedisRepo,
	}
}

type MiddlewareManager interface {
	RequiredAuthJWTMiddleware() gin.HandlerFunc
	OptionalAuthJWTMiddleware() gin.HandlerFunc
	WriteHistory() gin.HandlerFunc
	IncreaseView() gin.HandlerFunc
}
