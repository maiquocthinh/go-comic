package repository

import (
	"context"
	"github.com/maiquocthinh/go-comic/internal/user/models"
	"github.com/maiquocthinh/go-comic/pkg/common"
)

func (repo userRepo) GetHistoryView(ctx context.Context, userID int, paging *common.Paging) ([]*models.HistoryView, error) {
	err := repo.db.QueryRowxContext(ctx, "SELECT COUNT(*) FROM `history_view` WHERE `user_id`=?", userID).Scan(&paging.Total)
	if err != nil {
		return nil, err
	}
	paging.Sync()

	var histories []*models.HistoryView
	rows, err := repo.db.QueryxContext(
		ctx,
		"SELECT	c.id AS comic_id, c.name AS comic_name, c.other_name AS comic_other_name, c.image AS comic_image, "+
			"ch.id AS chapter_id, ch.name AS chapter_name, "+
			"hv.viewed_at "+
			"FROM history_view hv "+
			"JOIN comics c ON hv.comic_id = c.id "+
			"JOIN chapters ch ON hv.chapter_id = ch.id "+
			"WHERE hv.user_id = ? "+
			"ORDER BY viewed_at DESC;",
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var historyView models.HistoryView

		err := rows.Scan(
			&historyView.Comic.ID,
			&historyView.Comic.Name,
			&historyView.Comic.OtherName,
			&historyView.Comic.Image,
			&historyView.Chapter.ID,
			&historyView.Chapter.Name,
			&historyView.ViewedAt,
		)
		if err != nil {
			return nil, err
		}

		histories = append(histories, &historyView)
	}

	return histories, nil
}
