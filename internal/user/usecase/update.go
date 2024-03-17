package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/maiquocthinh/go-comic/internal/entities"
	"github.com/maiquocthinh/go-comic/internal/user/models"
	"github.com/maiquocthinh/go-comic/pkg/common"
	"github.com/maiquocthinh/go-comic/pkg/utils"
	"io"
	"path"
	"path/filepath"
	"strings"
)

const (
	avatarUploadPath = "/avatars"
)

func (uc *userUseCase) UpdateProfile(ctx context.Context, profileUpdate *models.UserProfileUpdate) (*entities.User, error) {
	if !profileUpdate.Validate() {
		return nil, common.NewBadRequestApiError(
			errors.New("UserProfileUpdate invalid"),
			"No thing to update",
		)
	}

	if err := uc.userRepo.UpdateProfile(ctx, profileUpdate); err != nil {
		return nil, err
	}
	return uc.userRepo.GetProfile(ctx, profileUpdate.ID)
}

func (uc *userUseCase) UpdateAvatar(ctx context.Context, userAvatarUpdate *models.UserAvatarUpdate) error {
	// open & get data of image
	fileHeader := userAvatarUpdate.FileHeader
	file, err := fileHeader.Open()
	if err != nil {
		return common.NewBadRequestApiError(err, "")
	}
	fileData, err := io.ReadAll(file)
	if err != nil {
		return common.NewInternalApiError(err, "")
	}

	extension := strings.ToLower(filepath.Ext(fileHeader.Filename))
	uploadFilePath := path.Join(avatarUploadPath, fmt.Sprintf("%s%s", userAvatarUpdate.Username, extension))

	// upload avatar
	if err := uc.uploadProvider.UploadImage(ctx, fileData, uploadFilePath); err != nil {
		return common.NewInternalApiError(err, "")
	}
	userAvatarUpdate.Avatar, err = uc.uploadProvider.GetShareLink(ctx, uploadFilePath)
	if err != nil {
		return common.NewInternalApiError(err, "")
	}

	userAvatarUpdate.Avatar = utils.DropboxShareLinkToDirectLink(userAvatarUpdate.Avatar)

	// update avatar
	if err := uc.userRepo.UpdateAvatar(ctx, userAvatarUpdate); err != nil {
		return err
	}

	return nil
}
