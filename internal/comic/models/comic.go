package models

import "github.com/maiquocthinh/go-comic/internal/entities"

type ComicDetail struct {
	entities.Comic
	Authors  []*entities.Author        `json:"authors,omitempty"`
	Genres   []*entities.GenreSimple   `json:"genres,omitempty"`
	Chapters []*entities.ChapterSimple `json:"chapters,omitempty"`
}
