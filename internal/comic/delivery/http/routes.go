package http

import "github.com/gin-gonic/gin"

func (h *comicHandlers) MapComicRotes(route *gin.RouterGroup) {
	route.GET(":comicID", h.GetComic())
	route.GET(":comicID/chapter/:chapterID", h.mm.WriteHistory(), h.GetChapterOfComic())
	route.GET("/search/", h.SearchComic())
	route.GET("", h.List())
}
