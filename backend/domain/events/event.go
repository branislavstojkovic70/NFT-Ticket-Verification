package domain

import (
	"time"

	"github.com/google/uuid"
)

type EventType string

const (
	Music      EventType = "music"
	Conference EventType = "conference"
)

type Event struct {
	UUID        uuid.UUID `json:"uuid" db:"uuid"`
	Location    string    `json:"location" db:"location"`
	Type        EventType `json:"type" db:"type"`
	DateStart   time.Time `json:"date_start" db:"date_start"`
	DateEnd     time.Time `json:"date_end" db:"date_end"`
	Description string    `json:"description" db:"description"`
	Title       string    `json:"title" db:"title"`
	Tags        []string  `json:"tags" db:"tags"`
}
