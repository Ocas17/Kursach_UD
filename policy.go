package Kursach_UD

import (
	"errors"
	"time"
)

type Policy struct {
	Id        int       `json:"id" db:"id"`
	ClientId  int       `json:"client_id" db:"client_id" binding:"required"`
	Type      string    `json:"type" db:"type" binding:"required"`
	StartDate time.Time `json:"start_date" db:"start_date" binding:"required"`
	EndDate   time.Time `json:"end_date" db:"end_date" binding:"required"`
	Price     float64   `json:"price" db:"price" binding:"required"`
	IsActive  bool      `json:"is_active" db:"is_active"`
}

type UpdatePolicyInput struct {
	Type      *string    `json:"type" db:"type"`
	StartDate *time.Time `json:"start_date" db:"start_date"`
	EndDate   *time.Time `json:"end_date" db:"end_date"`
	Price     *float64   `json:"price" db:"price"`
	IsActive  *bool      `json:"is_active" db:"is_active"`
}

func (i UpdatePolicyInput) Validate() error {
	if i.Type == nil && i.StartDate == nil && i.EndDate == nil && i.Price == nil && i.IsActive == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
