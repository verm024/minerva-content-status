package repository

import (
	"log"
	"minerva-content-status/models"
)

func (repo *Repository) RegisterNewUser(userData RegisterNewUserStruct) error {
	user := models.User{
		Username: userData.Username,
		Password: userData.Password,
		Email:    userData.Email,
	}
	result := repo.db.Create(&user)
	log.Println(result)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
