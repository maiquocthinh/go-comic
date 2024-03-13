package entities

import "database/sql"

type User struct {
	ID           int            `json:"id" db:"id"`
	Username     string         `json:"username" db:"username"`
	Email        string         `json:"email" db:"email"`
	HashPassword string         `json:"_" db:"hash_password"`
	Avatar       string         `json:"avatar" db:"avatar"`
	FirstName    sql.NullString `json:"first_name" db:"first_name"`
	LastName     sql.NullString `json:"last_name" db:"last_name"`
	Gender       sql.NullBool   `json:"gender" db:"gender"`
}
