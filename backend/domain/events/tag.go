package domain

import "github.com/google/uuid"

type Tag struct {
	UUID uuid.UUID `json:"uuid" db:"uuid"`
	Name string    `json:"name" db:"name"`
}
