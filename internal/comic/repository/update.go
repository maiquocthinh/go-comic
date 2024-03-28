package repository

import "context"

func (repo comicRepo) IncreaseComicView(ctx context.Context, comicID int) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE `comics` SET `viewed` = `viewed` + 1 WHERE `id`=?", comicID)
	if err != nil {
		return err
	}
	return nil
}

func (repo comicRepo) IncreaseChapterView(ctx context.Context, chapterID int) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE `chapters` SET `viewed` = `viewed` + 1 WHERE `id`=?", chapterID)
	if err != nil {
		return err
	}
	return nil
}
