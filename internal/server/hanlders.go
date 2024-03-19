package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maiquocthinh/go-comic/internal/middleware"
	"github.com/maiquocthinh/go-comic/pkg/uploadprovider"

	authHttp "github.com/maiquocthinh/go-comic/internal/auth/delivery/http"
	authRepository "github.com/maiquocthinh/go-comic/internal/auth/repository"
	authUseCase "github.com/maiquocthinh/go-comic/internal/auth/usecase"

	comicHttp "github.com/maiquocthinh/go-comic/internal/comic/delivery/http"
	comicRepository "github.com/maiquocthinh/go-comic/internal/comic/repository"
	comicUseCase "github.com/maiquocthinh/go-comic/internal/comic/usecase"

	userHttp "github.com/maiquocthinh/go-comic/internal/user/delivery/http"
	userRepository "github.com/maiquocthinh/go-comic/internal/user/repository"
	userUseCase "github.com/maiquocthinh/go-comic/internal/user/usecase"

	commentHttp "github.com/maiquocthinh/go-comic/internal/comment/delivery/http"
	commentRepository "github.com/maiquocthinh/go-comic/internal/comment/repository"
	commentUseCase "github.com/maiquocthinh/go-comic/internal/comment/usecase"
)

func (s *Server) mapHandlers() error {
	// Init providers
	dropboxProvider := uploadprovider.NewDropboxProvider(&s.config.Dropbox)

	// Init repositories
	authRepo := authRepository.NewAuthRepository(s.mysqlDB)
	authRedisRepo := authRepository.NewAuthRedisRepository(s.redisClient)
	comicRepo := comicRepository.NewComicRepository(s.mysqlDB)
	userRepo := userRepository.NewUserRepository(s.mysqlDB)
	commentRepo := commentRepository.CommentRepository(s.mysqlDB)

	// Init useCases
	authUC := authUseCase.NewAuthUseCase(s.config, authRepo, authRedisRepo)
	comicUC := comicUseCase.NewComicUseCase(comicRepo)
	userUC := userUseCase.NewUserUseCase(userRepo, dropboxProvider)
	commentUC := commentUseCase.NewCommentUseCase(commentRepo)

	// New middleware manager
	middlewareManager := middleware.NewMiddlewareManager(s.config, s.pubsub, authRedisRepo)

	// Init handlers
	authHandlers := authHttp.NewComicHandlers(middlewareManager, authUC)
	comicHandlers := comicHttp.NewComicHandlers(middlewareManager, comicUC)
	userHandlers := userHttp.NewUserHandlers(middlewareManager, userUC)
	commentHandlers := commentHttp.NewCommentHandlers(middlewareManager, commentUC)

	// Use middleware
	s.gin.Use(middleware.ErrorLogger(), middleware.Recovery()) // don't change order

	// Map Handlers
	v1 := s.gin.Group("/api/v1")

	authRoutes := v1.Group("/auth")
	authHandlers.MapComicRotes(authRoutes)

	comicRoutes := v1.Group("/comics")
	comicHandlers.MapComicRotes(comicRoutes)
	commentHandlers.MapComicRotes(comicRoutes)

	userRoutes := v1.Group("/user")
	userHandlers.MapComicRotes(userRoutes)

	s.gin.GET("/ping", func(ctx *gin.Context) {
		status := make(map[string]string)

		// check health of mysql & redis
		if err := s.mysqlDB.Ping(); err != nil {
			status["Mysql"] = fmt.Sprintf("MySQL connection error: %s", err.Error())
		} else {
			status["Mysql"] = "MySQL connection OK"
		}

		// check health of redis
		if _, err := s.redisClient.Ping(context.Background()).Result(); err != nil {
			status["Redis"] = fmt.Sprintf("Redis connection error: %s", err.Error())
		} else {
			status["Redis"] = "Redis connection OK"
		}

		ctx.JSON(http.StatusOK, &gin.H{
			"Status": status,
		})
	})

	return nil
}
