package repository

import (
	"context"
	"github.com/maiquocthinh/go-comic/internal/comment/models"
	"github.com/maiquocthinh/go-comic/pkg/common"
)

func (repo *commentRepo) GetCommentsByChapterID(ctx context.Context, chapterID int, paging *common.Paging) ([]*models.CommentDetail, error) {
	row := repo.db.QueryRowxContext(ctx, "SELECT COUNT(`id`) FROM `comments` WHERE `chapter_id`=? AND `parent_id` IS NULL", chapterID)
	if err := row.Scan(&paging.Total); err != nil {
		return nil, err
	}
	paging.Sync()

	var comments []*models.CommentDetail
	err := repo.db.SelectContext(
		ctx,
		&comments,
		"SELECT c.*, u.`username`, u.`avatar` AS user_avatar, "+
			"(SELECT COUNT(`id`) FROM `comments` WHERE `parent_id` = c.`id`) AS reply_num "+
			"FROM `comments` c "+
			"JOIN `users` u ON u.`id` = c.`user_id` "+
			"WHERE c.`chapter_id`=? AND c.`parent_id` IS NULL "+
			"ORDER BY c.`created_at` DESC",
		chapterID,
	)
	if err != nil {
		return nil, err
	}

	return comments, nil
}
