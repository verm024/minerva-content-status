package usecase

import (
	"minerva-content-status/repository"

	"golang.org/x/crypto/bcrypt"
)

func (uc *UseCase) RegisterNewUser(userData *RegisterNewUserStruct) error {
	hashedPass, hashErr := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)

	if hashErr != nil {
		return hashErr
	}

	uc.repo.RegisterNewUser(repository.RegisterNewUserStruct{Username: userData.Username, Email: userData.Email, Password: string(hashedPass)})

	return nil
}
