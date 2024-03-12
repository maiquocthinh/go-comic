package models

import (
	"encoding/json"
	"github.com/maiquocthinh/go-comic/internal/entities"
)

type ChapterImage struct {
	Original string `json:"original"`
	Cdn      string `json:"cdn,omitempty"`
}

type ChapterDetail struct {
	entities.Chapter
	Images []*ChapterImage `json:"images" `
}

func (c *ChapterDetail) ParseImages() error {
	return json.Unmarshal([]byte(c.ImagesJson), &c.Images)
}
