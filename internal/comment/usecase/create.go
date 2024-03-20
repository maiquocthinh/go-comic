package usecase

import (
	"context"
	"errors"
	"github.com/maiquocthinh/go-comic/internal/comment/models"
	"github.com/maiquocthinh/go-comic/internal/entities"
)

func (uc *commentUseCase) CreateComment(ctx context.Context, comicID int, commentCreate *models.CommentCreate) (*entities.Comment, error) {
	if !uc.commentRepo.IsChapterBelongComic(ctx, comicID, commentCreate.ChapterID) {
		return nil, errors.New("Chapter not found.")
	}

	if err := uc.commentRepo.CreateComment(ctx, commentCreate); err != nil {
		return nil, err
	}

	comment, err := uc.commentRepo.GetCommentByID(ctx, commentCreate.ID)
	if err != nil {
		return nil, err
	}

	return comment, nil
}
