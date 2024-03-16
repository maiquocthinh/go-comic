package repository

import (
	"context"
	"database/sql"
	"github.com/maiquocthinh/go-comic/internal/entities"
	"github.com/maiquocthinh/go-comic/pkg/common"
)

func (repo *userRepo) GetProfile(ctx context.Context, userID int) (*entities.User, error) {
	var user entities.User

	if err := repo.db.GetContext(ctx, &user, "SELECT * FROM `users` WHERE `id` = ?", userID); err != nil {
		if err == sql.ErrNoRows {
			return nil, common.NewNotFoundApiError(err, "User not found.")
		}
		return nil, err
	}

	return &user, nil
}
