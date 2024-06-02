package repository

import (
	"context"
)

func (repo *authRepo) ResetPassword(ctx context.Context, email string, hashedPassword string) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE `users` SET  `hash_password`=? WHERE `email`=?", hashedPassword, email)
	if err != nil {
		return err
	}
	return nil
}
