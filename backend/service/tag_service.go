package service

import (
	events "github.com/branislavstojkovic70/nft-ticket-verification/domain/events"
	"github.com/branislavstojkovic70/nft-ticket-verification/repository"
	"github.com/google/uuid"
)

type TagService interface {
	GetAllTags() ([]events.Tag, error)
	GetTagByID(id uuid.UUID) (*events.Tag, error)
	CreateTag(tag *events.Tag) error
	UpdateTag(tag *events.Tag) error
	DeleteTag(id uuid.UUID) error
}

type tagService struct {
	repo repository.TagRepository
}

func NewTagService(repo repository.TagRepository) TagService {
	return &tagService{repo: repo}
}

func (s *tagService) GetAllTags() ([]events.Tag, error) {
	return s.repo.GetAllTags()
}

func (s *tagService) GetTagByID(id uuid.UUID) (*events.Tag, error) {
	return s.repo.GetTagByID(id)
}

func (s *tagService) CreateTag(tag *events.Tag) error {
	return s.repo.CreateTag(tag)
}

func (s *tagService) UpdateTag(tag *events.Tag) error {
	return s.repo.UpdateTag(tag)
}

func (s *tagService) DeleteTag(id uuid.UUID) error {
	return s.repo.DeleteTag(id)
}
