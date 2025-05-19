package repository

import (
	"errors"

	domain "github.com/branislavstojkovic70/nft-ticket-verification/domain/events"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TagRepository interface {
	GetAllTags() ([]domain.Tag, error)
	GetTagByID(id uuid.UUID) (*domain.Tag, error)
	CreateTag(tag *domain.Tag) error
	UpdateTag(tag *domain.Tag) error
	DeleteTag(id uuid.UUID) error
}

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{db}
}

func (r *tagRepository) GetAllTags() ([]domain.Tag, error) {
	var tags []domain.Tag
	if err := r.db.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *tagRepository) GetTagByID(id uuid.UUID) (*domain.Tag, error) {
	var tag domain.Tag
	if err := r.db.First(&tag, "uuid = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &tag, nil
}

func (r *tagRepository) CreateTag(tag *domain.Tag) error {
	return r.db.Create(tag).Error
}

func (r *tagRepository) UpdateTag(tag *domain.Tag) error {
	return r.db.Save(tag).Error
}

func (r *tagRepository) DeleteTag(id uuid.UUID) error {
	return r.db.Delete(&domain.Tag{}, "uuid = ?", id).Error
}
