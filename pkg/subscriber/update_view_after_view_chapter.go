package subscriber

import (
	"context"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/maiquocthinh/go-comic/internal/comic/models"
	comicRepository "github.com/maiquocthinh/go-comic/internal/comic/repository"
)

func IncreaseViewAfterViewChapter(db *sqlx.DB) *consumerJob {
	return &consumerJob{
		Title: "Update view for comic & chapter  after user view chapter",
		Hld: func(ctx context.Context, msg string) error {
			repo := comicRepository.NewComicRepository(db)

			var increaseView models.IncreaseView
			if err := json.Unmarshal([]byte(msg), &increaseView); err != nil {
				return err
			}

			if err := repo.IncreaseComicView(ctx, increaseView.ComicID); err != nil {
				return err
			}
			if err := repo.IncreaseViewChapter(ctx, increaseView.ChapterID); err != nil {
				return err
			}

			return nil
		},
	}
}
