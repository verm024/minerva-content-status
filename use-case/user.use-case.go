package usecase

import (
	"minerva-content-status/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type createTokenParamStruct struct {
	username string
	email    string
}

func createToken(payload *createTokenParamStruct) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": payload.username, "email": payload.email, "exp": time.Now().Add(time.Hour * 24).Unix()})

	tokenStr, jwtSignError := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return tokenStr, jwtSignError
}

func (uc *UseCase) RegisterNewUser(userData *RegisterNewUserStruct) (string, error) {
	hashedPass, hashErr := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)

	if hashErr != nil {
		return "", hashErr
	}

	registeredUser, regError := uc.repo.RegisterNewUser(repository.RegisterNewUserStruct{Username: userData.Username, Email: userData.Email, Password: string(hashedPass)})

	if regError != nil {
		return "", regError
	}

	tokenStr, jwtSignError := createToken(&createTokenParamStruct{username: registeredUser.Username, email: registeredUser.Email})

	if jwtSignError != nil {
		return "", jwtSignError
	}
	return tokenStr, nil
}

func (uc *UseCase) Login(loginData *LoginParamStruct) (string, error) {
	oneUser, findUserErr := uc.repo.FindOneUserByUsername(loginData.Username)

	if findUserErr != nil {
		return "", findUserErr
	}

	hashedPass := oneUser.Password

	compareErr := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(loginData.Password))

	if compareErr != nil {
		return "", compareErr
	}

	tokenStr, jwtSignError := createToken(&createTokenParamStruct{username: oneUser.Username, email: oneUser.Email})

	if jwtSignError != nil {
		return "", jwtSignError
	}
	return tokenStr, nil
}
