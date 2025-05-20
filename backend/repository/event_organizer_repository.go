package repository

import (
	"errors"

	users "github.com/branislavstojkovic70/nft-ticket-verification/domain/users"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrganizerRepository interface {
	GetAllOrganizers() ([]users.Organizer, error)
	GetOrganizerByID(id uuid.UUID) (*users.Organizer, error)
	CreateOrganizer(org *users.Organizer) error
	UpdateOrganizer(org *users.Organizer) error
	DeleteOrganizer(id uuid.UUID) error
}

type organizerRepository struct {
	db *gorm.DB
}

func NewOrganizerRepository(db *gorm.DB) OrganizerRepository {
	return &organizerRepository{db}
}

func (r *organizerRepository) GetAllOrganizers() ([]users.Organizer, error) {
	var organizers []users.Organizer
	if err := r.db.Preload("Events").Find(&organizers).Error; err != nil {
		return nil, err
	}
	return organizers, nil
}

func (r *organizerRepository) GetOrganizerByID(id uuid.UUID) (*users.Organizer, error) {
	var organizer users.Organizer
	if err := r.db.Preload("Events").First(&organizer, "uuid = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &organizer, nil
}

func (r *organizerRepository) CreateOrganizer(org *users.Organizer) error {
	return r.db.Create(org).Error
}

func (r *organizerRepository) UpdateOrganizer(org *users.Organizer) error {
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(org).Error
}

func (r *organizerRepository) DeleteOrganizer(id uuid.UUID) error {
	return r.db.Delete(&users.Organizer{}, "uuid = ?", id).Error
}
