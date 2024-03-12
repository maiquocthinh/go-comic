package http

import "github.com/gin-gonic/gin"

func (h *comicHandlers) MapComicRotes(route *gin.RouterGroup) {
	route.GET(":id", h.GetComic())
	route.GET("", h.List())
}
