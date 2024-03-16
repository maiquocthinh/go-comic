package models

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
