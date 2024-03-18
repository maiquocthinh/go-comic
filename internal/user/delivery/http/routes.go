package http

import "github.com/gin-gonic/gin"

func (h *userHandlers) MapComicRotes(route *gin.RouterGroup) {
	route.GET("/profile", h.mm.AuthJWTMiddleware(), h.mm.VerifyJWTMiddleware(), h.GetProfile())
	route.PATCH("/profile", h.mm.AuthJWTMiddleware(), h.mm.VerifyJWTMiddleware(), h.UpdateProfile())
	route.PUT("/profile/avatar", h.mm.AuthJWTMiddleware(), h.mm.VerifyJWTMiddleware(), h.UpdateAvatar())
	route.PUT("/change-password", h.mm.AuthJWTMiddleware(), h.mm.VerifyJWTMiddleware(), h.ChangePassword())
	route.GET("/comments", h.GetComments())
}
