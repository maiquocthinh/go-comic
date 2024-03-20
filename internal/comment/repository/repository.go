package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/maiquocthinh/go-comic/internal/comment/models"
	"github.com/maiquocthinh/go-comic/internal/entities"
	"github.com/maiquocthinh/go-comic/pkg/common"
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
	IsChapterBelongComic(ctx context.Context, comicID, chapterID int) (bool, error)
	IsCommentBelongUser(ctx context.Context, commentID, userID int) (bool, error)
	GetCommentsByChapterID(ctx context.Context, chapterID int, paging *common.Paging) ([]*models.CommentDetail, error)
	GetCommentsReplies(ctx context.Context, commentID int, paging *common.Paging) ([]*models.CommentDetail, error)
	CreateComment(ctx context.Context, commentCreate *models.CommentCreate) error
	GetCommentByID(ctx context.Context, commentID int) (*entities.Comment, error)
	DeleteCommentByID(ctx context.Context, commentID int) error
}
