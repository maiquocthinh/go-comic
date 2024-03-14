package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/maiquocthinh/go-comic/config"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

const (
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
)

type Server struct {
	gin         *gin.Engine
	config      *config.Config
	mysqlDB     *sqlx.DB
	redisClient *redis.Client
	logger      *logrus.Logger
}

func NewServer(cfg *config.Config, mysqlDB *sqlx.DB, redisClient *redis.Client) *Server {
	return &Server{
		gin:         gin.Default(),
		config:      cfg,
		mysqlDB:     mysqlDB,
		redisClient: redisClient,
		logger:      logrus.New(),
	}
}

func (s *Server) Run() error {
	svr := &http.Server{
		Addr:           s.config.Server.Port,
		Handler:        s.gin,
		ReadTimeout:    time.Second * s.config.Server.ReadTimeout,
		WriteTimeout:   time.Second * s.config.Server.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		s.logger.Infof("Server is listening on PORT: %s", s.config.Server.Port)
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server listen error: %s\n", err)
		}
	}()

	// Map Handlers
	if err := s.mapHandlers(); err != nil {
		return err
	}

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	s.logger.Info("Server Exited Properly")
	return svr.Shutdown(ctx)
}
