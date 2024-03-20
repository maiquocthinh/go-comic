package repository

import "context"

func (repo *commentRepo) DeleteCommentByID(ctx context.Context, commentID int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM `comments` WHERE `id`=?", commentID)
	if err != nil {
		return err
	}
	return nil
}
