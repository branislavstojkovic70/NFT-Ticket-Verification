package service

import (
	users "github.com/branislavstojkovic70/nft-ticket-verification/domain/users"
	"github.com/branislavstojkovic70/nft-ticket-verification/repository"
	"github.com/google/uuid"
)

type UserService interface {
	GetAllUsers() ([]users.User, error)
	GetUserByID(id uuid.UUID) (*users.User, error)
	CreateUser(user *users.User) error
	UpdateUser(user *users.User) error
	DeleteUser(id uuid.UUID) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUsers() ([]users.User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) GetUserByID(id uuid.UUID) (*users.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) CreateUser(user *users.User) error {
	// Dodaj logiku validacije ako želiš (email format, itd.)
	return s.repo.CreateUser(user)
}

func (s *userService) UpdateUser(user *users.User) error {
	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(id uuid.UUID) error {
	return s.repo.DeleteUser(id)
}
