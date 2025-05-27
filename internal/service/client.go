package service

import (
	"github.com/Ocas17/Kursach_UD"
	"github.com/Ocas17/Kursach_UD/internal/repository"
)

type ClientService struct {
	repo repository.Client
}

func NewClientService(repo repository.Client) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) Create(client Kursach_UD.Client) (int, error) {
	return s.repo.Create(client)
}

func (s *ClientService) GetAll() ([]Kursach_UD.Client, error) {
	return s.repo.GetAll()
}

func (s *ClientService) GetById(id int) (Kursach_UD.Client, error) {
	return s.repo.GetById(id)
}

func (s *ClientService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *ClientService) Update(id int, input Kursach_UD.UpdateClientInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, input)
}
