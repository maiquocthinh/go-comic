package models

import "time"

type HistoryView struct {
	Comic    SimpleComic   `json:"comic"`
	Chapter  SimpleChapter `json:"chapter"`
	ViewedAt time.Time     `json:"viewed_at"`
}
