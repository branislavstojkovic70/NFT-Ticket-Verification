package repository

import (
	"errors"

	domain "github.com/branislavstojkovic70/nft-ticket-verification/domain/events"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventRepository interface {
	GetAllEvents() ([]domain.Event, error)
	GetEventByID(id uuid.UUID) (*domain.Event, error)
	CreateEvent(event *domain.Event) error
	UpdateEvent(event *domain.Event) error
	DeleteEvent(id uuid.UUID) error
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db}
}

func (r *eventRepository) GetAllEvents() ([]domain.Event, error) {
	var events []domain.Event
	if err := r.db.Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (r *eventRepository) GetEventByID(id uuid.UUID) (*domain.Event, error) {
	var event domain.Event
	if err := r.db.First(&event, "uuid = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &event, nil
}

func (r *eventRepository) CreateEvent(event *domain.Event) error {
	return r.db.Create(event).Error
}

func (r *eventRepository) UpdateEvent(event *domain.Event) error {
	return r.db.Save(event).Error
}

func (r *eventRepository) DeleteEvent(id uuid.UUID) error {
	return r.db.Delete(&domain.Event{}, "uuid = ?", id).Error
}
