package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maiquocthinh/go-comic/internal/middleware"
	"github.com/maiquocthinh/go-comic/internal/user/usecase"
)

type userHandlers struct {
	mm          middleware.MiddlewareManager
	userUseCase usecase.UserUseCase
}

func NewUserHandlers(mm middleware.MiddlewareManager, userUseCase usecase.UserUseCase) *userHandlers {
	return &userHandlers{
		mm:          mm,
		userUseCase: userUseCase,
	}
}

type UserHandlers interface {
	GetProfile() gin.HandlerFunc
	UpdateProfile() gin.HandlerFunc
	ChangePassword() gin.HandlerFunc
	GetComments() gin.HandlerFunc
}

func (h *userHandlers) GetProfile() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (h *userHandlers) UpdateProfile() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (h *userHandlers) ChangePassword() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (h *userHandlers) GetComments() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
