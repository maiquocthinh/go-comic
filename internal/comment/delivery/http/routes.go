package http

import "github.com/gin-gonic/gin"

func (h *commentHandlers) MapComicRotes(route *gin.RouterGroup) {
	route.GET(":comicID/chapter/:chapterID/comments", h.mm.OptionalAuthJWTMiddleware(), h.GetCommentsOfChapter())
	route.GET(":comicID/chapter/:chapterID/comments/:commentID/replies", h.mm.OptionalAuthJWTMiddleware(), h.GetRepliesOfComment())
	route.POST(":comicID/chapter/:chapterID/comments/:commentID/replies", h.mm.RequiredAuthJWTMiddleware(), h.PostReplyComment())
	route.POST(":comicID/chapter/:chapterID/comments", h.mm.RequiredAuthJWTMiddleware(), h.PostComment())
	route.DELETE(":comicID/chapter/:chapterID/comments/:commentID", h.mm.RequiredAuthJWTMiddleware(), h.DeleteComment())
}
