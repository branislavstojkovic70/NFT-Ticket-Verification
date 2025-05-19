package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type EventType string

const (
	Music      EventType = "music"
	Conference EventType = "conference"
)

type Event struct {
	ID          uuid.UUID      `json:"uuid" gorm:"type:uuid;primaryKey;column:uuid"`
	Location    string         `json:"location" gorm:"column:location"`
	Type        EventType      `json:"type" gorm:"column:type"`
	DateStart   time.Time      `json:"date_start" gorm:"column:date_start"`
	DateEnd     time.Time      `json:"date_end" gorm:"column:date_end"`
	Description string         `json:"description" gorm:"column:description"`
	Title       string         `json:"title" gorm:"column:title"`
	Tags        datatypes.JSON `json:"tags" gorm:"column:tags;type:json"`
}
