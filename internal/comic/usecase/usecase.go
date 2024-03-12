package usecase

import (
	"context"

	"github.com/maiquocthinh/go-comic/internal/comic/models"
	"github.com/maiquocthinh/go-comic/internal/comic/repository"
	"github.com/maiquocthinh/go-comic/internal/entities"
	"github.com/maiquocthinh/go-comic/pkg/common"
)

type comicUseCase struct {
	comicRepo repository.Repository
}

func NewComicUseCase(comicRepo repository.Repository) *comicUseCase {
	return &comicUseCase{comicRepo: comicRepo}
}

type ComicUseCase interface {
	List(ctx context.Context, filter *models.ComicFilter, paging *common.Paging) ([]*entities.Comic, error)
	GetComic(ctx context.Context, ID int) (*models.ComicDetail, error)
	GetChapterOfComic(ctx context.Context, comicID, chapterID int) (*models.ChapterDetail, error)
}
