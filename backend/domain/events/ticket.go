package domain

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID        uuid.UUID `json:"uuid" gorm:"type:uuid;primaryKey;column:uuid"`
	UserID    uuid.UUID `json:"user_id" gorm:"column:user_id"`
	EventID   uuid.UUID `json:"event_id" gorm:"column:event_id"`
	Price     float64   `json:"price" gorm:"column:price"`
	DateStart time.Time `json:"date_start" gorm:"column:date_start"`
	DateEnd   time.Time `json:"date_end" gorm:"column:date_end"`
	IsUsed    bool      `json:"is_used" gorm:"column:is_used"`
	Event     Event     `json:"event" gorm:"foreignKey:EventID;references:ID"`
}
