package service

import (
	events "github.com/branislavstojkovic70/nft-ticket-verification/domain/events"
	"github.com/branislavstojkovic70/nft-ticket-verification/repository"
	"github.com/google/uuid"
)

type TicketService interface {
	GetAllTickets() ([]events.Ticket, error)
	GetTicketByID(id uuid.UUID) (*events.Ticket, error)
	CreateTicket(ticket *events.Ticket) error
	UpdateTicket(ticket *events.Ticket) error
	DeleteTicket(id uuid.UUID) error
}

type ticketService struct {
	repo repository.TicketRepository
}

func NewTicketService(repo repository.TicketRepository) TicketService {
	return &ticketService{repo: repo}
}

func (s *ticketService) GetAllTickets() ([]events.Ticket, error) {
	return s.repo.GetAllTickets()
}

func (s *ticketService) GetTicketByID(id uuid.UUID) (*events.Ticket, error) {
	return s.repo.GetTicketByID(id)
}

func (s *ticketService) CreateTicket(ticket *events.Ticket) error {
	return s.repo.CreateTicket(ticket)
}

func (s *ticketService) UpdateTicket(ticket *events.Ticket) error {
	return s.repo.UpdateTicket(ticket)
}

func (s *ticketService) DeleteTicket(id uuid.UUID) error {
	return s.repo.DeleteTicket(id)
}
