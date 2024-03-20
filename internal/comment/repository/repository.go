package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/maiquocthinh/go-comic/internal/comment/models"
	"github.com/maiquocthinh/go-comic/internal/entities"
)

type commentRepo struct {
	db *sqlx.DB
}

func NewCommentRepository(db *sqlx.DB) *commentRepo {
	return &commentRepo{
		db: db,
	}
}

type CommentRepository interface {
	IsChapterBelongComic(ctx context.Context, comicID, chapterID int) bool
	CreateComment(ctx context.Context, commentCreate *models.CommentCreate) error
	GetCommentByID(ctx context.Context, commentID int) (*entities.Comment, error)
}