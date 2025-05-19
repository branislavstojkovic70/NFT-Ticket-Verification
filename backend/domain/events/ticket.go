package domain

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	UUID      uuid.UUID `json:"uuid" db:"uuid"`
	Price     float64   `json:"price" db:"price"`
	DateStart time.Time `json:"date_start" db:"date_start"`
	DateEnd   time.Time `json:"date_end" db:"date_end"`
	IsUsed    bool      `json:"is_used" db:"is_used"`
	Event     Event     `json:"event" db:"-"`
}
