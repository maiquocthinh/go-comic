package usecase

import (
	"context"
	"github.com/maiquocthinh/go-comic/internal/entities"
	"github.com/maiquocthinh/go-comic/internal/genre/repository"
)

type genreUseCase struct {
	genreRepo repository.GenreRepository
}

func NewGenreUseCase(genreRepo repository.GenreRepository) *genreUseCase {
	return &genreUseCase{
		genreRepo: genreRepo,
	}
}

type GenreUseCase interface {
	GetAllGenres(ctx context.Context) ([]*entities.Genre, error)
}
