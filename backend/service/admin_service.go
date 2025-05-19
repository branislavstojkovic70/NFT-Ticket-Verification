package service

import (
	users "github.com/branislavstojkovic70/nft-ticket-verification/domain/users"
	"github.com/branislavstojkovic70/nft-ticket-verification/repository"
	"github.com/google/uuid"
)

type AdminService interface {
	GetAllAdmins() ([]users.Admin, error)
	GetAdminByID(id uuid.UUID) (*users.Admin, error)
	CreateAdmin(admin *users.Admin) error
	UpdateAdmin(admin *users.Admin) error
	DeleteAdmin(id uuid.UUID) error
}

type adminService struct {
	repo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository) AdminService {
	return &adminService{repo: repo}
}

func (s *adminService) GetAllAdmins() ([]users.Admin, error) {
	return s.repo.GetAllAdmins()
}

func (s *adminService) GetAdminByID(id uuid.UUID) (*users.Admin, error) {
	return s.repo.GetAdminByID(id)
}

func (s *adminService) CreateAdmin(admin *users.Admin) error {
	return s.repo.CreateAdmin(admin)
}

func (s *adminService) UpdateAdmin(admin *users.Admin) error {
	return s.repo.UpdateAdmin(admin)
}

func (s *adminService) DeleteAdmin(id uuid.UUID) error {
	return s.repo.DeleteAdmin(id)
}
