package repository

import (
	"errors"

	domain "github.com/branislavstojkovic70/nft-ticket-verification/domain/users"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrganizerRepository interface {
	GetAllOrganizers() ([]domain.Organizer, error)
	GetOrganizerByID(id uuid.UUID) (*domain.Organizer, error)
	CreateOrganizer(org *domain.Organizer) error
	UpdateOrganizer(org *domain.Organizer) error
	DeleteOrganizer(id uuid.UUID) error
}

type organizerRepository struct {
	db *gorm.DB
}

func NewOrganizerRepository(db *gorm.DB) OrganizerRepository {
	return &organizerRepository{db}
}

func (r *organizerRepository) GetAllOrganizers() ([]domain.Organizer, error) {
	var organizers []domain.Organizer
	if err := r.db.Find(&organizers).Error; err != nil {
		return nil, err
	}
	return organizers, nil
}

func (r *organizerRepository) GetOrganizerByID(id uuid.UUID) (*domain.Organizer, error) {
	var organizer domain.Organizer
	if err := r.db.First(&organizer, "uuid = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &organizer, nil
}

func (r *organizerRepository) CreateOrganizer(org *domain.Organizer) error {
	return r.db.Create(org).Error
}

func (r *organizerRepository) UpdateOrganizer(org *domain.Organizer) error {
	return r.db.Save(org).Error
}

func (r *organizerRepository) DeleteOrganizer(id uuid.UUID) error {
	return r.db.Delete(&domain.Organizer{}, "uuid = ?", id).Error
}
