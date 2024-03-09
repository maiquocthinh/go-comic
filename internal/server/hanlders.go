package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) mapHandlers() error {

	// Init repositories

	// Init useCases

	// Init handlers

	// Use middleware

	// Map Handlers
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
