package repository

import (
	"errors"

	users "github.com/branislavstojkovic70/nft-ticket-verification/domain/users"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]users.User, error)
	GetUserByID(id uuid.UUID) (*users.User, error)
	CreateUser(user *users.User) error
	UpdateUser(user *users.User) error
	DeleteUser(id uuid.UUID) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAllUsers() ([]users.User, error) {
	var users []users.User
	if err := r.db.Preload("Tickets").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetUserByID(id uuid.UUID) (*users.User, error) {
	var user users.User
	if err := r.db.Preload("Tickets").First(&user, "uuid = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUser(user *users.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) UpdateUser(user *users.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) DeleteUser(id uuid.UUID) error {
	return r.db.Delete(&users.User{}, "uuid = ?", id).Error
}
