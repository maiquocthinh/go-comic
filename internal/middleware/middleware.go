package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/maiquocthinh/go-comic/config"
	authRepository "github.com/maiquocthinh/go-comic/internal/auth/repository"
)

type middlewareManager struct {
	cfg           *config.Config
	authRedisRepo authRepository.AuthRedisRepository
}

func NewMiddlewareManager(cfg *config.Config, authRedisRepo authRepository.AuthRedisRepository) *middlewareManager {
	return &middlewareManager{
		cfg:           cfg,
		authRedisRepo: authRedisRepo,
	}
}

type MiddlewareManager interface {
	AuthJWTMiddleware() gin.HandlerFunc
	VerifyJWTMiddleware() gin.HandlerFunc
}
