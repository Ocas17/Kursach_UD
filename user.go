package Kursach_UD

import (
	"errors"
)

type Client struct {
	Id       int    `json:"id" db:"id"`
	FullName string `json:"full_name" binding:"required" db:"full_name"`
	Email    string `json:"email" binding:"required" db:"email"`
	Phone    string `json:"phone" binding:"required" db:"phone"`
}

type UpdateClientInput struct {
	FullName *string `json:"full_name" db:"full_name"`
	Email    *string `json:"email" db:"email"`
	Phone    *string `json:"phone" db:"phone"`
}

func (i UpdateClientInput) Validate() error {
	if i.FullName == nil && i.Email == nil && i.Phone == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
