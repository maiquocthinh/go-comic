package repository

import (
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

type UserRepository interface {
}
