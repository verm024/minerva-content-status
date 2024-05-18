package repository

import (
	"minerva-content-status/dto"
	"minerva-content-status/models"
)

func (repo *Repository) RegisterNewUser(userData dto.RegisterNewUserDTO) (*models.User, error) {
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

func (repo *Repository) FindOneUserByUsername(username string) (*models.User, error) {
	user := models.User{}

	result := repo.db.Model(&user).Where("username = ?", username).Limit(1).Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
