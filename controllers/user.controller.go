package controllers

import (
	"fmt"
	"minerva-content-status/dto"
	helper_response "minerva-content-status/helper"
	"minerva-content-status/validators"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (cont *Controller) GetAllUsers(c echo.Context) error {
	fmt.Println("Executing Get All Users")
	return c.String(http.StatusOK, "Executing Get All Users")
}

func (cont *Controller) RegisterNewUser(c echo.Context) error {
	var reqBody = dto.RegisterNewUserRequestBody{}
	reqBodyErr := c.Bind(&reqBody)
	if reqBodyErr != nil {
		return helper_response.ErrorResponseHandler(c, reqBodyErr, 400)
	}

	validationErr := validators.ValidateRequest(reqBody)

	if validationErr != nil {
		return helper_response.ErrorResponseHandler(c, validationErr, 400)
	}

	token, registerError := cont.uc.RegisterNewUser(&dto.RegisterNewUserUseCaseStruct{Username: reqBody.Username, Email: reqBody.Email, Password: reqBody.Password})

	if registerError != nil {
		return helper_response.ErrorResponseHandler(c, registerError, 400)
	}
	return helper_response.ResponseHandler(c, map[string]interface{}{"token": token})
}

func (cont *Controller) Login(c echo.Context) error {
	var reqBody = dto.LoginRequestBody{}
	reqBodyErr := c.Bind(&reqBody)
	if reqBodyErr != nil {
		return helper_response.ErrorResponseHandler(c, reqBodyErr, 400)
	}

	validationErr := validators.ValidateRequest(reqBody)

	if validationErr != nil {
		return helper_response.ErrorResponseHandler(c, validationErr, 400)
	}

	token, loginUseCaseError := cont.uc.Login(&dto.LoginParamUseCaseStruct{Username: reqBody.Username, Password: reqBody.Password})

	if loginUseCaseError != nil {
		return helper_response.ErrorResponseHandler(c, loginUseCaseError, 400)
	}
	return helper_response.ResponseHandler(c, map[string]interface{}{"token": token})
}
