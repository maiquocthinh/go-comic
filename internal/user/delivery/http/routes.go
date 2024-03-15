package http

import "github.com/gin-gonic/gin"

func (h *userHandlers) MapComicRotes(route *gin.RouterGroup) {
	route.GET("/profile", h.GetProfile())
	route.POST("/profile", h.UpdateProfile())
	route.PATCH("/change-password", h.ChangePassword())
	route.GET("/comments", h.GetComments())
}
