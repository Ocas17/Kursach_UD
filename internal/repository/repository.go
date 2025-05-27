package repository

import (
	"github.com/Ocas17/Kursach_UD"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
	listsTable      = "lists"
	itemsTable      = "words"
	userslistsTable = "users_lists"
	listsItemsTable = "list_words"
	clientsTable    = "clients"
	policiesTable   = "policies"
	claimsTable     = "claims"
)


type Client interface {
	Create(client Kursach_UD.Client) (int, error)
	GetAll() ([]Kursach_UD.Client, error)
	GetById(id int) (Kursach_UD.Client, error)
	Delete(id int) error
	Update(id int, input Kursach_UD.UpdateClientInput) error
}

type Policy interface {
	Create(policy Kursach_UD.Policy) (int, error)
	GetAll(clientId int) ([]Kursach_UD.Policy, error)
	GetById(id int) (Kursach_UD.Policy, error)
	Delete(id int) error
	Update(id int, input Kursach_UD.UpdatePolicyInput) error
}

type Claim interface {
	Create(claim Kursach_UD.Claim) (int, error)
	GetAll(policyId int) ([]Kursach_UD.Claim, error)
	GetById(id int) (Kursach_UD.Claim, error)
	Delete(id int) error
	Update(id int, input Kursach_UD.UpdateClaimInput) error
}

type Repository struct {
	Client
	Policy
	Claim
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Client:        NewClientPostgres(db),
		Policy:        NewPolicyPostgres(db),
		Claim:         NewClaimPostgres(db),
	}
}
