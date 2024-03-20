package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maiquocthinh/go-comic/internal/comment/models"
	"github.com/maiquocthinh/go-comic/internal/comment/usecase"
	"github.com/maiquocthinh/go-comic/internal/middleware"
	"github.com/maiquocthinh/go-comic/pkg/common"
	"github.com/maiquocthinh/go-comic/pkg/utils"
	"net/http"
	"strconv"
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
		comicID, err := strconv.Atoi(ctx.Param("comicID"))
		if err != nil {
			panic(common.NewBadRequestApiError(err, "`comicID` must be int"))
		}
		chapterID, err := strconv.Atoi(ctx.Param("chapterID"))
		if err != nil {
			panic(common.NewBadRequestApiError(err, "`chapterID` must be int"))
		}

		var paging common.Paging
		if err := ctx.BindQuery(&paging); err != nil {
			common.HandleBindingErr(ctx, err)
			return
		}

		var userID int
		userClaims, err := utils.GetUserClaimsFromContext(ctx)
		if err == nil {
			userID = userClaims.UserID
		}

		comments, err := h.commentUseCase.GetComments(ctx.Request.Context(), comicID, chapterID, userID, &paging)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(comments, &paging))
	}
}

func (h *commentHandlers) GetRepliesOfComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		comicID, err := strconv.Atoi(ctx.Param("comicID"))
		if err != nil {
			panic(common.NewBadRequestApiError(err, "`comicID` must be int"))
		}
		chapterID, err := strconv.Atoi(ctx.Param("chapterID"))
		if err != nil {
			panic(common.NewBadRequestApiError(err, "`chapterID` must be int"))
		}
		commentID, err := strconv.Atoi(ctx.Param("commentID"))
		if err != nil {
			panic(common.NewBadRequestApiError(err, "`commentID` must be int"))
		}

		var userID int
		userClaims, err := utils.GetUserClaimsFromContext(ctx)
		if err == nil {
			userID = userClaims.UserID
		}

		var paging common.Paging
		if err := ctx.BindQuery(&paging); err != nil {
			common.HandleBindingErr(ctx, err)
			return
		}

		comments, err := h.commentUseCase.GetCommentReplies(ctx.Request.Context(), commentID, comicID, chapterID, userID, &paging)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(comments, &paging))
	}
}

func (h *commentHandlers) PostComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		comicID, err := strconv.Atoi(ctx.Param("comicID"))
		if err != nil {
			panic(common.NewBadRequestApiError(err, "`comicID` must be int"))
		}
		chapterID, err := strconv.Atoi(ctx.Param("chapterID"))
		if err != nil {
			panic(common.NewBadRequestApiError(err, "`chapterID` must be int"))
		}

		var commentCreate models.CommentCreate
		if err := ctx.BindJSON(&commentCreate); err != nil {
			common.HandleBindingErr(ctx, err)
			return
		}

		userClaims, err := utils.GetUserClaimsFromContext(ctx)
		if err != nil {
			panic(err)
		}

		commentCreate.UserID = userClaims.UserID
		commentCreate.ChapterID = chapterID

		comment, err := h.commentUseCase.CreateComment(ctx.Request.Context(), comicID, &commentCreate)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusCreated, common.SimpleDataSuccessResponse(comment))
	}
}

func (h *commentHandlers) DeleteComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		comicID, err := strconv.Atoi(ctx.Param("comicID"))
		if err != nil {
			panic(common.NewBadRequestApiError(err, "`comicID` must be int"))
		}
		chapterID, err := strconv.Atoi(ctx.Param("chapterID"))
		if err != nil {
			panic(common.NewBadRequestApiError(err, "`chapterID` must be int"))
		}
		commentID, err := strconv.Atoi(ctx.Param("commentID"))
		if err != nil {
			panic(common.NewBadRequestApiError(err, "`commentID` must be int"))
		}

		userClaims, err := utils.GetUserClaimsFromContext(ctx)
		if err != nil {
			panic(err)
		}

		err = h.commentUseCase.DeleteComment(ctx.Request.Context(), comicID, chapterID, commentID, userClaims.UserID)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleMessageSuccessResponse("Delete comment success."))
	}
}
