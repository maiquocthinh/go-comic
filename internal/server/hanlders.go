package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maiquocthinh/go-comic/internal/middleware"

	authHttp "github.com/maiquocthinh/go-comic/internal/auth/delivery/http"
	authRepository "github.com/maiquocthinh/go-comic/internal/auth/repository"
	authUseCase "github.com/maiquocthinh/go-comic/internal/auth/usecase"

	comicHttp "github.com/maiquocthinh/go-comic/internal/comic/delivery/http"
	comicRepository "github.com/maiquocthinh/go-comic/internal/comic/repository"
	comicUseCase "github.com/maiquocthinh/go-comic/internal/comic/usecase"
)

func (s *Server) mapHandlers() error {

	// Init repositories
	authRepo := authRepository.NewAuthRepository(s.mysqlDB)
	comicRepo := comicRepository.NewComicRepository(s.mysqlDB)

	// Init useCases
	authUC := authUseCase.NewAuthUseCase(authRepo)
	comicUC := comicUseCase.NewComicUseCase(comicRepo)

	// Init handlers
	authHandlers := authHttp.NewComicHandlers(authUC)
	comicHandlers := comicHttp.NewComicHandlers(comicUC)

	// Use middleware
	s.gin.Use(middleware.ErrorLogger(), middleware.Recovery()) // don't change order

	// Map Handlers
	v1 := s.gin.Group("/api/v1")

	authRoutes := v1.Group("/auth")
	authHandlers.MapComicRotes(authRoutes)

	comicRoutes := v1.Group("/comics")
	comicHandlers.MapComicRotes(comicRoutes)

	s.gin.GET("/ping", func(ctx *gin.Context) {
		if err := s.mysqlDB.Ping(); err != nil {
			ctx.JSON(http.StatusInternalServerError, &gin.H{
				"Message": fmt.Sprintf("MySQL Error: %s", err.Error()),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, &gin.H{
			"Message": "Mysql connected success",
		})
	})

	return nil
}
