package repository

import (
	"context"
	"errors"
	"github.com/maiquocthinh/go-comic/pkg/common"
)

func (repo *commentRepo) DeleteCommentByID(ctx context.Context, commentID int) error {
	result, err := repo.db.ExecContext(ctx, "DELETE FROM `comments` WHERE `id`=?", commentID)
	if err != nil {
		return err
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return common.NewBadRequestApiError(errors.New("Comment not found"), "Comment not found")
	}

	return nil
}
