package http

import "github.com/gin-gonic/gin"

func (h *commentHandlers) MapComicRotes(route *gin.RouterGroup) {
	route.GET(":comicID/chapter/:chapterID/comments", h.GetCommentsOfChapter())
	route.GET(":comicID/chapter/:chapterID/comments/:commentID/replies", h.GetRepliesOfComment())
	route.POST(":comicID/chapter/:chapterID/comments", h.mm.AuthJWTMiddleware(), h.mm.VerifyJWTMiddleware(), h.PostComment())
	route.DELETE(":comicID/chapter/:chapterID/comments/:commentID", h.mm.AuthJWTMiddleware(), h.mm.VerifyJWTMiddleware(), h.DeleteComment())
}
