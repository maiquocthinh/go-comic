package repository

import (
	"context"
	"github.com/maiquocthinh/go-comic/internal/auth/models"
)

func (repo *authRepo) CreateUser(ctx context.Context, userRegister *models.UserRegister) error {
	_, err := repo.db.NamedExecContext(
		ctx,
		"INSERT INTO `users` (`username`, `email`, `hash_password`) VALUES (:username, :email, :hash_password)",
		userRegister,
	)
	if err != nil {
		return err
	}

	return nil
}
