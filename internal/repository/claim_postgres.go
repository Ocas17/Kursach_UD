package repository

import (
	"fmt"
	"strings"

	"github.com/Ocas17/Kursach_UD"
	"github.com/jmoiron/sqlx"
)

type ClaimPostgres struct {
	db *sqlx.DB
}

func NewClaimPostgres(db *sqlx.DB) *ClaimPostgres {
	return &ClaimPostgres{db: db}
}

func (r *ClaimPostgres) Create(claim Kursach_UD.Claim) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (policy_id, incident_date, description, status) VALUES ($1, $2, $3, $4) RETURNING id", claimsTable)
	row := r.db.QueryRow(query, claim.PolicyId, claim.IncidentDate, claim.Description, claim.Status)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ClaimPostgres) GetAll(policyId int) ([]Kursach_UD.Claim, error) {
	var claims []Kursach_UD.Claim
	query := fmt.Sprintf("SELECT * FROM %s WHERE policy_id = $1", claimsTable)
	err := r.db.Select(&claims, query, policyId)
	return claims, err
}

func (r *ClaimPostgres) GetById(id int) (Kursach_UD.Claim, error) {
	var claim Kursach_UD.Claim
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", claimsTable)
	err := r.db.Get(&claim, query, id)
	return claim, err
}

func (r *ClaimPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", claimsTable)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *ClaimPostgres) Update(id int, input Kursach_UD.UpdateClaimInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.IncidentDate != nil {
		setValues = append(setValues, fmt.Sprintf("incident_date=$%d", argId))
		args = append(args, *input.IncidentDate)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Status != nil {
		setValues = append(setValues, fmt.Sprintf("status=$%d", argId))
		args = append(args, *input.Status)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", claimsTable, setQuery, argId)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}
