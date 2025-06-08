package repository

import (
	"github.com/cmartinez/GolangRestApiTesting/internal/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUsers() ([]domain.User, error) {
	var users []domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetUserByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	r.db.Model(&user).Association("Task").Find(&user.Task)
	return &user, nil
}

func (r *UserRepository) CreateUser(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) UpdateUser(user *domain.User) error {
	return r.db.Model(user).Updates(user).Error
}

func (r *UserRepository) DeleteUser(id uint) error {
	return r.db.Delete(&domain.User{}, id).Error
}
