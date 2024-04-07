package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/maiquocthinh/go-comic/internal/entities"
)

type genreRepo struct {
	db *sqlx.DB
}

func NewGenreRepository(db *sqlx.DB) GenreRepository {
	return &genreRepo{db: db}
}

type GenreRepository interface {
	GetAllGenres(ctx context.Context) ([]*entities.Genre, error)
}
