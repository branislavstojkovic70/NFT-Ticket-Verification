package domain

import (
	domain "github.com/branislavstojkovic70/nft-ticket-verification/domain/events"
	"github.com/google/uuid"
)

type User struct {
	UUID      uuid.UUID       `json:"uuid" db:"uuid"`
	Email     string          `json:"email" db:"email"`
	Password  string          `json:"password" db:"password"`
	Wallet    string          `json:"wallet" db:"wallet"`
	Age       int             `json:"age" db:"age"`
	Location  string          `json:"location" db:"location"`
	Interests []string        `json:"interests" db:"interests"`
	Gender    Gender          `json:"gender" db:"gender"`
	Name      string          `json:"name" db:"name"`
	Surname   string          `json:"surname" db:"surname"`
	Role      Role            `json:"role" db:"role"`
	Tickets   []domain.Ticket `json:"tickets" db:"-"`
}
