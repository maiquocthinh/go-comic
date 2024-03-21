package models

import "github.com/maiquocthinh/go-comic/internal/entities"

type Comic struct {
	entities.Comic
	LastedChapter string `json:"lasted_chapter" db:"lasted_chapter"`
}

type ComicDetail struct {
	entities.Comic
	Authors  []*entities.Author        `json:"authors,omitempty"`
	Genres   []*entities.GenreSimple   `json:"genres,omitempty"`
	Chapters []*entities.ChapterSimple `json:"chapters,omitempty"`
}
