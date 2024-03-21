package http

import "github.com/gin-gonic/gin"

func (h *genreHandlers) MapComicRotes(route *gin.RouterGroup) {
	route.GET("", h.GetAllGenres())
}
