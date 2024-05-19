package entities

import "time"

type ChapterSimple struct {
	ID        int        `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Viewed    int        `json:"viewed" db:"viewed"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}
type Chapter struct {
	ChapterSimple
	ImagesJson string `json:"-" db:"images_json"`
	ComicID    string `json:"-" db:"comic_id"`
	FolderPath string `json:"-" db:"folder_path"`
	OriginLink string `json:"-" db:"origin_link"`
	IsUploaded bool   `json:"-" db:"is_uploaded"`
}
