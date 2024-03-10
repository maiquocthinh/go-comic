package entities

import "time"

type Comic struct {
	ID          int        `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	OtherName   string     `json:"other_name" db:"other_name"`
	Image       string     `json:"image" db:"image"`
	Description string     `json:"description" db:"description"`
	IsFinished  bool       `json:"is_finished" db:"is_finished"`
	Viewed      int        `json:"viewed" db:"viewed"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
	FolderPath  string     `json:"_" db:"folder_path"`
	OriginLink  string     `json:"_" db:"origin_link"`
}
