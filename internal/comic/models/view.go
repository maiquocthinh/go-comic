package models

type IncreaseView struct {
	ComicID   int `json:"comic_id" db:"comic_id"`
	ChapterID int `json:"chapter_id" db:"chapter_id"`
}
