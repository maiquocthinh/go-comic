package repository

import (
	"context"
	"github.com/maiquocthinh/go-comic/internal/comic/models"
)

func (repo *comicRepo) WriteHistoryView(ctx context.Context, historyView *models.HistoryView) error {
	// validate chapter is in comic
	rows, err := repo.db.NamedQueryContext(ctx, "SELECT 1 FROM `chapters` WHERE `id`=:chapter_id AND `comic_id`=:comic_id", historyView)
	if err := rows.Err(); err != nil {
		return err
	}

	// crate or update history_view
	_, err = repo.db.NamedExecContext(
		ctx,
		"INSERT INTO `history_view`(`user_id`, `comic_id`, `chapter_id`, `viewed_at`) VALUES (:user_id, :comic_id, :chapter_id, :viewed_at) "+
			"ON DUPLICATE KEY UPDATE chapter_id = VALUES(chapter_id), viewed_at = VALUES(viewed_at)",
		historyView,
	)
	if err != nil {
		return err
	}

	return nil
}
