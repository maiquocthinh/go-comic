package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/maiquocthinh/go-comic/internal/middleware"
	"github.com/maiquocthinh/go-comic/internal/user/models"
	"github.com/maiquocthinh/go-comic/internal/user/usecase"
	"github.com/maiquocthinh/go-comic/pkg/common"
	"github.com/maiquocthinh/go-comic/pkg/utils"
	"net/http"
	"strings"
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
	UpdateAvatar() gin.HandlerFunc
	ChangePassword() gin.HandlerFunc
	GetComments() gin.HandlerFunc
}

func (h *userHandlers) GetProfile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userClaims, err := utils.GetUserTokenClaimsFromContext(ctx)
		if err != nil {
			panic(common.NewUnauthorizedApiError(err, ""))
		}

		user, err := h.userUseCase.GetProfile(ctx.Request.Context(), userClaims.UserID)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleDataSuccessResponse(user))
	}
}

func (h *userHandlers) UpdateProfile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var profileUpdate models.UserProfileUpdate

		if err := ctx.BindJSON(&profileUpdate); err != nil {
			common.HandleBindingErr(ctx, err)
			return
		}

		userClaims, err := utils.GetUserTokenClaimsFromContext(ctx)
		if err != nil {
			panic(common.NewUnauthorizedApiError(err, ""))
		}

		profileUpdate.ID = userClaims.UserID

		user, err := h.userUseCase.UpdateProfile(ctx.Request.Context(), &profileUpdate)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(user, "Update Profile success."))
	}
}

func (h *userHandlers) UpdateAvatar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fileHeader, err := ctx.FormFile("file")
		if err != nil {
			panic(common.NewBadRequestApiError(err, ""))
		}

		filetype := fileHeader.Header.Get("Content-Type")
		if !strings.HasPrefix(filetype, "image/") {
			panic(common.NewBadRequestApiError(errors.New("Only image file are allowed"), ""))
		}

		userClaims, err := utils.GetUserTokenClaimsFromContext(ctx)
		userAvatarUpdate := models.UserAvatarUpdate{
			ID:         userClaims.UserID,
			Username:   userClaims.Username,
			Avatar:     userClaims.Avatar,
			FileHeader: fileHeader,
		}

		if err := h.userUseCase.UpdateAvatar(ctx.Request.Context(), &userAvatarUpdate); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleDataSuccessResponse(&userAvatarUpdate))
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
