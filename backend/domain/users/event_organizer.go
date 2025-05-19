package domain

import "github.com/google/uuid"

type Organizer struct {
	ID       uuid.UUID `json:"uuid" gorm:"type:uuid;primaryKey;column:uuid"`
	Email    string    `json:"email" gorm:"column:email"`
	Password string    `json:"password" gorm:"column:password"`
	Wallet   string    `json:"wallet" gorm:"column:wallet"`
	Name     string    `json:"name" gorm:"column:name"`
	Surname  string    `json:"surname" gorm:"column:surname"`
	Gender   Gender    `json:"gender" gorm:"column:gender"`
	Role     Role      `json:"role" gorm:"column:role"`
}
