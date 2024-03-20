package usecase

import (
	"context"
	"github.com/maiquocthinh/go-comic/internal/comment/models"
	"github.com/maiquocthinh/go-comic/internal/comment/repository"
	"github.com/maiquocthinh/go-comic/internal/entities"
	"github.com/maiquocthinh/go-comic/pkg/common"
)

type commentUseCase struct {
	commentRepo repository.CommentRepository
}

func NewCommentUseCase(commentRepo repository.CommentRepository) *commentUseCase {
	return &commentUseCase{commentRepo: commentRepo}
}

type CommentUseCase interface {
	GetComments(ctx context.Context, comicID, chapterID, userID int, paging *common.Paging) ([]*models.CommentDetail, error)
	GetCommentReplies(ctx context.Context, commentID, comicID, chapterID, userID int, paging *common.Paging) ([]*models.CommentDetail, error)
	CreateComment(ctx context.Context, comicID int, commentCreate *models.CommentCreate) (*entities.Comment, error)
	CreateReplyComment(ctx context.Context, comicID int, commentCreate *models.CommentReplyCreate) (*entities.Comment, error)
	DeleteComment(ctx context.Context, comicID, chapterID, commentID, userID int) error
}
