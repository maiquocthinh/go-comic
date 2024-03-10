package http

import (
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
}

func (h *comicHandlers) List() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging
		var filter models.ComicFilter

		if err := ctx.BindQuery(&paging); err != nil {
			panic(err)
		}

		if err := ctx.BindQuery(&filter); err != nil {
			panic(err)
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
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}

		comicDetail, err := h.comicUseCase.GetComic(ctx.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(&comicDetail))
	}
}
