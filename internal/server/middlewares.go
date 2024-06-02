package server

import (
	"github.com/gin-contrib/cors"
	"github.com/maiquocthinh/go-comic/internal/middleware"
	"time"
)

func (s Server) useMiddlewares() {

	s.gin.Use(cors.New(cors.Config{
		AllowOrigins:     s.config.CorsPolicy.AllowOrigins,
		AllowMethods:     s.config.CorsPolicy.AllowMethods,
		AllowHeaders:     s.config.CorsPolicy.AllowHeaders,
		ExposeHeaders:    s.config.CorsPolicy.ExposeHeaders,
		AllowCredentials: s.config.CorsPolicy.AllowCredentials,
		MaxAge:           time.Duration(s.config.CorsPolicy.MaxAge) * time.Second,
	}))

	s.gin.Use(middleware.ErrorLogger(), middleware.Recovery()) // don't change order

}
