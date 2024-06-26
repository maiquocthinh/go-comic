package repository

import (
	"context"
	"github.com/maiquocthinh/go-comic/internal/comment/models"
)

func (repo *commentRepo) CreateComment(ctx context.Context, commentCreate *models.CommentCreate) error {
	result, err := repo.db.NamedExecContext(
		ctx,
		"INSERT INTO `comments` (`chapter_id`, `user_id`, `content`) VALUES (:chapter_id, :user_id, :content)",
		commentCreate,
	)
	if err != nil {
		return err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	commentCreate.ID = int(insertedID)

	return nil
}

func (repo *commentRepo) CreateReplyComment(ctx context.Context, commentCreate *models.CommentReplyCreate) error {
	result, err := repo.db.NamedExecContext(
		ctx,
		"INSERT INTO `comments` (`chapter_id`, `user_id`, `content`, `parent_id`) VALUES (:chapter_id, :user_id, :content, :comment_parent_id)",
		commentCreate,
	)
	if err != nil {
		return err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	commentCreate.ID = int(insertedID)

	return nil
}
