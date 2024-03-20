package models

import "github.com/maiquocthinh/go-comic/internal/entities"

type CommentDetail struct {
	entities.Comment
	ComicName   string `json:"comic_name" db:"comic_name"`
	ComicImage  string `json:"comic_image" db:"comic_image"`
	ChapterName string `json:"chapter_name" db:"chapter_name"`
}
