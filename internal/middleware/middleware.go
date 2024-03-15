package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/maiquocthinh/go-comic/config"
)

type middlewareManager struct {
	cfg *config.Config
}

func NewMiddlewareManager(cfg *config.Config) *middlewareManager {
	return &middlewareManager{
		cfg: cfg,
	}
}

type MiddlewareManager interface {
	AuthJWTMiddleware() gin.HandlerFunc
}
