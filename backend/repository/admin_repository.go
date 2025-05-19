package repository

import (
	"errors"

	domain "github.com/branislavstojkovic70/nft-ticket-verification/domain/users"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdminRepository interface {
	GetAllAdmins() ([]domain.Admin, error)
	GetAdminByID(id uuid.UUID) (*domain.Admin, error)
	CreateAdmin(admin *domain.Admin) error
	UpdateAdmin(admin *domain.Admin) error
	DeleteAdmin(id uuid.UUID) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db}
}

func (r *adminRepository) GetAllAdmins() ([]domain.Admin, error) {
	var admins []domain.Admin
	if err := r.db.Find(&admins).Error; err != nil {
		return nil, err
	}
	return admins, nil
}

func (r *adminRepository) GetAdminByID(id uuid.UUID) (*domain.Admin, error) {
	var admin domain.Admin
	if err := r.db.First(&admin, "uuid = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &admin, nil
}

func (r *adminRepository) CreateAdmin(admin *domain.Admin) error {
	return r.db.Create(admin).Error
}

func (r *adminRepository) UpdateAdmin(admin *domain.Admin) error {
	return r.db.Save(admin).Error
}

func (r *adminRepository) DeleteAdmin(id uuid.UUID) error {
	return r.db.Delete(&domain.Admin{}, "uuid = ?", id).Error
}
