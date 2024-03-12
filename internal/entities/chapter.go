package entities

import "time"

type ChapterSimple struct {
	ID        int        `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}
type Chapter struct {
	ChapterSimple
	ImagesJson string `json:"-" db:"images_json"`
	ComicID    string `json:"-" db:"comic_id"`
	Viewed     int    `json:"viewed" db:"viewed"`
	FolderPath string `json:"-" db:"folder_path"`
	OriginLink string `json:"-" db:"origin_link"`
	IsUploaded bool   `json:"-" db:"is_uploaded"`
}
