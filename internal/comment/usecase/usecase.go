package usecase

import (
	"context"
	"github.com/maiquocthinh/go-comic/internal/comment/models"
	"github.com/maiquocthinh/go-comic/internal/comment/repository"
	"github.com/maiquocthinh/go-comic/internal/entities"
)

type commentUseCase struct {
	commentRepo repository.CommentRepository
}

func NewCommentUseCase(commentRepo repository.CommentRepository) *commentUseCase {
	return &commentUseCase{commentRepo: commentRepo}
}

type CommentUseCase interface {
	CreateComment(ctx context.Context, comicID int, commentCreate *models.CommentCreate) (*entities.Comment, error)
}
