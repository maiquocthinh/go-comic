package models

type CommentCreate struct {
	ID        int    `json:"-" db:"id"`
	UserID    int    `json:"-" db:"user_id"`
	ChapterID int    `json:"-" db:"chapter_id"`
	Content   string `json:"content" db:"content" binding:"required,gte=8"`
}
