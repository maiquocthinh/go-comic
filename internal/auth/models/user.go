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

type UserSendCodeResetPassword struct {
	Email string `json:"email" binding:"required,email"`
}

type UserResetPassword struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"new_password" binding:"required,gte=6,lte=30"`
	Code     string `json:"code" binding:"required"`
}

type UserResetPasswordPubSub struct {
	Email     string  `json:"email"`
	Username  string  `json:"username"`
	Firstname *string `json:"firstname"`
	Code      string  `json:"code"`
}
