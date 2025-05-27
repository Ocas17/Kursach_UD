package repository

import (
	"fmt"
	"strings"

	"github.com/Ocas17/Kursach_UD"
	"github.com/jmoiron/sqlx"
)

type PolicyPostgres struct {
	db *sqlx.DB
}

func NewPolicyPostgres(db *sqlx.DB) *PolicyPostgres {
	return &PolicyPostgres{db: db}
}

func (r *PolicyPostgres) Create(policy Kursach_UD.Policy) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (client_id, type, start_date, end_date, price, is_active) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", policiesTable)
	row := r.db.QueryRow(query, policy.ClientId, policy.Type, policy.StartDate, policy.EndDate, policy.Price, policy.IsActive)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PolicyPostgres) GetAll(clientId int) ([]Kursach_UD.Policy, error) {
	var policies []Kursach_UD.Policy
	query := fmt.Sprintf("SELECT * FROM %s WHERE client_id = $1", policiesTable)
	err := r.db.Select(&policies, query, clientId)
	return policies, err
}

func (r *PolicyPostgres) GetById(id int) (Kursach_UD.Policy, error) {
	var policy Kursach_UD.Policy
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", policiesTable)
	err := r.db.Get(&policy, query, id)
	return policy, err
}

func (r *PolicyPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", policiesTable)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *PolicyPostgres) Update(id int, input Kursach_UD.UpdatePolicyInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Type != nil {
		setValues = append(setValues, fmt.Sprintf("type=$%d", argId))
		args = append(args, *input.Type)
		argId++
	}

	if input.StartDate != nil {
		setValues = append(setValues, fmt.Sprintf("start_date=$%d", argId))
		args = append(args, *input.StartDate)
		argId++
	}

	if input.EndDate != nil {
		setValues = append(setValues, fmt.Sprintf("end_date=$%d", argId))
		args = append(args, *input.EndDate)
		argId++
	}

	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, *input.Price)
		argId++
	}

	if input.IsActive != nil {
		setValues = append(setValues, fmt.Sprintf("is_active=$%d", argId))
		args = append(args, *input.IsActive)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", policiesTable, setQuery, argId)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}
