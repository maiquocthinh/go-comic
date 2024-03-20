package usecase

import (
	"context"
	"errors"
	"github.com/maiquocthinh/go-comic/internal/comment/models"
	"github.com/maiquocthinh/go-comic/pkg/common"
)

func (uc *commentUseCase) GetComments(ctx context.Context, comicID, chapterID, userID int, paging *common.Paging) ([]*models.CommentDetail, error) {
	paging.Fulfill()

	ok, err := uc.commentRepo.IsChapterBelongComic(ctx, comicID, chapterID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, common.NewBadRequestApiError(errors.New("Chapter not found."), "Chapter not found.")
	}

	comments, err := uc.commentRepo.GetCommentsByChapterID(ctx, chapterID, paging)

	for _, comment := range comments {
		if comment.UserID == userID {
			comment.IsOwner = true
		}
	}

	return comments, nil
}
