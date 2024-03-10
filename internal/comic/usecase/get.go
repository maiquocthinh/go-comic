package usecase

import (
	"context"

	"github.com/maiquocthinh/go-comic/internal/comic/models"
)

func (uc *comicUseCase) GetComic(ctx context.Context, ID int) (*models.ComicDetail, error) {
	return uc.comicRepo.GetComic(ctx, ID)
}
