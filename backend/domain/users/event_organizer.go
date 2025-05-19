package domain

import "github.com/google/uuid"

type Organizer struct {
	UUID     uuid.UUID `json:"uuid" db:"uuid"`
	Email    string    `json:"email" db:"email"`
	Password string    `json:"password" db:"password"`
	Wallet   string    `json:"wallet" db:"wallet"`
	Name     string    `json:"name" db:"name"`
	Surname  string    `json:"surname" db:"surname"`
	Gender   Gender    `json:"gender" db:"gender"`
	Role     Role      `json:"role" db:"role"`
}
