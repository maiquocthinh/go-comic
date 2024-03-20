package usecase

import (
	"context"
	"errors"
	"github.com/maiquocthinh/go-comic/pkg/common"
)

func (uc *commentUseCase) DeleteComment(ctx context.Context, comicID, chapterID, commentID, userID int) error {
	ok, err := uc.commentRepo.IsChapterBelongComic(ctx, comicID, chapterID)
	if err != nil {
		return err
	}
	if !ok {
		return common.NewBadRequestApiError(errors.New("Chapter not found."), "Chapter not found.")
	}

	ok, err = uc.commentRepo.IsCommentBelongUser(ctx, commentID, userID)
	if err != nil {
		return err
	}
	if !ok {
		return common.NewUnauthorizedApiError(errors.New("Comment not belong user"), "Comment not belong user")
	}

	if err := uc.commentRepo.DeleteCommentByID(ctx, commentID); err != nil {
		return err
	}

	return nil
}
