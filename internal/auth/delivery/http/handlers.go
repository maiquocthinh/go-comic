package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maiquocthinh/go-comic/internal/auth/models"
	"github.com/maiquocthinh/go-comic/internal/auth/usecase"
	"github.com/maiquocthinh/go-comic/pkg/common"
	"net/http"
)

type authHandlers struct {
	authUseCase usecase.AuthUseCase
}

func NewComicHandlers(authUseCase usecase.AuthUseCase) *authHandlers {
	return &authHandlers{authUseCase: authUseCase}
}

type ComicHandlers interface {
	Register() gin.HandlerFunc
	Login() gin.HandlerFunc
}

func (h *authHandlers) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userRegister models.UserRegister

		if err := ctx.BindJSON(&userRegister); err != nil {
			panic(common.NewBadRequestApiError(err, ""))
		}

		if err := h.authUseCase.Register(ctx.Request.Context(), &userRegister); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleMessageSuccessResponse("Register successfully."))
	}
}

func (h *authHandlers) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userLogin models.UserLogin

		if err := ctx.BindJSON(&userLogin); err != nil {
			panic(common.NewBadRequestApiError(err, ""))
		}

		token, err := h.authUseCase.Login(ctx.Request.Context(), &userLogin)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleDataSuccessResponse(gin.H{
			"token": token,
		}))
	}
}
