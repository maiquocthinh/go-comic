package repository

import (
	"context"
	"database/sql"
	"github.com/maiquocthinh/go-comic/internal/entities"
	"github.com/maiquocthinh/go-comic/pkg/common"
)

func (repo *authRepo) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	var user entities.User

	if err := repo.db.GetContext(ctx, &user, "SELECT * FROM `users` WHERE `email` = ?", email); err != nil {
		if err == sql.ErrNoRows {
			return nil, common.NewUnauthorizedApiError(err, "User not existed!")
		}
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
