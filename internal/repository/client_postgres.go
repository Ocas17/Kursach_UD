package repository

import (
	"fmt"
	"strings"

	"github.com/Ocas17/Kursach_UD"
	"github.com/jmoiron/sqlx"
)

type ClientPostgres struct {
	db *sqlx.DB
}

func NewClientPostgres(db *sqlx.DB) *ClientPostgres {
	return &ClientPostgres{db: db}
}

func (r *ClientPostgres) Create(client Kursach_UD.Client) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (full_name, email, phone) VALUES ($1, $2, $3) RETURNING id", clientsTable)
	row := r.db.QueryRow(query, client.FullName, client.Email, client.Phone)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ClientPostgres) GetAll() ([]Kursach_UD.Client, error) {
	var clients []Kursach_UD.Client
	query := fmt.Sprintf("SELECT * FROM %s", clientsTable)
	err := r.db.Select(&clients, query)
	return clients, err
}

func (r *ClientPostgres) GetById(id int) (Kursach_UD.Client, error) {
	var client Kursach_UD.Client
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", clientsTable)
	err := r.db.Get(&client, query, id)
	return client, err
}

func (r *ClientPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", clientsTable)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *ClientPostgres) Update(id int, input Kursach_UD.UpdateClientInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.FullName != nil {
		setValues = append(setValues, fmt.Sprintf("full_name=$%d", argId))
		args = append(args, *input.FullName)
		argId++
	}

	if input.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, *input.Email)
		argId++
	}

	if input.Phone != nil {
		setValues = append(setValues, fmt.Sprintf("phone=$%d", argId))
		args = append(args, *input.Phone)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", clientsTable, setQuery, argId)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}
