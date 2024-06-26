package usecase

import (
	"context"
	"github.com/maiquocthinh/go-comic/internal/comic/models"
)

func (uc *comicUseCase) GetComic(ctx context.Context, ID int) (*models.ComicDetail, error) {
	return uc.comicRepo.GetComic(ctx, ID)
}

func (uc *comicUseCase) GetChapterOfComic(ctx context.Context, comicID, chapterID int) (*models.ChapterDetail, error) {
	return uc.comicRepo.GetChapterOfComic(ctx, comicID, chapterID)
}
