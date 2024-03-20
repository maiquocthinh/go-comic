package usecase

import (
	"context"
	"errors"
	"github.com/maiquocthinh/go-comic/internal/comment/models"
	"github.com/maiquocthinh/go-comic/internal/entities"
	"github.com/maiquocthinh/go-comic/pkg/common"
)

func (uc *commentUseCase) CreateComment(ctx context.Context, comicID int, commentCreate *models.CommentCreate) (*entities.Comment, error) {
	ok, err := uc.commentRepo.IsChapterBelongComic(ctx, comicID, commentCreate.ChapterID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, common.NewBadRequestApiError(errors.New("Chapter not found."), "Chapter not found.")
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
