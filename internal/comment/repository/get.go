package repository

import (
	"context"
	"github.com/maiquocthinh/go-comic/internal/entities"
)

func (repo *commentRepo) IsChapterBelongComic(ctx context.Context, comicID, chapterID int) bool {
	row := repo.db.QueryRowxContext(ctx, "SELECT 1 FROM `chapters` WHERE `id`=? AND `comic_id`=?", chapterID, comicID)
	if err := row.Err(); err != nil {
		return false
	}
	return true
}

func (repo *commentRepo) IsCommentBelongUser(ctx context.Context, commentID, userID int) bool {
	row := repo.db.QueryRowxContext(ctx, "SELECT 1 FROM `comments` WHERE `id`=? AND `user_id`=?", commentID, userID)
	if err := row.Err(); err != nil {
		return false
	}
	return true
}

func (repo *commentRepo) GetCommentByID(ctx context.Context, commentID int) (*entities.Comment, error) {
	var comment entities.Comment

	err := repo.db.GetContext(ctx, &comment, "SELECT * FROM `comments` WHERE `id`=?", commentID)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}
