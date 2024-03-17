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

func (repo *userRepo) UpdateAvatar(ctx context.Context, userAvatarUpdate *models.UserAvatarUpdate) error {
	_, err := repo.db.NamedExecContext(ctx, "UPDATE `users` SET `avatar`=:avatar WHERE `id`=:id", userAvatarUpdate)
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepo) UpdatePassword(ctx context.Context, userID int, hashedPassword string) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE `users` SET  `hash_password`=? WHERE `id`=?", hashedPassword, userID)
	if err != nil {
		return err
	}
	return nil
}
