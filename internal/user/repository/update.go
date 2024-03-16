package repository

import (
	"context"
	"github.com/maiquocthinh/go-comic/internal/user/models"
)

func (repo *userRepo) UpdateProfile(ctx context.Context, profileUpdate *models.UserProfileUpdate) error {
	query := "UPDATE  `users` SET `id`=`id`"

	if profileUpdate.FirstName != nil {
		query += ", `first_name`=:first_name"
	}
	if profileUpdate.LastName != nil {
		query += ", `last_name`=:last_name"
	}
	if profileUpdate.Gender != nil {
		query += ", `gender`=:gender"
	}

	query += " WHERE `id`=:id;"

	_, err := repo.db.NamedExecContext(ctx, query, profileUpdate)
	return err
}
