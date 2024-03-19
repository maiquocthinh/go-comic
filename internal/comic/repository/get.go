package repository

import (
	"context"
	"database/sql"
	"github.com/maiquocthinh/go-comic/internal/comic/models"
	"github.com/maiquocthinh/go-comic/pkg/common"
)

func (repo *comicRepo) GetComic(ctx context.Context, ID int) (*models.ComicDetail, error) {
	var comicDetail models.ComicDetail

	// Prepared Statements
	comicStmt, err := repo.db.Preparex("SELECT * FROM `comics` WHERE `comics`.`id` = ? ")
	if err != nil {
		return nil, err
	}
	authorsStmt, err := repo.db.Preparex(
		"SELECT `authors`.* " +
			"FROM `comic_author` " +
			"LEFT JOIN `authors` ON `authors`.`id` = `comic_author`.`author_id`" +
			"WHERE `comic_author`.`comic_id` = ?",
	)
	if err != nil {
		return nil, err
	}
	genresStmt, err := repo.db.Preparex(
		"SELECT `genres`.`id`, `genres`.`name` " +
			"FROM `comic_genre` " +
			"LEFT JOIN `genres` ON `genres`.`id` = `comic_genre`.`genre_id`" +
			"WHERE `comic_genre`.`comic_id` = ?",
	)
	if err != nil {
		return nil, err
	}
	chaptersStmt, err := repo.db.Preparex(
		"SELECT `chapters`.`id`, `chapters`.`name`, `chapters`.`updated_at` " +
			"FROM `chapters` " +
			"WHERE `chapters`.`comic_id` = ?",
	)
	if err != nil {
		return nil, err
	}

	// Query
	if err = comicStmt.Get(&comicDetail, ID); err != nil {
		if err == sql.ErrNoRows {
			return nil, common.NewNotFoundApiError(err, "Comic not found")
		}
		return nil, err
	}

	if err = authorsStmt.Select(&comicDetail.Authors, ID); err != nil {
		return nil, err
	}

	if err = genresStmt.Select(&comicDetail.Genres, ID); err != nil {
		return nil, err
	}

	if err = chaptersStmt.Select(&comicDetail.Chapters, ID); err != nil {
		return nil, err
	}

	return &comicDetail, nil
}

func (repo *comicRepo) GetChapterOfComic(ctx context.Context, comicID, chapterID int) (*models.ChapterDetail, error) {
	var chapterDetail models.ChapterDetail

	err := repo.db.GetContext(
		ctx,
		&chapterDetail,
		"SELECT * FROM chapters WHERE id = ? AND comic_id = ?",
		chapterID,
		comicID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, common.NewBadRequestApiError(err, "Chapter not found.")
		}
		return nil, err
	}

	if err := chapterDetail.ParseImages(); err != nil {
		return nil, err
	}

	return &chapterDetail, nil
}
