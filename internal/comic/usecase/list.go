package usecase

import (
	"context"

	"github.com/maiquocthinh/go-comic/internal/comic/models"
	"github.com/maiquocthinh/go-comic/internal/entities"
	"github.com/maiquocthinh/go-comic/pkg/common"
)

func (uc *comicUseCase) List(ctx context.Context, filter *models.ComicFilter, paging *common.Paging) ([]*entities.Comic, error) {
	paging.Fulfill()

	return uc.comicRepo.List(ctx, filter, paging)
}
