package repository

import (
	"errors"

	events "github.com/branislavstojkovic70/nft-ticket-verification/domain/events"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TagRepository interface {
	GetAllTags() ([]events.Tag, error)
	GetTagByID(id uuid.UUID) (*events.Tag, error)
	CreateTag(tag *events.Tag) error
	UpdateTag(tag *events.Tag) error
	DeleteTag(id uuid.UUID) error
}

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{db}
}

func (r *tagRepository) GetAllTags() ([]events.Tag, error) {
	var tags []events.Tag
	if err := r.db.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *tagRepository) GetTagByID(id uuid.UUID) (*events.Tag, error) {
	var tag events.Tag
	if err := r.db.First(&tag, "uuid = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &tag, nil
}

func (r *tagRepository) CreateTag(tag *events.Tag) error {
	return r.db.Create(tag).Error
}

func (r *tagRepository) UpdateTag(tag *events.Tag) error {
	return r.db.Save(tag).Error
}

func (r *tagRepository) DeleteTag(id uuid.UUID) error {
	return r.db.Delete(&events.Tag{}, "uuid = ?", id).Error
}
