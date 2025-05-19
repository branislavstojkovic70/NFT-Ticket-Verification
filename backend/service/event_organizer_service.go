package service

import (
	users "github.com/branislavstojkovic70/nft-ticket-verification/domain/users"
	"github.com/branislavstojkovic70/nft-ticket-verification/repository"
	"github.com/google/uuid"
)

type OrganizerService interface {
	GetAllOrganizers() ([]users.Organizer, error)
	GetOrganizerByID(id uuid.UUID) (*users.Organizer, error)
	CreateOrganizer(org *users.Organizer) error
	UpdateOrganizer(org *users.Organizer) error
	DeleteOrganizer(id uuid.UUID) error
}

type organizerService struct {
	repo repository.OrganizerRepository
}

func NewOrganizerService(repo repository.OrganizerRepository) OrganizerService {
	return &organizerService{repo: repo}
}

func (s *organizerService) GetAllOrganizers() ([]users.Organizer, error) {
	return s.repo.GetAllOrganizers()
}

func (s *organizerService) GetOrganizerByID(id uuid.UUID) (*users.Organizer, error) {
	return s.repo.GetOrganizerByID(id)
}

func (s *organizerService) CreateOrganizer(org *users.Organizer) error {
	return s.repo.CreateOrganizer(org)
}

func (s *organizerService) UpdateOrganizer(org *users.Organizer) error {
	return s.repo.UpdateOrganizer(org)
}

func (s *organizerService) DeleteOrganizer(id uuid.UUID) error {
	return s.repo.DeleteOrganizer(id)
}
