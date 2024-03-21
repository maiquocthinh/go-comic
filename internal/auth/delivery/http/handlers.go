package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maiquocthinh/go-comic/internal/auth/models"
	"github.com/maiquocthinh/go-comic/internal/auth/usecase"
	"github.com/maiquocthinh/go-comic/internal/middleware"
	"github.com/maiquocthinh/go-comic/pkg/common"
	"github.com/maiquocthinh/go-comic/pkg/utils"
	"net/http"
)

type authHandlers struct {
	mm          middleware.MiddlewareManager
	authUseCase usecase.AuthUseCase
}

func NewComicHandlers(mm middleware.MiddlewareManager, authUseCase usecase.AuthUseCase) *authHandlers {
	return &authHandlers{
		mm:          mm,
		authUseCase: authUseCase,
	}
}

type AuthHandlers interface {
	Register() gin.HandlerFunc
	Login() gin.HandlerFunc
	Logout() gin.HandlerFunc
}

func (h *authHandlers) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userRegister models.UserRegister

		if err := ctx.BindJSON(&userRegister); err != nil {
			common.HandleBindingErr(ctx, err)
			return
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
			common.HandleBindingErr(ctx, err)
			return
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

func (h *authHandlers) Logout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userClaims, err := utils.GetUserClaimsFromContext(ctx)
		if err != nil {
			panic(common.NewUnauthorizedApiError(err, ""))
		}

		if err := h.authUseCase.Logout(ctx.Request.Context(), userClaims); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleMessageSuccessResponse("Logout success."))
	}
}
