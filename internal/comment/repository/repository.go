package repository

import "github.com/jmoiron/sqlx"

type commentRepo struct {
	db *sqlx.DB
}

func NewCommentRepository(db *sqlx.DB) *commentRepo {
	return &commentRepo{
		db: db,
	}
}

type CommentRepository interface {
}
