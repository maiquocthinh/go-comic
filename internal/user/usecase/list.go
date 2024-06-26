package usecase

import (
	"context"
	"github.com/maiquocthinh/go-comic/internal/user/models"
	"github.com/maiquocthinh/go-comic/pkg/common"
)

func (uc *userUseCase) GetHistoryView(ctx context.Context, userID int, paging *common.Paging) ([]*models.HistoryView, error) {
	paging.Fulfill()
	return uc.userRepo.GetHistoryView(ctx, userID, paging)
}

func (uc *userUseCase) GetComments(ctx context.Context, userID int, paging *common.Paging) ([]*models.CommentDetail, error) {
	paging.Fulfill()
	return uc.userRepo.GetComments(ctx, userID, paging)
}
