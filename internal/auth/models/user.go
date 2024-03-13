package models

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

type UserRegister struct {
	UserLogin
	Username string `json:"username" binding:"required,gte=5"`
}
