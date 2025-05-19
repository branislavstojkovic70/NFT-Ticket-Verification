package service

import (
	events "github.com/branislavstojkovic70/nft-ticket-verification/domain/events"
	"github.com/branislavstojkovic70/nft-ticket-verification/repository"
	"github.com/google/uuid"
)

type EventService interface {
	GetAllEvents() ([]events.Event, error)
	GetEventByID(id uuid.UUID) (*events.Event, error)
	CreateEvent(event *events.Event) error
	UpdateEvent(event *events.Event) error
	DeleteEvent(id uuid.UUID) error
}

type eventService struct {
	repo repository.EventRepository
}

func NewEventService(repo repository.EventRepository) EventService {
	return &eventService{repo: repo}
}

func (s *eventService) GetAllEvents() ([]events.Event, error) {
	return s.repo.GetAllEvents()
}

func (s *eventService) GetEventByID(id uuid.UUID) (*events.Event, error) {
	return s.repo.GetEventByID(id)
}

func (s *eventService) CreateEvent(event *events.Event) error {
	return s.repo.CreateEvent(event)
}

func (s *eventService) UpdateEvent(event *events.Event) error {
	return s.repo.UpdateEvent(event)
}

func (s *eventService) DeleteEvent(id uuid.UUID) error {
	return s.repo.DeleteEvent(id)
}
