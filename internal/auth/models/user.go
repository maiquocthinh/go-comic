package models

type UserLogin struct {
	Email    string `json:"email" binding:"required,email" db:"email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

type UserRegister struct {
	UserLogin
	Username       string `json:"username" binding:"required,gte=5" db:"username"`
	HashedPassword string `json:"-" db:"hash_password"`
}
