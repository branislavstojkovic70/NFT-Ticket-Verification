package repository

import (
	"errors"

	users "github.com/branislavstojkovic70/nft-ticket-verification/domain/users"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdminRepository interface {
	GetAllAdmins() ([]users.Admin, error)
	GetAdminByID(id uuid.UUID) (*users.Admin, error)
	CreateAdmin(admin *users.Admin) error
	UpdateAdmin(admin *users.Admin) error
	DeleteAdmin(id uuid.UUID) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db}
}

func (r *adminRepository) GetAllAdmins() ([]users.Admin, error) {
	var admins []users.Admin
	if err := r.db.Find(&admins).Error; err != nil {
		return nil, err
	}
	return admins, nil
}

func (r *adminRepository) GetAdminByID(id uuid.UUID) (*users.Admin, error) {
	var admin users.Admin
	if err := r.db.First(&admin, "uuid = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &admin, nil
}

func (r *adminRepository) CreateAdmin(admin *users.Admin) error {
	return r.db.Create(admin).Error
}

func (r *adminRepository) UpdateAdmin(admin *users.Admin) error {
	return r.db.Save(admin).Error
}

func (r *adminRepository) DeleteAdmin(id uuid.UUID) error {
	return r.db.Delete(&users.Admin{}, "uuid = ?", id).Error
}
