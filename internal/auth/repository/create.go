package repository

import (
	"context"
	"github.com/maiquocthinh/go-comic/internal/auth/models"
	"github.com/maiquocthinh/go-comic/pkg/utils"
)

func (repo *authRepo) CreateUser(ctx context.Context, userRegister *models.UserRegister) error {

	// hash password
	hashedPassword, err := utils.HashPassword(userRegister.Password)
	if err != nil {
		return err
	}

	// create new user
	_, err = repo.db.ExecContext(
		ctx,
		"INSERT INTO `users` (`username`, `email`, `hash_password`) VALUES (?, ?, ?)",
		userRegister.Username,
		userRegister.Email,
		hashedPassword,
	)
	if err != nil {
		return err
	}

	return nil
}
