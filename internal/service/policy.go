package service

import (
	"github.com/Ocas17/Kursach_UD"
	"github.com/Ocas17/Kursach_UD/internal/repository"
)

type PolicyService struct {
	repo repository.Policy
}

func NewPolicyService(repo repository.Policy) *PolicyService {
	return &PolicyService{repo: repo}
}

func (s *PolicyService) Create(policy Kursach_UD.Policy) (int, error) {
	return s.repo.Create(policy)
}

func (s *PolicyService) GetAll(clientId int) ([]Kursach_UD.Policy, error) {
	return s.repo.GetAll(clientId)
}

func (s *PolicyService) GetById(id int) (Kursach_UD.Policy, error) {
	return s.repo.GetById(id)
}

func (s *PolicyService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *PolicyService) Update(id int, input Kursach_UD.UpdatePolicyInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, input)
}
