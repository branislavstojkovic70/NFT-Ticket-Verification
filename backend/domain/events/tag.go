package domain

import "github.com/google/uuid"

type Tag struct {
	ID   uuid.UUID `json:"uuid" gorm:"type:uuid;primaryKey;column:uuid"`
	Name string    `json:"name" gorm:"column:name"`
}
