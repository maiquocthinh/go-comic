package entities

import "time"

type Comment struct {
	ID        int        `json:"id" db:"id"`
	UserID    int        `json:"-" db:"user_id"`
	ChapterID int        `json:"-" db:"chapter_id"`
	Content   string     `json:"content" db:"content"`
	ParentID  *int       `json:"-" db:"parent_id"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
}
