package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/maiquocthinh/go-comic/internal/comic/models"
	"github.com/maiquocthinh/go-comic/pkg/common"
)

type comicRepo struct {
	db *sqlx.DB
}

func NewComicRepository(db *sqlx.DB) ComicRepository {
	return &comicRepo{db: db}
}

type ComicRepository interface {
	GetComic(ctx context.Context, ID int) (*models.ComicDetail, error)
	ListComic(ctx context.Context, filter *models.ComicFilter, paging *common.Paging) ([]*models.Comic, error)
	GetChapterOfComic(ctx context.Context, comicID, chapterID int) (*models.ChapterDetail, error)
	SearchComic(ctx context.Context, keyword string, paging *common.Paging) ([]*models.Comic, error)
	SearchChapterOfComic(ctx context.Context, comicID int, keyword string, paging *common.Paging) ([]*models.ChapterSimple, error)
	WriteHistoryView(ctx context.Context, historyView *models.HistoryView) error
	IncreaseComicView(ctx context.Context, comicID int) error
	IncreaseChapterView(ctx context.Context, chapterID int) error
}
