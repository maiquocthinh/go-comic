package usecase

import (
	"context"
	"github.com/maiquocthinh/go-comic/internal/entities"
)

func (uc genreUseCase) GetAllGenres(ctx context.Context) ([]*entities.Genre, error) {
	return uc.genreRepo.GetAllGenres(ctx)
}
