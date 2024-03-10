package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/maiquocthinh/go-comic/internal/comic/models"
	"github.com/maiquocthinh/go-comic/internal/entities"
	"github.com/maiquocthinh/go-comic/pkg/common"
)

type comicRepo struct {
	db *sqlx.DB
}

func NewComicRepository(db *sqlx.DB) *comicRepo {
	return &comicRepo{db: db}
}

type Repository interface {
	GetComic(ctx context.Context, ID int) (*models.ComicDetail, error)
	List(ctx context.Context, filter *models.ComicFilter, paging *common.Paging) ([]*entities.Comic, error)
}
