package repository

import (
	"context"
	"github.com/maiquocthinh/go-comic/internal/entities"
)

func (repo *authRepo) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	var user entities.User

	if err := repo.db.GetContext(ctx, &user, "SELECT * FROM `users` WHERE `email` = ?", email); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *authRepo) GetUserByUsername(ctx context.Context, username string) (*entities.User, error) {
	var user entities.User

	if err := repo.db.GetContext(ctx, &user, "SELECT * FROM `users` WHERE `username` = ?", username); err != nil {
		return nil, err
	}

	return &user, nil
}
