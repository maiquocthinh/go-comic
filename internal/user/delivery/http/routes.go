package http

import "github.com/gin-gonic/gin"

func (h *userHandlers) MapComicRotes(route *gin.RouterGroup) {
	route.GET("/profile", h.mm.RequiredAuthJWTMiddleware(), h.GetProfile())
	route.PATCH("/profile", h.mm.RequiredAuthJWTMiddleware(), h.UpdateProfile())
	route.PUT("/profile/avatar", h.mm.RequiredAuthJWTMiddleware(), h.UpdateAvatar())
	route.PUT("/change-password", h.mm.RequiredAuthJWTMiddleware(), h.ChangePassword())
	route.GET("/histories", h.mm.RequiredAuthJWTMiddleware(), h.GetHistoryView())
	route.GET("/comments", h.mm.RequiredAuthJWTMiddleware(), h.GetComments())
}
