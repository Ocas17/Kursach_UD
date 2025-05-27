package service

import (
	"github.com/Ocas17/Kursach_UD"
	"github.com/Ocas17/Kursach_UD/internal/repository"
)

type ClaimService struct {
	repo repository.Claim
}

func NewClaimService(repo repository.Claim) *ClaimService {
	return &ClaimService{repo: repo}
}

func (s *ClaimService) Create(claim Kursach_UD.Claim) (int, error) {
	return s.repo.Create(claim)
}

func (s *ClaimService) GetAll(policyId int) ([]Kursach_UD.Claim, error) {
	return s.repo.GetAll(policyId)
}

func (s *ClaimService) GetById(id int) (Kursach_UD.Claim, error) {
	return s.repo.GetById(id)
}

func (s *ClaimService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *ClaimService) Update(id int, input Kursach_UD.UpdateClaimInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, input)
}
