package controllers

import (
	"minerva-content-status/dto"
	helper_response "minerva-content-status/helper"
	usecase "minerva-content-status/use-case"
	"minerva-content-status/validators"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	uc usecase.UserUseCaseInterface
}

func InitializeUserController(uc usecase.UserUseCaseInterface) *UserController {
	cont := UserController{uc}
	return &cont
}

func (cont *UserController) RegisterNewUser(c echo.Context) error {
	var reqBody = dto.RegisterNewUserRequestDTO{}
	reqBodyErr := c.Bind(&reqBody)
	if reqBodyErr != nil {
		return helper_response.ErrorResponseHandler(c, reqBodyErr, 400)
	}

	validationErr := validators.ValidateRequest(reqBody)

	if validationErr != nil {
		return helper_response.ErrorResponseHandler(c, validationErr, 400)
	}

	token, registerError := cont.uc.RegisterNewUser(&dto.RegisterNewUserDTO{Username: reqBody.Username, Email: reqBody.Email, Password: reqBody.Password})

	if registerError != nil {
		return helper_response.ErrorResponseHandler(c, registerError, 400)
	}
	return helper_response.ResponseHandler(c, map[string]interface{}{"token": token})
}

func (cont *UserController) Login(c echo.Context) error {
	var reqBody = dto.LoginRequestDTO{}
	reqBodyErr := c.Bind(&reqBody)
	if reqBodyErr != nil {
		return helper_response.ErrorResponseHandler(c, reqBodyErr, http.StatusBadRequest)
	}

	validationErr := validators.ValidateRequest(reqBody)

	if validationErr != nil {
		return helper_response.ErrorResponseHandler(c, validationErr, http.StatusBadRequest)
	}

	token, loginUseCaseError := cont.uc.Login(&dto.LoginDTO{Username: reqBody.Username, Password: reqBody.Password})

	if loginUseCaseError != nil {
		return helper_response.ErrorResponseHandler(c, loginUseCaseError, http.StatusBadRequest)
	}
	return helper_response.ResponseHandler(c, map[string]interface{}{"token": token})
}
