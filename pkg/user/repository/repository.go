package repository

import (
	"github.com/Shakezidin/pkg/model"
	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func (a *AdminRepository) FindAdminRepo(username string) (*model.Admin, error) {
	var admin model.Admin
	if err := a.db.Where("username = ?", username).First(&admin).Error; err != nil {
		return nil, err
	}

	return &admin, nil
}

func (a *AdminRepository) CreateUser(user model.User) (*model.User, error) {
	if err := a.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func NewAdminRepo(db *gorm.DB) AdminRepository {
	return AdminRepository{
		db: db,
	}
}
