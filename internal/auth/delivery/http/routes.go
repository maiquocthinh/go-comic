package http

import "github.com/gin-gonic/gin"

func (h *authHandlers) MapComicRotes(route *gin.RouterGroup) {
	route.POST("/register", h.Register())
	route.POST("/login", h.Login())
}