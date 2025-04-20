package users

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetUserByID(id int) (*User, error) {
	var user User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByUUID(uid uuid.UUID) (*User, error) {
	var user User
	if err := r.DB.Where("uuid = ?", uid).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*User, error) {
	var user User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(user *User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepository) DeleteUserByID(id int) error {
	result := r.DB.Delete(&User{}, id)
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return result.Error
}

func (r *UserRepository) DeleteUserByUUID(uid uuid.UUID) error {
	result := r.DB.Where("uuid = ?", uid).Delete(&User{})
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return result.Error
}
