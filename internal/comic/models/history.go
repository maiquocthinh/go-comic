package models

import "time"

type HistoryView struct {
	UserID    int        `json:"user_id" db:"user_id"`
	ComicID   int        `json:"comic_id" db:"comic_id"`
	ChapterID int        `json:"chapter_id" db:"chapter_id"`
	ViewedAt  *time.Time `json:"viewed_at" db:"viewed_at"`
}
