package entities

type User struct {
	ID           int     `json:"id" db:"id"`
	Username     string  `json:"username" db:"username"`
	Email        string  `json:"email" db:"email"`
	HashPassword string  `json:"-" db:"hash_password"`
	Avatar       string  `json:"avatar" db:"avatar"`
	FirstName    *string `json:"first_name" db:"first_name"`
	LastName     *string `json:"last_name" db:"last_name"`
	Gender       *bool   `json:"gender" db:"gender"`
}
