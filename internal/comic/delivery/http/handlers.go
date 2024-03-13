package http

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maiquocthinh/go-comic/internal/comic/models"
	"github.com/maiquocthinh/go-comic/internal/comic/usecase"
	"github.com/maiquocthinh/go-comic/pkg/common"
)

type comicHandlers struct {
	comicUseCase usecase.ComicUseCase
}

func NewComicHandlers(comicUseCase usecase.ComicUseCase) *comicHandlers {
	return &comicHandlers{comicUseCase: comicUseCase}
}

type ComicHandlers interface {
	List() gin.HandlerFunc
	GetComic() gin.HandlerFunc
	GetChapterOfComic() gin.HandlerFunc
	SearchComic() gin.HandlerFunc
}

func (h *comicHandlers) List() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging
		var filter models.ComicFilter

		if err := ctx.BindQuery(&paging); err != nil {
			panic(common.NewBadRequestApiError(err, ""))
		}

		if err := ctx.BindQuery(&filter); err != nil {
			panic(common.NewBadRequestApiError(err, ""))
		}

		listComic, err := h.comicUseCase.List(ctx.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(listComic, paging))
	}
}

func (h *comicHandlers) GetComic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("comicID"))
		if err != nil {
			panic(common.NewBadRequestApiError(err, "`comicID` must be int"))
		}

		comicDetail, err := h.comicUseCase.GetComic(ctx.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleDataSuccessResponse(&comicDetail))
	}
}

func (h *comicHandlers) GetChapterOfComic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		comicID, err := strconv.Atoi(ctx.Param("comicID"))
		if err != nil {
			panic(common.NewBadRequestApiError(err, "`comicID` must be int"))
		}
		chapterID, err := strconv.Atoi(ctx.Param("chapterID"))
		if err != nil {
			panic(common.NewBadRequestApiError(err, "`chapterID` must be int"))
		}

		chapterDetail, err := h.comicUseCase.GetChapterOfComic(ctx.Request.Context(), comicID, chapterID)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleDataSuccessResponse(chapterDetail))
	}
}

func (h *comicHandlers) SearchComic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var search models.ComicSearch
		var paging common.Paging

		if err := ctx.BindQuery(&search); err != nil || search.Keyword == "" {
			panic(common.NewBadRequestApiError(
				errors.New("`keyword` must not empty"),
				"`keyword` must not empty",
			))
		}

		if err := ctx.BindQuery(&paging); err != nil {
			panic(common.NewBadRequestApiError(err, ""))
		}

		listComic, err := h.comicUseCase.SearchComic(ctx.Request.Context(), search.Keyword, &paging)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(listComic, paging))
	}
}
