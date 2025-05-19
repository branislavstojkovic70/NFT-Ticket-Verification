package repository

import (
	"errors"

	events "github.com/branislavstojkovic70/nft-ticket-verification/domain/events"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventRepository interface {
	GetAllEvents() ([]events.Event, error)
	GetEventByID(id uuid.UUID) (*events.Event, error)
	CreateEvent(event *events.Event) error
	UpdateEvent(event *events.Event) error
	DeleteEvent(id uuid.UUID) error
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db}
}

func (r *eventRepository) GetAllEvents() ([]events.Event, error) {
	var events []events.Event
	if err := r.db.Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (r *eventRepository) GetEventByID(id uuid.UUID) (*events.Event, error) {
	var event events.Event
	if err := r.db.First(&event, "uuid = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &event, nil
}

func (r *eventRepository) CreateEvent(event *events.Event) error {
	return r.db.Create(event).Error
}

func (r *eventRepository) UpdateEvent(event *events.Event) error {
	return r.db.Save(event).Error
}

func (r *eventRepository) DeleteEvent(id uuid.UUID) error {
	return r.db.Delete(&events.Event{}, "uuid = ?", id).Error
}
