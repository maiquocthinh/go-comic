package models

import (
	"encoding/json"
	"github.com/maiquocthinh/go-comic/internal/entities"
	"time"
)

type ChapterImage struct {
	Original string `json:"original"`
	Cdn      string `json:"cdn,omitempty"`
}

type ChapterSimple struct {
	ID        int        `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Viewed    int        `json:"viewed" db:"viewed"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type ChapterDetail struct {
	entities.Chapter
	Images []*ChapterImage `json:"images" `
}

func (c *ChapterDetail) ParseImages() error {
	return json.Unmarshal([]byte(c.ImagesJson), &c.Images)
}
