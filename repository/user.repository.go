package repository

import (
	"log"
	"minerva-content-status/models"
)

func (repo *Repository) RegisterNewUser(userData RegisterNewUserStruct) (*models.User, error) {
	user := models.User{
		Username: userData.Username,
		Password: userData.Password,
		Email:    userData.Email,
	}
	result := repo.db.Create(&user)
	log.Println(result)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo *Repository) FindOneUserByUsername(username string) (*models.User, error) {
	user := models.User{}

	result := repo.db.Where("username = ?", username).Limit(1).Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
