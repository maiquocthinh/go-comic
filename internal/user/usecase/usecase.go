package usecase

import "github.com/maiquocthinh/go-comic/internal/user/repository"

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) *userUseCase {
	return &userUseCase{userRepo: userRepo}
}

type UserUseCase interface {
}
