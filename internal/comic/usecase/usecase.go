package usecase

import (
	"context"

	"github.com/maiquocthinh/go-comic/internal/comic/models"
	"github.com/maiquocthinh/go-comic/internal/comic/repository"
	"github.com/maiquocthinh/go-comic/pkg/common"
)

type comicUseCase struct {
	comicRepo repository.ComicRepository
}

func NewComicUseCase(comicRepo repository.ComicRepository) ComicUseCase {
	return &comicUseCase{comicRepo: comicRepo}
}

type ComicUseCase interface {
	List(ctx context.Context, filter *models.ComicFilter, paging *common.Paging) ([]*models.Comic, error)
	GetComic(ctx context.Context, ID int) (*models.ComicDetail, error)
	GetChapterOfComic(ctx context.Context, comicID, chapterID int) (*models.ChapterDetail, error)
	SearchComic(ctx context.Context, keyword string, paging *common.Paging) ([]*models.Comic, error)
	SearchChapterOfComic(ctx context.Context, comicID int, keyword string, paging *common.Paging) ([]*models.ChapterSimple, error)
}
