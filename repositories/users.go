package repositories

import (
	"waysbuck/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	CreateUser(models.User) (models.User, error)
	UpdateUser(models.User) (models.User, error)
	DeleteUser(models.User) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User

	err := r.db.First(&user, ID).Error

	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) DeleteUser(user models.User) (models.User, error) {
	err := r.db.Delete(&user).Error

	return user, err
}
