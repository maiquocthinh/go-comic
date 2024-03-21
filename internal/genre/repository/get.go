package repository

import (
	"context"
	"github.com/maiquocthinh/go-comic/internal/entities"
)

func (repo genreRepo) GetAllGenres(ctx context.Context) ([]*entities.Genre, error) {
	var genres []*entities.Genre

	if err := repo.db.SelectContext(ctx, &genres, "SELECT * FROM `genres`"); err != nil {
		return nil, err
	}

	return genres, nil
}
