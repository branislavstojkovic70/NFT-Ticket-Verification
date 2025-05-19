package repository

import (
	"errors"

	events "github.com/branislavstojkovic70/nft-ticket-verification/domain/events"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketRepository interface {
	GetAllTickets() ([]events.Ticket, error)
	GetTicketByID(id uuid.UUID) (*events.Ticket, error)
	CreateTicket(ticket *events.Ticket) error
	UpdateTicket(ticket *events.Ticket) error
	DeleteTicket(id uuid.UUID) error
}

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{db}
}

func (r *ticketRepository) GetAllTickets() ([]events.Ticket, error) {
	var tickets []events.Ticket
	if err := r.db.Preload("Event").Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *ticketRepository) GetTicketByID(id uuid.UUID) (*events.Ticket, error) {
	var ticket events.Ticket
	if err := r.db.Preload("Event").First(&ticket, "uuid = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &ticket, nil
}

func (r *ticketRepository) CreateTicket(ticket *events.Ticket) error {
	return r.db.Create(ticket).Error
}

func (r *ticketRepository) UpdateTicket(ticket *events.Ticket) error {
	return r.db.Save(ticket).Error
}

func (r *ticketRepository) DeleteTicket(id uuid.UUID) error {
	return r.db.Delete(&events.Ticket{}, "uuid = ?", id).Error
}
