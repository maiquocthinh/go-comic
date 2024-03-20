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
		userClaims, err := utils.GetUserClaimsFromContext(ctx)
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

		userClaims, err := utils.GetUserClaimsFromContext(ctx)
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

		userClaims, err := utils.GetUserClaimsFromContext(ctx)
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
	return func(ctx *gin.Context) {
		var userChangePassword models.UserChangePassword

		if err := ctx.BindJSON(&userChangePassword); err != nil {
			common.HandleBindingErr(ctx, err)
			return
		}

		userClaims, err := utils.GetUserClaimsFromContext(ctx)
		if err != nil {
			panic(common.NewUnauthorizedApiError(err, ""))
		}

		userChangePassword.ID = userClaims.UserID

		if err := h.userUseCase.ChangePassword(ctx.Request.Context(), &userChangePassword); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleMessageSuccessResponse("Change password success."))
	}
}

func (h *userHandlers) GetHistoryView() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging
		if err := ctx.BindQuery(&paging); err != nil {
			common.HandleBindingErr(ctx, err)
			return
		}

		userClaims, err := utils.GetUserClaimsFromContext(ctx)
		if err != nil {
			panic(common.NewUnauthorizedApiError(err, ""))
		}

		histories, err := h.userUseCase.GetHistoryView(ctx.Request.Context(), userClaims.UserID, &paging)
		if err != nil {
			panic(common.NewInternalApiError(err, ""))
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(histories, paging))
	}
}

func (h *userHandlers) GetComments() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging
		if err := ctx.BindQuery(&paging); err != nil {
			common.HandleBindingErr(ctx, err)
			return
		}

		userClaims, err := utils.GetUserClaimsFromContext(ctx)
		if err != nil {
			panic(common.NewUnauthorizedApiError(err, ""))
		}

		comments, err := h.userUseCase.GetComments(ctx.Request.Context(), userClaims.UserID, &paging)

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(comments, paging))
	}
}
