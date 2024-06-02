package http

import "github.com/gin-gonic/gin"

func (h *authHandlers) MapComicRotes(route *gin.RouterGroup) {
	route.POST("/register", h.Register())
	route.POST("/login", h.Login())
	route.DELETE("/logout", h.mm.RequiredAuthJWTMiddleware(), h.Logout())
	route.POST("/reset-password", h.ResetPassword())
	route.POST("/reset-password/send-code", h.SendCodeResetPassword())
}
