package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maiquocthinh/go-comic/internal/genre/usecase"
	"github.com/maiquocthinh/go-comic/pkg/common"
	"net/http"
)

type genreHandlers struct {
	genreUseCase usecase.GenreUseCase
}

func NewGenreHandlers(genreUseCase usecase.GenreUseCase) GenreHandlers {
	return &genreHandlers{
		genreUseCase: genreUseCase,
	}
}

type GenreHandlers interface {
	MapComicRotes(route *gin.RouterGroup)
	GetAllGenres() gin.HandlerFunc
}

func (h genreHandlers) GetAllGenres() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		genres, err := h.genreUseCase.GetAllGenres(ctx.Request.Context())
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleDataSuccessResponse(&genres))
	}
}
