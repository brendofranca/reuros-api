package repositories

import (
	"reuros-api/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *models.UserModel) error {
	if err := r.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
