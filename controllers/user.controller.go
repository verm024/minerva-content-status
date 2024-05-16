package controllers

import (
	"fmt"
	helper_response "minerva-content-status/helper"
	usecase "minerva-content-status/use-case"
	"minerva-content-status/validators"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (cont *Controller) GetAllUsers(c echo.Context) error {
	fmt.Println("Executing Get All Users")
	return c.String(http.StatusOK, "Executing Get All Users")
}

func (cont *Controller) RegisterNewUser(c echo.Context) error {
	var errors []error

	var reqBody = validators.RegisterNewUserRequestBody{}
	reqBodyErr := c.Bind(&reqBody)
	if reqBodyErr != nil {
		errors = append(errors, reqBodyErr)
	}

	validationErr := validators.ValidateRequest(reqBody)

	if validationErr != nil {
		errors = append(errors, validationErr)
	}

	token, registerError := cont.uc.RegisterNewUser(&usecase.RegisterNewUserStruct{Username: reqBody.Username, Email: reqBody.Email, Password: reqBody.Password})

	if registerError != nil {
		errors = append(errors, registerError)
	}
	return helper_response.ResponseHandler(c, map[string]interface{}{"token": token}, errors)
}

func (cont *Controller) Login(c echo.Context) error {
	var errors []error

	var reqBody = validators.LoginRequestBody{}
	reqBodyErr := c.Bind(&reqBody)
	if reqBodyErr != nil {
		errors = append(errors, reqBodyErr)
	}

	validationErr := validators.ValidateRequest(reqBody)

	if validationErr != nil {
		errors = append(errors, validationErr)
	}

	token, loginUseCaseError := cont.uc.Login(&usecase.LoginParamStruct{Username: reqBody.Username, Password: reqBody.Password})

	if loginUseCaseError != nil {
		errors = append(errors, loginUseCaseError)
	}
	return helper_response.ResponseHandler(c, map[string]interface{}{"token": token}, errors)
}
