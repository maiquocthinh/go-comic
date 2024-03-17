package models

import "mime/multipart"

type UserProfileUpdate struct {
	ID        int     `json:"-" db:"id"`
	FirstName *string `json:"first_name" db:"first_name" binding:"omitempty,gte=2"`
	LastName  *string `json:"last_name" db:"last_name" binding:"omitempty,gte=2"`
	Gender    *bool   `json:"gender" db:"gender" binding:"omitempty"`
}

func (p *UserProfileUpdate) Validate() bool {
	if p.FirstName == nil && p.LastName == nil && p.Gender == nil {
		return false
	}
	return true
}

type UserAvatarUpdate struct {
	ID         int                   `json:"-" db:"id"`
	Username   string                `json:"-" db:"username"`
	Avatar     string                `json:"avatar" db:"avatar"`
	FileHeader *multipart.FileHeader `json:"-"`
}
