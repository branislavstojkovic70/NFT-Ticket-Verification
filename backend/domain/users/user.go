package domain

import (
	domain "github.com/branislavstojkovic70/nft-ticket-verification/domain/events"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type User struct {
	ID        uuid.UUID       `json:"uuid" gorm:"type:uuid;primaryKey;column:uuid"`
	Email     string          `json:"email" gorm:"column:email"`
	Password  string          `json:"password" gorm:"column:password"`
	Wallet    string          `json:"wallet" gorm:"column:wallet"`
	Age       int             `json:"age" gorm:"column:age"`
	Location  string          `json:"location" gorm:"column:location"`
	Interests datatypes.JSON  `json:"interests" gorm:"column:interests;type:json"`
	Gender    Gender          `json:"gender" gorm:"column:gender"`
	Name      string          `json:"name" gorm:"column:name"`
	Surname   string          `json:"surname" gorm:"column:surname"`
	Role      Role            `json:"role" gorm:"column:role"`
	Tickets   []domain.Ticket `json:"tickets" gorm:"foreignKey:UserID;references:ID"`
}

type UserRequest struct {
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
	Wallet   string `json:"wallet" gorm:"column:wallet"`
	Age      int    `json:"age" gorm:"column:age"`
	Location string `json:"location" gorm:"column:location"`
	Gender   Gender `json:"gender" gorm:"column:gender"`
	Name     string `json:"name" gorm:"column:name"`
	Surname  string `json:"surname" gorm:"column:surname"`
	Role     Role   `json:"role" gorm:"column:role"`
}
