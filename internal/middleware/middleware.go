package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/maiquocthinh/go-comic/config"
	authRepository "github.com/maiquocthinh/go-comic/internal/auth/repository"
	"github.com/maiquocthinh/go-comic/pkg/pubsub"
)

type middlewareManager struct {
	cfg           *config.Config
	pubsub        pubsub.PubSub
	authRedisRepo authRepository.AuthRedisRepository
}

func NewMiddlewareManager(cfg *config.Config, pubsub pubsub.PubSub, authRedisRepo authRepository.AuthRedisRepository) *middlewareManager {
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
