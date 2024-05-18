package usecase

import (
	"minerva-content-status/dto"
	"minerva-content-status/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repo *repository.UserRepository
}

func InitializeUserUseCase(repo *repository.UserRepository) *UserUseCase {
	uc := UserUseCase{repo}
	return &uc
}

type createTokenParamStruct struct {
	username string
	email    string
	user_id  uint64
	role     string
}

func createToken(payload *createTokenParamStruct) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": payload.username, "email": payload.email, "exp": time.Now().Add(time.Hour * 24).Unix(), "user_id": payload.user_id, "role": payload.role})

	tokenStr, jwtSignError := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return tokenStr, jwtSignError
}

func (uc *UserUseCase) RegisterNewUser(userData *dto.RegisterNewUserDTO) (string, error) {
	hashedPass, hashErr := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)

	if hashErr != nil {
		return "", hashErr
	}

	registeredUser, regError := uc.repo.RegisterNewUser(dto.RegisterNewUserDTO{Username: userData.Username, Email: userData.Email, Password: string(hashedPass)})

	if regError != nil {
		return "", regError
	}

	tokenStr, jwtSignError := createToken(&createTokenParamStruct{username: registeredUser.Username, email: registeredUser.Email, user_id: registeredUser.UserID, role: registeredUser.Role})

	if jwtSignError != nil {
		return "", jwtSignError
	}
	return tokenStr, nil
}

func (uc *UserUseCase) Login(loginData *dto.LoginDTO) (string, error) {
	oneUser, findUserErr := uc.repo.FindOneUserByUsername(loginData.Username)

	if findUserErr != nil {
		return "", findUserErr
	}

	hashedPass := oneUser.Password

	compareErr := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(loginData.Password))

	if compareErr != nil {
		return "", compareErr
	}

	tokenStr, jwtSignError := createToken(&createTokenParamStruct{username: oneUser.Username, email: oneUser.Email, user_id: oneUser.UserID, role: oneUser.Role})

	if jwtSignError != nil {
		return "", jwtSignError
	}
	return tokenStr, nil
}
