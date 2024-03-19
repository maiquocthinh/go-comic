package usecase

import "github.com/maiquocthinh/go-comic/internal/comment/repository"

type commentUseCase struct {
	commentRepo repository.CommentRepository
}

func NewCommentUseCase(commentRepo repository.CommentRepository) *commentUseCase {
	return &commentUseCase{commentRepo: commentRepo}
}

type CommentUseCase interface {
}
