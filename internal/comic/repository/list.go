package repository

import (
	"context"
	"fmt"
	"github.com/maiquocthinh/go-comic/pkg/utils"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/maiquocthinh/go-comic/internal/comic/models"
	"github.com/maiquocthinh/go-comic/pkg/common"
)

func (repo *comicRepo) ListComic(ctx context.Context, filter *models.ComicFilter, paging *common.Paging) ([]*models.Comic, error) {
	listComic := make([]*models.Comic, 0)
	query := "SELECT *, ( SELECT `name` FROM `chapters` WHERE `comic_id` = `comics`.`id` ORDER BY `id` DESC LIMIT 1 ) AS lasted_chapter " +
		"FROM `comics` WHERE 1=1 "
	queryInit := query
	queryCount := "SELECT COUNT(*) FROM `comics` WHERE 1=1 "

	// filter
	if len(filter.Genres) > 0 {
		query += fmt.Sprintf("AND id IN ( "+
			"SELECT `comic_id` FROM `comic_genre` "+
			"WHERE `genre_id` IN (:genre_ids) "+
			"GROUP BY comic_id "+
			"HAVING COUNT(DISTINCT genre_id) >= %d "+
			") ", len(filter.Genres))
	}

	if filter.Author > 0 {
		query += "AND id IN (SELECT DISTINCT `comic_id` FROM `comic_author` WHERE `author_id` = :author_id) "
	}

	if filter.Status == models.StatusFinished || filter.Status == models.StatusOnGoing {
		query += "AND `is_finished` = :status "
	}

	if filter.MinChapter > 0 {
		query += "AND id IN (SELECT DISTINCT `comic_id` FROM `chapters` GROUP BY `comic_id` HAVING COUNT(*) >= :min_chapter) "
	}

	var sortType string
	if filter.Sort == models.SortAscending {
		sortType = "ASC"
	} else {
		sortType = "DESC"
	}

	switch filter.SortBy {
	case models.SortByNewComic:
		query += fmt.Sprintf("AND `comics`.`name` <> '' ORDER BY `comics`.`updated_at` %s ", sortType)
	case models.SortByMostView:
		query += fmt.Sprintf("ORDER BY `comics`.`viewed` %s ", sortType)
	case models.SortByNewChapter:
		query += fmt.Sprintf("ORDER BY (SELECT MAX(`updated_at`) FROM `chapters` WHERE `chapters`.`comic_id` = `comics`.`id`) %s ", sortType)
	case models.SortByMostChapter:
		query += fmt.Sprintf("ORDER BY (SELECT COUNT(*) FROM `chapters` WHERE `chapters`.`comic_id` = `comics`.`id`) %s ", sortType)
	default:
		query += fmt.Sprintf("ORDER BY `comics`.`id` %s ", sortType)
	}

	query, args, err := sqlx.Named(query, filter)
	if err != nil {
		return nil, err
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	//paging
	queryCount = strings.Replace(query, queryInit, queryCount, 1)
	if err := repo.db.QueryRowxContext(ctx, queryCount, args...).
		Scan(&paging.Total); err != nil {
		return nil, err
	}
	paging.Sync()

	query += "LIMIT ? OFFSET ?;"
	args = append(args, paging.PageSize, (paging.Page-1)*paging.PageSize)

	rows, err := repo.db.QueryxContext(
		ctx,
		query,
		args...,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comic models.Comic
		if err := rows.StructScan(&comic); err != nil {
			return nil, err
		}
		listComic = append(listComic, &comic)
	}

	// change base url of image
	for _, comic := range listComic {
		utils.ChangeDomain(&comic.Image, "st.nettruyenbb.com")
	}

	return listComic, nil
}

func (repo *comicRepo) SearchComic(ctx context.Context, keyword string, paging *common.Paging) ([]*models.Comic, error) {
	listComic := make([]*models.Comic, 0)

	if err := repo.db.QueryRowxContext(
		ctx,
		"SELECT COUNT(*) FROM `comics` WHERE MATCH ( `name`, `other_name` ) AGAINST ( ? )",
		keyword,
	).Scan(&paging.Total); err != nil {
		return nil, err
	}
	paging.Sync()

	rows, err := repo.db.Unsafe().QueryxContext(
		ctx,
		"SELECT	*, "+
			" ( SELECT `name` FROM `chapters` WHERE `comic_id` = `comics`.`id` ORDER BY `id` DESC LIMIT 1 ) AS `lasted_chapter`, "+
			" MATCH ( `name`, `other_name` ) AGAINST ( ? ) as `relative` "+
			"FROM `comics` "+
			"WHERE MATCH ( `name`, `other_name` ) AGAINST ( ? ) "+
			"ORDER BY `relative` DESC "+
			"LIMIT ? OFFSET ? ;",
		keyword, keyword,
		paging.PageSize, (paging.Page-1)*paging.PageSize,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comic models.Comic
		if err := rows.StructScan(&comic); err != nil {
			return nil, err
		}
		listComic = append(listComic, &comic)
	}

	// change base url of image
	for _, comic := range listComic {
		utils.ChangeDomain(&comic.Image, "st.nettruyenbb.com")
	}

	return listComic, nil
}

func (repo *comicRepo) SearchChapterOfComic(ctx context.Context, comicID int, keyword string, paging *common.Paging) ([]*models.ChapterSimple, error) {
	listChapter := make([]*models.ChapterSimple, 0)

	if err := repo.db.QueryRowxContext(
		ctx,
		"SELECT COUNT(*) FROM `chapters` WHERE `chapters`.`comic_id` = ? AND MATCH ( `name` ) AGAINST ( ? )",
		comicID, keyword,
	).Scan(&paging.Total); err != nil {
		return nil, err
	}
	paging.Sync()

	rows, err := repo.db.Unsafe().QueryxContext(
		ctx,
		"SELECT `chapters`.`id`, `chapters`.`name`, `chapters`.`viewed`, `chapters`.`updated_at` "+
			"FROM `chapters` "+
			"WHERE `chapters`.`comic_id` = ? "+
			"AND MATCH ( `name` ) AGAINST ( ? ) "+
			"ORDER BY `chapters`.`id` DESC "+
			"LIMIT ? OFFSET ? ;",
		comicID, keyword,
		paging.PageSize, (paging.Page-1)*paging.PageSize,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var chapter models.ChapterSimple
		if err := rows.StructScan(&chapter); err != nil {
			return nil, err
		}
		listChapter = append(listChapter, &chapter)
	}

	return listChapter, nil
}
