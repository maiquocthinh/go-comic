package subscriber

import (
	"context"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/maiquocthinh/go-comic/internal/comic/models"
	comicRepository "github.com/maiquocthinh/go-comic/internal/comic/repository"
)

func WriteHistoryAfterViewChapter(db *sqlx.DB) *consumerJob {
	return &consumerJob{
		Title: "Write history after user view chapter",
		Hld: func(ctx context.Context, msg string) error {
			repo := comicRepository.NewComicRepository(db)

			var historyView models.HistoryView
			if err := json.Unmarshal([]byte(msg), &historyView); err != nil {
				return err
			}

			if err := repo.WriteHistoryView(ctx, &historyView); err != nil {
				return err
			}

			return nil
		},
	}
}
