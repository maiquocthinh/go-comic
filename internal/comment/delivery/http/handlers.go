package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maiquocthinh/go-comic/internal/comment/usecase"
	"github.com/maiquocthinh/go-comic/internal/middleware"
)

type commentHandlers struct {
	mm             middleware.MiddlewareManager
	commentUseCase usecase.CommentUseCase
}

func NewCommentHandlers(mm middleware.MiddlewareManager, commentUseCase usecase.CommentUseCase) *commentHandlers {
	return &commentHandlers{
		mm:             mm,
		commentUseCase: commentUseCase,
	}
}

type CommentHandlers interface {
	GetCommentsOfChapter() gin.HandlerFunc
	GetRepliesOfComment() gin.HandlerFunc
	PostComment() gin.HandlerFunc
	DeleteComment() gin.HandlerFunc
}

func (h *commentHandlers) GetCommentsOfChapter() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (h *commentHandlers) GetRepliesOfComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (h *commentHandlers) PostComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (h *commentHandlers) DeleteComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
