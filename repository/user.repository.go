package repository

import (
	"minerva-content-status/dto"
	"minerva-content-status/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func InitializeUserRepository(db *gorm.DB) *UserRepository {
	repo := UserRepository{db}
	return &repo
}

func (repo *UserRepository) RegisterNewUser(userData dto.RegisterNewUserDTO) (*models.User, error) {
	user := models.User{
		Username: userData.Username,
		Password: userData.Password,
		Email:    userData.Email,
	}
	result := repo.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo *UserRepository) FindOneUserByUsername(username string) (*models.User, error) {
	user := models.User{}

	result := repo.db.Model(&user).Where("username = ?", username).Limit(1).Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
