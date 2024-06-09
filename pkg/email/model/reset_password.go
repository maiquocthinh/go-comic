package model

type ResetPassword struct {
	Firstname string
	Username  string
	Code      string
	ExpiredIn string
}
