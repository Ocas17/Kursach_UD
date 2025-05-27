package Kursach_UD

import (
	"errors"
	"time"
)

type Claim struct {
	Id           int       `json:"id" db:"id"`
	PolicyId     int       `json:"policy_id" db:"policy_id" binding:"required"`
	IncidentDate time.Time `json:"incident_date" db:"incident_date" binding:"required"`
	Description  string    `json:"description" db:"description" binding:"required"`
	Status       string    `json:"status" db:"status"`
}

type UpdateClaimInput struct {
	IncidentDate *time.Time `json:"incident_date" db:"incident_date"`
	Description  *string    `json:"description" db:"description"`
	Status       *string    `json:"status" db:"status"`
}

func (i UpdateClaimInput) Validate() error {
	if i.IncidentDate == nil && i.Description == nil && i.Status == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
