package models

import "github.com/maiquocthinh/go-comic/internal/entities"

type CommentCreate struct {
	ID        int    `json:"-" db:"id"`
	UserID    int    `json:"-" db:"user_id"`
	ChapterID int    `json:"-" db:"chapter_id"`
	Content   string `json:"content" db:"content" binding:"required,gte=8"`
}

type CommentDetail struct {
	entities.Comment
	Username   string `json:"username" db:"username"`
	UserAvatar string `json:"user_avatar" db:"user_avatar"`
	IsOwner    bool   `json:"is_owner" db:"is_owner"`
	ReplyNum   int    `json:"reply_num" db:"reply_num"`
}
